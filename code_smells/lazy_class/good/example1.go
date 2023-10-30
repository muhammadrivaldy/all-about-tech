package main

type Example1_Contact struct {
	Email       string
	PhoneNumber string
}

func (c *Example1_Contact) IsValidPhoneNumber() bool {
	// the process of validating phone number
	return true // << it's just for example, it can be true | false
}

type Example1_User struct {
	ID     int
	Name   string
	Status int
	Example1_Contact
}

func (u *Example1_User) IsUserActive() bool {
	return u.Status == 1
}
