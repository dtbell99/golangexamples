package main

type truck struct {
	Make  string
	Model string
}

type bike struct {
	Brand string
}

type vehicle interface {
	honk() string
	getMake() string
	getModel() string
}

type ride interface {
	bell() string
}

func (t truck) honk() string {
	return "Truck Honked"
}

func (t truck) getMake() string {
	return t.Make
}

func (t truck) getModel() string {
	return t.Model
}

func (b bike) bell() string {
	return "bell digned"
}
