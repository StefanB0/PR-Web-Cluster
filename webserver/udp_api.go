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
		reaponseStr := UDP_OK
		udpServer.WriteTo([]byte(reaponseStr), addr)
		log.Println("UDP internal communication. Ping")
	}
}

func udping(addr string) (status string){
	status = UDP_DEAD
	udpServer, err := net.ResolveUDPAddr("udp", addr + UDP_PING_PORT)
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

	conn.Write([]byte("Ping"))
	
	
	//data buffer
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		log.Printf("Read data failed: %s", err.Error())
		return
	}

	status = string(received)
	return status
}
