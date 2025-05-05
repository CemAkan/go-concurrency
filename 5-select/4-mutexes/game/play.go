package game

import (
	"bombgame/conf"
	"bombgame/model"
	_ "bombgame/network" //only init calling
	"bombgame/ui"
	"encoding/gob"
	"log"
)

// bosss, I swear. I am not play a game, it is a part of my job vallahi diyom

func StartGame() {
	log.Println("Game stated for ", conf.PlayerName)

	//creating encoder & decoder via tcp socket interface (conn)
	enc := gob.NewEncoder(conf.GameConn)
	dec := gob.NewDecoder(conf.GameConn)

	var bomb *model.Bomb

	if conf.PlayerStatus == "host" { //only host one can create a bomb structure because of the solving conflicts (rand holder & time)
		bomb = model.NewBomb()
		log.Println("Hey, sweety we have a newborn bomb. ", bomb)

		if enc.Encode(bomb) != nil {
			log.Fatalln("Bomb encoding fatal error")
		}
	}

	ui.ShowTurnInfo(bomb.WhoHold())

}
