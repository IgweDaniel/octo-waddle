package moves

import (
	"fmt"
	"strconv"
)

type Move struct {
	origin      int
	destination int
	// Queen = 2 Bishops = 3 Knights = 4 Rooks = 5
	promotion int
}

func NewMove(singleMove [3]int) *Move {
	mv := &Move{origin: singleMove[0], destination: singleMove[1], promotion: singleMove[2]}
	return mv
}

func AlgebriacToIndex(algebraic string) (int, error) {
	if algebraic == "-" {
		return 64, nil
	}
	column := string(algebraic[0])
	row, _ := strconv.Atoi(string(algebraic[1]))
	row--
	var square, file, rank int
	switch column {
	case "a":
		file = 0
	case "b":
		file = 1
	case "c":
		file = 2
	case "d":
		file = 3
	case "e":
		file = 4
	case "f":
		file = 5
	case "g":
		file = 6
	case "h":
		file = 7
	default:
		return 64, fmt.Errorf("invalid algebraic notation")
	}
	rank = ((7 - row) * 8)
	square = file + rank
	return square, nil
}
