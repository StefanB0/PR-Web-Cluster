package webserver

import (
	"log"
	"math/rand"
	"time"
)

func randomizeSlice(s []string) []string {
	newS := make([]string, len(s))
	copy(newS, s)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(newS), func(i, j int) { newS[i], newS[j] = newS[j], newS[i] })

	return newS
}

func pruneSlice(s []string, item string) []string {
	newS := []string{}
	for _, el := range s {
		if el != item {
			newS = append(newS, el)
		}
	}
	return newS
}

func printRequest(addr, protocol, method, fcheck string) {
	netw := "internal"
	if fcheck == "true" {
		netw = "external"
	}
	
	if protocol == "HTTP" {
		log.Printf("Server %s, %s %s request , Type: %s", addr, netw, protocol, method)
	}
}
