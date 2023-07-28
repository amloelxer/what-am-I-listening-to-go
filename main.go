package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"whatisalexlisteningto.com/api/utils"
)

func main() {
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
