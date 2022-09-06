package proxx

import (
	"errors"
	"math/rand"
	"time"
)

// Proxx represents the game. It exports the methods for game logic implementation
type Proxx struct {
	board      [][]*cell
	stepsCount int
	startTime  time.Time
}

// NewGame creates the game with populated board of specified size and black holes amount
func NewGame(rowsNum, columnsNum, holesNum int) (*Proxx, error) {
	rand.Seed(time.Now().UnixNano())

	if rowsNum < 1 || columnsNum < 1 {
		return nil, errors.New("invalid board size - number of columns and rows can't be less than 1")
	}

	if rowsNum*columnsNum <= holesNum {
		return nil, errors.New("holes number should be less than the total number of cells")
	}

	if holesNum < 1 {
		return nil, errors.New("need to place at least one hole")
	}

	b, err := generateBoard(rowsNum, columnsNum, holesNum)
	if err != nil {
		return nil, err
	}

	game := &Proxx{
		board: b,
	}

	return game, nil
}

func (p *Proxx) OpenCell(row, column int) {
	p.stepsCount++
	p.revealCell(row, column)
}

func (p *Proxx) revealCell(row, column int) {
	clickedCell := p.board[row][column]
	clickedCell.IsOpened = true
	if clickedCell.nearHole() {
		return
	}

	for _, surroundingCell := range clickedCell.SurroundingCells {
		if surroundingCell.nearHole() {
			surroundingCell.IsOpened = true
		} else {
			if !surroundingCell.IsOpened {
				p.revealCell(surroundingCell.Row, surroundingCell.Column)
			}
		}
	}
}

func (p *Proxx) Start() {
	p.startTime = time.Now()
}

func (p *Proxx) IsHole(row int, column int) bool {
	return p.board[row][column].IsHole
}

func (p *Proxx) Opened(row int, column int) bool {
	return p.board[row][column].IsOpened
}

func (p *Proxx) StepsCount() int {
	return p.stepsCount
}

func (p *Proxx) StartTime() time.Time {
	return p.startTime
}
