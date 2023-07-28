package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"whatisalexlisteningto.com/api/models"
	"whatisalexlisteningto.com/api/utils"
)

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://alex:password@127.0.0.1/what-am-i-listening-to?tLSMode=0")
	if err != nil {
		utils.GetLogger().Fatal("Die")
	}
	return db
}

func main() {
	ctx := context.Background()
	customLogger := utils.GetLogger()
	customLogger.Infoln("Hello World from sugar logger!")
	err := godotenv.Load()
	if err != nil {
		customLogger.Fatal("Error loading .env file")
	}
	tempArray := []string{"one", "two", "three"}
	fmt.Println(utils.ArrayToString(tempArray))

	fakeSecretKey := os.Getenv("FAKE_SECRET_KEY")
	customLogger.Infoln("The fake secret key is", fakeSecretKey)

	db := connectDB()
	boil.SetDB(db)

	artists, err := models.Artists().All(ctx, db)
	fmt.Println(artists)
	fmt.Println(err.Error())
	r := gin.Default()
	// albumsArray := []models.Album{{Name: "Stadium Arcadium", Artist: "Red Hot Chili Peppers"}}
	r.GET("/albums", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"okay": true,
			},
		})
	})
	const address = "localhost:3001"
	r.Run(address)
	customLogger.Infoln("Running on address", "localhost:3001")
}
