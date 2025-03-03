package secClient

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	applicationError "github.com/changchanghwang/wdtb_back/pkg/application-error"
)

type SecClient struct {
	client *http.Client
}

type UserAgentTransport struct {
	Transport http.RoundTripper
	UserAgent string
}

func (t *UserAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.UserAgent)
	return t.Transport.RoundTrip(req)
}

func New() *SecClient {
	transport := &UserAgentTransport{
		Transport: http.DefaultTransport,
		UserAgent: "wdwb/1.0 (window95pill@gmail.com)",
	}

	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	return &SecClient{
		client: httpClient,
	}
}

func (s *SecClient) GetFilings(cik string) ([]*FilingDTO, error) {
	if len(cik) < 10 {
		cik = strings.Repeat("0", 10-len(cik)) + cik
	} else if len(cik) > 10 {
		return nil, applicationError.New(http.StatusBadRequest, fmt.Sprintf("CIK must be 10 digits : %s", cik), "")
	}

	resp, err := s.client.Get(fmt.Sprintf("https://data.sec.gov/submissions/CIK%s.json", cik))
	if err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to get filings for CIK: %s\n%v", cik, err.Error()), "")
	}
	defer resp.Body.Close()

	var data struct {
		CIK     string `json:"cik"`
		Name    string `json:"name"`
		Filings struct {
			Recent struct {
				AccessionNumber []string `json:"accessionNumber"`
				FilingDate      []string `json:"filingDate"`
				ReportDate      []string `json:"reportDate"`
				Form            []string `json:"form"`
			} `json:"recent"`
		} `json:"filings"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to decode response body for CIK: %s\n%v", cik, err.Error()), "")
	}

	filingsMap := make(map[string]*FilingDTO)
	for i, form := range data.Filings.Recent.Form {
		if form == "13F-HR" || form == "13F-HR/A" {
			reportDate := data.Filings.Recent.ReportDate[i]
			filingDate := data.Filings.Recent.FilingDate[i]
			if reportDate == "" {
				continue
			}
			existing, exists := filingsMap[reportDate]
			if !exists || filingDate > existing.FilingDate {
				filingsMap[reportDate] = &FilingDTO{
					AccessionNumber: data.Filings.Recent.AccessionNumber[i],
					FilingDate:      filingDate,
					ReportDate:      reportDate,
					Form:            form,
				}
			}
		}
	}

	filings := make([]*FilingDTO, 0, len(filingsMap))
	for _, filing := range filingsMap {
		filings = append(filings, filing)
	}

	chunkSize := 5
	for i := 0; i < len(filings); i += chunkSize {
		end := i + chunkSize
		if end > len(filings) {
			end = len(filings)
		}

		filingsChunk := filings[i:end]
		for _, filing := range filingsChunk {
			baseUrl := fmt.Sprintf("https://www.sec.gov/Archives/edgar/data/%s/%s", cik, strings.ReplaceAll(filing.AccessionNumber, "-", ""))

			res, err := s.client.Get(fmt.Sprintf("%s/%s-index.html", baseUrl, filing.AccessionNumber))
			if err != nil {
				return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to get info table link for CIK: %s\n%v", cik, err.Error()), "")
			}
			defer res.Body.Close()

			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to parse response body for CIK: %s\n%v", cik, err.Error()), "")
			}

			var infoTableLink string
			var nonPrimaryDocLink string
			var foundInfoTable bool

			doc.Find("table.tableFile tr").Each(func(i int, s *goquery.Selection) {
				if i == 0 {
					return
				}

				if s.Text() != "" && strings.Contains(s.Text(), "Complete submission text file") {
					return
				}

				rowText := s.Text()
				if strings.Contains(rowText, "INFORMATION TABLE") {
					if !foundInfoTable {
						link, exists := s.Find("a").Attr("href")
						if exists {
							infoTableLink = link
							foundInfoTable = true
						}
					}
				}

				if nonPrimaryDocLink == "" {
					s.Find("a").Each(func(j int, a *goquery.Selection) {
						linkText := a.Text()
						link, exists := a.Attr("href")
						if exists && !strings.Contains(linkText, "primary_doc") {
							nonPrimaryDocLink = link
						}
					})
				}
			})

			if foundInfoTable {
				splittedFoundInfoTableLink := strings.Split(infoTableLink, "/")
				filing.InfoTableLink = fmt.Sprintf("%s/%s", baseUrl, splittedFoundInfoTableLink[len(splittedFoundInfoTableLink)-1])
			} else if nonPrimaryDocLink != "" {
				splittedNonPrimaryDocLink := strings.Split(nonPrimaryDocLink, "/")
				filing.InfoTableLink = fmt.Sprintf("%s/%s", baseUrl, splittedNonPrimaryDocLink[len(splittedNonPrimaryDocLink)-1])
			} else {
				return nil, fmt.Errorf("조건에 맞는 파일 링크를 찾을 수 없습니다")
			}
		}
		time.Sleep(1 * time.Second)
	}

	return filings, nil
}

func (s *SecClient) ParseInfoTable(url string) ([]HoldingDto, error) {
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, e := io.ReadAll(resp.Body)
	if e != nil {
		return nil, e
	}

	if strings.Contains(url, ".xml") {
		return s.parseXml(body)
	}

	if strings.Contains(url, ".txt") {
		return s.parseText(body)
	}

	return nil, applicationError.New(http.StatusBadRequest, "Invalid URL", "")
}

type InformationTable struct {
	XMLName   xml.Name    `xml:"informationTable,ns1:informationTable"`
	InfoTable []InfoTable `xml:"infoTable,ns1:infoTable"`
}

type InfoTable struct {
	NameOfIssuer         string            `xml:"nameOfIssuer,ns1:nameOfIssuer"`
	TitleOfClass         string            `xml:"titleOfClass,ns1:titleOfClass"`
	Cusip                string            `xml:"cusip,ns1:cusip"`
	Value                string            `xml:"value,ns1:value"`
	ShrsOrPrnAmt         SharesOrPrincipal `xml:"shrsOrPrnAmt,ns1:shrsOrPrnAmt"`
	InvestmentDiscretion string            `xml:"investmentDiscretion,ns1:investmentDiscretion"`
	VotingAuthority      VotingAuthority   `xml:"votingAuthority,ns1:votingAuthority"`
}

type SharesOrPrincipal struct {
	SshPrnamt     string `xml:"sshPrnamt,ns1:sshPrnamt"`
	SshPrnamtType string `xml:"sshPrnamtType,ns1:sshPrnamtType"`
}

type VotingAuthority struct {
	Sole   string `xml:"Sole,ns1:Sole"`
	Shared string `xml:"Shared,ns1:Shared"`
	None   string `xml:"None,ns1:None"`
}

func (s *SecClient) parseXml(infoTable []byte) ([]HoldingDto, error) {
	var informationTable InformationTable

	err := xml.Unmarshal(infoTable, &informationTable)
	if err != nil {
		return nil, applicationError.New(http.StatusInternalServerError, fmt.Sprintf("XML 파싱 오류: %v", err.Error()), "")
	}

	holdings := make([]HoldingDto, 0, len(informationTable.InfoTable))
	for _, infoTable := range informationTable.InfoTable {
		value, err := strconv.Atoi(strings.ReplaceAll(infoTable.Value, ",", ""))
		if err != nil {
			return nil, err
		}

		stockShares, err := strconv.Atoi(strings.ReplaceAll(infoTable.ShrsOrPrnAmt.SshPrnamt, ",", ""))
		if err != nil {
			return nil, err
		}

		holdings = append(holdings, HoldingDto{
			CompanyName: infoTable.NameOfIssuer,
			Cusip:       infoTable.Cusip,
			Value:       value,
			StockShares: stockShares,
		})
	}
	return holdings, nil
}

func (s *SecClient) parseText(infoTable []byte) ([]HoldingDto, error) {
	textData := string(infoTable)
	lines := strings.Split(textData, "\n")

	var headerLine, dashedLine string
	var dataLines []string
	inTable := false
	dashedFound := false

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, "<TABLE>") {
			inTable = true
			continue
		}

		if !inTable {
			continue
		}

		if headerLine == "" && strings.Contains(trimmed, "NAME OF ISSUER") {
			headerLine = trimmed
			continue
		}
		if !dashedFound && strings.Contains(trimmed, "-----") && strings.Count(trimmed, "-") > 10 {
			dashedLine = trimmed
			dashedFound = true
			for j := i + 1; j < len(lines); j++ {
				dataLine := strings.TrimSpace(lines[j])
				if strings.Contains(dataLine, "</TABLE>") {
					break
				}
				if dataLine != "" {
					dataLines = append(dataLines, dataLine)
				}
			}
			break
		}
	}

	if headerLine == "" || !dashedFound {
		return nil, fmt.Errorf("테이블 헤더나 dashed line을 찾지 못했습니다")
	}

	boundaries := computeBoundaries(dashedLine)
	if len(boundaries) == 0 {
		return nil, fmt.Errorf("dashed line으로부터 필드 경계를 계산할 수 없습니다")
	}

	var holdings []HoldingDto
	for _, line := range dataLines {
		if len(line) < boundaries[len(boundaries)-1].end {
			line += strings.Repeat(" ", boundaries[len(boundaries)-1].end-len(line))
		}

		fields := make([]string, len(boundaries))
		for i, b := range boundaries {
			var field string
			if len(line) < b.end {
				field = strings.TrimSpace(line[b.start:])
			} else {
				field = strings.TrimSpace(line[b.start:b.end])
			}
			fields[i] = field
		}

		if len(fields) < 5 {
			continue
		}

		companyName := fields[0]
		cusip := fields[2]
		valueStr := fields[3]
		sharesStr := fields[4]
		sharesParts := strings.Fields(sharesStr)
		if len(sharesParts) > 0 {
			sharesStr = sharesParts[0]
		}

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return nil, fmt.Errorf("value 파싱 오류: %v", err)
		}

		shares, err := strconv.Atoi(sharesStr)
		if err != nil {
			return nil, fmt.Errorf("shares 파싱 오류: %v", err)
		}

		holdings = append(holdings, HoldingDto{
			CompanyName: companyName,
			Cusip:       cusip,
			Value:       value,
			StockShares: shares,
		})
	}

	if len(holdings) == 0 {
		return nil, fmt.Errorf("파싱된 데이터가 없습니다")
	}

	return holdings, nil
}
