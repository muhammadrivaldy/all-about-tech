package sub

import "time"

type user struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	Address   string
	Status    status
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u user) userToResponseUser() responseUser {
	return responseUser{
		Name:    u.Name,
		Email:   u.Email,
		Phone:   u.Phone,
		Address: u.Address,
	}
}

type responseUser struct {
	Name    string
	Email   string
	Phone   string
	Address string
}
