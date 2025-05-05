package ui

import (
	"bombgame/conf"
	"fmt"
	"os"
	"time"
)

func ShowGameResult(loserStatus string) {

	clearScreen()

	resultText := "WINNER WINNER CHICKEN DINNER"

	if loserStatus == conf.PlayerStatus {
		resultText = "YOU LOSE HA HA HA"
	}

	fmt.Println("============================")
	fmt.Println("============================\n")

	fmt.Print(resultText)

	fmt.Println("============================\n")
	fmt.Println("============================")

	time.Sleep(time.Second * time.Duration(conf.SleepTime))

	os.Exit(0)
}

func HoldingTimeShower(holdingTime float64) {
	clearScreen()
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=")
	infoTitleColor.Println("You hold it for ", holdingTime, " seconds.")
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=")
}
