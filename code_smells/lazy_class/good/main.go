package main

type Type string

const TypeWagon Type = "wagon"
const TypeCrossover Type = "crossover"

type Car struct {
	Type     Type
	Engine   string
	Power    int
	MaxSpeed int
	Color    string
}

func (c *Car) BuildCar() {
	/* {The code of building car} */
}
