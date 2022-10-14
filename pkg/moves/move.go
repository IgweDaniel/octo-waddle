package moves

import (
	"fmt"
	"strconv"
)

/*


type Move struct {
	piece       int8
	origin      int
	destination int
	// Queen = 2 Bishops = 3 Knights = 4 Rooks = 5
	promotion  int
	isCapture  bool
	enpassant  bool
	isCastling bool
}


         binary move bits                               hexidecimal constants

   0000 0000 0000 0000 0000 0011 1111    origin square       0x3f
   0000 0000 0000 0000 1111 1100 0000    dest square       0xfc0
   0000 0000 0000 1111 0000 0000 0000    attackpiece               0xf000
   0000 0000 1111 0000 0000 0000 0000    promotedpiece      0xf0000
   0000 0001 0000 0000 0000 0000 0000    capture flag        0x100000
   0000 0010 0000 0000 0000 0000 0000    enpassant flag    0x200000
   0000 0100 0000 0000 0000 0000 0000    castling flag      0x400000


*/

const (
	originSqMask      = 0x3f
	destMaskSqMask    = 0xfc0
	attackPieceMask   = 0xf000
	promotedPieceMask = 0xf0000
	captureflagMask   = 0x100000
	enpassantflagMask = 0x200000
	castlingflagMask  = 0x400000
)

type Move uint64

func boolToInt(bflag bool) int {

	if bflag {
		return 1
	} else {
		return 0
	}
}

type encodeMoveOpt struct {
	origin      int
	dest        int
	attackPiece int
	// Queen = 2 Bishops = 3 Knights = 4 Rooks = 5
	promotedPiece int
	isCapture     bool
	enpassant     bool
	isCastling    bool
}

func encodeMove(opts encodeMoveOpt) Move {
	return Move(opts.origin | opts.dest<<6 | opts.attackPiece<<12 | opts.promotedPiece<<16 | boolToInt(opts.isCapture)<<20 | boolToInt(opts.enpassant)<<21 | boolToInt(opts.isCastling)<<22)
}

func (mv Move) Dest() int {
	return int((mv & destMaskSqMask) >> 6)
}
func (mv Move) Origin() int {
	return int(mv & originSqMask)
}
func (mv Move) Piece() int {
	return int((mv & attackPieceMask) >> 12)
}
func (mv Move) PromotedPiece() int {
	return int((mv & promotedPieceMask) >> 16)
}
func (mv Move) IsCapture() bool {
	return mv&captureflagMask != 0
}
func (mv Move) IsCastling() bool {

	return mv&castlingflagMask != 0
}
func (mv Move) Enpassant() bool {
	return mv&enpassantflagMask != 0
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
