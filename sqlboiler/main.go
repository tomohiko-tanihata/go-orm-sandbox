package main

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"sqlboiler/models"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {
	// Open handle to database like normal
	db, err := sql.Open("postgres", "dbname=postgres user=postgres password=postgres host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	boil.SetDB(db)
	ctx := context.Background()

	prepareData(db)

	student, _ := models.Students(
		qm.Load(qm.Rels(models.StudentRels.Class, models.ClassRels.Grade)),
		qm.InnerJoin("classes on students.class_id = classes.id"),
		qm.InnerJoin("grades on classes.grade_id = grades.id"),
	).All(ctx, db)

	for _, s := range student {
		println(s.Name)
		println(s.R.Class.Name)
		println(s.R.Class.R.Grade.Name)
	}

}

func prepareData(db *sql.DB) {
	for i := 0; i < 3; i++ {
		grade := &models.Grade{
			Name: "Grade " + strconv.Itoa(i),
		}
		err := InsertGrade(db, grade)
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < 5; i++ {
		class := &models.Class{
			Name:    "Class" + strconv.Itoa(i),
			GradeID: i/2 + 1,
		}
		err := InsertClass(db, class)
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < 12; i++ {
		student := &models.Student{
			Name:    "Student" + strconv.Itoa(i),
			ClassID: i/3 + 1,
		}
		err := InsertStudent(db, student)
		if err != nil {
			panic(err)
		}
	}
}

func InsertGrade(db *sql.DB, grade *models.Grade) error {
	err := grade.Insert(context.Background(), db, boil.Infer())
	return err
}

func InsertClass(db *sql.DB, class *models.Class) error {
	err := class.Insert(context.Background(), db, boil.Infer())
	return err
}

func InsertStudent(db *sql.DB, student *models.Student) error {
	err := student.Insert(context.Background(), db, boil.Infer())
	return err
}
