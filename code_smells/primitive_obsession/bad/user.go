package example

const USER_STATUS_ACTIVE = 1
const USER_STATUS_NOT_ACTIVE = 2

type User struct {
	Name        string
	PhoneNumber string
	Status      int
}

func (u *User) CheckingFormatName() {
	// code for checking format name
}

func (u *User) CheckingPhoneNumberPrefix() {
	// code for checking phone number prefix
}
