package db

import (
	"fmt"
	"log"

	"github.com/changchanghwang/wdwb_back/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	filingDomain "github.com/changchanghwang/wdwb_back/internal/services/filings/domain"
	holdingDomain "github.com/changchanghwang/wdwb_back/internal/services/holdings/domain"
	investorDomain "github.com/changchanghwang/wdwb_back/internal/services/investors/domain"
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

	db.AutoMigrate(&holdingDomain.Holding{}, &investorDomain.Investor{}, &filingDomain.Filing{})

	return db
}
