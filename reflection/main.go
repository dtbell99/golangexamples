package main

import (
	"fmt"
	"reflect"
)

type Car struct {
	Make  string
	Model string
}

type boat struct {
	Make string
}

func main() {
	c1 := Car{Make: "Ford", Model: "Fusion"}
	c2 := Car{Make: "Ford", Model: "Escape"}
	fmt.Println(reflect.TypeOf(c1))
	fmt.Println(reflect.TypeOf(c2))

	b1 := boat{Make: "Bayliner"}
	fmt.Println(reflect.TypeOf(b1))
}
