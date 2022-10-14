package position

import (
	"github.com/igwedaniel/dolly/pkg/attacks"
	"github.com/igwedaniel/dolly/pkg/bitboard"
	"github.com/igwedaniel/dolly/pkg/moves"
)

func (p *Position) GenerateMoves() moves.MoveList {

	moves := moves.NewList()

	generatePawnPushes(p, &moves)
	generatePawnCaptures(p, &moves)
	generateKnightMoves(p, &moves)
	generateBishopMoves(p, &moves)
	generateRookMoves(p, &moves)
	generateQueenMoves(p, &moves)
	moves.Print()

	return moves
}

func generatePawnCaptures(p *Position, moves *moves.MoveList) {
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

func generateKnightMoves(p *Position, moves *moves.MoveList) {
	activeSideOccupancy := p.bitboards[p.side][OccupancySq]
	for knights := p.bitboards[p.side][Knight]; !knights.IsEmpty(); {

		origin := knights.LsbIdx()
		knights.RemoveBit(origin)

		for attacks := attacks.Knights[origin] & ^activeSideOccupancy; !attacks.IsEmpty(); {
			dest := attacks.LsbIdx()
			attacks.RemoveBit(dest)
			if p.bitboards[p.side^1][OccupancySq].BitIsSet(dest) {
				moves.AddCapture(Knight, origin, dest)
			} else {
				moves.Add(Knight, origin, dest)
			}
		}

	}
}

func generateBishopMoves(p *Position, moves *moves.MoveList) {
	activeSideOccupancy := p.bitboards[p.side][OccupancySq]
	for bishops := p.bitboards[p.side][Bishop]; !bishops.IsEmpty(); {

		origin := bishops.LsbIdx()
		bishops.RemoveBit(origin)

		for attacks := p.getBishopAttacks(origin) & ^activeSideOccupancy; !attacks.IsEmpty(); {
			dest := attacks.LsbIdx()
			attacks.RemoveBit(dest)
			if p.bitboards[p.side^1][OccupancySq].BitIsSet(dest) {
				moves.AddCapture(Bishop, origin, dest)
			} else {
				moves.Add(Bishop, origin, dest)
			}
		}

	}
}
func generateRookMoves(p *Position, moves *moves.MoveList) {
	activeSideOccupancy := p.bitboards[p.side][OccupancySq]
	for rooks := p.bitboards[p.side][Rook]; !rooks.IsEmpty(); {

		origin := rooks.LsbIdx()
		rooks.RemoveBit(origin)

		for attacks := p.getRookAttacks(origin) & ^activeSideOccupancy; !attacks.IsEmpty(); {
			dest := attacks.LsbIdx()
			attacks.RemoveBit(dest)
			if p.bitboards[p.side^1][OccupancySq].BitIsSet(dest) {
				moves.AddCapture(Rook, origin, dest)
			} else {
				moves.Add(Rook, origin, dest)
			}
		}

	}
}
func generateQueenMoves(p *Position, moves *moves.MoveList) {
	activeSideOccupancy := p.bitboards[p.side][OccupancySq]
	for queens := p.bitboards[p.side][Queen]; !queens.IsEmpty(); {

		origin := queens.LsbIdx()
		queens.RemoveBit(origin)

		for attacks := p.getQueenAttacks(origin) & ^activeSideOccupancy; !attacks.IsEmpty(); {
			dest := attacks.LsbIdx()
			attacks.RemoveBit(dest)
			if p.bitboards[p.side^1][OccupancySq].BitIsSet(dest) {
				moves.AddCapture(Queen, origin, dest)
			} else {
				moves.Add(Queen, origin, dest)
			}
		}

	}
}

func generateKingMoves(p *Position, moves *moves.MoveList) {

}
