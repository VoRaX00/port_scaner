package main

import (
	"fmt"
	"net"
	"sync"
)

func isPortOpened(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	addr := fmt.Sprintf("scanme.nmap.org:%d", port)
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		return
	}

	conn.Close()
	fmt.Printf("Port %d open\n", port)
}

func main() {
	var wg sync.WaitGroup
	for port := 1; port < 1024; port++ {
		wg.Add(1)
		go isPortOpened(port, &wg)
	}
	wg.Wait()
}
