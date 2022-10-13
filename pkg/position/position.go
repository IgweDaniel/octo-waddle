package position

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/igwedaniel/dolly/pkg/attacks"
	"github.com/igwedaniel/dolly/pkg/bitboard"
	"github.com/igwedaniel/dolly/pkg/moves"
)

const (
	White int = iota
	Black
)

const (
	OccupancySq = iota
	King
	Queen
	Bishop
	Knight
	Rook
	Pawn
)
const (
	WhiteKingside  = 1
	WhiteQueenside = 2
	BlackKingside  = 4
	BlackQueenside = 8
)

type Position struct {
	bitboards      [2][]bitboard.Bitboard
	castlingRights int
	side           int
	enPassanteSq   int
	moveCount      int
	halfMoveCount  int
}

func tokenizeFenString(fen string) (string, int, int, int, int, int) {
	fenTokens := strings.Split(fen, " ")
	var (
		side         int
		castleRights int
	)

	moveCount, err := strconv.Atoi(fenTokens[4])
	if err != nil {
		log.Fatal(err)
	}
	halfMoveCount, err := strconv.Atoi(fenTokens[5])
	if err != nil {
		log.Fatal(err)
	}
	enpassantSquare, err := moves.AlgebriacToIndex(fenTokens[3])
	if err != nil {
		log.Fatal(err)
	}

	castling := fenTokens[2]
	for i := 0; i < len(castling); i++ {
		char := string(castling[i])
		switch char {
		case "K":
			castleRights |= WhiteKingside
		case "Q":
			castleRights |= WhiteQueenside
		case "k":
			castleRights |= BlackKingside
		case "q":
			castleRights |= BlackQueenside
		case "-":
		}
	}
	switch fenTokens[1] {
	case "w":
		side = White
	case "b":
		side = Black
	default:
		log.Fatal(errors.New("invalid side"))
	}

	return fenTokens[0], side, castleRights, enpassantSquare, moveCount, halfMoveCount
}

func NewFenPosition(fen string) *Position {
	p := new(Position)
	p.bitboards[0] = make([]bitboard.Bitboard, 7)
	p.bitboards[1] = make([]bitboard.Bitboard, 7)
	fenPosition, side, castlingRights, enPassanteSq, moveCount, halfMoveCount := tokenizeFenString(fen)

	p.setBitboardsFromFen(fenPosition)
	p.side = side
	p.moveCount = moveCount
	p.halfMoveCount = halfMoveCount
	p.castlingRights = castlingRights
	p.enPassanteSq = enPassanteSq
	return p
}

func (p *Position) setBitboardsFromFen(fenPosition string) {
	var sqIdx int

	for charIdx := 0; charIdx < len(fenPosition); charIdx++ {
		char := string(fenPosition[charIdx])

		switch char {
		case "/":
			sqIdx--
		case "p":
			p.bitboards[Black][Pawn].SetBit(sqIdx)
		case "P":
			p.bitboards[White][Pawn].SetBit(sqIdx)
		case "n":
			p.bitboards[Black][Knight].SetBit(sqIdx)
		case "N":
			p.bitboards[White][Knight].SetBit(sqIdx)
		case "b":
			p.bitboards[Black][Bishop].SetBit(sqIdx)
		case "B":
			p.bitboards[White][Bishop].SetBit(sqIdx)
		case "r":
			p.bitboards[Black][Rook].SetBit(sqIdx)
		case "R":
			p.bitboards[White][Rook].SetBit(sqIdx)
		case "q":
			p.bitboards[Black][Queen].SetBit(sqIdx)
		case "Q":
			p.bitboards[White][Queen].SetBit(sqIdx)
		case "k":
			p.bitboards[Black][King].SetBit(sqIdx)
		case "K":
			p.bitboards[White][King].SetBit(sqIdx)
		case "1":

		case "2":
			sqIdx++
		case "3":
			sqIdx += 2
		case "4":
			sqIdx += 3
		case "5":
			sqIdx += 4
		case "6":
			sqIdx += 5
		case "7":
			sqIdx += 6
		case "8":
			sqIdx += 7

		}
		sqIdx++
	}
	p.setOccupancy(White)
	p.setOccupancy(Black)
}

