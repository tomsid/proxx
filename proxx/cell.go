package proxx

type cell struct {
	IsHole           bool
	IsOpened         bool
	Row              int
	Column           int
	SurroundingCells []*cell
}

func (c *cell) nearHole() bool {
	for _, surroundingCell := range c.SurroundingCells {
		if surroundingCell.IsHole {
			return true
		}
	}

	return false
}

// surroundingBlackHolesNum don't store the number of adjacent black holes
// that approach was chosen because it's not a heavy operation, it's always at most 8 iterations
// but eliminate storing of another source of truth. This way you don't have to think of keeping the references
// and the count in sync and can implement logic of moving holes around without calculating the count
func (c *cell) surroundingBlackHolesNum() int {
	num := 0
	for _, surroundingCell := range c.SurroundingCells {
		if surroundingCell.IsHole {
			num++
		}
	}

	return num
}
