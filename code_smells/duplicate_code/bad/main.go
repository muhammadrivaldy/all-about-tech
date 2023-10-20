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
		/* {This is the code of creating machine of wagon} */
	} else if request.Type == "crossover" {
		/* {This is the code of creating machine of crossover} */
	}

	panic("the type is not eligible yet")

}
