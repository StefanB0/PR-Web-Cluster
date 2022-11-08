package main

import webserver "PR/Web_Cluster/web-server"

const (
	address = ":3042"
)

func main() {
	cluster := webserver.NewCluster("Hivemind", []string{address})
	server := webserver.NewWebServer(address, address, true, cluster)
	server.StartServer()
}