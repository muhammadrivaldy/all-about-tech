package main

type Example2_User struct {
	ID          int
	Name        string
	Status      int
	Email       string
	PhoneNumber string
}

func (u *Example2_User) IsUserActive() bool {
	return u.Status == 1
}
