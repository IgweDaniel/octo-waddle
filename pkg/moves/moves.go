package moves

import "fmt"

type MoveList []Move

func NewList() MoveList {
	return make([]Move, 0, 100)
}

func (m *MoveList) Add(piece, origin, dest int) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:      origin,
		dest:        dest,
		attackPiece: piece,
	}))
}

func (m *MoveList) AddEnpassantCapture(piece, origin, dest int) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:      origin,
		dest:        dest,
		attackPiece: piece,
		isCapture:   true,
		enpassant:   true,
	}))
}

func (m *MoveList) AddCapture(piece, origin, dest int) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:      origin,
		dest:        dest,
		attackPiece: piece,
		isCapture:   true,
	}))
}

func (m *MoveList) AddPromotion(piece, origin, dest, promotion int, isCapture bool) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:        origin,
		dest:          dest,
		attackPiece:   piece,
		promotedPiece: promotion,
		isCapture:     isCapture,
	}))

}
func (m *MoveList) AddCastling(piece, origin, dest int) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:      origin,
		dest:        dest,
		attackPiece: piece,
		isCastling:  true,
	}))
}

func (m *MoveList) Print() {
	var pieceMaps = map[int]string{1: "King",
		2: "Queen",
		3: "Bishop",
		4: "Knight",
		5: "Rook",
		6: "Pawn",
	}
	for _, move := range *m {
		fmt.Printf("%s at %v can move to %v x:%v c:%v e:%v p:%v\n",
			pieceMaps[move.Piece()], IndexToAlgebraic(move.Origin()), IndexToAlgebraic(move.Dest()),
			move.IsCapture(), move.IsCastling(), move.Enpassant(), move.IsPromotion())
	}
}
