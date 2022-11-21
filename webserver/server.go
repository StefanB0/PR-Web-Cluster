package webserver

import (
	"PR/Web_Cluster/database"
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	REFRESH = time.Second * 4
)

type WebServer struct {
	id            int
	port          string
	addressSelf   string
	leaderAddress string
	network       []string
	isLeader      bool
	serverAlive   bool

	memory database.DatabaseInstance
	http.Server
}

func NewWebServer(_id int, _address string, _port string, _network []string) *WebServer {
	return &WebServer{
		id:          _id,
		port:        _port,
		addressSelf: _address,
		network:     _network,
		serverAlive: true,
		memory:      database.NewDatabase(),
	}
}

func (s *WebServer) StartServer() {
	s.network = pruneSlice(s.network, s.addressSelf)
	go s.serverRun()
	s.initHandlers()
	s.initListen()
}

func (s *WebServer) initListen() {
	err := http.ListenAndServe(s.port, nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server closed \n")
	} else if err != nil {
		log.Printf("error starting server %s\n", err)
		os.Exit(1)
	}
	s.serverAlive = false
}

func (s *WebServer) serverRun() {
	// if s.isLeader {
	// 	s.memory.Create("Hello", []byte("World"))
	// 	time.Sleep(1)
	// }

	for s.serverAlive {
		// if s.isLeader {
		// 	s.periodicSync()
		// }
		log.Printf("%s: internal memory: %+v", s.addressSelf, s.memory)
		time.Sleep(REFRESH)
	}
}

func (s *WebServer) periodicSync() {
	for _, address := range s.network {
		keys, values := s.memory.GetKeyValuePairs()
		s.overwriteRequest(address, keys, values)
	}
}

func (s *WebServer) SetLeader(address string, _isLeader bool) {
	s.leaderAddress = address
	s.isLeader = _isLeader
}
