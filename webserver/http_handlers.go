package webserver

import (
	"PR/Web_Cluster/database"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func (s *WebServer) initHandlers() {
	http.HandleFunc("/", getHello)

	http.HandleFunc("/create", s.createElement)
	http.HandleFunc("/read", s.readElement)
	http.HandleFunc("/update", s.updateElement)
	http.HandleFunc("/syncRequest", s.syncRequestHandler)
	http.HandleFunc("/delete", s.deleteElement)
	http.HandleFunc("/kill", s.kill)
	http.HandleFunc("/overwrite", s.overwriteMemory)
	http.HandleFunc("/handshake", s.handshake)
}

func (s *WebServer) initProxy() {
	s.proxylist = make(map[string]*httputil.ReverseProxy)
	for _, addr := range s.network {
		rUrl, _ := url.Parse(HTTP_PREFIX + addr + ":3000")
		s.proxylist[addr] = httputil.NewSingleHostReverseProxy(rUrl)
	}
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

	if s.isLeader && r.Header.Get("Forward") != "true" && len(s.ledger[key]) > 0 {
		randProxy := rand.Intn(len(s.ledger[key]))
		proxy := s.proxylist[s.ledger[key][randProxy]]
		r.Header.Add("Forward", "true")
		proxy.ServeHTTP(w, r)
	} else {
		w.Write(s.memory.Read(key))
	}
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

func (s *WebServer) syncRequestHandler(w http.ResponseWriter, r *http.Request) {
	server_name := r.Header.Get("server_name")
	snapshot := &Snapshot{database.NewDatabase(), s.ledger}

	for key, serverSet := range s.ledger {
		if checkSlice(serverSet, server_name) {
			snapshot.Memory.Create(key, s.memory.Read(key))
		}
	}

	jsondata, err := json.Marshal(*snapshot)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsondata)
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
