package webserver

import (
	"PR/Web_Cluster/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getHello(w http.ResponseWriter, h *http.Request) {
	log.Println("Hello world!")
	fmt.Fprint(w, "hello\n")
}

func (s *WebServer) createElement(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Server could not read request body: %s\n", err)
	}

	var p database.Pair
	json.Unmarshal(reqBody, &p)

	s.memory.Create(p)
}

func (s *WebServer) readElement(w http.ResponseWriter, r *http.Request) {
	reqKey := r.Header.Get("Key")
	w.Write(s.memory.Read(reqKey))
}

func (s *WebServer) updateElement(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Server could not read request body: %s\n", err)
	}

	var p database.Pair
	json.Unmarshal(reqBody, &p)

	s.memory.Update(p)
}

func (s *WebServer) deleteElement(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Server could not read request body: %s\n", err)
	}

	var p database.Pair
	json.Unmarshal(reqBody, &p)

	s.memory.Delete(p.Key)
}

func (s *WebServer) overwriteMemory(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Server could not read request body: %s\n", err)
	}

	var db map[string][]byte
	json.Unmarshal(reqBody, &db)

	s.memory.OverwriteMemory(db)
}
