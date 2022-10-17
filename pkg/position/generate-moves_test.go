package position

import (
	"fmt"
	"testing"

	"github.com/igwedaniel/dolly/pkg/moves"
)

func TestMoveGeneration(t *testing.T) {

	// p := NewFenPosition("8/8/8/2p1pP2/1P1P4/8/8/8 w KQkq e6 0 1")

	// tricky position
	p := NewFenPosition("rnbqkbnr/pppp1p1p/8/4p2P/4PPp1/8/PPPP2P1/RNBQKBNR b KQkq f3 0 4")

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

	// fen := "r1b1kbnr/pppp1ppp/8/8/1n2q3/8/PPPPBPPP/RNBQK1NR b KQkq - 0 1"
	// // double check
	// fen = "r1b1kbr1/pppp3p/5p2/5Nn1/2B5/8/PPPP1P1P/RNBQ1RK1 b Qkq - 0 1"
	// fen = "rnb1kbnr/pppp1ppp/8/8/4q3/8/PPPPBPPP/RNBQK1NR w KQkq - 0 1"
	// p = NewFenPosition("rnbq1rk1/pp1p1pPp/3b4/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR b KQ e6 0 1")
	p = NewFenPosition("8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 1")
	fmt.Println("old Pos")
	p.Print()
	// fmt.Println(p.prevPosition)

	fmt.Println("======legal moves======")
	legalmoves := onlyLegalMoves(*p)
	legalmoves.Print()
	fmt.Println("======legal moves======")

	fmt.Println("======pseudo legal moves======")
	moves := p.GenerateMoves()
	moves.Print()
	fmt.Println("======pseudo legal moves======")

	// fmt.Println(p.prevPosition)
	// move white kingside
	// validMove, err := findMove(moves, "e1", "c1")
	// validMove, err := findMove(moves, "e8", "g8")

	// chop move
	// validMove, err := findMove(moves, "b7", "a8")
	validMove, err := findMove(moves, "a5", "b6")
	// validMove, err := findMove(moves, "e2", "f3")

	if err != nil {
		fmt.Println("move invalid! Here are avalable moves")
		// moves.Print()
	} else {
		// fmt.Println("before")
		// p.getOccupancy().Print()

		p.MakeMove(validMove)
		kingSqBB := p.bitboards[p.side^1][King]
		kingSqIdx := kingSqBB.LsbIdx()

		if p.IsSquareAttackedBy(kingSqIdx, p.side) {
			fmt.Printf("King is grave in danger by attack from %s undoing move \n", colorMap[p.side])
		}

		p.UnMakeMove()
		p.Print()
		// fmt.Println("after")
		// p.getOccupancy().Print()
		// kingSqBB := p.bitboards[p.side^1][King]
		// kingSqIdx := kingSqBB.LsbIdx()
		// kingSqBB.Print()

		// fmt.Printf("curr side %s opp side %s\n", colorMap[p.side], colorMap[p.side^1])

		// fmt.Println(kingSqIdx)
		// fmt.Println("is attacked", p.IsSquareAttackedBy(kingSqIdx, p.side))

		// if p.IsSquareAttackedBy(kingSqIdx, p.side) {
		// 	p.UnMakeMove()
		// 	fmt.Printf("King is grave in danger by attack from %s undoing move \n", colorMap[p.side])
		// }
		// fmt.Println("prev pos empty", p.prevPosition == nil)
		// p.Print()

		// moves = p.GenerateMoves()
		// validMove, err = findMove(moves, "f8", "g7")
		// if err != nil {
		// 	fmt.Println("move invalid! Here are avalable moves")
		// 	moves.Print()
		// } else {
		// 	p.MakeMove(validMove)
		// 	p.UnMakeMove()
		// 	p.UnMakeMove()
		// 	p.UnMakeMove()
		// 	fmt.Println("nEw Pos")
		// 	fmt.Println("prev pos empty", p.prevPosition == nil)
		// 	p.Print()
		// }
	}

	// fmt.Println("previous position", !reflect.DeepEqual(p.prevPosition, p))

}

