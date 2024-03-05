package migration

import (
	"PresentationProject/internal/model"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.User{},
	); err != nil {
		return err
	}
	return nil
}
