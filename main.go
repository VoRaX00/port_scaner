package main

import (
	"fmt"
	"net"
	"sort"

	"github.com/VoRaX00/ProgressBar/progressbar"
)

func worker(ports, results chan int, bar *progressbar.ProgressBar, count *int) {
	for port := range ports {
		addr := fmt.Sprintf("scanme.nmap.org:%d", port)
		conn, err := net.Dial("tcp", addr)

		bar.Play(int64(*count))
		(*count)++
		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
		results <- port
	}
}

func main() {
	fmt.Println("Scanning ports...")
	var bar progressbar.ProgressBar

	// for i := 0; i <= 100; i++ {
	// 	time.Sleep(time.Second * 1)
	// 	bar.Play(int64(i))
	// }
	// bar.Finish()
	countPorts := 65536
	count := 1
	bar.NewOption(0, int64(countPorts))
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, &bar, &count)
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
	bar.Finish()
	fmt.Println("Here are your open ports:")
	for _, port := range openports {
		fmt.Printf("Port %d open\n", port)
	}
}
