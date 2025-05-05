package game

import (
	"bombgame/conf"
	"bombgame/model"
	_ "bombgame/network" //only init calling
	"bombgame/ui"
	"encoding/gob"
	"github.com/eiannone/keyboard"
	"log"
	"time"
)

type bombWrapper struct {
	*model.Bomb
}

var (
	enc *gob.Encoder
	dec *gob.Decoder
)

func init() {

	//global keyboard opening & closing
	err := keyboard.Open()

	if err != nil {
		log.Fatalln("Keyboard open error", err)
	}

	defer keyboard.Close()

	//creating encoder & decoder via tcp socket interface (conn)
	enc = gob.NewEncoder(conf.GameConn)
	dec = gob.NewDecoder(conf.GameConn)
}

// bosss, I swear. I am not play a game, it is a part of my job vallahi diyom

func StartGame() {
	log.Println("Game stated for ", conf.PlayerName)

	if conf.PlayerStatus == "host" { //only host one can create a bomb structure because of the solving conflicts (rand holder & time)
		bomb := &bombWrapper{model.NewBomb()} //initial create
		log.Println("Hey, sweety we have a newborn bomb. ", bomb)

		if enc.Encode(bomb) != nil { //sending and error check
			log.Fatalln("Bomb encoding fatal error")
		}
	}
	for {
		var bomb bombWrapper // all for loop it will be reset

		if dec.Decode(&bomb) != nil { // receiving bomb and error check
			ui.ShowWarningMessage("CONNECTION LOST :(")
			log.Fatalln("I want to a bomb, but you gave me a decoding error whyyy :(")
		}

		ui.ShowTurnInfo(bomb.WhoHold()) // ui module turn info shower on terminal

		bomb.holdSpaceAndDecreaseTime()

		if bomb.IsExploded() {
			ui.ShowGameResult(bomb.WhoHold())
			return
		}
	}
}

func (bomb *bombWrapper) holdSpaceAndDecreaseTime() {
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

					if enc.Encode(bomb) != nil { //sending and error check
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

				if enc.Encode(bomb) != nil { //sending and error check
					log.Fatalln("Bomb encoding fatal error")
				}

				return
			}
		}
	}
}
