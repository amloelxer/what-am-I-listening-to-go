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
	"go.uber.org/zap"
	"whatisalexlisteningto.com/api/models"
	"whatisalexlisteningto.com/api/utils"
)

type person struct {
	name string
	age  int
}

func (inputPerson *person) ToString() string {
	myString := fmt.Sprintf("The name of my person is %s and the age is %v", inputPerson.name, inputPerson.age)
	return myString
}

func connectDB(logger *zap.SugaredLogger) *sql.DB {
	db, err := sql.Open("postgres", "postgresql://alex:password@127.0.0.1/what-am-i-listening-to?sslmode=disable")
	if err != nil {
		logger.Fatal("Die")
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

	myFakePerson := &person{name: "Alex", age: 29}
	customLogger.Infoln(myFakePerson.ToString())

	fakeSecretKey := os.Getenv("FAKE_SECRET_KEY")
	customLogger.Infoln("The fake secret key is", fakeSecretKey)

	db := connectDB(customLogger)
	boil.SetDB(db)

	artists, err := models.Artists().All(ctx, db)
	if err != nil {
		customLogger.Panicln("The heinous error is", err.Error())
	}

	for _, artist := range artists {
		fmt.Println("The artist name is", artist.Name.String)
		fmt.Println("The artist biography is", artist.Biography.String)
	}

	fmt.Println("The value of the artists is", artists)
	// fmt.Println(err.Error())
	r := gin.Default()
	// albumsArray := []models.Album{{Name: "Stadium Arcadium", Artist: "Red Hot Chili Peppers"}}
	r.GET("/artists", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": artists,
		})
	})
	const address = "localhost:3001"
	r.Run(address)
	customLogger.Infoln("Running on address", "localhost:3001")
}
