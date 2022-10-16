package moves

import "fmt"

type Moves []Move

func NewList() Moves {
	return make([]Move, 0, 100)
}

func (m *Moves) Add(piece, origin, dest int) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:      origin,
		dest:        dest,
		attackPiece: piece,
	}))
}

func (m *Moves) AddEnpassantCapture(piece, origin, dest int) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:      origin,
		dest:        dest,
		attackPiece: piece,
		isCapture:   true,
		enpassant:   true,
	}))
}

func (m *Moves) AddCapture(piece, origin, dest int) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:      origin,
		dest:        dest,
		attackPiece: piece,
		isCapture:   true,
	}))
}

func (m *Moves) AddPromotion(piece, origin, dest, promotion int, isCapture bool) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:        origin,
		dest:          dest,
		attackPiece:   piece,
		promotedPiece: promotion,
		isCapture:     isCapture,
	}))

}
func (m *Moves) AddCastling(piece, origin, dest int) {
	*m = append(*m, encodeMove(encodeMoveOpt{
		origin:      origin,
		dest:        dest,
		attackPiece: piece,
		isCastling:  true,
	}))
}

func (m *Moves) Print() {
	var pieceMaps = map[int]string{1: "King",
		2: "Queen",
		3: "Bishop",
		4: "Knight",
		5: "Rook",
		6: "Pawn",
	}

	/*

		Pawn at f4 can capture piece at f3 via enpassante and promote

	*/
	if len(*m) == 0 {
		fmt.Println("No moves for this position")
	}
	for _, move := range *m {
		origin, dest := move.Origin(), move.Dest()
		isCapture := move.IsCapture()
		piece := pieceMaps[move.Piece()]
		var movestr string
		if move.IsPromotion() {
			movestr = fmt.Sprintf("%s at %v can promote via move to %v\n",
				piece,
				IndexToAlgebraic(origin),
				IndexToAlgebraic(dest),
			)
			if isCapture {
				movestr = fmt.Sprintf("%s at %v can promote via capture at %v\n",
					piece,
					IndexToAlgebraic(origin),
					IndexToAlgebraic(dest),
				)
			}
		} else if isCapture {
			movestr = fmt.Sprintf("%s at %v can capture piece at %v\n",
				piece,
				IndexToAlgebraic(origin),
				IndexToAlgebraic(dest))

			if move.Enpassant() {
				movestr = fmt.Sprintf("%s at %v can capture piece at %v via enpassante\n",
					piece,
					IndexToAlgebraic(origin),
					IndexToAlgebraic(dest),
				)
			}
		} else if move.IsCastling() {

			if dest > origin {
				fmt.Printf("%s at %v can castle King side to %v\n",
					pieceMaps[move.Piece()], IndexToAlgebraic(origin),
					IndexToAlgebraic(dest))
			} else {
				fmt.Printf("%s at %v can castle Queen side to %v\n",
					pieceMaps[move.Piece()], IndexToAlgebraic(origin),
					IndexToAlgebraic(dest))
			}

		} else {
			fmt.Printf("%s at %v can move to %v\n",
				pieceMaps[move.Piece()], IndexToAlgebraic(origin),
				IndexToAlgebraic(dest))
		}
		fmt.Print(movestr)
	}
	fmt.Printf("==found %v moves==\n", len(*m))
}
