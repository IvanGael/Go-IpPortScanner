package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func scanHost(ip string, ports chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for port := range ports {
		address := ip + ":" + strconv.Itoa(port)
		conn, err := net.DialTimeout("tcp", address, time.Second)
		if err != nil {
			fmt.Printf("Port %d on %s is closed\n", port, ip)
			continue
		}
		defer conn.Close()
		fmt.Printf("Port %d on %s is open\n", port, ip)
	}
}

func main() {
	var wg sync.WaitGroup
	ip := "192.168.137.1" // Change this to the IP range you want to scan
	ports := make(chan int, 100)

	// Start worker goroutines
	for i := 0; i < cap(ports); i++ {
		wg.Add(1)
		go scanHost(ip, ports, &wg)
	}

	// Queue up ports to scan
	for i := 1; i <= 1024; i++ {
		ports <- i
	}
	close(ports)

	// Wait for all scans to finish
	wg.Wait()
}
