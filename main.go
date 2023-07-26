package main

import (
	"fmt"

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
}
