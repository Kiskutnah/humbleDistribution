package main

import (
	"net/http"

	"main/src/connections"
	"main/src/db"

	"github.com/gin-gonic/gin"
)

// // TableName overrides the table name used by User to `profiles`
// func (testTable) TableName() string {
// 	return "testTable"
// }
// func JSONMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         c.Writer.Header().Set("Content-Type", "application/json")
//         c.Next()
//     }
// }

func main() {
	app := gin.Default()
	// app.Use(JSONMiddleware())
	dbConn := connections.DBInit()
	dbConn.AutoMigrate(&db.HumbleGame{}, &db.AccessKey{}, &db.GameKey{})

	app.GET("/", func(c *gin.Context) {
		games := db.GetGames(*dbConn)
		c.JSON(http.StatusOK, gin.H{
			"message": games,
		})
	})
	app.Run("127.0.0.1:8080")
}
