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

	flag.StringVar(&host, "e", ":11123", "server address")
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("failed to create socket: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("listenting for time requests: (udp) %s\n", conn.LocalAddr())

	_, raddr, err := conn.ReadFromUDP(make([]byte, 48))
	if err != nil {
		fmt.Println("err getting request: ", err)
		os.Exit(1)
	}

	if raddr == nil {
		fmt.Println("request missing remote addr")
		os.Exit(1)
	}

	secs, fracs := getNTPSeconds(time.Now())

	rsp := make([]byte, 48)
	// write seconds (as uint32) in buffer at [40:43]
	binary.BigEndian.PutUint32(rsp[40:], uint32(secs))
	// write seconds (as uint32) in buffer at [44:47]
	binary.BigEndian.PutUint32(rsp[44:], uint32(fracs))

	if _, err := conn.WriteToUDP(rsp, raddr); err != nil {
		fmt.Println("err sending data: ", err)
		os.Exit(1)
	}

}

// getNTPSecs decompose current time as NTP seconds
func getNTPSeconds(t time.Time) (int64, int64) {
	// convert time to total # of secs since 1970
	// add NTP epoch offets as total #secs between 1900-1970
	secs := t.Unix() + int64(getNTPOffset())
	fracs := t.Nanosecond()
	return secs, int64(fracs)
}

// getNTPOffset returns the 70yrs between Unix epoch
// and NTP epoch (1970-1900) in seconds
func getNTPOffset() float64 {
	ntpEpoch := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	unixEpoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	offset := unixEpoch.Sub(ntpEpoch).Seconds()
	return offset
}
