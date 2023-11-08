package employee

import (
	"code_smells/inappropriate_intimacy/good/salary"
	"fmt"
)

type Employee struct {
	ID      int
	Name    string
	Address string
}

func (e Employee) PrintEmployeeSalary(salary salary.Salary) {
	fmt.Printf("Employee %s will receive Rp. %d for %d days\n", e.Name, salary.TotalSalary(), salary.Day)
}
