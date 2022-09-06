package proxx

import "math/rand"

// The first idea was to use a simple two-dimensional slice, but then I realized it'd be hard to do
// the traversal of adjacent cells. That's why a graph approach was chosen, where every cell knows about near located cells
type board [][]*cell

// generateBoard generate the board with randomly located black holes
// It includes generation the slices, populating the reference of adjacent cells for each cell and placing the black holes randomly
func generateBoard(rowsNum, columnsNum, numberOfHoles int) (board, error) {
	b := make(board, rowsNum)
	for i := range b {
		b[i] = make([]*cell, columnsNum)
	}

	type coordinate struct {
		row, column int
	}
	var coordinates []coordinate
	for i := 0; i < rowsNum; i++ {
		for j := 0; j < columnsNum; j++ {
			coordinates = append(coordinates, coordinate{row: i, column: j})
			b[i][j] = &cell{Row: i, Column: j}
		}
	}

	for i := 0; i < rowsNum; i++ {
		for j := 0; j < columnsNum; j++ {
			b.setSurroundingCells(i, j)
		}
	}

	i := 0
	for i < numberOfHoles {
		cellIndex := rand.Intn(len(coordinates))
		cellCoordinate := coordinates[cellIndex]
		b[cellCoordinate.row][cellCoordinate.column].IsHole = true
		coordinates = append(coordinates[:cellIndex], coordinates[cellIndex+1:]...)
		i++
	}

	return b, nil
}

func (b board) setSurroundingCells(row int, column int) {
	for i := row - 1; i <= row+1; i++ {
		for j := column - 1; j <= column+1; j++ {
			if i == row && j == column {
				continue //the cell we're looking adjacent black holes for
			}
			if i < 0 || j < 0 || i > len(b)-1 || j > len(b[0])-1 {
				continue // out of bound
			}

			b[row][column].SurroundingCells = append(b[row][column].SurroundingCells, b[i][j])
		}
	}
}
