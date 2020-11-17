package main

import "fmt"

type Car struct {
	Make  string
	Model string
}

func main() {
	c1 := Car{Make: "Ford", Model: "Fusion"}

	c2 := c1
	c2.Model = "Escape"

	c3 := &c1
	c3.Make = "Chevy"

	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Printf("%p\n", &c1)
	fmt.Printf("%p\n", &c2)
	fmt.Printf("%p\n", c3)

	fmt.Printf("%t\n", &c1 == c3)

}
