package position

import (
	"fmt"
	"testing"

	"github.com/igwedaniel/dolly/pkg/bitboard"
)

var colorMap = map[int]string{White: "white", Black: "Black"}

type PrintOption struct {
	position Position
	piece    int
}

func printAttackBitboards(id string, bitboards [64]bitboard.Bitboard) {
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

const (
	fileA = bitboard.Bitboard(0x0101010101010101)
	fileB = bitboard.Bitboard(0x0202020202020202)
	fileG = bitboard.Bitboard(0x4040404040404040)
	fileH = bitboard.Bitboard(0x8080808080808080)
)

func PrintPositionRookAttacks(p Position) {
	for color := White; color <= Black; color++ {
		rooks := p.bitboards[color][Rook]
		rooks.Print()
		fmt.Printf("rooks default\n")
		for !rooks.IsEmpty() {

			rookIdx := rooks.LsbIdx()

			p.RookAttacks(rookIdx).Print()
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

			p.BishopAttacks(bishopIdx).Print()
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

			p.QueenAttacks(queenIdx).Print()
			fmt.Printf("%s queen  attacks on %v\n", colorMap[color], queenIdx)
			queen.RemoveBit(queenIdx)
		}

	}

}

func TestPositionParseFen(t *testing.T) {
	var p *Position
	/* 	p = NewFenPosition("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	   	PrintPositionRookAttacks(*p)
	   	p.Print()
	   	p = NewFenPosition("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2")
	   	p.Print()
	   	PrintPositionRookAttacks(*p)
	   	p = NewFenPosition("rnbqkbnr/pp5p/2pppp2/6p1/2B1P3/2N2P1N/PPPP2PP/R1BQK1R1 b Qkq - 1 6")
	   	p.Print()
	   	PrintPositionRookAttacks(*p) */

	// p = NewFenPosition("rnbqkbnr/pp5p/2pppp2/6p1/2B1P3/2N2P1N/PPPP2PP/R1BQK1R1 b Qkq - 1 6")
	p = NewFenPosition("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1")
	// PrintPositionQueenAttacks(*p)
	p = NewFenPosition("rnbqkbnr/pppppppp/8/4R3/8/8/PPPPPPPP/1NBQKBNR w KQkq - 0 1")
	// PrintPositionQueenAttacks(*p)
	p = NewFenPosition("rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1")
	p.Print()

	fmt.Println()

}
