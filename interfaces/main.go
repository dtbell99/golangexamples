package main

import "fmt"

type dog struct {
	Name string
}

type cat struct {
	Name string
}

type animal interface {
	speak() string
}

func (a dog) speak() string {
	return fmt.Sprintf("%s says Bark", a.Name)
}

func (a cat) speak() string {
	return fmt.Sprintf("%s says Meow", a.Name)
}

func talk(a animal) {
	fmt.Println(a.speak())
}

func main() {
	d1 := dog{Name: "Sparky"}
	c1 := cat{Name: "Whiskers"}

	talk(d1)
	talk(c1)
}
