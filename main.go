package main

import (
	"PR/Web_Cluster/webserver"
	"flag"
)

var (
	idPtr       *int
	portPtr     *string
	addressPtr  *string
	leaderPtr   *string
	isLeaderPtr *bool

	addressSet []string
)

func readConfig() {
	idPtr = flag.Int("id", 0, "server id")
	addressPtr = flag.String("address", "minion1", "server address")
	portPtr = flag.String("port", ":3000", "chooses one of the predifined ports")
	leaderPtr = flag.String("leader", "leader0", "cluster leader address")
	isLeaderPtr = flag.Bool("isLeader", false, "defines if server is cluster leader at creation")
	flag.Parse()

	addressSet = []string{"leader0", "minion1", "minion2"}
}

func startServer() {
	server := webserver.NewWebServer(*idPtr, *addressPtr, *portPtr, addressSet)
	server.SetLeader(*leaderPtr, *isLeaderPtr)
	server.StartServer()
}

func main() {
	readConfig()
	startServer()
}
