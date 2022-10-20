package position

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/igwedaniel/dolly/pkg/moves"
)

func onlyLegalMoves(p Position) moves.Moves {

	legalmoves := moves.NewList()

	for _, move := range p.GenerateMoves() {

		p.MakeMove(move)

		kingSqBB := p.bitboards[p.side^1][King]
		kingSqIdx := kingSqBB.LsbIdx()

		if !p.IsSquareAttackedBy(kingSqIdx, p.side) {
			legalmoves = append(legalmoves, move)
		}
		p.UnMakeMove()
	}
	return legalmoves
}

func PerftLegalWithPrint(position *Position, depth int, parentMv moves.Move, pstats *PositionStats) int {
	nodes := 0
	legalmoves := onlyLegalMoves(*position)

	if len(legalmoves) == 0 {
		pstats.checkmates += 1
	}
	if depth == 0 {
		return 1
	}

	for _, move := range legalmoves {
		countStats(move, pstats)
		position.MakeMove(move)
		nodes += PerftLegalWithPrint(position, depth-1, move, pstats)
		position.UnMakeMove()

	}

	return nodes
}

func countStats(move moves.Move, pstats *PositionStats) {
	if move.IsCastling() {

		pstats.castles += 1
	}
	if move.Enpassant() {

		pstats.epCaptures += 1
	}
	if move.IsCapture() {

		pstats.captures += 1
	}

	if move.IsPromotion() {

		pstats.promotions += 1
	}
}

func PerftRoot(p *Position, depth int) PositionStats {
	legalmoves := onlyLegalMoves(*p)
	var pstats PositionStats
	for _, move := range legalmoves {
		p.MakeMove(move)
		nodes := PerftLegalWithPrint(p, depth-1, move, &pstats)
		pstats.totalnodes += nodes
		countStats(move, &pstats)
		p.UnMakeMove()
	}

	return pstats

}

type PositionStats struct {
	totalnodes int
	captures   int
	epCaptures int
	castles    int
	promotions int
	checkmates int
}

func TestMoveGeneration(t *testing.T) {
	type test struct {
		name    string
		fen     string
		depth   int
		results PositionStats
	}
	tests := []test{
		{
			"Position 4",
			"r2q1rk1/pP1p2pp/Q4n2/bbp1p3/Np6/1B3NBn/pPPP1PPP/R3K2R b KQ - 0 1 ",
			3,
			PositionStats{9467, 1108, 4, 6, 168, 22},
		},
		{
			"Kiwipete ",
			"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 0",
			3,
			PositionStats{97862, 17461, 46, 3255, 0, 1},
		},
		{
			"Position 3",
			"8/2p5/3p4/KP5r/1R3p1k/8/4P1P1/8 w - - 0 0",
			4,
			PositionStats{43238, 3572, 125, 0, 0, 17},
		},
		{
			"starting Position",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			4,
			PositionStats{197281, 1610, 0, 0, 0, 8},
		},
	}

	for _, tt := range tests {

		func(tt test) {

			t.Run(fmt.Sprintf("%s at depth %d", tt.name, tt.depth), func(t *testing.T) {
				t.Parallel()
				p := NewFenPosition(tt.fen)
				results := PerftRoot(p, tt.depth)
				if !reflect.DeepEqual(tt.results, results) {
					t.Errorf("incorrect node count at depth %v  expected %v, got %v", tt.depth, tt.results, results)
				}

			})
		}(tt)
	}
}
