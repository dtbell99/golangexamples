package main

import "fmt"

type car struct {
	Make  string
	Model string
	Value float32
}

func (c car) summary() string {
	return fmt.Sprintf("%s %s :: %f", c.Make, c.Model, c.Value)
}

func main() {
	c1 := car{Make: "Ford", Model: "Fusion", Value: 25000.99}
	fmt.Println(c1.summary())
}
