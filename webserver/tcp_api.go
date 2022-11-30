package webserver

import (
	"context"
	"log"
	"math"
	"net"
	"time"
)

func (s *WebServer) listenTCP() {
	tcpListener, err := net.Listen("tcp", s.addressSelf+TCP_ID_PORT)
	if err != nil {
		log.Printf("error making tcp listener: %s", err.Error())
		return
	}
	defer tcpListener.Close()

	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			log.Printf("error listening to tcp: %s", err.Error())
			return
		}
		go s.handleIdRequest(conn)
	}
}

func (s *WebServer) handleIdRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Printf("error reading tcp buffer: %s", err.Error())
		return
	}

	conn.Write([]byte{byte(s.id)})
	conn.Close()
}

func (s *WebServer) idRequest(address string) int {
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", address+TCP_ID_PORT)
	if err != nil {
		log.Printf("error dialing tcp with context: %s", err.Error())
		return math.MaxInt
	}
	defer conn.Close()

	conn.Write([]byte("ID Request"))

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		log.Printf("error reading response: %s", err.Error())
		return math.MaxInt
	}

	return int(buffer[0])
}
