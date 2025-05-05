package game

import (
	"bombgame/conf"
	_ "bombgame/network" //only init calling
	"log"
)

// bosss, I swear. I am not play a game, it is a part of my job vallahi diyom

func StartGame() {
	log.Println("Game stated for ", conf.PlayerName)
}
