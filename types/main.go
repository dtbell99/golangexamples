package main

import "fmt"

func testHorn(v vehicle) {
	fmt.Printf("%s %s : %s\n", v.getMake(), v.getModel(), v.honk())
}

func testBell(r ride) {
	fmt.Printf("%v : %s\n", r, r.bell())
}

func main() {
	c := car{Make: "Ford", Model: "Fusion"}
	t := truck{Make: "Chevy", Model: "Tahoe"}
	b := bike{Brand: "Specialized"}
	testHorn(c)
	testHorn(t)
	//testHorn(b) b doesn't have honk so you cant compile this line.
	testBell(b)
}
