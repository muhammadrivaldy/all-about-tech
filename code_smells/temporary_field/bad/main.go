package main

import "fmt"

type area struct {
	side_a float32
	side_b float32
	high   float32
}

func (a area) calculateSquare() {
	fmt.Println(a.side_a * a.side_b)
}

func (a area) calculateRectangle() {
	fmt.Println(a.side_a * a.side_b)
}

func (a area) calculateTrapezoid() {
	fmt.Println((a.side_a + a.side_b) * 0.5 * a.high)
}

func main() {
	req := area{
		side_a: 2,
		side_b: 4,
		high:   10,
	}

	req.calculateSquare()
	req.calculateRectangle()
	req.calculateTrapezoid()
}
