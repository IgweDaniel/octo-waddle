package bitboard

import "testing"

func TestBitboardSetBit(t *testing.T) {
	tests := []struct {
		b   Bitboard
		pos int
	}{
		{0x0, 5},
		{0xF, 63},
		{0x0, 0},
		{0x1, 0},
		{0xFFFF, 63},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tt.b.SetBit(tt.pos)
			if !tt.b.BitIsSet(tt.pos) {
				t.Fatalf("set(%v) gives %v in %b. Want %v", tt.pos, false, tt.b, true)
			}

			tt.b.RemoveBit(tt.pos)
			if tt.b.BitIsSet(tt.pos) {
				t.Fatalf("remove(%v) gives %v in %b. Want %v", tt.pos, true, tt.b, false)
			}

		})
	}

}

func TestBitboardIsEmpty(t *testing.T) {
	bitboard := Bitboard(0)
	if !bitboard.IsEmpty() {
		t.Errorf("incorrect result: expected %v, got %v", true, false)
	}
	bitboard = Bitboard(0x80)
	if bitboard.IsEmpty() {
		t.Errorf("incorrect result: expected %v, got %v", false, true)
	}
}

func TestBitboardPopCount(t *testing.T) {
	tests := []struct {
		b     Bitboard
		count int
	}{
		{0x510, 3},
		{0x1, 1},
		{0xFFFF, 16},
		{0x0, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			bitCount := tt.b.PopCount()
			if bitCount != tt.count {
				t.Errorf("incorrect result: expected %v, got %v", tt.count, bitCount)
			}

		})
	}

}

func TestBitboardLsbIndex(t *testing.T) {
	tests := []struct {
		b     Bitboard
		index int
	}{
		{0x20000000, 29},
		{0x1, 0},
		{0xcc, 2},
		{0x0, -1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			ls1bIndex := tt.b.LsbIndex()
			if ls1bIndex != tt.index {
				t.Errorf("incorrect result: expected %v, got %v", tt.index, ls1bIndex)
			}

		})
	}

}
