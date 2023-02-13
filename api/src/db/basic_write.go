package db

import (
	"crypto/rand"
	"encoding/hex"
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

func NewAccessKey(key *AccessKey, conn gorm.DB) *AccessKey {
	var newKey, err = generateAccessKey(64)
	if err != nil {
		panic(err)
	}
	for {
		if CheckAccessKey(newKey, conn) {
			break
		} else {
			newKey, err = generateAccessKey(64)
			if err != nil {
				panic(err)
			}
		}
	}
	key.Key = newKey

	res := conn.Create(&key)

	fmt.Println(res.Error)
	fmt.Println(res.RowsAffected)
	fmt.Println(key)

	return key
}

func GenerateAccessKeyWrapper() (string, error) {
	return generateAccessKey(64)
}
func generateAccessKey(len int) (string, error) {
	bytes := make([]byte, len)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func DeleteAccessKey(key string, conn gorm.DB) AccessKey {
	deletedAccessKey := AccessKey{}
	conn.Where(&AccessKey{Key: key}).Delete(&deletedAccessKey)

	return deletedAccessKey
}

func DeleteGameKey(ID uint, conn gorm.DB) HumbleGame {
	deletedHubmleGame := HumbleGame{}
	conn.Where(&HumbleGame{ID: ID}).Delete(&deletedHubmleGame)

	return deletedHubmleGame
}
