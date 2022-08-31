package t36

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows := make([][]bool, 9, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make([]bool, 9, 9)
	}

	columns := make([][]bool, 9, 9)
	for i := 0; i < 9; i++ {
		columns[i] = make([]bool, 9, 9)
	}

	cells := make([][]bool, 9, 9)
	for i := 0; i < 9; i++ {
		cells[i] = make([]bool, 9, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			value := int(board[i][j] - '1')
			if rows[i][value] == true {
				return false
			}
			rows[i][value] = true

			if columns[j][value] == true {
				return false
			}
			columns[j][value] = true

			k := (i/3)*3 + j/3
			if cells[k][value] == true {
				return false
			}
			cells[k][value] = true
		}
	}
	return true
}

func ExampleIsValidSudoku() {
	fmt.Println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))

	// Output:
	// true
}
