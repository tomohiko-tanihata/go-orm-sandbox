package bun

import "github.com/uptrace/bun"

type Grade struct {
	bun.BaseModel `bun:"table:grades,alias:g"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull"`
}

type Class struct {
	bun.BaseModel `bun:"table:classes,alias:c"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull"`
	GradeID       int64  `bun:"grade_id,notnull"`
	Grade         *Grade `bun:"rel:belongs-to"`
}

type Student struct {
	bun.BaseModel `bun:"table:students,alias:s"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull"`
	ClassID       int64  `bun:"class_id,notnull"`
	Class         *Class `bun:"rel:belongs-to"`
}
