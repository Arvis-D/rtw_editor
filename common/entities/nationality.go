package entities

import "gorm.io/gorm"

type Nationality struct {
	gorm.Model
	Name string
}
