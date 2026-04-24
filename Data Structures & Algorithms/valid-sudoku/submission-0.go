type Seen struct {
	rows	[9][10]bool
	cols	[9][10]bool
	boxes	[9][10]bool
}

func (s *Seen) CheckAndMark(row, col, digit int) bool {
	box := (row/3)*3 + col/3

	if s.rows[row][digit] || s.cols[col][digit] || s.boxes[box][digit] {
		return false
	}

	s.rows[row][digit] = true
	s.cols[col][digit] = true
	s.boxes[box][digit] = true
	
	return true
}

func isValidSudoku(board [][]byte) bool {
	seen := Seen{}
	
	for row, line := range board {
		for col, cell := range line {
			if cell == '.' {
				continue
			}

			digit := int(cell - '0')
			if !seen.CheckAndMark(row, col, digit) {
				return false
			}
		}
	}

	return true
}
