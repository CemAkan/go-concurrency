package network

import (
	"log"
	"net"
)

func StartHostTCP() net.Conn {
	listener, err := net.Listen("tcp", ":0")

	if err != nil {
		log.Fatal("Tcp port listening error")
	}

}
