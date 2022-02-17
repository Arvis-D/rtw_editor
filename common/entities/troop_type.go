package entities

import "gorm.io/gorm"

type TroopType struct {
	gorm.Model
	Name         string
	MoraleBonus  int8
	AttackBonus  int8
	DefenceBonus int8
	MissileBonus int8
}
