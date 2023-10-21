package square

type Square struct {
	side int
}

func (s *Square) SetSide(side int) {
	s.side = side
}

func (s *Square) GetSide() int {
	return s.side
}
