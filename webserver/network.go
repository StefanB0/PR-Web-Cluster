package webserver

import (
	"log"
	"time"
)

type ServerRefference struct {
	id      int
	address string
}

func (s *WebServer) checkNetwork() {
	for _, addr := range s.network {
		status := udping(addr)
		if status == UDP_OK {
			continue
		} else if status == UDP_DEAD {
			s.pruneDeadServer(addr)
		}
	}
}

func (s *WebServer) checkLeader() {
	for s.serverAlive {
		time.Sleep(REFRESH)
		if s.isLeader {
			continue
		}

		status := udping(s.leaderAddress)
		if status == UDP_OK {
			continue
		} else if status == UDP_DEAD {
			log.Println("Partition leader died")
			s.pruneDeadServer(s.leaderAddress)
			s.chooseLeader()
			log.Printf("New partition leader: %s\n", s.leaderAddress)
		}
	}
}

func (s *WebServer) chooseLeader() {
	leaderAddress := s.addressSelf
	leaderID := s.id
	for _, candidateAddress := range s.network {
		candidateID := s.idRequest(candidateAddress)
		if candidateID < leaderID {
			leaderAddress = candidateAddress
			leaderID = candidateID
		}
	}
	_isLeader := (s.id == leaderID)
	s.SetLeader(leaderAddress, _isLeader)

	if s.isLeader {
		for key, addressSet := range s.ledger {
			if !checkSlice(addressSet, s.addressSelf) {
				s.memory.Create(key, s.getValue(addressSet[0], key))
			} else {
				s.ledger[key] = pruneSlice(s.ledger[key], s.addressSelf)
			}
		}
	}
}

func (s *WebServer) pruneDeadServer(addr string) {
	log.Println("Server died:", addr)
	s.network = pruneSlice(s.network, addr)
	for k, v := range s.ledger {
		s.ledger[k] = pruneSlice(v, addr)
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
