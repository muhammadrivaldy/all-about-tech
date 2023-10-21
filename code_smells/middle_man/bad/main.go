package main

import "fmt"

type User struct {
	ID   int
	Name string
}

// Service is a middle man that delegates class to the Repository.
type Service struct {
	repo *Repository
}

func NewService() *Service {
	return &Service{
		repo: NewRepository(),
	}
}

func (s *Service) GetUserByID(id int) User {
	user := s.repo.GetUserByID(id)
	return user
}

// Repository is the actual class responsible for data access.
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
	service := NewService()
	fmt.Println(service.GetUserByID(1))
}
