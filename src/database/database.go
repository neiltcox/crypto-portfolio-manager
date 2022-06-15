package database

import (
	"fmt"
	"log"

	"github.com/jaksonkallio/coinbake/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var handle *gorm.DB

func Connect(cfg config.Config) error {
	log.Printf("Connecting to database at %s", cfg.Database.Host)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Name,
	)
	newHandle, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	handle = newHandle

	return nil
}

func Handle() *gorm.DB {
	return handle
}
