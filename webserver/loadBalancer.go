package webserver

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []string
}

func NewLoadBalancer(port string, servers []string) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}