func onlyLegalMoves(p Position) moves.Moves {
	var pieceMaps = map[int]string{1: "King",
		2: "Queen",
		3: "Bishop",
		4: "Knight",
		5: "Rook",
		6: "Pawn",
	}
	legalmoves := moves.NewList()

	// fmt.Println("only legal", &p.prevPosition)
	for _, move := range p.GenerateMoves() {

		p.MakeMove(move)

		kingSqBB := p.bitboards[p.side^1][King]
		kingSqIdx := kingSqBB.LsbIdx()

		if !p.IsSquareAttackedBy(kingSqIdx, p.side) {
			legalmoves = append(legalmoves, move)
		} else {
			fmt.Printf("can't move %s %s at %s to %s\n", colorMap[p.side^1], pieceMaps[move.Piece()], moves.IndexToAlgebraic(move.Origin()), moves.IndexToAlgebraic(move.Dest()))
		}
		p.UnMakeMove()
	}
	return legalmoves
}

/*


	u64 Perft(int depth)
		{
		MOVE move_list[256];
		int n_moves, i;
		u64 nodes = 0;

		if (depth == 0)
			return 1ULL;

		n_moves = GenerateMoves(move_list);
		for (i = 0; i < n_moves; i++) {
			MakeMove(move_list[i]);
			if (!IsIncheck())
			nodes += Perft(depth - 1);
			UndoMove(move_list[i]);
		}
		return nodes;

*/
//

// p := NewFenPosition("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")/

var captures = 0

var checks = 0
var epCap = 0
var castles = 0
var prom = 0

// func Perft(position *Position, depth int) int {

// 	nodes := 0
// 	legalmoves := onlyLegalMoves(*position)
// 	if depth == 0 {
// 		return 1
// 	}

// 	for _, move := range legalmoves {
// 		if move.Enpassant() {
// 			epCap += 1
// 		}
// 		if move.IsCapture() {
// 			captures += 1
// 		}
// 		if move.IsCastling() {
// 			castles += 1
// 		}
// 		position.MakeMove(move)
// 		nodes += Perft(position, depth-1)
// 		position.UnMakeMove()

// 	}
// 	return nodes
// }

func Perft(position *Position, depth int) int {

	nodes := 0
	legalmoves := position.GenerateMoves()
	if depth == 0 {
		return 1
	}

	for _, move := range legalmoves {

		position.MakeMove(move)
		kingSqBB := position.bitboards[position.side^1][King]
		kingSqIdx := kingSqBB.LsbIdx()

		if !position.IsSquareAttackedBy(kingSqIdx, position.side) {
			if move.Enpassant() {
				epCap += 1
			}
			if move.IsCapture() {
				captures += 1
			}
			if move.IsCastling() {
				castles += 1
			}
			if move.IsPromotion() {
				prom += 1
			}
			nodes += Perft(position, depth-1)
		} else {
			checks += 1
		}
		position.UnMakeMove()

	}
	return nodes
}

/*
	8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - -  wrong at depth 6
	r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1  wrong at depth 4
	r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1	wrong at depth 3
*/

func TestMovePerft(t *testing.T) {

	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	fen = "r3k2r/Pppp1ppp/1b3nbN/nP6/BBP1P3/q4N2/Pp1P2PP/R2Q1RK1 w kq - 0 1"
	p := NewFenPosition(fen)
	depth := 3
	fmt.Println("===================================================")
	p.Print()
	nodes := Perft(p, depth)
	fmt.Printf("Nodes for %s at depth: %d is %d nodes %d captures %d enpassant captures %d castles %d promotions \n", fen, depth, nodes, captures, epCap, castles, prom)
	// fmt.Printf("No of Nodes for %s at depth: %d is %d nodes  %d captures %d checks and %d enpassante\n", fen, depth, nodes, captures, checks, epCap)

	fmt.Println("====================================================")
}

// tests := []struct {
// 	fen   string
// 	depth int
// 	nodes int
// }{

