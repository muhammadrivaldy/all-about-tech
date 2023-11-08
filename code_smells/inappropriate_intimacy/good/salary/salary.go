package salary

type Salary struct {
	Day int
	Fee int
}

func (s Salary) TotalSalary() int {
	return s.Day * s.Fee
}
