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
	var path string
	flag.StringVar(&path, "e", "/tmp/time.sock", "NTP client sock endpoint")
	flag.Parse()

	req := make([]byte, 48)

	req[0] = 0x1B

	rsp := make([]byte, 48)

	raddr, err := net.ResolveUnixAddr("unixgram", path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	laddr := &net.UnixAddr{Name: fmt.Sprintf("%s-client", raddr.Name), Net: "unixgram"}

	conn, err := net.DialUnix("unixgram", laddr, raddr)
	if err != nil {
		fmt.Printf("failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("failed while closing connection:", err)
		}
	}()

	fmt.Printf("time from (uinxgram) (%s)\n", conn.RemoteAddr())

	if _, err = conn.Write(req); err != nil {
		fmt.Printf("failed to send request: %v\n", err)
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
