package models

import "gorm.io/gorm"

type PokemonType struct {
	gorm.Model
	Slot uint
	Type string
}

type Pokemon struct {
	gorm.Model
	Name string
	//PkdxId uint
	//Order  uint
	//Types  []PokemonType `gorm:"many2many:pokemon_type;"`
}
