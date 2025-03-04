package database

import (
	"fmt"
	"gin-boiler-plate/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(driver, dsn string, autoMigrate bool) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch driver {
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", driver)
	}

	if autoMigrate {
		if err := db.AutoMigrate(model.NewModels()...); err != nil {
			return nil, fmt.Errorf("failed to auto migrate: %s", err)
		}
	}

	return db, err
}
