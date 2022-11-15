package main

import (
	"PR/Web_Cluster/webserver"
	"flag"
	"fmt"
)

func main() {
	idPtr := flag.Int("id", 0, "server id")
	addressPtr := flag.String("address", "http:minion", "server address")
	portPtr := flag.String("port", ":3000", "chooses one of the predifined ports")
	leaderPtr := flag.String("leader", "http:leader0:3000", "cluster leader address")
	isLeaderPtr := flag.Bool("isLeader", false, "defines if server is cluster leader at creation")
	flag.Parse()

	addressSet := []string{fmt.Sprint(*addressPtr, *idPtr, *portPtr)}
	cluster := webserver.NewCluster("Hivemind", addressSet[:])
	server := webserver.NewWebServer(*idPtr, *addressPtr, *portPtr)
	server.SetLeader(*leaderPtr, *isLeaderPtr)
	server.SetCluster(cluster)

	server.StartServer()
}

//done get request
//todo cluster logic
//todo sync-backup of part leader

//todo client application
//todo unit tests

//done input from CLI
//todo docker
//todo docker compose
