package main

import "fmt"

type Vehicle struct {
	Name string
}

func (v Vehicle) StartMachine() {
	fmt.Printf("%s's start the machine!\n", v.Name)
}

func (v Vehicle) GoForward() {
	fmt.Printf("%s go forward\n", v.Name)
}

type Bicycle struct {
	Vehicle
}

func main() {
	veh := Vehicle{Name: "Car"}
	bic := Bicycle{Vehicle{Name: "Bicycle"}}

	veh.StartMachine() // Output: Car's start the machine!
	bic.StartMachine() // Output: Bicycle's start the machine! (it's still work but not make sense because bicycle doesn't have machine)
	veh.GoForward()    // Output: Car go forward
	bic.GoForward()    // Output: Bicycle go forward
}
