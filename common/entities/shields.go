package entities

import "gorm.io/gorm"

type Shield struct {
	gorm.Model
	Name         string
	DefenceBonus uint8
	Cost         uint8
	Weight       uint8
	Coverage     uint8
}
