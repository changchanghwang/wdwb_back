package events

import (
	calendarDate "github.com/changchanghwang/wdwb_back/internal/libs/calendar-date"
	"github.com/google/uuid"
)

type FilingRegisteredEvent struct {
	FilingId        int                       `json:"filingId"`
	InvestorId      uuid.UUID                 `json:"investorId"`
	FilingType      string                    `json:"filingType"`
	AccessionNumber string                    `json:"accessionNumber"`
	FilledOn        calendarDate.CalendarDate `json:"filledOn"`
	ReportedOn      calendarDate.CalendarDate `json:"reportedOn"`
	Link            string                    `json:"link"`
	Year            int                       `json:"year"`
	Quarter         int                       `json:"quarter"`
}
