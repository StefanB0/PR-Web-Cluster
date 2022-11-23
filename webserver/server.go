package webserver

import (
	"PR/Web_Cluster/database"
	"net"
	"net/http"
	"time"
)

const (
	REFRESH = time.Second * 4

	HTTP_PREFIX   = "http://"
	HTTP_PORT     = ":3000"
	UDP_PING_PORT = ":3001"

	UDP_OK   = "OK"
	UDP_DEAD = "DEAD"
)

type WebServer struct {
	id            int
	port          string
	addressSelf   string
	leaderAddress string
	network       []string
	isLeader      bool
	serverAlive   bool
	ledger        map[string][]string

	udpServer net.PacketConn
	memory    database.DatabaseInstance
	http.Server
}

func NewWebServer(_id int, _address string, _port string, _network []string) *WebServer {
	return &WebServer{
		id:          _id,
		port:        _port,
		addressSelf: _address,
		network:     _network,
		serverAlive: true,
		ledger:      make(map[string][]string),
		memory:      database.NewDatabase(),
	}
}

func (s *WebServer) StartServer() {
	s.network = pruneSlice(s.network, s.addressSelf)
	go s.udpListen()
	s.initHandlers()
	go s.initListen()
	s.serverRun()
}

func (s *WebServer) serverRun() {

	for s.serverAlive {
		if s.isLeader {
			s.checkNetwork()
		}
		// log.Printf("%s: internal memory: %+v", s.addressSelf, s.memory)
		time.Sleep(REFRESH)
	}

	s.udpServer.Close()
}
