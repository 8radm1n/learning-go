package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	host string
)

func main() {
	flag.StringVar(&host, "host", "localhost", "host name to resolve")
	flag.Parse()

	// Specifiy use of the internal go dns resolver
	res := net.Resolver{PreferGo: true}
	addrs, err := res.LookupHost(context.Background(), host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(addrs)
}
