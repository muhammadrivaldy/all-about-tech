package main

import (
	"code_smells/long_method/good/sub"
	"fmt"
)

func main() {
	services := sub.NewServices(sub.NewRepo())

	responseUser, err := services.GetUser(1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response user: %+v", responseUser)
}
