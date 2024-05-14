package main

import (
	"fmt"
	"net"
)

func main() {
	for port := 1; port < 1024; port++ {
		addr := fmt.Sprintf("scanme.nmap.org:%d", port)
		conn, err := net.Dial("tcp", addr)

		if err != nil {
			continue
		}

		conn.Close()
		fmt.Printf("Port %d open\n", port)
	}
}
