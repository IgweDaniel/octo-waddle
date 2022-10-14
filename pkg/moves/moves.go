package moves

type MoveList []Move

func NewList() MoveList {
	return make([]Move, 0, 100)
}

func (m *MoveList) AddMove(piece, origin, dest int) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:      origin,
		dest:        dest,
		attackPiece: piece,
	}))
}

func (m *MoveList) AddCaptureMove(piece, origin, dest int) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:      origin,
		dest:        dest,
		attackPiece: piece,
		isCapture:   true,
	}))
}

func (m *MoveList) AddPromotionMove(piece, origin, dest, promotion int, isCapture bool) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:        origin,
		dest:          dest,
		attackPiece:   piece,
		promotedPiece: promotion,
		isCapture:     isCapture,
	}))

}
func (m *MoveList) AddCastlingMove(piece, origin, dest, promotion int) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:      origin,
		dest:        dest,
		attackPiece: piece,
		isCastling:  true,
	}))
}
