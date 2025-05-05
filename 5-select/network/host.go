package network

import (
	"bombgame/conf"
	"bombgame/ui"
	"fmt"
	"log"
	"net"
)

func startHostTCP() {
	listener, err := net.Listen("tcp", ":0")

	if err != nil {
		log.Fatal("Tcp port listening error")
	}
	addr := listener.Addr().(*net.TCPAddr) // getting port and assert the type to interface

	hostInfo := fmt.Sprintf("%s:%d", getLocalIP(), addr.Port)

	ui.HostInfoShowMenu(hostInfo)

	conf.GameAddress = hostInfo

	conn, err := listener.Accept()

	if err != nil || conn == nil {
		log.Fatal("Captain, we have a big problem. Our tcp socket can not accept client's request :'( ")
	}

	conf.GameConn = conn
}

func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Println("ERROR: UDP connection can not establish with google for getting ip")
		return "localhost"
	}
	return conn.LocalAddr().(*net.UDPAddr).IP.String()
}
