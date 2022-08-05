package db

import (
	"e-mochi-app/app/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//unused
func Init() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Mochi{})

	return db
}
