package entities

import "gorm.io/gorm"

type Base struct {
	gorm.Model
	Name  string
	Value uint8
}
