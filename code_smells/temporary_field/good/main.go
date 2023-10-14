package main

import "fmt"

type areaOfSquareFamily struct {
	side_a float32
	side_b float32
}

func (as areaOfSquareFamily) calculateSquare() {
	fmt.Println(as.side_a * as.side_b)
}

func (as areaOfSquareFamily) calculateRectangle() {
	fmt.Println(as.side_a * as.side_b)
}

func (as areaOfSquareFamily) calculateTrapezoid(at areaOfTrapezoid) {
	at.calculateTrapezoid(as)
}

type areaOfTrapezoid struct {
	high float32
}

func (at areaOfTrapezoid) calculateTrapezoid(as areaOfSquareFamily) {
	fmt.Println((as.side_a + as.side_b) * 0.5 * at.high)
}

func main() {
	req := areaOfSquareFamily{
		side_a: 2,
		side_b: 4,
	}

	req.calculateSquare()
	req.calculateRectangle()
	req.calculateTrapezoid(areaOfTrapezoid{high: 10})
}
