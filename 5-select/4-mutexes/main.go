package main

import (
	"bombgame/conf"
	"bombgame/game"
	"bombgame/network"
	"bombgame/ui"
)

func main() {
	conf.LogFileInit()

	ui.UserInputStart()

	network.NetwokStart()

	game.StartGame()

	conf.SleepTime = 4 //global sleep time assertion

}
