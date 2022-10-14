package position

import (
	"testing"

	"github.com/igwedaniel/dolly/pkg/moves"
)

func TestPositionParseFen(t *testing.T) {

	tests := []struct {
		fen            string
		halfMoveCount  int
		moveCount      int
		turn           int
		castlingRights int
		enPassanteSq   int
	}{
		{"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", 0, 1, White, 15, 64},
		{"rnbqkbnr/pppppppp/8/4R3/8/8/PPPPPPPP/1NBQKBNR w KQkq - 0 1", 0, 1, White, 15, 64},
		{"rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1", 0, 1, White, 15, 20},
		{"rnbqkbnr/pp5p/2pppp2/6p1/2B1P3/2N2P1N/PPPP2PP/R1BQK1R1 b Qkq - 1 6", 1, 6, Black, 14, 64},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			p := NewFenPosition(tt.fen)

			if p.halfMoveCount != tt.halfMoveCount {
				t.Errorf("incorrect halfmove count: expected %v, got %v", tt.halfMoveCount, p.halfMoveCount)
			}

			if p.moveCount != tt.moveCount {
				t.Errorf("incorrect move count: expected %v, got %v", tt.moveCount, p.moveCount)
			}
			if p.enPassanteSq != tt.enPassanteSq {
				t.Errorf("incorrect enpassante square: expected %v, got %v", tt.enPassanteSq, p.enPassanteSq)
			}
			if p.side != tt.turn {
				t.Errorf("incorrect side to move: expected %v, got %v", colorMap[tt.turn], colorMap[p.side])
			}
			if p.castlingRights != tt.castlingRights {
				t.Errorf("incorrect side to move: expected %v, got %v", tt.castlingRights, p.castlingRights)
			}

		})
	}

}

func TestPositionIsAttacked(t *testing.T) {
	tests := []struct {
		square     string
		side       int
		isAttacked bool
	}{
		{"a1", White, false},
		{"d7", White, true},
		{"e5", Black, false},
		{"g2", Black, true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			p := NewFenPosition("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1")
			square, _ := moves.AlgebriacToIndex(tt.square)
			isAttacked := p.IsSquareAttackedBy(square, tt.side)

			if isAttacked != tt.isAttacked {
				t.Errorf("incorrect attack status for: %s expected %v, got %v", tt.square, tt.isAttacked, isAttacked)
			}

		})
	}

}

func TestMoveGeneration(t *testing.T) {
	// rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1
	// p := NewFenPosition("8/8/8/2p1pP2/1P1P4/8/8/8 w KQkq e6 0 1")

	/* // tricky position
	p := NewFenPosition("rnbqkbnr/pppp1p1p/8/4p2P/4PPp1/8/PPPP2P1/RNBQKBNR b KQkq f3 0 4") */

	// killer position
	p := NewFenPosition("rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1")
	// p := NewFenPosition("8/8/8/2p1pP2/1P1P4/8/8/8 w KQkq e6 0 1")
	// p := NewFenPosition("rnbqkbnr/ppppp1pp/8/5p2/8/3P4/PPP1PPPP/RNBQKBNR w KQkq - 0 1")
	p.GenerateMoves()
	p.Print()
}
