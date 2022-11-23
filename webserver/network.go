package webserver

import "log"

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
