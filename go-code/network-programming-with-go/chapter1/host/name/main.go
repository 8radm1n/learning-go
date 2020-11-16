package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	addr string
)

func main() {
	flag.StringVar(&addr, "addr", "127.0.0.1", "host name to resolve")
	flag.Parse()

	addrs, err := net.LookupAddr(addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(addrs)
}
