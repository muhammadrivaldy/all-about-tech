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
