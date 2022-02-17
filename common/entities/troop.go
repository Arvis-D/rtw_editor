package entities

import "gorm.io/gorm"

type Troop struct {
	gorm.Model
	Name            string
	Mount           Mount
	Shield          Shield
	MeleeWeapon     MeleeWeapon
	RangedWeapon    RangedWeapon
	Nationalities   []Nationality `gorm:"many2many:troop_nationalities;"`
	Factions        []Faction     `gorm:"many2many:troop_factions;"`
	Armour          Armour
	TroopType       TroopType
	AmmoCount       uint8
	Skill           uint8
	TypicalUnitSize uint8
}
