package webserver

import (
	"log"
	"net"
)

func (s *WebServer) udpListen() {
	udpServer, err := net.ListenPacket("udp", UDP_PING_PORT)
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
	udpServer, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Printf("ResolveUDPAddress failed: %s", err.Error())
		return
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		log.Printf("Listen failed: %s", err.Error())
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Ping"))
	
	//data buffer
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		log
	}

}
