package models

import (
	"log"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		log.Printf("[ERROR]failed to auto migrate: %v", err)
		return err
	}
	return nil
}
