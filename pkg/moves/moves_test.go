package moves

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveEncodingAndDecoding(t *testing.T) {
	assert := assert.New(t)
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
		t.Run(fmt.Sprintf("origin:%v dest: %v", tt.origin, tt.dest), func(t *testing.T) {

			move := encodeMove(tt)

			assert.Equal(move.Origin(), tt.origin, "move origin  not valid")
			assert.Equal(move.Dest(), tt.dest, "move dest  not valid")
			assert.Equal(move.Piece(), tt.attackPiece, "move piece  not valid")
			assert.Equal(move.IsCastling(), tt.isCapture, "move castling  not valid")
			assert.Equal(move.IsCapture(), tt.isCastling, "move capture  not valid")
			assert.Equal(move.Enpassant(), tt.enpassant, "move enpassant  not valid")
			assert.Equal(move.PromotedPiece(), tt.promotedPiece, "move enpassant  not valid")
		})
	}

}
