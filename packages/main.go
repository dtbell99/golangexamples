package main

import (
	"fmt"

	"github.com/dtbell99/golangexamples/packages/jokes"
)

func main() {
	fmt.Printf("Joke One: %s\n", jokes.GetJokeOne())
	fmt.Printf("Joke Two: %s\n", jokes.GetJokeTwo())
}
