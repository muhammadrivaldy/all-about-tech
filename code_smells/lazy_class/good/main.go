package main

type Contact struct {
	Email       string
	PhoneNumber string
}

func (c *Contact) IsValidPhoneNumber() bool {
	// the process of validating phone number
	return true // << it's just for example, it can be true | false
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
