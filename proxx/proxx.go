package proxx

import (
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
	game := &Proxx{
		board: board{},
	}

	return game, nil
}

func (p *Proxx) Start() {
	p.startTime = time.Now()
}

func (p *Proxx) StepsCount() int {
	return p.stepsCount
}

func (p *Proxx) StartTime() time.Time {
	return p.startTime
}
