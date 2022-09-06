package proxx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	a := assert.New(t)
	rNum, cNum, hNum := 4, 5, 6
	game, err := NewGame(rNum, cNum, hNum)
	a.NoError(err)
	a.Equal(len(game.board), rNum)

	for i, row := range game.board {
		a.Equal(len(row), cNum, "row %d have incorrect length", i)
	}

	actualHolesNum := 0
	for i := 0; i < rNum; i++ {
		for j := 0; j < cNum; j++ {
			if game.board[i][j].IsHole {
				actualHolesNum++
			}
			a.False(game.board[i][j].IsOpened, "cell [%d %d] is opened. All cells should be closed initially", i, j)
		}
	}

	a.Equal(actualHolesNum, hNum)
	a.Equal(game.stepsCount, 0)

	game.OpenCell(2, 1)
	a.True(game.board[2][1].IsOpened)

	_, err = NewGame(-1, 10, 4)
	a.Error(err)

	_, err = NewGame(10, -10, 4)
	a.Error(err)

	_, err = NewGame(10, 10, 150)
	a.Error(err)

	_, err = NewGame(10, 10, 0)
	a.Error(err)
}

func TestGame(t *testing.T) {
	a := assert.New(t)
	rNum, cNum, hNum := 5, 5, 5
	game, err := NewGame(rNum, cNum, hNum)
	a.NoError(err)

	a.True(game.StartTime().IsZero())
	game.Start()
	a.False(game.StartTime().IsZero())

	//reset the cell and set holes manually to eliminate randomness of holes placement
	for _, row := range game.board {
		for _, c := range row {
			c.IsHole = false
		}
	}

	game.board[0][3].IsHole = true
	game.board[0][4].IsHole = true
	game.board[1][4].IsHole = true
	game.board[2][4].IsHole = true
	game.board[2][1].IsHole = true

	a.True(game.IsHole(0, 3))
	a.False(game.IsHole(4, 4))
	a.False(game.Opened(4, 4))

	//check if the cells show the correct number of surrounding black holes
	a.Equal(game.board[1][3].surroundingBlackHolesNum(), 4)
	a.Equal(game.board[2][3].surroundingBlackHolesNum(), 2)
	a.Equal(game.board[3][2].surroundingBlackHolesNum(), 1)
	a.Equal(game.board[4][2].surroundingBlackHolesNum(), 0)

	game.OpenCell(4, 0)
	expectedOpenedCells := make(map[int]map[int]struct{})
	expectedOpenedCells[4] = map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}, 4: {}}
	expectedOpenedCells[3] = map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}, 4: {}}
	for _, row := range game.board {
		for _, c := range row {
			cellShouldBeOpened := false
			if openedRow, found := expectedOpenedCells[c.Row]; found {
				_, cellShouldBeOpened = openedRow[c.Column]
			}

			if cellShouldBeOpened {
				a.True(c.IsOpened, "Cell [%d %d] should be opened but it's closed", c.Row, c.Column)
			} else {
				a.False(c.IsOpened, "Cell [%d %d] should be closed but it's opened", c.Row, c.Column)
			}
		}
	}

	a.False(game.Won())
	game.OpenCell(0, 0)
	a.False(game.Won())
	game.OpenCell(1, 2)
	a.False(game.Won())
	game.OpenCell(1, 3)
	a.False(game.Won())
	game.OpenCell(2, 3)
	a.False(game.Won())
	game.OpenCell(2, 2)
	a.False(game.Won())
	game.OpenCell(2, 0)
	a.True(game.Won())
	a.Equal(7, game.StepsCount())
}
