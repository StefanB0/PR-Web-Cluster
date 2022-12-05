package webserver

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *WebServer) sendSyncRequest() {
	log.Println(s.addressSelf, "sync request")
	addr := HTTP_PREFIX + s.leaderAddress + HTTP_PORT + "/syncRequest"
	req, err := http.NewRequest(http.MethodGet, addr, nil)
	if err != nil {
		log.Printf("sendSyncRequest: could not create request%s\n", err)
		return
	}

	req.Header.Add("server_name", s.addressSelf)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("syncRequest: error making http request: %s\n", err)
		return
	}

	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("syncRequest: could not read response body %s\n", err)
		return
	}
	var snapshot Snapshot
	json.Unmarshal(resbody, &snapshot)

	s.memory = snapshot.Memory
	s.ledger = snapshot.Ledger
	// log.Printf("snapshot %+v", snapshot)
}

func (s *WebServer) getValue(address string, key string) []byte {
	addr := HTTP_PREFIX + address + HTTP_PORT + "/read"
	req, err := http.NewRequest(http.MethodGet, addr, nil)
	if err != nil {
		log.Printf("getValue: could not create request%s\n", err)
		return []byte{0x0}
	}

	req.Header.Add("Key", key)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("handshake: error making http request: %s\n", err)
		return []byte{0x0}
	}

	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("handshake: could not read response body %s\n", err)
		return[]byte{0x0}
	}
	return resbody
}

func (s *WebServer) overwriteRequest(address string, keys []string, values [][]byte) {
	addr := HTTP_PREFIX + address + HTTP_PORT
	jsonData, err := json.Marshal(s.memory)
	if err != nil {
		log.Printf("overwriteRequest: could not marshal memory data%s\n", err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, addr+"/overwrite", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("overwriteRequest: could not create request%s\n", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("overwriteRequest: error making http request: %s\n", err)
		return
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("overwriteRequest: could not read response body %s\n", err)
		return
	}
}

func (s *WebServer) forwardRequest(r *http.Request, reqBody []byte, target []string, endpoint string) {
	for _, address := range target {
		addr := HTTP_PREFIX + address + HTTP_PORT
		forwReq, err := http.NewRequest(r.Method, addr+endpoint, bytes.NewBuffer(reqBody))
		if err != nil {
			log.Printf("forwardCreateRequest: could not create request%s\n", err)
			return
		}
		forwReq.Header.Add("Forward", "true")

		_, err = http.DefaultClient.Do(forwReq)
		if err != nil {
			log.Printf("forwardCreateRequest: error making http request: %s\n", err)
			return
		}
	}
}

func handshakeRequest(address, message string) {
	addr := HTTP_PREFIX + address + HTTP_PORT
	req, err := http.NewRequest(http.MethodGet, addr+"/handshake", nil)
	if err != nil {
		log.Printf("handshake: could not create request%s\n", err)
		return
	}

	req.Header.Add("message", message)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("handshake: error making http request: %s\n", err)
		return
	}

	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("handshake: could not read response body %s\n", err)
		return
	}

	log.Printf("handshake: response body:%s", resbody)
}
