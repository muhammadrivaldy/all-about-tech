package example

type User struct {
	Name        string
	PhoneNumber PhoneNumber
	Status      Status
}

func (u *User) CheckingFormatName() {
	// code for checking format name
}
