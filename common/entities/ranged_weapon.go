package entities

import "gorm.io/gorm"

type RangedWeapon struct {
	gorm.Model
	Name     string
	Damage   uint8
	Cost     uint8
	Weight   uint8
	Range    uint8
	AP       bool
	AmmoCost float32
}
