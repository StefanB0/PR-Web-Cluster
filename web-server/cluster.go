package webserver

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	REFRESH = time.Second * 30
)

type Cluster struct {
	name      string
	serverSet []ServerRefference
}

type ServerRefference struct {
	address string
	version uint64
}

func NewCluster(_name string, addressSet []string) Cluster {
	_serverSet := make([]ServerRefference, len(addressSet))
	for i, _address := range addressSet {
		_serverSet[i] = ServerRefference{address: _address, version: 0}
	}
	cluster := Cluster{name: _name, serverSet: _serverSet}
	return cluster
}

func (c *Cluster) AddToCluster(reff ServerRefference) {
	c.serverSet = append(c.serverSet, reff)
}

func (s *WebServer) periodicSync() {
	defer s.wg.Done()
	for s.serverAlive {
		time.Sleep(REFRESH)
		for _, reff := range s.cluster.serverSet {
			s.syncRequest(reff.address)
		}
	}
}

func (s *WebServer) syncRequest(address string) {
	body, err := json.Marshal(s.memory)
	if err != nil {
		log.Printf("server: error making json copy of database: %s\n", err)
	}
	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPut, address, bodyReader)
	if err != nil {
		log.Printf("server: error making http request: %s\n", err)
	}
	http.DefaultClient.Do(req)
}
