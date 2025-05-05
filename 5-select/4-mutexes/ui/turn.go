package ui

import (
	"bombgame/conf"
	"fmt"
)

func ShowTurnInfo(holderStatus string) {
	var turnText string

	if holderStatus == conf.PlayerStatus {
		turnText = ">> It is your turn. Hold space key <<"
	} else {
		turnText = ">> It is your friend's turn, wait please <<"
	}

	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=")
	infoTitleColor.Println(turnText)
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=")
}
