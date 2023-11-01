package main

import (
	"fmt"
	"time"
)

type job struct {
	title string
	join  time.Time
}

type user struct {
	name   string
	addres string
	job
}

func (u *user) PrintUserInformation() {
	fmt.Printf("My name is %s and I live in %s, I have been joining this company start from %d. My title at this company is %s\n",
		u.name,
		u.addres,
		u.join.Year(),
		u.title)
}

func main() {
	job := job{
		title: "Staff officer",
		join:  time.Date(1997, 10, 10, 0, 0, 0, 0, time.Local),
	}

	user := user{
		name:   "Rivaldy",
		addres: "Jakarta",
		job:    job,
	}

	user.PrintUserInformation()
}
