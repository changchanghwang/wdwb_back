package secClient

import calendarDate "github.com/changchanghwang/wdwb_back/internal/libs/calendar-date"

type FilingDTO struct {
	AccessionNumber string
	FilingDate      calendarDate.CalendarDate
	ReportDate      calendarDate.CalendarDate
	Form            string
	InfoTableLink   string
}

type HoldingDto struct {
	CompanyName  string
	TitleOfClass string
	Cusip        string
	Value        int
	StockShares  int
}

type CompanyDto struct {
	CIK            string   `json:"cik"`
	Name           string   `json:"name"`
	Sic            string   `json:"sic"`
	SicDescription string   `json:"sicDescription"`
	OwnerOrg       string   `json:"ownerOrg"`
	Tickers        []string `json:"tickers"`
	Exchanges      []string `json:"exchanges"`
	Filings        struct {
		Recent struct {
			AccessionNumber []string                    `json:"accessionNumber"`
			FilingDate      []calendarDate.CalendarDate `json:"filingDate"`
			ReportDate      []calendarDate.CalendarDate `json:"reportDate"`
			Form            []string                    `json:"form"`
		} `json:"recent"`
	} `json:"filings"`
	Symbols []symbol
}

type symbol struct {
	Ticker   string
	Exchange string
}

type StockDTO struct {
	Name     string
	Ticker   string
	CIK      string
	CUSIP    string
	Exchange string
	Sic      string
	Industry string
	Sector   string
}
