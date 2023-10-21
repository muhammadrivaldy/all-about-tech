package main

import (
	"code_smells/inappropriate_intimacy/good/area"
	"code_smells/inappropriate_intimacy/good/square"
)

func main() {
	squareObj := square.Square{}
	squareObj.SetSide(5)

	areaObj := area.Area{}
	areaObj.SetSquare(squareObj)
	areaObj.GetArea()
}
