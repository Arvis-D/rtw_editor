package entities

import "gorm.io/gorm"

type Mount struct {
	gorm.Model
	Speed       uint8
	ArmourBonus uint8
	ChargeBonus uint8
	BaseUpkeep  uint8
	Cost        uint8
}
