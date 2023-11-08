package salary

type Salary struct {
	day int
	fee int
}

func (s Salary) totalSalary() int {
	return s.day * s.fee
}
