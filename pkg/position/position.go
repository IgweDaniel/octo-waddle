package position

import (
	"errors"
	"fmt"
	"log"
	"math"
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
	// remeber to flip the ranks according to ur visual board printout
	rank8 = bitboard.Bitboard(0x00000000000000FF)
	rank7 = bitboard.Bitboard(0x000000000000FF00)
	rank5 = bitboard.Bitboard(0x00000000FF000000)
	rank4 = bitboard.Bitboard(0x000000FF00000000)
	rank2 = bitboard.Bitboard(0x00FF000000000000)
	rank1 = bitboard.Bitboard(0xFF00000000000000)
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
	bitboards      [2][7]bitboard.Bitboard
	castlingRights int
	side           int
	enPassanteSq   int
	moveCount      int
	halfMoveCount  int
	prevPosition   *Position
}

func tokenizeFenString(fen string) (string, int, int, int, int, int) {
	fenTokens := strings.Split(fen, " ")
	var (
		side         int
		castleRights int
	)

	moveCount, err := strconv.Atoi(fenTokens[5])
	if err != nil {
		log.Fatal(err)
	}
	halfMoveCount, err := strconv.Atoi(fenTokens[4])
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
	// p.bitboards[0] = make([]bitboard.Bitboard, 7)
	// p.bitboards[1] = make([]bitboard.Bitboard, 7)
	fenPosition, side, castlingRights, enPassanteSq, moveCount, halfMoveCount := tokenizeFenString(fen)

	p.setBitboardsFromFen(fenPosition)
	p.side = side
	p.moveCount = moveCount
	p.halfMoveCount = halfMoveCount
	p.castlingRights = castlingRights
	p.enPassanteSq = enPassanteSq
	return p
}

func (p Position) copy() Position {
	pCopied := p

	if p.prevPosition != nil {
		prevPos := *p.prevPosition
		pCopied.prevPosition = &prevPos

	}

	return pCopied
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

func (p *Position) getOccupancy() bitboard.Bitboard {
	return p.bitboards[White][OccupancySq] | p.bitboards[Black][OccupancySq]
}

func (p *Position) getRookAttacks(square int) bitboard.Bitboard {
	rookAttacks := bitboard.NewMask(0)
	// North Attacks
	rookAttacks |= attacks.NorthRay[square]
	blockers := p.getOccupancy() & attacks.NorthRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.MsbIdx()
		rookAttacks &= ^attacks.NorthRay[blockerIdx]
	}
	// South Attacks
	rookAttacks |= attacks.SouthRay[square]
	blockers = p.getOccupancy() & attacks.SouthRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.LsbIdx()
		rookAttacks ^= attacks.SouthRay[blockerIdx]
	}

	// West Attacks
	rookAttacks |= attacks.WestRay[square]
	blockers = p.getOccupancy() & attacks.WestRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.MsbIdx()
		rookAttacks ^= attacks.WestRay[blockerIdx]
	}

	// East Attacks
	rookAttacks |= attacks.EastRay[square]
	blockers = p.getOccupancy() & attacks.EastRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.LsbIdx()
		rookAttacks &= ^attacks.EastRay[blockerIdx]
	}

	return rookAttacks
}

func (p *Position) getBishopAttacks(square int) bitboard.Bitboard {
	bishopAttacks := bitboard.NewMask(0)
	// North WestAttacks
	bishopAttacks |= attacks.NorthWestRay[square]
	blockers := p.getOccupancy() & attacks.NorthWestRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.MsbIdx()
		bishopAttacks &= ^attacks.NorthWestRay[blockerIdx]
	}

	// North EastAttacks
	blockers = p.getOccupancy() & attacks.NorthEastRay[square]
	bishopAttacks |= attacks.NorthEastRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.MsbIdx()
		bishopAttacks &= ^attacks.NorthEastRay[blockerIdx]
	}
	// South East
	blockers = p.getOccupancy() & attacks.SouthEastRay[square]
	bishopAttacks |= attacks.SouthEastRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.LsbIdx()
		bishopAttacks &= ^attacks.SouthEastRay[blockerIdx]
	}

	// South West
	blockers = p.getOccupancy() & attacks.SouthWestRay[square]
	bishopAttacks |= attacks.SouthWestRay[square]
	if !blockers.IsEmpty() {
		blockerIdx := blockers.LsbIdx()
		bishopAttacks &= ^attacks.SouthWestRay[blockerIdx]
	}
	return bishopAttacks
}

