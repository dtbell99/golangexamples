package main

type car struct {
	Make  string
	Model string
}

func (c car) honk() string {
	return "Car Honked"
}

func (c car) getMake() string {
	return c.Make
}

func (c car) getModel() string {
	return c.Model
}
