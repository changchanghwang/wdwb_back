package domain

import (
	calendarDate "github.com/changchanghwang/wdwb_back/internal/libs/calendar-date"
	"github.com/changchanghwang/wdwb_back/internal/libs/ddd"
	"github.com/changchanghwang/wdwb_back/internal/services/filings/domain/events"
	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/google/uuid"
)

type Filing struct {
	ddd.Aggregate
	Id              int                       `json:"id" gorm:"column:id; primaryKey; autoIncrement:true"`
	InvestorId      uuid.UUID                 `json:"investorId" gorm:"type:varchar(36); column:investorId; not null;"`
	Type            string                    `json:"type" gorm:"type:varchar(20); not null;"`
	AccessionNumber string                    `json:"accessionNumber" gorm:"type:varchar(50); column:accessionNumber; not null;"`
	FilledOn        calendarDate.CalendarDate `json:"filledOn" gorm:"type:date; column:filledOn; not null;"`
	ReportedOn      calendarDate.CalendarDate `json:"reportedOn" gorm:"type:date; column:reportedOn; not null;"`
	Link            string                    `json:"link" gorm:"type:varchar(255); not null;"`
	Year            int                       `json:"year" gorm:"type:int; not null;"`
	Quarter         int                       `json:"quarter" gorm:"type:int; not null;"`
}

func (f *Filing) TableName() string {
	return "filing"
}

func New(filingType, accessionNumber, link string, investorId uuid.UUID, filledOn, reportedOn calendarDate.CalendarDate) (*Filing, error) {
	year, err := reportedOn.Year()
	if err != nil {
		return nil, applicationError.Wrap(err)
	}
	quarter, err := reportedOn.Quarter()
	if err != nil {
		return nil, applicationError.Wrap(err)
	}

	filing := &Filing{
		InvestorId:      investorId,
		Type:            filingType,
		AccessionNumber: accessionNumber,
		FilledOn:        filledOn,
		ReportedOn:      reportedOn,
		Year:            year,
		Quarter:         quarter,
		Link:            link,
	}

	filing.PublishEvent(
		ddd.NewEvent(
			"FilingRegisteredEvent",
			&events.FilingRegisteredEvent{
				FilingId:        filing.Id,
				InvestorId:      filing.InvestorId,
				FilingType:      filing.Type,
				AccessionNumber: filing.AccessionNumber,
				FilledOn:        filing.FilledOn,
				ReportedOn:      filing.ReportedOn,
				Link:            filing.Link,
			},
		),
	)

	return filing, nil
}
