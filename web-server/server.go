package webserver

import (
	"PR/Web_Cluster/database"
	"errors"
	"log"
	"net/http"
	"os"
	"sync"
)

type WebServer struct {
	http.Server
	serverAlive bool

	port string

	isLeader      bool
	leaderAddress string

	cluster Cluster
	memory  database.DatabaseInstance
	wg      sync.WaitGroup
}

func NewWebServer(_port string, _leaderAddress string, _isLeader bool, _cluster Cluster) *WebServer {
	return &WebServer{
		serverAlive:   true,
		port:          _port,
		leaderAddress: _leaderAddress,
		isLeader:      _isLeader,
		cluster:       _cluster,
		memory:        database.NewDatabase(),
	}
}

func (s *WebServer) StartServer() {
	s.initHandlers()

	s.wg.Add(1)
	go s.initListen()

	if s.isLeader {
		s.wg.Add(1)
		go s.periodicSync()
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
