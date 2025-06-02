package validsudoku

// Time complexity: O(1) - The board size is fixed at 9x9, so the operations are constant time.
// Space complexity: O(1) - No additional space is used that scales with input size.
func IsValidSudokuBruteForce(board [][]byte) bool {
	check := func(nums []byte) bool {
		count := [9]int{}
		for _, ch := range nums {
			if ch != '.' {
				idx := ch - '1'
				if count[idx] > 0 {
					return false
				}
				count[idx]++
			}
		}
		return true
	}

	// Check rows
	for i := 0; i < 9; i++ {
		if !check(board[i]) {
			return false
		}
	}

	// Check columns
	for j := 0; j < 9; j++ {
		col := make([]byte, 9)
		for i := 0; i < 9; i++ {
			col[i] = board[i][j]
		}
		if !check(col) {
			return false
		}
	}

	// Check 3x3 sub-boxes
	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			box := make([]byte, 0, 9)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					box = append(box, board[boxRow*3+i][boxCol*3+j])
				}
			}
			if !check(box) {
				return false
			}
		}
	}
	return true
}

// Time complexity: O(1) - The board size is fixed at 9x9, so the operations are constant time.
// Space complexity: O(1) - No additional space is used that scales with input size.
func IsValidSudokuHashSet(board [][]byte) bool {
	rows := [9]map[byte]bool{}
	cols := [9]map[byte]bool{}
	boxes := [9]map[byte]bool{}

	for i := range rows {
		rows[i] = map[byte]bool{}
		cols[i] = map[byte]bool{}
		boxes[i] = map[byte]bool{}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}
			boxIdx := (i/3)*3 + j/3

			if rows[i][val] || cols[j][val] || boxes[boxIdx][val] {
				return false
			}
			rows[i][val] = true
			cols[j][val] = true
			boxes[boxIdx][val] = true
		}
	}
	return true
}

// Time complexity: O(1) - The board size is fixed at 9x9, so the operations are constant time.
// Space complexity: O(1) - No additional space is used that scales with input size.
func IsValidSudokuBitMask(board [][]byte) bool {
	var rows, cols, boxes [9]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}
			bit := 1 << (val - '1')
			boxIdx := (i/3)*3 + j/3

			if (rows[i]&bit) != 0 || (cols[j]&bit) != 0 || (boxes[boxIdx]&bit) != 0 {
				return false
			}

			rows[i] |= bit
			cols[j] |= bit
			boxes[boxIdx] |= bit
		}
	}
	return true
}
