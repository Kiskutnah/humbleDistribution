package db

import (
	"fmt"

	"gorm.io/gorm"
)

func GetGames(conn gorm.DB) []HumbleGame {
	games := []HumbleGame{}
	conn.Find(&games)

	fmt.Println("games:", games)

	return games
}

func CheckAccessKey(generatedKey string, conn gorm.DB) bool {
	testForUnique := AccessKey{}
	conn.Where(&AccessKey{Key: generatedKey}).First(&testForUnique)
	// if there's no ID (as in ID == 0), we know it's unique
	// tbh, what are teh chances they aren't unique, but may aswell check
	if testForUnique.ID != 0 {
		return true
	}
	return false
}

func GetGamekey(gameID int, conn gorm.DB) HumbleGame {
	gameData := HumbleGame{}
	conn.Where(&HumbleGame{ID: uint(gameID)}).First(&gameData)
	return gameData
}
