package main

import "fmt"

type User struct {
	ID   int
	Name string
}

// Repository is responsible for data access.
type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetUserByID(id int) User {
	/* {Code for getting data user from the database} */
	return User{
		ID:   1,
		Name: "Rivaldy",
	}
}

func main() {
	repo := NewRepository()
	fmt.Println(repo.GetUserByID(1))
}
