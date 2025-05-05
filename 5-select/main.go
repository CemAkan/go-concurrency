package main

import (
	"bombgame/conf"
	"bombgame/game"
	"bombgame/model"
	"bombgame/network"
	"bombgame/ui"
	"encoding/gob"
)

func init() {
	gob.Register(&model.Bomb{}) // Mutlaka pointer ile
}

func main() {
	conf.LogFileInit()

	ui.UserInputStart()

	network.NetwokStart()

	game.StartGame()

	conf.SleepTime = 4 //global sleep time assertion

}
