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
		fmt.Printf("%s bitboard at %v", id, idx)
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

func TestPositionParseFen(t *testing.T) {
	/* 	var p *Position
	   	p = NewFenPosition("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	   	p.Print()

	   	p = NewFenPosition("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2")
	   	p.Print()
	   	p = NewFenPosition("rnbqkbnr/pp5p/2pppp2/6p1/2B1P3/2N2P1N/PPPP2PP/R1BQK1R1 b Qkq - 1 6")
	   	// PrintAllBoards(PrintOption{position: *p, piece: Pawn})
	   	fmt.Println("castling rights", p.castlingRights)
	   	fmt.Println("enpassantSquare", p.enPassanteSq)
	   	// fmt.Println("pieceMap", p.pieceMap)
	   	p.Print()
	*/
	// fmt.Println("chess_gyphicons", chess_gyphicons)

	/* b := bitboard.Bitboard(0)
	pos, _ := moves.AlgebriacToIndex("c6")
	b.SetBit(pos)
	// b.SetBit(63)
	b.Print() */
	// return

	// printBitboards("white pawn", attacks.LookupTable.PawnAttacks[White])
	// printBitboards("black pawn attcks", attacks.LookupTable.PawnAttacks[Black])
	// printAttackBitboards("knight attcks", attacks.LookupTable.KnightAttacks)
	printAttackBitboards("knight attacks", attacks.LookupTable.KingAttacks)
	fmt.Println("End")
	// attacks.LookupTable.KnightAttacks[63].Print()
}
