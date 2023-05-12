package migration

import (
	admin "backend-ekkn/modules/admin/domain"
	group "backend-ekkn/modules/group/domain"
	lecture "backend-ekkn/modules/lecturer/domain"
	lecturerRegistration "backend-ekkn/modules/lecturer_registration/domain"
	logbook "backend-ekkn/modules/logbook/domain"
	period "backend-ekkn/modules/period/domain"
	student "backend-ekkn/modules/student/domain"
	studentRegistration "backend-ekkn/modules/student_registration/domain"
	village "backend-ekkn/modules/village/domain"
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
		&village.Village{},
		&logbook.Logbook{},
		&admin.Admin{},
		&lecture.Lecturer{},
		&lecturerRegistration.LecturerRegistration{},
	)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migration Success")
}
