package ui

import (
	"bombgame/conf"
	"fmt"
)

func ShowGameResult(loserStaus string) {

	resultText := "WINNER WINNER CHICKEN DINNER"

	if loserStaus == conf.PlayerStatus {
		resultText = "YOU LOSE HA HA HA"
	}

	fmt.Println("============================")
	fmt.Println("============================\n")

	fmt.Print(resultText)

	fmt.Println("============================\n")
	fmt.Println("============================")

}
