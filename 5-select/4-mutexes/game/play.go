package game

import (
	"bombgame/conf"
	"bombgame/model"
	"bombgame/ui"
	"encoding/gob"
	"github.com/eiannone/keyboard"
	"log"
	"time"
)

var (
	enc *gob.Encoder
	dec *gob.Decoder
)

func StartGame() {

	//keyboard opening & closing
	err := keyboard.Open()

	if err != nil {
		log.Fatalln("Keyboard open error", err)
	}
	defer keyboard.Close()

	//creating encoder & decoder via tcp socket interface (conn)
	enc = gob.NewEncoder(conf.GameConn)
	dec = gob.NewDecoder(conf.GameConn)

	//initial conn check
	if conf.GameConn == nil {
		log.Fatal("StartGame içinde bağlantı yok (conf.GameConn nil)")
	}

	log.Println("Game stated for ", conf.PlayerName)

	if conf.PlayerStatus == "host" { //only host one can create a bomb structure because of the solving conflicts (rand holder & time)
		bomb := model.NewBomb() //initial create
		log.Println("Hey, sweety we have a newborn bomb. ", bomb)

		err := enc.Encode(bomb)
		if err != nil { //sending and error check
			log.Fatalln("Bomb encoding fatal error")
		}

	}

	isFirstTurn := true

	for {
		var bomb model.Bomb

		if !(isFirstTurn && conf.PlayerStatus == "host") {
			err := dec.Decode(&bomb)
			if err != nil { // receiving bomb and error check
				ui.ShowWarningMessage("CONNECTION LOST :(")
				log.Fatalln("I want to a bomb, but you gave me a decoding error whyyy :( ", err)
			}
		}
		isFirstTurn = false

		ui.ShowTurnInfo(bomb.WhoHold()) // ui module turn info shower on terminal

		if bomb.WhoHold() == conf.PlayerStatus {
			holdSpaceAndDecreaseTime(&bomb)
		}

		if bomb.IsExploded() {
			ui.ShowGameResult(bomb.WhoHold())
			return
		}
	}
}

func holdSpaceAndDecreaseTime(bomb *model.Bomb) {
	start := time.Now()
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			_, key, err := keyboard.GetKey()
			if err != nil {
				log.Println("Keyboard read error: ", err)
				ui.ShowWarningMessage("Keyboard input error.")
				return
			}

			if key == keyboard.KeySpace {
				bomb.DecreaseTime(0.1)
				if bomb.IsExploded() {
					log.Println("Bomb exploded in ", conf.PlayerName, "'s hand")

					err := enc.Encode(bomb)

					if err != nil { //sending and error check
						log.Fatalln("Bomb encoding fatal error")
					}

					ui.ShowGameResult(conf.PlayerStatus)
					return
				}
			} else {
				held := time.Since(start).Seconds()

				ui.HoldingTimeShower(held)

				bomb.SwitchHolder()

				log.Println("Turn switched")

				err := enc.Encode(bomb)

				if err != nil { //sending and error check
					log.Fatalln("Bomb encoding fatal error")
				}

				return
			}
		}
	}
}
