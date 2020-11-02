package main

import (
	"fmt"
	"strings"
)

func main() {
	c := make(chan string)
	go process("DaveBell", c)
	for msg := range c {
		fmt.Printf("Message:%s\n", msg)
	}
}

func process(s string, c chan string) {
	for i := 1; i <= 5; i++ {
		s2 := strings.ToUpper(s)
		msg := fmt.Sprintf("%s :: %d", s2, i)
		c <- msg
	}

	close(c) // Comment this out to see deadlock msg
}
