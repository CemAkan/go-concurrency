package conf

import (
	"fmt"
	"log"
	"os"
)

func init() {

	//open if it is exist or create new one
	logFile, err := os.OpenFile("bombgame.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("FATAL ERROR: LOG FILE CAN NOT OPEN/CREATE")
	}

	log.SetOutput(logFile)

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("Log file initialized")
}
