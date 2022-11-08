package webserver

import "time"

const (
	RUNSPEED = time.Second * 30
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
		time.Sleep(RUNSPEED)
		for _, reff := range s.cluster.serverSet {
			s.syncRequest(reff.address)
		}
	}
}

func (s *WebServer) syncRequest(address string) //TODO
