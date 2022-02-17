package entities

import "gorm.io/gorm"

type MeleeWeapon struct {
	gorm.Model
	Name        string
	Damage      uint8
	Cost        uint8
	Weight      uint8
	AP          bool
	ChargeBonus uint8
}
