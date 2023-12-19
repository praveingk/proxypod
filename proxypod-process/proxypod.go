package main

import (
	"flag"
	"log"

	"inet.af/tcpproxy"
)

func main() {
	listenPort := flag.String("port", "", "Local port to listen to")
	target := flag.String("target", "", "target service to redirect connections to")
	flag.Parse()
	var p tcpproxy.Proxy
	p.AddRoute(":"+*listenPort, tcpproxy.To(*target))
	log.Printf("Proxy created - from :" + *listenPort + " to " + *target)
	log.Fatal(p.Run())
}
