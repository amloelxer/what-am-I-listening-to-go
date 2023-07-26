package main

import (
	"fmt"
	"os"

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
}
