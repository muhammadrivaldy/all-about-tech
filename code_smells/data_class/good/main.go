package main

import "fmt"

type AreaOfCircle struct {
	Phi    float32
	Radius float32 // in cm
}

func (a AreaOfCircle) GetArea() float32 {
	return a.Phi * (a.Radius * a.Radius)
}

func main() {
	object := AreaOfCircle{Phi: 3.14, Radius: 7}
	fmt.Println(object.GetArea()) // Output: 153.86
}
