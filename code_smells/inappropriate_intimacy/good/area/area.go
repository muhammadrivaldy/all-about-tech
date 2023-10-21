package area

import (
	"code_smells/inappropriate_intimacy/good/square"
	"fmt"
)

type Area struct {
	square square.Square
}

func (a *Area) SetSquare(square square.Square) {
	a.square = square
}

func (a *Area) GetArea() {
	fmt.Println(a.square.GetSide() * a.square.GetSide())
}
