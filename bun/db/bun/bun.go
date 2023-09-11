package bun

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"strconv"
)

func Run() {
	ctx := context.Background()
	db := bun.NewDB(
		sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"))),
		pgdialect.New(),
	)

	prepareData(db)

	// Joinの実行
	var students []Student
	err := db.NewSelect().
		Model(&students).
		Relation("Class").
		Relation("Class.Grade").
		Where("class__grade.id = ?", 1).
		Scan(ctx)

	if err != nil {
		panic(err)
	}

	// 結果を表示
	for _, student := range students {
		fmt.Printf("Student: %s, Class: %s, Grade: %s\n", student.Name, student.Class.Name, student.Class.Grade.Name)
	}

}

func prepareData(db *bun.DB) {
	for i := 0; i < 3; i++ {
		grade := &Grade{
			Name: "Grade " + strconv.Itoa(i),
		}
		err := InsertGrade(db, grade)
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < 5; i++ {
		class := &Class{
			Name:    "Class" + strconv.Itoa(i),
			GradeID: int64(i)/2 + 1,
		}
		err := InsertClass(db, class)
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < 12; i++ {
		student := &Student{
			Name:    "Student" + strconv.Itoa(i),
			ClassID: int64(i)/3 + 1,
		}
		err := InsertStudent(db, student)
		if err != nil {
			panic(err)
		}
	}
}
