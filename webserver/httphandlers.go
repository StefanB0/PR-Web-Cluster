package webserver

import (
	"PR/Web_Cluster/database"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *WebServer) initHandlers() {
	http.HandleFunc("/", getHello)

	http.HandleFunc("/create", s.createElement)
	http.HandleFunc("/read", s.readElement)
	http.HandleFunc("/update", s.updateElement)
	http.HandleFunc("/delete", s.deleteElement)
	http.HandleFunc("/overwrite", s.overwriteMemory)
	http.HandleFunc("/handshake", s.handshake)
}

func getHello(w http.ResponseWriter, h *http.Request) {
	log.Println("Hello world!")
}

func (s *WebServer) createElement(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Server could not read request body: %s\n", err)
	}

	var p database.Pair
	json.Unmarshal(reqBody, &p)

	s.memory.Create(p.Key, p.Value)
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

	s.memory.Update(p.Key, p.Value)
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

	var db database.DatabaseInstance
	json.Unmarshal(reqBody, &db)

	s.memory = db
}

func (s *WebServer) handshake(w http.ResponseWriter, r *http.Request) {
	message := r.Header.Get("message")
	log.Printf("Server nr %d. Message received\n\"%s\"", s.id, message)
	w.Write([]byte("Message received"))
}
