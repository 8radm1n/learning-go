package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var host string

	flag.StringVar(&host, "e", "us.pool.ntp.org:123", "NTP host")
	flag.Parse()

	req := make([]byte, 48)

	req[0] = 0x1B

	rsp := make([]byte, 48)

	raddr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		fmt.Printf("failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("fialed while closing connection: ", err)
		}
	}()

	fmt.Printf("time from (udp) %s\n", conn.RemoteAddr())

	if _, err = conn.Write(req); err != nil {
		fmt.Printf("failed to send request: %v", err)
		os.Exit(1)
	}

	read, err := conn.Read(rsp)
	if err != nil {
		fmt.Printf("failed to receive response: %v\n", err)
		os.Exit(1)
	}

	if read != 48 {
		fmt.Println("did not get all expected bytes from server")
		os.Exit(1)
	}

	secs := binary.BigEndian.Uint32(rsp[40:])
	frac := binary.BigEndian.Uint32(rsp[44:])

	ntpEpoch := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	unixEpoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	offset := unixEpoch.Sub(ntpEpoch).Seconds()
	now := float64(secs) - offset
	fmt.Printf("%v\n", time.Unix(int64(now), int64(frac)))
}
