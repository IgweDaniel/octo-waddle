package position

import (
	"github.com/igwedaniel/dolly/pkg/attacks"
	"github.com/igwedaniel/dolly/pkg/bitboard"
	"github.com/igwedaniel/dolly/pkg/moves"
)

func (p *Position) GenerateMoves() moves.MoveList {

	moves := moves.NewList()
	generatePawnAttacks(p, &moves)
	generatePawnPushes(p, &moves)
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

func generatePawnPushes(p *Position, moves *moves.MoveList) {

	var step, direction = 8, -1
	var pawnStartRanks = bitboard.NewMask(uint64(rank2))

	if p.side == Black {
		direction = 1
		pawnStartRanks = bitboard.NewMask(uint64(rank7))
	}

	pawns := p.bitboards[p.side][Pawn]
	occupancy := p.getOccupancy()

	for !pawns.IsEmpty() {
		origin := pawns.LsbIdx()
		pawns.RemoveBit(origin)
		dest := origin + step*direction

		if !occupancy.BitIsSet(dest) {
			if pawnStartRanks.BitIsSet(origin) {
				dblPushDest := origin + 2*step*direction
				moves.Add(Pawn, origin, dblPushDest)
			}
			moves.Add(Pawn, origin, dest)

		}

	}

}
