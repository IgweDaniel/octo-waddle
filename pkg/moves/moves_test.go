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
				t.Errorf("incorrect origin square: expected %v, got %v", tt.origin, move.Origin())
			}
			if move.Dest() != tt.dest {
				t.Errorf("incorrect destination square: expected %v, got %v", tt.dest, move.Dest())
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

func TestMoveListAddMove(t *testing.T) {
	moveList := NewList()
	var origin, dest int = 40, 30
	// The Queen
	attackPiece := 2
	moveList.Add(attackPiece, origin, dest)
	if len(moveList) != 1 {
		t.Errorf("incorrect length of move list: expected %v, got %v", 1, len(moveList))
	}
	move := moveList[0]

	if move.Origin() != origin {
		t.Errorf("incorrect origin square: expected %v, got %v", origin, move.Origin())
	}
	if move.Dest() != dest {
		t.Errorf("incorrect destination square: expected %v, got %v", dest, move.Dest())
	}
	if move.Piece() != attackPiece {
		t.Errorf("incorrect attack piece : expected %v, got %v", attackPiece, move.Piece())
	}

}

func TestMoveListAddCapture(t *testing.T) {
	moveList := NewList()
	var origin, dest int = 40, 30
	// The Queen
	attackPiece := 2
	moveList.AddCapture(attackPiece, origin, dest)
	if len(moveList) != 1 {
		t.Errorf("incorrect length of move list: expected %v, got %v", 1, len(moveList))
	}
	move := moveList[0]

	if move.Origin() != origin {
		t.Errorf("incorrect origin square: expected %v, got %v", origin, move.Origin())
	}
	if move.Dest() != dest {
		t.Errorf("incorrect destination square: expected %v, got %v", dest, move.Dest())
	}
	if move.Piece() != attackPiece {
		t.Errorf("incorrect attack piece : expected %v, got %v", attackPiece, move.Piece())
	}
	if !move.IsCapture() {
		t.Errorf("incorrect capture status : expected %v, got %v", true, move.IsCapture())
	}

}

func TestMoveListAddPromotion(t *testing.T) {
	moveList := NewList()
	var origin, dest int = 40, 30
	// The Queen
	attackPiece, promotedPiece := 8, 2
	moveList.AddPromotion(attackPiece, origin, dest, promotedPiece, true)
	if len(moveList) != 1 {
		t.Errorf("incorrect length of move list: expected %v, got %v", 1, len(moveList))
	}
	move := moveList[0]

	if move.Origin() != origin {
		t.Errorf("incorrect origin square: expected %v, got %v", origin, move.Origin())
	}
	if move.Dest() != dest {
		t.Errorf("incorrect destination square: expected %v, got %v", dest, move.Dest())
	}
	if move.Piece() != attackPiece {
		t.Errorf("incorrect attack piece : expected %v, got %v", attackPiece, move.Piece())
	}
	if !move.IsCapture() {
		t.Errorf("incorrect capture status : expected %v, got %v", true, move.IsCapture())
	}
	if move.PromotedPiece() != promotedPiece {
		t.Errorf("incorrect capture status : expected %v, got %v", promotedPiece, move.PromotedPiece())
	}

}

func TestMoveListAddCastling(t *testing.T) {
	moveList := NewList()
	var origin, dest int = 40, 30
	// The King
	attackPiece := 1
	moveList.AddCastling(attackPiece, origin, dest)
	if len(moveList) != 1 {
		t.Errorf("incorrect length of move list: expected %v, got %v", 1, len(moveList))
	}
	move := moveList[0]

	if !move.IsCastling() {
		t.Errorf("incorrect catling status : expected %v, got %v", true, move.IsCastling())
	}

	if move.Origin() != origin {
		t.Errorf("incorrect origin square: expected %v, got %v", origin, move.Origin())
	}
	if move.Dest() != dest {
		t.Errorf("incorrect destination square: expected %v, got %v", dest, move.Dest())
	}
	if move.Piece() != attackPiece {
		t.Errorf("incorrect attack piece : expected %v, got %v", attackPiece, move.Piece())
	}

	if move.IsCapture() {
		t.Errorf("incorrect capture status : expected %v, got %v", false, move.IsCapture())
	}
	if move.IsPromotion() {
		t.Errorf("incorrect promotion status : expected %v, got %v", false, move.IsPromotion())
	}
}

func TestMoveListAddEnpassantCapture(t *testing.T) {
	moveList := NewList()
	var origin, dest int = 40, 30
	// The King
	attackPiece := 1
	moveList.AddEnpassantCapture(attackPiece, origin, dest)
	if len(moveList) != 1 {
		t.Errorf("incorrect length of move list: expected %v, got %v", 1, len(moveList))
	}
	move := moveList[0]
	if !move.IsCapture() {
		t.Errorf("incorrect capture status : expected %v, got %v", true, move.IsCapture())
	}
	if !move.Enpassant() {
		t.Errorf("incorrect enpassant status : expected %v, got %v", true, move.IsCapture())
	}

	if move.Origin() != origin {
		t.Errorf("incorrect origin square: expected %v, got %v", origin, move.Origin())
	}
	if move.Dest() != dest {
		t.Errorf("incorrect destination square: expected %v, got %v", dest, move.Dest())
	}
	if move.Piece() != attackPiece {
		t.Errorf("incorrect attack piece : expected %v, got %v", attackPiece, move.Piece())
	}

	if move.IsPromotion() {
		t.Errorf("incorrect promotion status : expected %v, got %v", false, move.IsPromotion())
	}

	if move.IsCastling() {
		t.Errorf("incorrect catling status : expected %v, got %v", false, move.IsCastling())
	}
}
