package webserver

import (
	"PR/Web_Cluster/database"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func (s *WebServer) initHandlers() {
	http.HandleFunc("/", getHello)

	http.HandleFunc("/create", s.createElement)
	http.HandleFunc("/read", s.readElement)
	http.HandleFunc("/update", s.updateElement)
	http.HandleFunc("/delete", s.deleteElement)
	http.HandleFunc("/kill", s.kill)
	http.HandleFunc("/overwrite", s.overwriteMemory)
	http.HandleFunc("/handshake", s.handshake)
}

func (s *WebServer) initListen() {
	err := http.ListenAndServe(HTTP_PORT, nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server closed \n")
	} else if err != nil {
		log.Printf("error starting server %s\n", err)
		os.Exit(1)
	}
	s.serverAlive = false
}

func getHello(w http.ResponseWriter, h *http.Request) {
	log.Println("Hello world!")
}

func (s *WebServer) createElement(w http.ResponseWriter, r *http.Request) {
	printRequest(s.addressSelf, "HTTP", r.Method, r.Header.Get("Forward"))
	
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Server could not read request body: %s\n", err)
		return
	}

	var p database.Pair
	json.Unmarshal(reqBody, &p)
	if s.memory.Read(p.Key) != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	s.memory.Create(p.Key, p.Value)

	if r.Header.Get("Forward") != "true" {
		target := randomizeSlice(s.network)
		target = target[:(len(target)+1)/2]
		s.ledger[p.Key] = target
		s.forwardRequest(r, reqBody, target, "/create")
	}
}

func (s *WebServer) readElement(w http.ResponseWriter, r *http.Request) {
	printRequest(s.addressSelf, "HTTP", r.Method, r.Header.Get("Forward"))
	
	key := r.Header.Get("Key")
	if s.memory.Read(key) == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.Header.Get("Forward") != "true" {
		s.forwardRequest(r, nil, s.ledger[key], "/delete")
		delete(s.ledger, key)
	}
	w.Write(s.memory.Read(key))
}

func (s *WebServer) updateElement(w http.ResponseWriter, r *http.Request) {
	printRequest(s.addressSelf, "HTTP", r.Method, r.Header.Get("Forward"))

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Server could not read request body: %s\n", err)
	}

	var p database.Pair
	json.Unmarshal(reqBody, &p)

	if s.memory.Read(p.Key) == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	s.memory.Update(p.Key, p.Value)

	if r.Header.Get("Forward") != "true" {
		s.forwardRequest(r, reqBody, s.ledger[p.Key], "/update")
	}
}

func (s *WebServer) deleteElement(w http.ResponseWriter, r *http.Request) {
	printRequest(s.addressSelf, "HTTP", r.Method, r.Header.Get("Forward"))
	
	key := r.Header.Get("Key")
	if s.memory.Read(key) == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	s.memory.Delete(key)

	if r.Header.Get("Forward") != "true" {
		s.forwardRequest(r, nil, s.ledger[key], "/delete")
		delete(s.ledger, key)
	}
}

func (s *WebServer) kill(w http.ResponseWriter, r *http.Request) {
	printRequest(s.addressSelf, "HTTP", r.Method, r.Header.Get("Forward"))

	s.serverAlive = false
	if err := s.Shutdown(context.TODO()); err != nil {
		os.Exit(1)
		panic(err)
	}
}

func (s *WebServer) overwriteMemory(w http.ResponseWriter, r *http.Request) {
	printRequest(s.addressSelf, "HTTP", r.Method, r.Header.Get("Forward"))
	
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Server could not read request body: %s\n", err)
	}

	var db database.DatabaseInstance
	json.Unmarshal(reqBody, &db)

	s.memory = db
}

func (s *WebServer) handshake(w http.ResponseWriter, r *http.Request) {
	printRequest(s.addressSelf, "HTTP", r.Method, r.Header.Get("Forward"))

	message := r.Header.Get("message")
	log.Printf("Server nr %d. Message received\n\"%s\"", s.id, message)
	w.Write([]byte("Message received"))
}
