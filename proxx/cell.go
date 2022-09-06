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

func (c *cell) surroundingBlackHolesNum() int {
	num := 0
	for _, surroundingCell := range c.SurroundingCells {
		if surroundingCell.IsHole {
			num++
		}
	}

	return num
}
