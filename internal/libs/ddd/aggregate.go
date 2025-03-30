package ddd

import (
	"time"

	"gorm.io/gorm"
)

type Aggregate struct {
	CreatedAt time.Time `json:"-" gorm:"column:createdAt;autoCreateTime:nano; not null;"`
	UpdatedAt time.Time `json:"-" gorm:"column:updatedAt;autoUpdateTime:nano; not null;"`

	Events []*Event `gorm:"-"`
}

func (a *Aggregate) GetPublishedEvents() []*Event {
	if a.Events == nil {
		return []*Event{}
	}

	return a.Events
}

func (a *Aggregate) PublishEvent(event *Event) {
	if a.Events == nil {
		a.Events = []*Event{}
	}

	a.Events = append(
		a.Events,
		event,
	)
}

type SoftDeletableAggregate struct {
	Aggregate
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"column:deletedAt"`
}
