package game

import (
	"bombgame/conf"
	_ "bombgame/network" //only init calling
	"encoding/gob"
	"log"
)

// bosss, I swear. I am not play a game, it is a part of my job vallahi diyom

func StartGame() {
	log.Println("Game stated for ", conf.PlayerName)

	//creating encoder & decoder via tcp socket interface (conn)
	ecn := gob.NewEncoder(conf.GameConn)
	dec := gob.NewDecoder(conf.GameConn)
}
