package position

import (
	"fmt"
	"testing"
)

func TestMoveGeneration(t *testing.T) {

	// p := NewFenPosition("8/8/8/2p1pP2/1P1P4/8/8/8 w KQkq e6 0 1")

	// tricky position
	// p := NewFenPosition("rnbqkbnr/pppp1p1p/8/4p2P/4PPp1/8/PPPP2P1/RNBQKBNR b KQkq f3 0 4")

	// killer position
	// p := NewFenPosition("rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1")

	// Knight position
	// p := NewFenPosition("rnbqkb1r/ppppppp1/5n2/7p/3PP3/5N2/PPP2PPP/RNBQKB1R b KQkq - 1 3")
	// p = NewFenPosition("rnbqkb1r/ppppppp1/8/7p/3Pn3/5N2/PPP2PPP/RNBQKB1R w KQkq - 0 4")

	// Bishop position
	// p := NewFenPosition("r3k2r/p1pp1pb1/bn1qpnp1/3PN3/1p2PQ2/2N4p/PPPBBPPP/R3K2R w KQkq - 2 2")

	// start position
	// p = NewFenPosition("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	// castle position
	// p = NewFenPosition("4k2r/p1pp1pb1/1n2pnp1/3PN1q1/1p2PQ2/1PNb3P/PrPB1P1P/R3K2R w KQk - 1 4")

	// castle chop position and move
	// p = NewFenPosition("rnb1kb1r/pB1p2Qp/1p3p2/4p3/3PP3/2N1BNn1/qPP2PPP/R3K2R w KQkq - 0 1")

	fen := "rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1"
	// p = NewFenPosition("rnbq1rk1/pp1p1pPp/3b4/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR b KQ e6 0 1")
	p := NewFenPosition(fen)
	fmt.Println("old Pos")
	fmt.Println("prev pos empty", p.prevPosition == nil)
	p.Print()
	moves := p.GenerateMoves()
	// move white kingside
	// validMove, err := findMove(moves, "e1", "c1")
	// validMove, err := findMove(moves, "e8", "g8")

	// chop move
	// validMove, err := findMove(moves, "b7", "a8")
	validMove, err := findMove(moves, "e2", "e4")
	if err != nil {
		fmt.Println("move invalid! Here are avalable moves")
		moves.Print()
	} else {
		p.MakeMove(validMove)
		fmt.Println("prev pos empty", p.prevPosition == nil)
		p.Print()

		moves = p.GenerateMoves()
		validMove, err = findMove(moves, "f8", "g7")
		if err != nil {
			fmt.Println("move invalid! Here are avalable moves")
			moves.Print()
		} else {
			p.MakeMove(validMove)
			p.UnMakeMove()
			p.UnMakeMove()
			p.UnMakeMove()
			fmt.Println("nEw Pos")
			fmt.Println("prev pos empty", p.prevPosition == nil)
			p.Print()
		}
	}

	// fmt.Println("previous position", !reflect.DeepEqual(p.prevPosition, p))

}
