package main

import (
	"fmt"
	"jokes"
	"strings"
)

func processArrayOption1(dta []string) []string {
	newDta := make([]string, 0)
	for _, itm := range dta {
		itm := strings.ToUpper(itm)
		newDta = append(newDta, itm)
	}
	return newDta
}

func processArrayOption2(dta []string) {
	for p, itm := range dta {
		itm = strings.ToUpper(itm)
		dta[p] = itm
	}
}

func changeString(s *string) {
	*s = strings.ToUpper(*s)
}

func changeString2(s string) string {
	s2 := strings.ToUpper(s)
	return s2
}

func printData(dta []string, title string) {
	fmt.Println(title)
	for _, itm := range dta {
		fmt.Println(itm)
	}
}

func main() {
	fmt.Println("Pointers Example\n----------------")

	dta := []string{"dave", "bell"}
	dta = processArrayOption1(dta)
	printData(dta, "Option 1 : Return new array")

	dta2 := []string{"bob", "smith"}
	processArrayOption2(dta2[:])
	printData(dta2, "Option 2 : Pointer Manipulation")

	s := "brody"
	changeString(&s)
	fmt.Printf("Value of Brody\n%s\n", s)

	s2 := "samantha"
	s3 := changeString2(s2)
	fmt.Printf("Value of %s is now %s\n", s2, s3)

	jokeOne := jokes.GetJokeOne()
}
