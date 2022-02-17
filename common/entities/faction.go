package entities

import "gorm.io/gorm"

type Faction struct {
	gorm.Model
	Name string
}
