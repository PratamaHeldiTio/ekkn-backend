package migration

import (
	period "backend-ekkn/modules/period/domain"
	student "backend-ekkn/modules/student/domain"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(&period.Period{}, student.Student{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migration Success")
}
