package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var ports = []string{"20", "21", "22", "23", "25", "53", "80", "110", "119", "123", "143", "161", "194", "443", "5900"}

var machines = make([]string, 0)

func writeLog() {
	now := time.Now()
	day := now.Day()
	month := now.Month()
	year := now.Year()
	hour := now.Hour()
	minute := now.Minute()
	seconds := now.Second()

	fName := fmt.Sprintf("Machines_%d-%d-%d_%d-%d-%d.log", month, day, year, hour, minute, seconds)
	file, err := os.OpenFile(fName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	datawriter := bufio.NewWriter(file)
	for _, data := range machines {
		_, _ = datawriter.WriteString(data + "\n")
	}
	datawriter.Flush()
}

func rawConnect(host string, ports []string) []string {
	found := make([]string, 0)
	for _, port := range ports {
		fmt.Printf("%s::%s\n", host, port)
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 3*time.Second)
		if err == nil && conn != nil {
			defer conn.Close()
			hl, _ := net.LookupAddr(host)
			foundStr := fmt.Sprintf("%s :: %s", net.JoinHostPort(host, port), hl)
			found = append(found, foundStr)
		}
	}
	return found
}

func findByPort(ip string) {
	found := rawConnect(ip, ports)
	for _, f := range found {
		machines = append(machines, f)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 255; i++ {
		classC := "10.0.0"
		ip := fmt.Sprintf("%s.%d", classC, i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			findByPort(ip)
		}()
	}
	wg.Wait()
	writeLog()
}
