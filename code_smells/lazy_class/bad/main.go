package main

type Type struct {
	// the object doesn't have any attributes
}

func (t *Type) GetTypeOfWagon() string {
	return "wagon"
}

func (t *Type) GetTypeOfCrossover() string {
	return "crossover"
}

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
