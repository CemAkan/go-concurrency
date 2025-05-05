package network

import (
	"bombgame/conf"
	"log"
	"net"
)

func joinHostTCP() {

	log.Println("Client try to establish with host at ", conf.GameAddress)

	conn, err := net.Dial("tcp", conf.GameAddress)
	if err != nil {
		log.Fatalln("I can not establish a tcp connection with host, help me I am a poor client :( ")
	}

	log.Println("Yeaahhh, client can successfully establish a tcp connection with host at ", conf.GameAddress)

	conf.GameConn = conn
}
