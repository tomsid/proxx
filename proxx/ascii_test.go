package proxx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestASCIIPrint(t *testing.T) {
	a := assert.New(t)
	rNum, cNum, hNum := 5, 5, 5
	game, err := NewGame(rNum, cNum, hNum)
	a.NoError(err)

	//reset the cell and set holes manually to eliminate randomness of holes placement
	for _, row := range game.board {
		for _, c := range row {
			c.IsHole = false
		}
	}

	game.board[3][2].IsHole = true
	game.board[1][4].IsHole = true
	game.board[4][0].IsHole = true
	game.board[4][4].IsHole = true
	game.board[0][0].IsHole = true

	a.Equal(`	0	1	2	3	4	

0	X	X	X	X	X	
1	X	X	X	X	X	
2	X	X	X	X	X	
3	X	X	X	X	X	
4	X	X	X	X	X	

`, game.ASCIIBoard(false))

	game.OpenCell(2, 0)

	a.Equal(`	0	1	2	3	4	

0	X	X	X	X	X	
1	1	1	X	X	X	
2	 	1	X	X	X	
3	1	2	X	X	X	
4	X	X	X	X	X	

`, game.ASCIIBoard(false))

	game.OpenAllCells()
	a.Equal(`	0	1	2	3	4	

0	@	1	 	1	1	
1	1	1	 	1	@	
2	 	1	1	2	1	
3	1	2	@	2	1	
4	@	2	1	2	@	

`, game.ASCIIBoard(false))
}
