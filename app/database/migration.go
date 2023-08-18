package database

import (
	"log"

	questionaireData "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/data"
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		questionaireData.Question{},
		questionaireData.Choice{},
	)

	if err != nil {
		log.Println("Error Migration")
	}
	log.Println("Migration Done")
}
