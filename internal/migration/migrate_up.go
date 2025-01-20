package migration

import (
	"github.com/RhoNit/budgetapi/internal/models"
	"gorm.io/gorm"
)

func MigrateDBUp(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}
