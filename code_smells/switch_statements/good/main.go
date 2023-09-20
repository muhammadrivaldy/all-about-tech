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
		createNormalOrder()
	case "multiple":
		createMultipleOrder()
	case "older":
		createOlderOrder()
	case "duplicate":
		createDuplicateOrder()
	default:
		return errors.New("unexpected order types")
	}

	return nil

}

func createNormalOrder() {
	// the process of creating normal order
}

func createMultipleOrder() {
	// the process of creating multiple order
}

func createOlderOrder() {
	// the process of creating older order
}

func createDuplicateOrder() {
	// the process of creating duplicate order
}
