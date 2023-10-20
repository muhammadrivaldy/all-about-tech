package main

type Car struct {
	Type     string
	Engine   string
	Power    int
	MaxSpeed int
	Color    string
}

func BuildCar(request Car) {

	// before we build a car, we should validate the payload to make sure
	// all of the requirements are eligible to process
	// because if one of the requirements are not eligible
	// will breaking the system
	if request.Engine != "vohc" {
		panic("the eligible engine type is only vohc")
	}

	if request.MaxSpeed < 2500 {
		panic("the maximum speed is only eligible under 2500cc")
	}

	/* This is another process of build car, we will skip it because it's only example */

}
