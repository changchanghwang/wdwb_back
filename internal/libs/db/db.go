package db

import (
	"fmt"
	"log"

	"github.com/changchanghwang/wdwb_back/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// logger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // Slow SQL threshold
	// 		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	// 		Colorful:                  true,
	// 		LogLevel:                  logger.Info, // Log level
	// 	},
	// )

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger,
	})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func ApplyOptions(options *FindOptions, orderOptions *OrderOptions) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if options != nil {
			if options.Offset != 0 {
				db = db.Offset(options.Offset)
			}
			if options.Limit != 0 {
				db = db.Limit(options.Limit)
			}
			if options.GroupBy != "" {
				db = db.Group(options.GroupBy)
			}
		}

		if orderOptions != nil {
			db = db.Order(fmt.Sprintf("%s %s", orderOptions.OrderBy, orderOptions.Direction))
		}

		return db
	}
}
