package main

import (
	webserver "PR/Web_Cluster/web-server"
	"flag"
)

var (
	addressSet = [3]string{":3042", ":3043", ":3044"}
)

func main() {
	addressPtr := flag.Int("address", 0, "chooses one of the predifined ports")
	mainAddressPtr := flag.Int("mainAddress", 0, "chooses cluster leader from predefined ports")
	leaderPtr := flag.String("leader", "false", "determines if the server is cluster leader")
	flag.Parse()

	cluster := webserver.NewCluster("Hivemind", addressSet[:])
	server := webserver.NewWebServer(addressSet[*addressPtr], addressSet[*mainAddressPtr], *leaderPtr=="true", cluster)
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
