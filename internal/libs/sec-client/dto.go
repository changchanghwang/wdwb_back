package secClient

type FilingDTO struct {
	AccessionNumber string
	FilingDate      string
	ReportDate      string
	Form            string
	InfoTableLink   string
}

type HoldingDto struct {
	CompanyName string
	Cusip       string
	Value       int
	StockShares int
}
