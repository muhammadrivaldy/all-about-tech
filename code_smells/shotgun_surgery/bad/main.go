package main

import "fmt"

type statusUser int

const STATUS_USER_ACTIVE statusUser = 1
const STATUS_USER_NON_ACTIVE statusUser = 1

type User struct {
	status statusUser
}

func (u *User) CheckIn() {
	if u.status == STATUS_USER_ACTIVE {
		fmt.Println("User check in")
	}
}

func (u *User) DoWork() {
	if u.status == STATUS_USER_ACTIVE {
		fmt.Println("User do work")
	}
}

func (u *User) CheckOut() {
	if u.status == STATUS_USER_ACTIVE {
		fmt.Println("User check out")
	}
}
