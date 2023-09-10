package bun

import (
	"context"
	"github.com/uptrace/bun"
)

func InsertGrade(db *bun.DB, grade *Grade) error {
	_, err := db.NewInsert().Model(grade).Exec(context.Background())
	return err
}

func InsertClass(db *bun.DB, class *Class) error {
	_, err := db.NewInsert().Model(class).Exec(context.Background())
	return err
}

func InsertStudent(db *bun.DB, student *Student) error {
	_, err := db.NewInsert().Model(student).Exec(context.Background())
	return err
}
