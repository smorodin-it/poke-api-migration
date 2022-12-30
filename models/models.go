package models

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model
	Name   string
	PkdxId uint
	Types  pgtype.JSONB
}
