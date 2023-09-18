package main

import "code_smells/large_class/bad/sub"

func main() {
	repositories := sub.NewRepo()
	repositories.GetUser(1)  // for getting user
	repositories.GetOrder(1) // for getting order
}
