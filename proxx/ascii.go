package proxx

import (
	"fmt"
	"strconv"
)

func (p *Proxx) ASCIIBoard(colored bool) string {
	out := "\t"                 //put the column numbers in line with columns since the first column will be the row numbers
	for i := range p.board[0] { //put column numbers
		out += strconv.Itoa(i) + "\t"
	}
	out += "\n" //first line is the header with column numbers
	out += "\n"

	for i, row := range p.board {
		out += strconv.Itoa(i) + "\t" //row number
		for _, boardCell := range row {
			if boardCell.IsOpened {
				if boardCell.IsHole {
					if colored {
						out += red("@")
					} else {
						out += "@"
					}
				} else {
					// show the number of holes if there are surrounding
					surroundingBlackHoles := boardCell.surroundingBlackHolesNum()
					if surroundingBlackHoles > 0 {
						if colored {
							out += orange(strconv.Itoa(boardCell.surroundingBlackHolesNum()))
						} else {
							out += strconv.Itoa(boardCell.surroundingBlackHolesNum())
						}
					} else {
						out += " "
					}
				}
			} else {
				if colored {
					out += blue("X")
				} else {
					out += "X"
				}
			}
			out += "\t" //separate cells horisontally with a tab
		}
		out += "\n" //separate rows with newline
	}
	out += "\n" //separate the board with a newline

	return out
}

func red(s string) string {
	return fmt.Sprintf("\033[91m%s\033[0m", s)
}

func blue(s string) string {
	return fmt.Sprintf("\033[94m%s\033[0m", s)
}

func orange(s string) string {
	return fmt.Sprintf("\033[93m%s\033[0m", s)
}
