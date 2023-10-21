package main

func main() {
	mysql := NewMySQL()
	mysql.selectUserByID(123)
}

type User interface {
	selectUserByID(id int) error
	selectUserByName(name string) error
	selectUserByCity(city string) error
}

type mysql struct{}

func NewMySQL() User {
	return &mysql{}
}

func (m *mysql) selectUserByID(id int) error {
	return nil
}

func (m *mysql) selectUserByName(name string) error {
	return nil
}

func (m *mysql) selectUserByCity(city string) error {
	return nil
}
