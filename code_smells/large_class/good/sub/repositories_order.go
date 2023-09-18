package sub

import (
	"database/sql"

	goutil "github.com/muhammadrivaldy/go-util"
)

const (
	selectOrderByID = "select id, item_id, price, total_item, created_at, updated_at from orders where id = ?"
)

type repositoriesOrder struct {
	clientDB *sql.DB
}

func NewRepoOrder() *repositoriesOrder {
	clientDB, _ := goutil.NewMySQL("root", "root", "localhost:3306", "testing", nil)
	return &repositoriesOrder{
		clientDB: clientDB,
	}
}

func (repo *repositoriesOrder) GetOrder(orderID int) (order, error) {
	var order order

	if err := repo.clientDB.QueryRow(selectOrderByID, orderID).Scan(
		&order.ID,
		&order.ItemID,
		&order.Price,
		&order.TotalItem,
		&order.CreatedAt,
		&order.UpdatedAt,
	); err != nil {
		return order, err
	}

	return order, nil
}
