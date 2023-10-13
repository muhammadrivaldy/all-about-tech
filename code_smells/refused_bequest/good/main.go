package main

import "fmt"

type Handling struct{}

func (h Handling) GoForward(name string) {
	fmt.Printf("%s go forward", name)
}

type Vehicle struct {
	Name string
	Handling
}

func (v Vehicle) StartMachine() {
	fmt.Printf("%s's start the machine!\n", v.Name)
}

type Bicycle struct {
	Name string
	Handling
}

func main() {
	veh := Vehicle{Name: "Car"}
	bic := Bicycle{Name: "Bicycle"}

	veh.StartMachine()
	veh.GoForward(veh.Name) // Output: Vehicle go forward
	bic.GoForward(bic.Name) // Output: Bicycle go forward
}
