package ormbenchmark

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Gormcrud struct {
}

var gormdb *gorm.DB

func init() {
	gormDB, err := gorm.Open("mysql", "root:Tomato-123@/student")
	if err != nil {
		log.Panic(err)
	}

	gormdb = gormDB
}

func (Gormcrud) SelectAll(sql string) {
	var students []Student
	gormdb.Find(&students)
}

func (Gormcrud) SelectLeftJoin(sql string) {
	var sg []StudentGrade
	gormdb.Table("student").Select("*").Joins("left join grade on student.id = grade.student_id").Find(&sg)
}

func (Gormcrud) Insert(sql string) {
	tx := gormdb.Begin()
	student := Student{
		0, "JINX", 1, "Japan",
	}
	gormdb.Save(&student)
	tx.Commit()
}

func (Gormcrud) Update(sql string) {
	tx := gormdb.Begin()
	student := Student{
		3, "NGINX", 0, "LONDON",
	}
	gormdb.Save(&student)
	tx.Commit()
}

func (Gormcrud) Delete(sql string) {
	tx := gormdb.Begin()
	student := Student{
		Id: 100,
	}
	gormdb.Delete(&student)
	tx.Commit()
}
