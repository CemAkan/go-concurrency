package network

import "bombgame/conf"

func init() {
	if conf.PlayerStatus == "host" {
		startHostTCP()
	} else {
		joinHostTCP()
	}
}
