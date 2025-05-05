package main

import (
	"bombgame/conf"
	"bombgame/game"
	"bombgame/ui"
)

func main() {
	ui.UserInputStart()

	game.StartGame()

	conf.SleepTime = 4 //global sleep time assertion

}
