package log

import (
	"fmt"
	"os"
)

var logFile *os.File //global file var

func init() {
	var err error

	logFile, err = os.OpenFile("bombgame.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("FATAL ERROR: LOG FILE CAN NOT OPEN/CREATE")
	}
}
