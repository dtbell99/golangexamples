package main

import (
	"fmt"
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
}
