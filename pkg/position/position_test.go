package position

import (
	"fmt"
	"testing"

	"github.com/igwedaniel/dolly/pkg/attacks"
	"github.com/igwedaniel/dolly/pkg/bitboard"
)

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
		colorMap := map[int]string{White: "white", Black: "Black"}
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

func getRookAttacksAtIndex(square int) bitboard.Bitboard {
	return attacks.NorthRay[square] |
		attacks.SouthRay[square] |
		attacks.EastRay[square] |
		attacks.WestRay[square]
}
func getBishopAttacksAtIndex(square int) bitboard.Bitboard {
	return attacks.NorthWestRay[square] |
		attacks.SouthEastRay[square] |
		attacks.NorthEastRay[square] |
		attacks.SouthWestRay[square]
}

func getQueenAttacks(square int) bitboard.Bitboard {
	return getBishopAttacksAtIndex(square) | getRookAttacksAtIndex(square)
}

func TestPositionParseFen(t *testing.T) {
	var p *Position
	p = NewFenPosition("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	p.Print()
	p = NewFenPosition("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2")
	p.Print()
	p = NewFenPosition("rnbqkbnr/pp5p/2pppp2/6p1/2B1P3/2N2P1N/PPPP2PP/R1BQK1R1 b Qkq - 1 6")
	p.Print()
	// p.RookAttacks(0).Print()
	// PrintAllBoards(PrintOption{position: *p, piece: Pawn})

	// printAttackBitboards("knight attcks", attacks.Knights)
	// printAttackBitboards("knight attacks", attacks.Kings)
	// printAttackBitboards("black pawn attcks", attacks.Pawns[Black])
	// printAttackBitboards("white pawn", attacks.Pawns[White])
	// printAttackBitboards("southRay", attacks.SouthRay)
	// printAttackBitboards("northRay", attacks.NorthRay)

	// printAttackBitboards("southRay", attacks.SouthRay)
	// printAttackBitboards("eastArray", attacks.EastRay)
	// printAttackBitboards("westArray", attacks.WestRay)
	// printAttackBitboards("northEastArray", attacks.NorthEastRay)
	// printAttackBitboards("northWestArray", attacks.NorthWestRay)
	// printAttackBitboards("southWestArray", attacks.SouthWestRay)
	// printAttackBitboards("southEastArray", attacks.SouthEastRay)

	square := 35
	// fmt.Printf("attack at %v", square)
	sqBb := bitboard.New(square)
	sqBb.SetBit(53)
	// sqBb = bitboard.NewMask(0)
	sqBb.Print()
	fmt.Println("trailing zeros count: ", sqBb.LsbIdx())
	fmt.Println("leading zeros count: ", sqBb.MsbIdx())

	// attacks := getQueenAttacks(square)
	// attacks.Print()

	fmt.Println()

}
