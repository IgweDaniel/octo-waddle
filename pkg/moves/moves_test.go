package moves

import (
	"fmt"
	"testing"
)

func TestMoveEncodingAndDecoding(t *testing.T) {

	tests := []encodeMoveOpt{
		{
			origin:      30,
			dest:        5,
			attackPiece: 3,
		},
		{
			origin:      20,
			dest:        6,
			attackPiece: 3,
			isCapture:   true,
		},
		{
			origin:        40,
			dest:          41,
			isCastling:    true,
			promotedPiece: 2,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("move with origin %v dest: %v", tt.origin, tt.dest), func(t *testing.T) {

			move := encodeMove(tt)

			if move.Origin() != tt.origin {
				t.Errorf("incorrect result: expected %v, got %v", tt.origin, move.Origin())
			}
			if move.Dest() != tt.dest {
				t.Errorf("incorrect result: expected %v, got %v", tt.dest, move.Dest())
			}
			if move.Piece() != tt.attackPiece {
				t.Errorf("incorrect attack piece : expected %v, got %v", tt.attackPiece, move.Piece())
			}

			if move.PromotedPiece() != tt.promotedPiece {
				t.Errorf("incorrect promoted piece: expected %v, got %v", tt.promotedPiece, move.PromotedPiece())
			}

			if move.IsCapture() != tt.isCapture {
				t.Errorf("incorrect capture status: expected %v, got %v", tt.isCapture, move.IsCapture())
			}
			if move.IsCastling() != tt.isCastling {
				t.Errorf("incorrect castle status: expected %v, got %v", tt.isCastling, move.IsCastling())
			}
			if move.Enpassant() != tt.enpassant {
				t.Errorf("incorrect enpassant status: expected %v, got %v", tt.enpassant, move.Enpassant())
			}

		})
	}

}
