package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	goutil "github.com/muhammadrivaldy/go-util"
)

func main() {
	services := newServices(newRepo())

	responseUser, err := services.getUser(1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response user: %+v", responseUser)
}

const selectUserByID = "select id, name, email, phone, address, status, created_at, updated_at from users where id = ?"

type repositories struct {
	clientDB *sql.DB
}

func newRepo() *repositories {
	clientDB, _ := goutil.NewMySQL("root", "root", "localhost:3306", "testing", nil)
	return &repositories{
		clientDB: clientDB,
	}
}

func (repo *repositories) getUser(userID int) (user, error) {
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

type services struct {
	repo *repositories
}

func newServices(repo *repositories) *services {
	return &services{
		repo: repo,
	}
}

func (s *services) getUser(userID int) (responseUser, error) {
	response, err := s.repo.getUser(userID)
	if err != nil {
		return responseUser{}, err
	}

	if !response.Status.isActive() {
		return responseUser{}, errors.New("user not active")
	}

	return response.userToResponseUser(), nil
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
