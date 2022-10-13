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

var Pawns [2][64]bitboard.Bitboard

var Knights,
	Kings,
	NorthRay,
	SouthRay,
	EastRay,
	WestRay,
	NorthEastRay,
	NorthWestRay,
	SouthEastRay,
	SouthWestRay [64]bitboard.Bitboard

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

func generateNorthEastRayAtIndex(square int) bitboard.Bitboard {

	attacks := bitboard.NewMask(0)
	sqBb := bitboard.New(square)

	square -= 7

	for ; square%8 != 0; square -= 7 {
		sqBb >>= 7
		attacks |= sqBb
	}
	return attacks

}
func generateNorthWestRayAtIndex(square int) bitboard.Bitboard {

	attacks := bitboard.NewMask(0)
	sqBb := bitboard.New(square)

	for ; square%8 != 0; square -= 9 {
		sqBb >>= 9
		attacks |= sqBb
	}

	return attacks

}
func generateSouthWestRayAtIndex(square int) bitboard.Bitboard {

	attacks := bitboard.NewMask(0)
	sqBb := bitboard.New(square)
	for ; square%8 != 0; square += 7 {
		sqBb <<= 7
		attacks |= sqBb
	}

	return attacks

}
func generateSouthEastRayAtIndex(square int) bitboard.Bitboard {

	attacks := bitboard.NewMask(0)
	sqBb := bitboard.New(square)
	square += 9

	for ; square%8 != 0; square += 9 {
		sqBb <<= 9
		attacks |= sqBb
	}
	return attacks

}

func init() {

	for index := 0; index < 64; index++ {
		Pawns[0][index] = generateWhitePawnAttackAtIndex(index)
		Pawns[1][index] = generateBlackPawnAttackAtIndex(index)
		Knights[index] = generateKnightAttackAtIndex(index)
		Kings[index] = generateKingAttackAtIndex(index)

		// South Ray
		southMask := bitboard.NewMask(0x0101010101010100)
		SouthRay[index] = southMask << index

		// North Ray
		northMask := bitboard.NewMask(0x0080808080808080)
		NorthRay[63-index] = northMask >> index

		// east array
		eastMask := bitboard.NewMask(1)
		EastRay[index] = bitboard.Bitboard(2 * ((eastMask << (index | 7)) - (eastMask << index)))

		// east array
		westMask := bitboard.NewMask(1)
		WestRay[index] = bitboard.Bitboard((westMask << index) - (westMask << (index & 56)))

		//NorthEast
		NorthEastRay[index] = generateNorthEastRayAtIndex(index)

		// NorthWest
		NorthWestRay[index] = generateNorthWestRayAtIndex(index)

		// SouthWest
		SouthWestRay[index] = generateSouthWestRayAtIndex(index)

		// SouthWest
		SouthEastRay[index] = generateSouthEastRayAtIndex(index)

	}
}
