package main

import "fmt"

/*
	# Structure codes:

	- Car
		- Toyota
		- Honda
		- Tesla
*/

type ICar interface {
	BuildCar()
}

type Car struct {
	Type string
}

type Toyota struct {
	Car
}

func NewToyota(car Car) ICar {
	return &Toyota{
		Car: car,
	}
}

func (t *Toyota) BuildCar() {
	fmt.Printf("We build Toyota %s\n", t.Type)
}

type Honda struct {
	Car
}

func NewHonda(car Car) ICar {
	return &Honda{
		Car: car,
	}
}

func (t *Honda) BuildCar() {
	fmt.Printf("We build Honda %s\n", t.Type)
}

type Tesla struct {
	Car
}

func NewTesla(car Car) ICar {
	return &Tesla{
		Car: car,
	}
}

func (t *Tesla) BuildCar() {
	fmt.Printf("We build Tesla %s\n", t.Type)
}

func main() {
	car := Car{Type: "crossover"}
	toyota := NewToyota(car)

	toyota.BuildCar()
}
