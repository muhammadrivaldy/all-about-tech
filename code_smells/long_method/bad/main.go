package main

import (
	"errors"
	"fmt"
	"time"

	goutil "github.com/muhammadrivaldy/go-util"
)

func main() {
	const selectUserByID = "select id, name, email, phone, address, status, created_at, updated_at from users where id = ?"

	clientDB, _ := goutil.NewMySQL("root", "root", "localhost:3306", "testing", nil)

	var user user

	if err := clientDB.QueryRow(selectUserByID, 1).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Address,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		panic(err)
	}

	if !user.Status.isActive() {
		panic(errors.New("user not active"))
	}

	responseUser := responseUser{
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		Address: user.Address,
	}

	fmt.Printf("response user: %+v", responseUser)
}

type status int

const (
	userActive    status = 1
	userNotActive status = 0
)

func (s status) isActive() bool {
	return s == userActive
}

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

type responseUser struct {
	Name    string
	Email   string
	Phone   string
	Address string
}
