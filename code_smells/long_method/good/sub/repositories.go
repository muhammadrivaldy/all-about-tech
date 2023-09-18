package sub

import (
	"database/sql"

	goutil "github.com/muhammadrivaldy/go-util"
)

const selectUserByID = "select id, name, email, phone, address, status, created_at, updated_at from users where id = ?"

type repositories struct {
	clientDB *sql.DB
}

func NewRepo() *repositories {
	clientDB, _ := goutil.NewMySQL("root", "root", "localhost:3306", "testing", nil)
	return &repositories{
		clientDB: clientDB,
	}
}

func (repo *repositories) GetUser(userID int) (user, error) {
	var user user

	if err := repo.clientDB.QueryRow(selectUserByID, userID).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Address,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return user, err
	}

	return user, nil
}
