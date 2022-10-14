package position

import (
	"github.com/igwedaniel/dolly/pkg/attacks"
	"github.com/igwedaniel/dolly/pkg/bitboard"
	"github.com/igwedaniel/dolly/pkg/moves"
)

func (p *Position) GenerateMoves() moves.MoveList {

	moves := moves.NewList()
	generatePawnAttacks(p, &moves)
	moves.Print()
	return moves
}

func generatePawnAttacks(p *Position, moves *moves.MoveList) {
	side := p.side
	pawns := p.bitboards[side][Pawn]

	for !pawns.IsEmpty() {
		origin := pawns.LsbIdx()
		pawns.RemoveBit(origin)
		captures := attacks.Pawns[side][origin] & p.bitboards[side^1][OccupancySq]
		for !captures.IsEmpty() {
			// check for promotion capture
			promoteRanks := bitboard.NewMask(uint64(rank1) | uint64(rank8))
			dest := captures.LsbIdx()
			if promoteRanks.BitIsSet(dest) {
				moves.AddPromotion(Pawn, origin, dest, Queen, true)
				moves.AddPromotion(Pawn, origin, dest, Rook, true)
				moves.AddPromotion(Pawn, origin, dest, Bishop, true)
				moves.AddPromotion(Pawn, origin, dest, Knight, true)

			} else {
				moves.AddCapture(Pawn, origin, dest)
			}
			captures.RemoveBit(dest)
		}
		if p.enPassanteSq != 64 {
			enpassantAttack := bitboard.New(p.enPassanteSq) & attacks.Pawns[side][origin]
			if !enpassantAttack.IsEmpty() {
				enpassanteSq := enpassantAttack.LsbIdx()
				moves.AddEnpassantCapture(Pawn, origin, enpassanteSq)
			}
		}

	}
}

// func generatePawnPushes(p *Position, moves *moves.MoveList) {

// 	pawns := p.bitboards[p.side][Pawn]
// 	occupancy := p.getOccupancy()

// }
