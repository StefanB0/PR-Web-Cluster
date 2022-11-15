package webserver

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	REFRESH = time.Second * 5
)

type Cluster struct {
	name      string
	serverSet []ServerRefference
}

type ServerRefference struct {
	id      int
	address string
}

func NewCluster(_name string, addressSet []string) Cluster {
	_serverSet := make([]ServerRefference, len(addressSet))
	for i, _address := range addressSet {
		_serverSet[i] = ServerRefference{address: _address}
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
		log.Println("Flag1")
		for _, reff := range s.cluster.serverSet {
			s.syncRequest("http://127.0.0.1" + reff.address + "/sync")
		}
		log.Println("It just works")
	}
}

func (s *WebServer) clusterSync(address string) {
	body, err := json.Marshal(ServerRefference{id: s.id, address: s.address})
	if err != nil {
		log.Printf("server: error making json copy of custom struct id + address: %s\n", err)
	}

	bodyReader := bytes.NewReader(body)

	resp, err := http.DefaultClient.Post(address+"/contact", "struct", bodyReader)
	respBody, err := ioutil.ReadAll(resp.Body)
	
	var nCluster Cluster
	json.Unmarshal(respBody, &nCluster)
	s.cluster = nCluster
	
	err = resp.Body.Close()
	if err != nil {
		log.Printf("server: error closing response body: %s\n", err)
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

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("server: error sending new request: %s\n", err)
	}
}
