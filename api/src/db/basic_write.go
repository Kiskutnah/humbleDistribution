package db

import (
	"fmt"

	"gorm.io/gorm"
)

func NewGame(game *HumbleGame, conn gorm.DB) *HumbleGame {
	res := conn.Create(&game)

	fmt.Println(res.Error)
	fmt.Println(res.RowsAffected)
	fmt.Println(game)

	return game
}
