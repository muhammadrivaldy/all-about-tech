package sub

import "time"

type order struct {
	ID        int
	ItemID    int
	Price     int
	TotalItem int
	CreatedAt time.Time
	UpdatedAt time.Time
}
