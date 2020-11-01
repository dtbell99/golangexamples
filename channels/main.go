package main

import (
	"fmt"
	"strings"
	"time"
)

func upit(s string, t time.Duration, c chan string) {
	fmt.Printf("sleeping on %s\n", s)
	time.Sleep(t * time.Millisecond)
	c <- strings.ToUpper(s)
	fmt.Printf("finished processing %s\n", s)
}

func main() {
	c := make(chan string)
	s := []string{"dave", "bell", "bob", "smith"}
	for _, itm := range s {
		fmt.Printf("Calling go upit on %s\n", itm)
		go upit(itm, 2000, c)
		fmt.Printf("Post go upit %s\n", itm)
	}

	s1, s2, s3, s4 := <-c, <-c, <-c, <-c
	fmt.Printf("%s %s %s %s\n", s1, s2, s3, s4)

}
