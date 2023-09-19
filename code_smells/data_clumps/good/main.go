package main

import "fmt"

type area struct {
	side side
	high float32
}

type side struct {
	a float32
	b float32
}

func calculateSquare(req side) {
	fmt.Println(req.a * req.b)
}

func calculateRectangle(req side) {
	fmt.Println(req.a * req.b)
}

func calculateTrapezoid(req area) {
	fmt.Println((req.side.a + req.side.b) * 0.5 * req.high)
}

func main() {
	req := area{
		side: side{
			a: 2,
			b: 4,
		},
		high: 10,
	}

	calculateSquare(req.side)
	calculateRectangle(req.side)
	calculateTrapezoid(req)
}
