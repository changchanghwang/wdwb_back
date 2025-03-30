package ddd

import (
	"encoding/json"
	"time"
)

type Event struct {
	Id         int       `json:"id" gorm:"column:id; primaryKey; autoIncrement:true"`
	Type       string    `json:"type" gorm:"type:varchar(255); not null;"`
	Data       string    `json:"data" gorm:"type:text; not null;"`
	OccurredAt time.Time `json:"occurredAt" gorm:"type:datetime; column:occurredAt; not null;"`
}

func (e *Event) TableName() string {
	return "event"
}

func NewEvent(eventType string, data any) *Event {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	return &Event{
		Type:       eventType,
		Data:       string(jsonData),
		OccurredAt: time.Now(),
	}
}
