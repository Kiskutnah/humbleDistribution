package db

import (
	"fmt"

	"gorm.io/gorm"
)

// // TableName overrides the table name used by User to `profiles`
// func (testTable) TableName() string {
// 	return "testTable"
// }

func GetGames(conn gorm.DB) []HumbleGame {
	games := []HumbleGame{}
	conn.Find(&games)

	fmt.Println("games:", games)

	return games
}
