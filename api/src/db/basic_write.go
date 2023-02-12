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
	testForUnique := AccessKey{}
	for {
		conn.Where(&AccessKey{Key: newKey}).First(&testForUnique)
		if testForUnique.ID == 0 {
			break
		}
	}
	key.Key = newKey

	res := conn.Create(&key)

	fmt.Println(res.Error)
	fmt.Println(res.RowsAffected)
	fmt.Println(key)

	return key
}

func generateAccessKey(len int) (string, error) {
	bytes := make([]byte, len)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
