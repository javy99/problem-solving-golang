package validsudoku

import "testing"

func TestIsValidSudoku(t *testing.T) {
	cases := []struct {
		name     string
		fn       func([][]byte) bool
		board    [][]byte
		expected bool
	}{
		// IsValidSudokuBruteForce
		{"Brute Force - Valid", IsValidSudokuBruteForce, validSudokuBoard(), true},
		{"Brute Force - Invalid Row", IsValidSudokuBruteForce, invalidRowSudokuBoard(), false},
		{"Brute Force - Invalid Column Only", IsValidSudokuBruteForce, invalidColOnlySudokuBoard(), false},
		{"Brute Force - Invalid Box Only", IsValidSudokuBruteForce, invalidBoxOnlySudokuBoard(), false},

		// IsValidSudokuHashSet - can use original or new ones, HashSet checks "all at once"
		{"Hash Set - Valid", IsValidSudokuHashSet, validSudokuBoard(), true},
		{"Hash Set - Invalid Row", IsValidSudokuHashSet, invalidRowSudokuBoard(), false},
		{"Hash Set - Invalid Column", IsValidSudokuHashSet, originalInvalidColSudokuBoard(), false}, // or invalidColOnlySudokuBoard()
		{"Hash Set - Invalid Box", IsValidSudokuHashSet, originalInvalidBoxSudokuBoard(), false},    // or invalidBoxOnlySudokuBoard()

		// IsValidSudokuBitMask - similar to HashSet
		{"BitMask - Valid", IsValidSudokuBitMask, validSudokuBoard(), true},
		{"BitMask - Invalid Row", IsValidSudokuBitMask, invalidRowSudokuBoard(), false},
		{"BitMask - Invalid Column", IsValidSudokuBitMask, originalInvalidColSudokuBoard(), false}, // or invalidColOnlySudokuBoard()
		{"BitMask - Invalid Box", IsValidSudokuBitMask, originalInvalidBoxSudokuBoard(), false},    // or invalidBoxOnlySudokuBoard()
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.fn(c.board)
			if got != c.expected {
				t.Errorf("%s failed: expected %v, got %v for board:\n%v", c.name, c.expected, got, boardToString(c.board))
			}
		})
	}
}

// Helper to print board for easier debugging
func boardToString(board [][]byte) string {
	s := ""
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s += string(board[i][j]) + " "
		}
		s += "\n"
	}
	return s
}

func validSudokuBoard() [][]byte {
	return [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
}

func invalidRowSudokuBoard() [][]byte {
	board := validSudokuBoard()
	board[0][1] = '5' // Duplicate 5 in first row
	return board
}

func invalidColOnlySudokuBoard() [][]byte {
	board := validSudokuBoard()
	board[1][0] = '8' // Col 0: 5,8,.,8,... (invalid); Row 1: 8,.,.,1,9,5,... (valid)
	return board
}

func invalidBoxOnlySudokuBoard() [][]byte {
	board := validSudokuBoard()
	board[1][1] = '8' // Box 0,0: contains two '8's; Row 1 and Col 1 remain valid
	return board
}

func originalInvalidColSudokuBoard() [][]byte {
	board := validSudokuBoard()
	board[1][0] = '5'
	return board
}

func originalInvalidBoxSudokuBoard() [][]byte {
	board := validSudokuBoard()
	board[1][1] = '9'
	return board
}
