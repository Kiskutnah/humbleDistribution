package main

import (
	"fmt"
	"net/http"

	"main/src/connections"
	"main/src/db"
	"main/src/requests"

	"github.com/gin-gonic/gin"
)

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func main() {
	app := gin.Default()
	app.Use(JSONMiddleware())
	dbConn := connections.DBInit()
	dbConn.AutoMigrate(&db.HumbleGame{}, &db.AccessKey{}, &db.GameKey{})

	app.GET("/games", func(c *gin.Context) {
		games := db.GetGames(*dbConn)
		c.JSON(http.StatusOK, gin.H{
			"message": games,
		})
	})
	app.PUT("/games", func(c *gin.Context) {
		var newGame db.HumbleGame
		if err := c.BindJSON(&newGame); err != nil {
			return
		}
		fmt.Println(newGame)
		whatThis := db.NewGame(&newGame, *dbConn)
		fmt.Println(whatThis, newGame)
		c.IndentedJSON(http.StatusCreated, whatThis)
	})
	app.PUT("/keys", func(c *gin.Context) {
		var newKey db.AccessKey
		if err := c.BindJSON(&newKey); err != nil {
			return
		}
		fmt.Println(newKey)
		db.NewAccessKey(&newKey, *dbConn)
		c.IndentedJSON(http.StatusCreated, newKey)

	})
	app.POST("/redeem", func(c *gin.Context) {
		var redemptionRequest requests.ReedemKey
		if err := c.BindJSON(&redemptionRequest); err != nil {
			return
		}
		fmt.Println(redemptionRequest)
		isValidKey := db.CheckAccessKey(redemptionRequest.AccessKey, *dbConn)
		// if it's not good we're aborting here
		if !isValidKey {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Invalid Access key",
			})
			return
		}
		steamKey := db.GetGamekey(redemptionRequest.GameID, *dbConn)
		c.JSON(http.StatusAccepted, gin.H{
			"key": steamKey,
		})
		db.DeleteGameKey(uint(redemptionRequest.GameID), *dbConn)
		db.DeleteAccessKey(redemptionRequest.AccessKey, *dbConn)
	})
	app.Run("127.0.0.1:8080")
}
