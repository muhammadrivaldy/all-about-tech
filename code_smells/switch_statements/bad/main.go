package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := createOrder("normal"); err != nil {
		fmt.Print(err)
	}
}

func createOrder(types string) error {

	switch types {
	case "normal":
		// the process of creating normal order
	case "multiple":
		// the process of creating multiple order
	case "older":
		// the process of creating older order
	case "duplicate":
		// the process of creating duplicate order
	default:
		return errors.New("unexpected order types")
	}

	return nil

}
