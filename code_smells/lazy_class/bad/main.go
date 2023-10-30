package main

type Contact struct {
	Email       string
	PhoneNumber string
}

type User struct {
	ID     int
	Name   string
	Status int
	Contact
}

func (u *User) IsUserActive() bool {
	return u.Status == 1
}
