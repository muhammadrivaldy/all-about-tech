package main

import "fmt"

type Car struct {
	Type     string
	Engine   string
	Power    int
	MaxSpeed int
	Color    string
}

func BuildCar(request Car) {

	/* {This is the code of validating} */

	if request.Type == "wagon" {
		BuildMachine("wagon family")
	} else if request.Type == "crossover" {
		BuildMachine("crossover family")
	}

	panic("the type is not eligible yet")

}

func BuildMachine(typeCar string) {
	fmt.Printf("We build a machine for %s\n", typeCar)
}
