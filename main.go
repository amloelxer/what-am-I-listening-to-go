package main

import (
	"fmt"

	"whatisalexlisteningto.com/api/utils"
)

func main() {
	tempArray := []string{"one", "two", "three"}
	fmt.Println(utils.ArrayToString(tempArray))
	customLogger := utils.GetLogger()
	customLogger.Infoln("Hello World from sugar logger!")
}
