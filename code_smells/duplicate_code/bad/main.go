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
		fmt.Println("We build a machine for wagon family") // Let's assume that the code more complex than this
	} else if request.Type == "crossover" {
		fmt.Println("We build a machine for crossover family") // Let's assume that the code more complex than this
	}

	panic("the type is not eligible yet")

}
