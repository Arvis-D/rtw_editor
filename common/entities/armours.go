package entities

import "gorm.io/gorm"

type Armour struct {
	gorm.Model
	Name        string
	ArmourBonus uint8
	Cost        uint8
	Weight      uint8
}