func (p *Position) setOccupancy(color int) {
	p.bitboards[color][OccupancySq] = bitboard.Bitboard(
		p.bitboards[color][Pawn] |
			p.bitboards[color][Rook] |
			p.bitboards[color][Knight] |
			p.bitboards[color][Bishop] |
			p.bitboards[color][King] |
			p.bitboards[color][Queen])
}

func (p *Position) Print() {
	chess_gyphicons := [2][]string{}
	chess_gyphicons[White] = strings.Split(",♚,♛,♝,♞,♜,♟︎", ",")
	chess_gyphicons[Black] = strings.Split(",♔,♕,♗,♘,♖,♙", ",")
	p.Occupancy().Print()

	for sq := 0; sq < 64; sq++ {

		var colorIdx, piece int

		for i := range p.bitboards {
			for idx, bitboard := range p.bitboards[i] {
				if bitboard.BitIsSet(sq) {
					piece = idx
					colorIdx = i
				}
			}
		}

		if piece == 0 {
			fmt.Print(" . ")
		} else {
			fmt.Printf(" %s ", chess_gyphicons[colorIdx][piece])
		}
		if ((sq + 1) % 8) == 0 {
			fmt.Println("")
		}
	}
	fmt.Println("castling rights", p.castlingRights)
	fmt.Println("enpassantSquare", p.enPassanteSq)
	fmt.Println("")
}

func (p *Position) Occupancy() bitboard.Bitboard {
	return p.bitboards[White][OccupancySq] | p.bitboards[Black][OccupancySq]
}

func (p *Position) RookAttacks(square int) bitboard.Bitboard {
	rookAttacks := bitboard.NewMask(0)
	// North Attacks
	rookAttacks |= attacks.NorthRay[square]
	blockers := p.Occupancy() & attacks.NorthRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.MsbIdx()
		rookAttacks &= ^attacks.NorthRay[blockerIdx]
	}
	// South Attacks
	rookAttacks |= attacks.SouthRay[square]
	blockers = p.Occupancy() & attacks.SouthRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.LsbIdx()
		rookAttacks ^= attacks.SouthRay[blockerIdx]
	}

	// West Attacks
	rookAttacks |= attacks.WestRay[square]
	blockers = p.Occupancy() & attacks.WestRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.MsbIdx()
		rookAttacks ^= attacks.WestRay[blockerIdx]
	}

	// East Attacks
	rookAttacks |= attacks.EastRay[square]
	blockers = p.Occupancy() & attacks.EastRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.LsbIdx()
		rookAttacks &= ^attacks.EastRay[blockerIdx]
	}

	return rookAttacks
}

func (p *Position) BishopAttacks(square int) bitboard.Bitboard {
	bishopAttacks := bitboard.NewMask(0)
	// North WestAttacks
	bishopAttacks |= attacks.NorthWestRay[square]
	blockers := p.Occupancy() & attacks.NorthWestRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.MsbIdx()
		bishopAttacks &= ^attacks.NorthWestRay[blockerIdx]
	}

	// North EastAttacks
	blockers = p.Occupancy() & attacks.NorthEastRay[square]
	bishopAttacks |= attacks.NorthEastRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.LsbIdx()
		bishopAttacks &= ^attacks.NorthEastRay[blockerIdx]
	}
	// South East
	blockers = p.Occupancy() & attacks.SouthEastRay[square]
	bishopAttacks |= attacks.SouthEastRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.LsbIdx()
		bishopAttacks &= ^attacks.SouthEastRay[blockerIdx]
	}

	// South West
	blockers = p.Occupancy() & attacks.SouthWestRay[square]
	bishopAttacks |= attacks.SouthWestRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.MsbIdx()
		bishopAttacks &= ^attacks.SouthWestRay[blockerIdx]
	}
	return bishopAttacks
}

func (p *Position) QueenAttacks(square int) bitboard.Bitboard {
	return p.RookAttacks(square) | p.BishopAttacks(square)
}
