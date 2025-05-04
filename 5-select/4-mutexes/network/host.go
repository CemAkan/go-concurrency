package network

import (
	"bombgame/ui"
	"log"
	"net"
)

func StartHostTCP() {
	listener, err := net.Listen("tcp", ":0")

	if err != nil {
		log.Fatal("Tcp port listening error")
	}
	addr := listener.Addr().(*net.TCPAddr) // getting port and assert the type to interface

	ui.HostInfoShowMenu(getLocalIP(), addr.Port)

}

func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Println("ERROR: UDP connection can not establish with google for getting ip")
		return "localhost"
	}
	return conn.LocalAddr().(*net.TCPAddr).IP.String()
}
