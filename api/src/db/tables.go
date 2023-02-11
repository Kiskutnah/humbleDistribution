package db

import (
	"gorm.io/gorm"
)

type HumbleGame struct {
	gorm.Model
	ID   uint `gorm:"primary_key"`
	Name string
	Key  string
}

type AccessKey struct {
	gorm.Model
	ID   uint `gorm:"primary_key"`
	Key  int  // will do a hexadecimal number
	Name string
}

type GameKey struct {
	ID     uint `gorm:"primary_key"`
	Key    int
	GameID int
	Game   HumbleGame
}
