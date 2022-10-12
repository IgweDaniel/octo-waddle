package attacks

import (
	"github.com/igwedaniel/dolly/pkg/bitboard"
)

const (
	rank1 = bitboard.Bitboard(0x00000000000000FF)
	rank4 = bitboard.Bitboard(0x00000000FF000000)
	rank5 = bitboard.Bitboard(0x000000FF00000000)
	rank8 = bitboard.Bitboard(0xFF00000000000000)
	fileA = bitboard.Bitboard(0x0101010101010101)
	fileB = bitboard.Bitboard(0x0202020202020202)
	fileG = bitboard.Bitboard(0x4040404040404040)
	fileH = bitboard.Bitboard(0x8080808080808080)
)

type AttackTable struct {
	PawnAttacks        [2][64]bitboard.Bitboard
	KnightAttacks      [64]bitboard.Bitboard
	KingAttacks        [64]bitboard.Bitboard
	NorthAttackRay     [64]bitboard.Bitboard
	SouthAttackRay     [64]bitboard.Bitboard
	EastAttackRay      [64]bitboard.Bitboard
	WestAttackRay      [64]bitboard.Bitboard
	NorthEastAttackRay [64]bitboard.Bitboard
	NorthWestAttackRay [64]bitboard.Bitboard
	SouthEastAttackRay [64]bitboard.Bitboard
	SouthWestAttackRay [64]bitboard.Bitboard
}

func generateKnightAttackAtIndex(square int) bitboard.Bitboard {
	attacks := bitboard.Bitboard(0)

	sqBb := bitboard.Bitboard(0)
	sqBb.SetBit(square)

	attacks |= (sqBb >> 17) & ^fileH
	attacks |= (sqBb >> 15) & ^fileA
	attacks |= (sqBb >> 10) & ^(fileH | fileG)
	attacks |= (sqBb >> 6) & ^(fileA | fileB)
	attacks |= (sqBb << 17) & ^(fileA)
	attacks |= (sqBb << 15) & ^(fileH)
	attacks |= (sqBb << 10) & ^(fileA | fileB)
	attacks |= (sqBb << 6) & ^(fileH | fileG)

	return attacks
}

func generateWhitePawnAttackAtIndex(square int) bitboard.Bitboard {
	attacks := bitboard.Bitboard(0)

	sqBb := bitboard.Bitboard(0)
	sqBb.SetBit(square)

	attacks |= (sqBb >> 7) & ^fileA
	attacks |= (sqBb >> 9) & ^fileH

	return attacks
}
func generateBlackPawnAttackAtIndex(square int) bitboard.Bitboard {
	attacks := bitboard.Bitboard(0)

	sqBb := bitboard.Bitboard(0)
	sqBb.SetBit(square)

	attacks |= (sqBb << 7) & ^fileH
	attacks |= (sqBb << 9) & ^fileA

	return attacks
}
func generateKingAttackAtIndex(square int) bitboard.Bitboard {
	attacks := bitboard.Bitboard(0)

	sqBb := bitboard.Bitboard(0)
	sqBb.SetBit(square)

	attacks |= (sqBb >> 8)
	attacks |= (sqBb >> 9) & ^fileH
	attacks |= (sqBb >> 7) & ^fileA
	attacks |= (sqBb >> 1) & ^fileH
	attacks |= (sqBb << 8)
	attacks |= (sqBb << 9) & ^fileA
	attacks |= (sqBb << 7) & ^fileH
	attacks |= (sqBb << 1) & ^fileA

	return attacks
}

func generateLookupTables() *AttackTable {
	attackTable := new(AttackTable)

	for index := 0; index < 64; index++ {
		attackTable.PawnAttacks[0][index] = generateWhitePawnAttackAtIndex(index)
		attackTable.PawnAttacks[1][index] = generateBlackPawnAttackAtIndex(index)
		attackTable.KnightAttacks[index] = generateKnightAttackAtIndex(index)
		attackTable.KingAttacks[index] = generateKingAttackAtIndex(index)

	}
	return attackTable
}

var LookupTable = generateLookupTables()