func (p *Position) getQueenAttacks(square int) bitboard.Bitboard {
	return p.getRookAttacks(square) | p.getBishopAttacks(square)
}

func (p *Position) GetFen() string {
	var board [64]string
	for index := range board {
		for piece, pieceBB := range p.bitboards[White] {
			if pieceBB.BitIsSet(index) {
				switch piece {
				case Pawn:
					board[index] = "P"
				case Knight:
					board[index] = "N"
				case King:
					board[index] = "K"
				case Queen:
					board[index] = "Q"
				case Rook:
					board[index] = "R"
				case Bishop:
					board[index] = "B"
				}
			}
		}
		for piece, pieceBB := range p.bitboards[Black] {
			if pieceBB.BitIsSet(index) {
				switch piece {
				case Pawn:
					board[index] = "p"
				case Knight:
					board[index] = "n"
				case King:
					board[index] = "k"
				case Queen:
					board[index] = "q"
				case Rook:
					board[index] = "r"
				case Bishop:
					board[index] = "b"
				}
			}
		}
		// }
	}

	fstr := ""

	for rank := 0; rank < 8; rank++ {
		var empty = 0
		for file := 0; file < 8; file++ {

			// fmt.Println(empty)
			square := rank*8 + file

			// fmt.Println(board[9])
			if board[square] != "" {
				if empty > 0 {
					fstr += strconv.Itoa(empty)
					empty = 0
				}
				fstr += board[square]
			} else {
				empty += 1
			}

		}
		if empty > 0 {
			fstr += strconv.Itoa(empty)
		}
		if rank != 7 {
			fstr += "/"

		}
	}
	fstr += " "
	switch p.side {
	case White:
		fstr += "w"
	case Black:
		fstr += "b"
	}

	fstr += " "

	if p.castlingRights != 0 {
		if p.castlingRights&WhiteKingside == WhiteKingside {
			fstr += "K"
		}
		if p.castlingRights&WhiteQueenside == WhiteQueenside {
			fstr += "Q"
		}
		if p.castlingRights&BlackKingside == BlackKingside {
			fstr += "k"
		}
		if p.castlingRights&BlackQueenside == BlackQueenside {
			fstr += "q"
		}
	} else {
		fstr += "-"
	}

	fstr += " "

	switch p.enPassanteSq {
	case 64:
		fstr += "-"
	default:
		fstr += moves.IndexToAlgebraic(p.enPassanteSq)

	}
	fstr += " "
	fstr += strconv.Itoa(p.halfMoveCount)

	fstr += " "
	fstr += strconv.Itoa(p.moveCount)

	return fstr
}
func (p *Position) IsSquareAttackedBy(square, side int) bool {

	pawnAttacks := attacks.Pawns[side^1][square] & p.bitboards[side][Pawn]

	if !pawnAttacks.IsEmpty() {
		return true
	}
	knightAttacks := attacks.Knights[square] & p.bitboards[side][Knight]
	if !knightAttacks.IsEmpty() {
		return true
	}
	bishopAttacks := p.getBishopAttacks(square) & p.bitboards[side][Bishop]
	if !bishopAttacks.IsEmpty() {
		return true
	}
	rookAttacks := p.getRookAttacks(square) & p.bitboards[side][Rook]
	if !rookAttacks.IsEmpty() {
		return true
	}

	queenAttacks := p.getQueenAttacks(square) & p.bitboards[side][Queen]

	if !queenAttacks.IsEmpty() {
		return true
	}

	kingAttacks := attacks.Kings[square] & p.bitboards[side][King]
	// (pawnAttacks | knightAttacks | bishopAttacks | rookAttacks | queenAttacks | queenAttacks).Print()
	return !kingAttacks.IsEmpty()
}

var colorMap = map[int]string{White: "white", Black: "Black"}

