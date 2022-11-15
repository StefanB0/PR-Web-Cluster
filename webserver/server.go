package webserver

import (
	"PR/Web_Cluster/database"
	"errors"
	"log"
	"net/http"
	"os"
	"sync"
)

const (
	VIRTUAL_PORT     = ":3000"
	ADDRESS_TEMPLATE = "http:minion"
)

type WebServer struct {
	id      int
	address string
	port    string

	isLeader      bool
	leaderAddress string

	cluster Cluster
	memory  database.DatabaseInstance
	wg      sync.WaitGroup

	http.Server
	serverAlive bool
}

func NewWebServer(_id int, _address string, _port string) *WebServer {
	return &WebServer{
		id:          _id,
		address:     _address,
		port:        _port,
		serverAlive: true,
		memory:      database.NewDatabase(),
	}
}

func (s *WebServer) StartServer() {
	s.initHandlers()

	s.wg.Add(1)
	go s.initListen()

	if s.isLeader {
		s.wg.Add(1)
		go s.periodicSync()
	} else {
		s.clusterSync(s.leaderAddress)
	}

	s.wg.Wait()
}

func (s *WebServer) initHandlers() {
	http.HandleFunc("/", getHello)
	http.HandleFunc("/create", s.createElement)
	http.HandleFunc("/read", s.readElement)
	http.HandleFunc("/update", s.updateElement)
	http.HandleFunc("/delete", s.deleteElement)
	http.HandleFunc("/sync", s.overwriteMemory)
	http.HandleFunc("/contact", s.resolveContact)
}

func (s *WebServer) initListen() {
	defer s.wg.Done()
	err := http.ListenAndServe(s.port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server closed \n")
	} else if err != nil {
		log.Printf("error starting server %s\n", err)
		os.Exit(1)
	}
}

func (s *WebServer) SetLeader(address string, _isLeader bool) {
	s.leaderAddress = address
	s.isLeader = _isLeader
}

func (s *WebServer) SetCluster(_cluster Cluster) {
	s.cluster = _cluster
}
