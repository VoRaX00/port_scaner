package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for port := range ports {
		addr := fmt.Sprintf("scanme.nmap.org:%d", port)
		conn, err := net.Dial("tcp", addr)

		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
		results <- port
	}
}

func main() {
	countPorts := 65536
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= countPorts; i++ {
			ports <- i
		}
	}()

	for i := 0; i < countPorts; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("Port %d open\n", port)
	}
}
