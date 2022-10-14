package bitboard

import (
	"fmt"
	"math/bits"
)

type Bitboard uint64

func New(pos int) Bitboard {
	bitboard := Bitboard(0)
	bitboard.SetBit(pos)
	return bitboard
}
func NewMask(mask uint64) Bitboard {
	return Bitboard(mask)
}

func (b *Bitboard) SetBit(pos int) {
	*b |= Bitboard(uint64(1) << uint(pos))
}

func (b Bitboard) BitIsSet(pos int) bool {
	return (b & Bitboard(uint64(1)<<uint(pos))) != 0
}

func (b *Bitboard) RemoveBit(pos int) {
	*b &= Bitboard(^(uint64(1) << uint(pos)))
}

func (b Bitboard) IsEmpty() bool {
	return b == Bitboard(0)
}

func (b Bitboard) PopCount() int {
	count := 0
	for !b.IsEmpty() {
		count++
		b &= b - 1
	}
	return count
}

func (b Bitboard) LsbIdx() int {
	if !b.IsEmpty() {
		return bits.TrailingZeros(uint(b))

	} else {
		return -1
	}
}

func (b Bitboard) MsbIdx() int {
	if !b.IsEmpty() {
		return 63 - bits.LeadingZeros(uint(b))

	} else {
		return -1
	}
}

func (b Bitboard) Print() {
	var i int
	fmt.Println("")
	for i = 0; i < 64; i++ {
		var sq int
		if b.BitIsSet(i) {
			sq = 1
		}
		fmt.Print(sq)
		if ((i + 1) % 8) == 0 {
			fmt.Println("")
		}
	}
	fmt.Println("")
}
