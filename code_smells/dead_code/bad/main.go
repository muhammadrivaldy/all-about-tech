package main

import "fmt"

type area struct {
	side_a float32
	side_b float32
	high   float32
}

func calculateSquare(req area) {
	fmt.Println(req.side_a * req.side_b)
}

func calculateRectangle(req area) {
	fmt.Println(req.side_a * req.side_b)
}

func main() {
	req := area{
		side_a: 2,
		side_b: 4,
		high:   10,
	}

	calculateSquare(req)
	calculateRectangle(req)
}
