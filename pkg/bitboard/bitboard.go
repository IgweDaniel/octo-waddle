package bitboard

import "fmt"

type Bitboard uint64

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

func (b Bitboard) LsbIndex() int {
	if !b.IsEmpty() {
		bitboard := Bitboard((b & -b) - 1)
		return bitboard.PopCount()

	} else {
		return -1
	}
}

func (b *Bitboard) Print() {
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
	fmt.Println("")
}
