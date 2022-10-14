package position

import (
	"fmt"
	"testing"

	"github.com/igwedaniel/dolly/pkg/bitboard"
	"github.com/igwedaniel/dolly/pkg/moves"
	"github.com/stretchr/testify/assert"
)

var colorMap = map[int]string{White: "white", Black: "Black"}

type PrintOption struct {
	position Position
	piece    int
}

func PrintAttackBitboards(id string, bitboards [64]bitboard.Bitboard) {
	fmt.Println(id)

	for idx, bitB := range bitboards {
		bitB.Print()
		idxBB := bitboard.Bitboard(0)
		idxBB.SetBit(idx)
		idxBB.Print()
		fmt.Printf("%s bitboard at %v\n", id, idx)
	}
}

func PrintAllBoards(opts PrintOption) {
	for colorIndex, color := range opts.position.bitboards {

		fmt.Printf("color: %s bitboards\n", colorMap[colorIndex])
		for piece, bitboards := range color {
			if piece == opts.piece {
				bitboards.Print()
			} else if piece == -1 {
				bitboards.Print()
			}
		}
	}
}

func PrintPositionRookAttacks(p Position) {
	for color := White; color <= Black; color++ {
		rooks := p.bitboards[color][Rook]
		rooks.Print()
		fmt.Printf("rooks default\n")
		for !rooks.IsEmpty() {

			rookIdx := rooks.LsbIdx()

			p.getRookAttacks(rookIdx).Print()
			fmt.Printf("%s rooks  attacks on %v\n", colorMap[color], rookIdx)
			rooks.RemoveBit(rookIdx)
		}

	}

}
func PrintPositionBishopAttacks(p Position) {
	for color := White; color <= Black; color++ {
		bishops := p.bitboards[color][Bishop]
		bishops.Print()
		fmt.Printf("bishops default\n")
		for !bishops.IsEmpty() {

			bishopIdx := bishops.LsbIdx()

			p.getBishopAttacks(bishopIdx).Print()
			fmt.Printf("%s bishops  attacks on %v\n", colorMap[color], bishopIdx)
			bishops.RemoveBit(bishopIdx)
		}

	}

}
func PrintPositionQueenAttacks(p Position) {
	for color := White; color <= Black; color++ {
		queen := p.bitboards[color][Queen]
		queen.Print()
		fmt.Printf("%s queen default\n", colorMap[color])
		for !queen.IsEmpty() {

			queenIdx := queen.LsbIdx()

			p.getQueenAttacks(queenIdx).Print()
			fmt.Printf("%s queen  attacks on %v\n", colorMap[color], queenIdx)
			queen.RemoveBit(queenIdx)
		}

	}

}

func TestPositionParseFen(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		fen            string
		halfMoveCount  int
		moveCount      int
		turn           int
		castlingRights int
		enpassantSq    int
	}{
		{"r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1", 0, 1, White, 15, 64},
		{"rnbqkbnr/pppppppp/8/4R3/8/8/PPPPPPPP/1NBQKBNR w KQkq - 0 1", 0, 1, White, 15, 64},
		{"rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1", 0, 1, White, 15, 20},
		{"rnbqkbnr/pp5p/2pppp2/6p1/2B1P3/2N2P1N/PPPP2PP/R1BQK1R1 b Qkq - 1 6", 1, 6, Black, 14, 64},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			p := NewFenPosition(tt.fen)
			assert.Equal(p.halfMoveCount, tt.halfMoveCount, "halfmove count not valid")
			assert.Equal(p.moveCount, tt.moveCount, "move count not valid")
			assert.Equal(p.enPassanteSq, tt.enpassantSq, "enpassante square not valid")
			assert.Equal(p.side, tt.turn, "side to move invalid")
			assert.Equal(p.castlingRights, tt.castlingRights, "castling rights invalid")

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
			assert.Equal(t, p.IsSquareAttackedBy(square, tt.side), tt.isAttacked, "")

		})
	}

}
