package ddd

import "gorm.io/gorm"

type ApplicationService struct {
	Manager *gorm.DB
}
