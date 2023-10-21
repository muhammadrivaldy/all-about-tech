package main

import "fmt"

type AreaOfCircle struct {
	Phi    float32
	Radius float32 // in cm
}

func main() {
	object := AreaOfCircle{Phi: 3.14, Radius: 7}
	fmt.Println(object.Phi * (object.Radius * object.Radius)) // Output: 153.86
}