// 	{
// 		depth: 1,
// 		nodes: 8,
// 		fen:   "r6r/1b2k1bq/8/8/7B/8/8/R3K2R b KQ - 3 2",
// 	},
// 	{
// 		depth: 1,
// 		nodes: 8,
// 		fen:   "8/8/8/2k5/2pP4/8/B7/4K3 b - d3 0 3",
// 	},
// 	{
// 		depth: 1,
// 		nodes: 19,
// 		fen:   "r1bqkbnr/pppppppp/n7/8/8/P7/1PPPPPPP/RNBQKBNR w KQkq - 2 2",
// 	},
// 	{
// 		depth: 1,
// 		nodes: 5,
// 		fen:   "r3k2r/p1pp1pb1/bn2Qnp1/2qPN3/1p2P3/2N5/PPPBBPPP/R3K2R b KQkq - 3 2",
// 	},
// 	{
// 		depth: 1,
// 		nodes: 44,
// 		fen:   "2kr3r/p1ppqpb1/bn2Qnp1/3PN3/1p2P3/2N5/PPPBBPPP/R3K2R b KQ - 3 2",
// 	},
// 	{
// 		depth: 1,
// 		nodes: 39,
// 		fen:   "rnb2k1r/pp1Pbppp/2p5/q7/2B5/8/PPPQNnPP/RNB1K2R w KQ - 3 9",
// 	},
// 	{
// 		depth: 1,
// 		nodes: 9,
// 		fen:   "2r5/3pk3/8/2P5/8/2K5/8/8 w - - 5 4",
// 	},
// 	{
// 		depth: 3,
// 		nodes: 62379,
// 		fen:   "rnbq1k1r/pp1Pbppp/2p5/8/2B5/8/PPP1NnPP/RNBQK2R w KQ - 1 8",
// 	},
// 	{
// 		depth: 3,
// 		nodes: 89890,
// 		fen:   "r4rk1/1pp1qppp/p1np1n2/2b1p1B1/2B1P1b1/P1NP1N2/1PP1QPPP/R4RK1 w - - 0 10",
// 	},
// 	{
// 		depth: 6,
// 		nodes: 1134888,
// 		fen:   "3k4/3p4/8/K1P4r/8/8/8/8 b - - 0 1",
// 	},
// 	{
// 		depth: 6,
// 		nodes: 1015133,
// 		fen:   "8/8/4k3/8/2p5/8/B2P2K1/8 w - - 0 1",
// 	},
// 	{
// 		depth: 6,
// 		nodes: 1440467,
// 		fen:   "8/8/1k6/2b5/2pP4/8/5K2/8 b - d3 0 1",
// 	},
// 	{
// 		depth: 6,
// 		nodes: 661072,
// 		fen:   "5k2/8/8/8/8/8/8/4K2R w K - 0 1",
// 	},
// 	{
// 		depth: 6,
// 		nodes: 803711,
// 		fen:   "3k4/8/8/8/8/8/8/R3K3 w Q - 0 1",
// 	},
// 	{
// 		depth: 4,
// 		nodes: 1274206,
// 		fen:   "r3k2r/1b4bq/8/8/8/8/7B/R3K2R w KQkq - 0 1",
// 	},
// 	{
// 		depth: 4,
// 		nodes: 1720476,
// 		fen:   "r3k2r/8/3Q4/8/8/5q2/8/R3K2R b KQkq - 0 1",
// 	},
// 	{
// 		depth: 6,
// 		nodes: 3821001,
// 		fen:   "2K2r2/4P3/8/8/8/8/8/3k4 w - - 0 1",
// 	},
// 	{
// 		depth: 5,
// 		nodes: 1004658,
// 		fen:   "8/8/1P2K3/8/2n5/1q6/8/5k2 b - - 0 1",
// 	},
// 	{
// 		depth: 6,
// 		nodes: 217342,
// 		fen:   "4k3/1P6/8/8/8/8/K7/8 w - - 0 1",
// 	},
// 	{
// 		depth: 6,
// 		nodes: 92683,
// 		fen:   "8/P1k5/K7/8/8/8/8/8 w - - 0 1",
// 	},
// 	{
// 		depth: 6,
// 		nodes: 2217,
// 		fen:   "K1k5/8/P7/8/8/8/8/8 w - - 0 1",
// 	},
// 	{
// 		depth: 7,
// 		nodes: 567584,
// 		fen:   "8/k1P5/8/1K6/8/8/8/8 w - - 0 1",
// 	},
// 	{
// 		depth: 4,
// 		nodes: 23527,
// 		fen:   "8/8/2k5/5q2/5n2/8/5K2/8 b - - 0 1",
// 	},
// }

// for _, tt := range tests {
// 	t.Run("", func(t *testing.T) {
// 		p := NewFenPosition(tt.fen)
// 		nodes := Perft(p, tt.depth)
// 		if tt.nodes != nodes {
// 			t.Errorf("incorrect node count at %v for: %s expected %v, got %v", tt.depth, tt.fen, tt.nodes, nodes)
// 		}

// 	})
// }
// }
