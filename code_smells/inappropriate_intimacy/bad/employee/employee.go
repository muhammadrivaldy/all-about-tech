package employee

import (
	"code_smells/inappropriate_intimacy/bad/salary"
	"fmt"
)

type Employee struct {
	id      int
	name    string
	address string
}

func (e Employee) printEmployeeSalary(salary salary.Salary) {
	fmt.Printf("Employee %s will receive Rp. %d for %d days\n", e.name, salary.totalSalary(), salary.day) // this is expected error because golang not supported use private method or attribute
}
