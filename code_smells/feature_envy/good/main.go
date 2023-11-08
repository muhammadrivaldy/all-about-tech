package main

import "fmt"

type salary struct {
	day int
	fee int
}

func (s salary) totalSalary() int {
	return s.day * s.fee
}

type employee struct {
	id      int
	name    string
	address string
}

func (e employee) printEmployeeSalary(salary salary) {
	fmt.Printf("Employee %s will receive Rp. %d for %d days\n", e.name, salary.totalSalary(), salary.day)
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
