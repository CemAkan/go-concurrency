package network

import "bombgame/conf"

func NetwokStart() {
	if conf.PlayerStatus == "host" {
		startHostTCP()
	} else {
		joinHostTCP()
	}
}
