package main

import "fmt"

type Engine struct {
	Type string
}

func (e *Engine) BuildEngine() {
	fmt.Printf("Engine: Build an engine for %s\n", e.Type)
}

type Car struct {
	Engine   Engine
	Power    int
	MaxSpeed int
	Color    string
}

func (c *Car) BuildCar() {
	fmt.Println("Building a car with attributes:")
	fmt.Printf("Power: %d\n", c.Power)
	fmt.Printf("Max Speed: %d\n", c.MaxSpeed)
	fmt.Printf("Color: %s\n", c.Color)
}

func (c *Car) BuildEngine() {
	c.Engine.BuildEngine()
}

func main() {
	car := Car{
		Engine: Engine{
			Type: "crossover",
		},
		Power:    400,
		MaxSpeed: 250,
		Color:    "Black",
	}

	car.BuildCar()
	car.BuildEngine()
}
