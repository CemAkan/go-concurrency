package ui

import (
	"bombgame/conf"
	"fmt"
)

func ShowGAmeResult(loserStaus string) {

	var resultText string

	if loserStaus == conf.PlayerStatus {
		resultText = "YOU LOSE HA HA HA"
	} else {

		resultText = "WINNER WINNER CHICKEN DINNER"
	}

	fmt.Println("======================================")
	fmt.Println("======================================\n")

	fmt.Print(resultText)

	fmt.Println("======================================\n")
	fmt.Println("======================================")

}
