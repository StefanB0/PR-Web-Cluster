package webserver

import (
	"log"
	"net"
)

func (s *WebServer) udpListen() {
	udpServer, err := net.ListenPacket("udp", UDP_PORT)
	if err != nil {
		log.Printf("udp start server: error listening to packets%s\n", err)
	}
	defer udpServer.Close()

	for s.serverAlive {
		buf := make([]byte, 1024)
		_, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			continue
		}
		reaponseStr := "alive"
		udpServer.WriteTo([]byte(reaponseStr), addr)
	}
}

func (s *WebServer) udping(addr string) {
	
}
