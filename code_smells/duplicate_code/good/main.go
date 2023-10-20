package main

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
		BuildMachine()
	} else if request.Type == "crossover" {
		BuildMachine()
	}

	panic("the type is not eligible yet")

}

func BuildMachine() {
	/* {This is the code of creating machine of crossover} */
}
