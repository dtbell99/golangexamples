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

func (c *car) updateMake(newMake string) {
	c.Make = newMake
}

// Won't work.. but will compile. It must be a pointer to modify.
func (c car) updateModel(newModel string) {
	c.Model = newModel
}

func main() {
	c1 := car{Make: "Ford", Model: "Fusion", Value: 25000.99}
	fmt.Println(c1.summary())

	c1.updateMake("chevy")
	fmt.Println(c1.summary())

	c1.updateModel("tohoe")
	fmt.Println(c1.summary())
}
