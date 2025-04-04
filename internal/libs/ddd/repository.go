package ddd

import (
	"fmt"
	"net/http"

	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
	"github.com/changchanghwang/wdwb_back/pkg/util"
	"gorm.io/gorm"
)

type EventHolder interface {
	GetPublishedEvents() []*Event
}

type Repository[T EventHolder] struct {
	Manager *gorm.DB
}

func (r *Repository[T]) Save(db *gorm.DB, entities []T) error {
	if db == nil {
		db = r.Manager
	}

	if len(entities) == 0 {
		return nil
	}

	if err := db.Save(entities).Error; err != nil {
		return applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to save. %s", err.Error()), "")
	}

	util.Map(entities, func(entity T) []*Event {
		events := entity.GetPublishedEvents()
		if len(events) > 0 {
			db.Save(events)
		}
		return nil
	})

	return nil
}
