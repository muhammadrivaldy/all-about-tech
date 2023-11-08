package main

import "fmt"

type salary struct {
	day int
	fee int
}

type employee struct {
	id      int
	name    string
	address string
}

func (e employee) printEmployeeSalary(salary salary) {
	totalSalary := salary.day * salary.fee
	fmt.Printf("Employee %s will receive Rp. %d for %d days\n", e.name, totalSalary, salary.day)
}

func main() {
	employee := employee{
		id:      1,
		name:    "Rival",
		address: "Jl. Pancar Kav. 1",
	}

	salary := salary{
		day: 20,
		fee: 65000,
	}

	employee.printEmployeeSalary(salary)
}
