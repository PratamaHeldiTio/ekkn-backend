package migration

import (
	group "backend-ekkn/modules/group/domain"
	period "backend-ekkn/modules/period/domain"
	student "backend-ekkn/modules/student/domain"
	studentRegistration "backend-ekkn/modules/student_registration/domain"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&period.Period{},
		&student.Student{},
		&studentRegistration.StudentRegistration{},
		&group.Group{},
	)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migration Success")
}