func (p *Position) Print() {
	var castleMap = map[int]string{
		7:  "Wks Wqs Bks",
		11: "Wks Wqs Bqs",
		13: "Wks Bks Bqs",
		14: "Wqs Bks Bqs",
		12: "Bks Bqs",
		3:  "Wks Wqs",
		15: "Wks Wqs Bks Bqs",
		4:  "only Bks",
		8:  "only Bqs",
		2:  "only Wqs",
		1:  "only Wks",
	}
	chess_gyphicons := [2][]string{}
	chess_gyphicons[White] = strings.Split(",♚,♛,♝,♞,♜,♟︎", ",")
	chess_gyphicons[Black] = strings.Split(",♔,♕,♗,♘,♖,♙", ",")

	for sq := 0; sq < 64; sq++ {
		var colorIdx, piece int
		if ((sq) % 8) == 0 {
			fmt.Printf("%v ", 8-(sq/8))
		}
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
	fmt.Println("   a  b  c  d  e  f  g  h ")
	// p.getOccupancy().Print()
	fmt.Printf("%s to move\n", colorMap[p.side])

	/*
		7=>KQk
		11>KQq
		13=>Kkq
		14=>Qkq
		12=>kq
		3=>KQ
		15=>KQkq
	*/

	fmt.Println("castling rights", p.castlingRights, castleMap[p.castlingRights])
	fmt.Println("enpassantSquare", moves.IndexToAlgebraic(p.enPassanteSq))
	fmt.Println("")

}

func (p *Position) MakeMove(move moves.Move) {
	// Copy position before any alterations
	prevPos := p.copy()

	origin, dest := move.Origin(), move.Dest()

	piece, promotedPiece := move.Piece(), move.PromotedPiece()
	pieceBb := &p.bitboards[p.side][piece]
	isDoublePawnPush := piece == Pawn && math.Abs(float64(dest-origin)) == 16
	pieceBb.RemoveBit(origin)
	pieceBb.SetBit(dest)
	// reset enpassant
	p.enPassanteSq = 64

	if move.IsCapture() {
		var capturePiece int
		var capturePieceBb *bitboard.Bitboard

		for capturePiece = 0; capturePiece < len(p.bitboards[p.side^1]); capturePiece++ {
			capturePieceBb = &p.bitboards[p.side^1][capturePiece]
			if capturePiece != OccupancySq && capturePieceBb.BitIsSet(dest) {
				capturePieceBb.RemoveBit(dest)
				break
			}
		}

		if capturePiece == Rook {
			// In the case of capture of a promoted rook e.g r3k2r/Pppp1ppp/1b3nbN/nPP5/BB2P3/q4N2/P2P2PP/r2Q1RK1 w kq - 0 2
			if dest == 0 {
				p.revokeKingSideCastle(Black)
			} else if dest == 7 {
				p.revokeQueenSideCastle(Black)
			} else if dest == 63 {
				p.revokeKingSideCastle(White)
			} else if dest == 56 {
				p.revokeQueenSideCastle(White)
			}
		}

		if move.Enpassant() {

			if p.side == Black {
				capturePieceBb.RemoveBit(dest - 8)
			} else {
				capturePieceBb.RemoveBit(dest + 8)
			}
		}

	}
	if move.IsPromotion() {
		/*
			since we've already updated its new
			location at the beginnig of the function, this undos it
		*/
		pieceBb.RemoveBit(dest)
		p.bitboards[p.side][promotedPiece].SetBit(dest)
	}

	if isDoublePawnPush {

		p.enPassanteSq = dest + 8
		if p.side == Black {
			p.enPassanteSq = dest - 8
		}
	}

	if !move.IsCastling() && piece == King {
		if origin == 60 || origin == 4 {
			p.revokeKingSideCastle(p.side)
			p.revokeQueenSideCastle(p.side)
		}
	}

	if !move.IsCastling() && piece == Rook {

		if origin == 7 || origin == 63 {
			p.revokeKingSideCastle(p.side)

		} else if origin == 0 || origin == 56 {
			p.revokeQueenSideCastle(p.side)
		}
	}

	if move.IsCastling() {
		// fmt.Println("castling")
		// fmt.Println(p.GetFen())
		if dest == 2 || dest == 58 {
			p.bitboards[p.side][Rook].SetBit(dest + 1)
			p.bitboards[p.side][Rook].RemoveBit(dest - 2)

		} else if dest == 6 || dest == 62 {

			p.bitboards[p.side][Rook].SetBit(dest - 1)
			p.bitboards[p.side][Rook].RemoveBit(dest + 1)
		}
		p.revokeKingSideCastle(p.side)
		p.revokeQueenSideCastle(p.side)

	}
	p.setOccupancy(Black)
	p.setOccupancy(White)
	p.side ^= 1
	p.prevPosition = &prevPos
}

func (p *Position) UnMakeMove() {
	if p.prevPosition != nil {
		*p = *p.prevPosition
	}
}

func (p *Position) revokeKingSideCastle(side int) {
	if side == Black {
		p.castlingRights &= ^BlackKingside
	} else {
		p.castlingRights &= ^WhiteKingside
	}
}

func (p *Position) revokeQueenSideCastle(side int) {
	if side == Black {
		p.castlingRights &= ^BlackQueenside
	} else {
		p.castlingRights &= ^WhiteQueenside
	}
}
