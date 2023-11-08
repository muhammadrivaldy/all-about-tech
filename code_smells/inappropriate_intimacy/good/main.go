package main

import (
	"code_smells/inappropriate_intimacy/good/employee"
	"code_smells/inappropriate_intimacy/good/salary"
)

func main() {
	employee := employee.Employee{
		ID:      1,
		Name:    "Rival",
		Address: "Jl. Pancar Kav. 1",
	}

	salary := salary.Salary{
		Day: 20,
		Fee: 65000,
	}

	employee.PrintEmployeeSalary(salary)
}
