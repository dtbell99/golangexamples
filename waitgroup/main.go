package main

import (
	"fmt"
	"sync"
	"time"
)

func process(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		process("Hello")
		wg.Done()
	}()

	fmt.Println("wg.Wait()")
	wg.Wait()
	fmt.Println("Done")

}
