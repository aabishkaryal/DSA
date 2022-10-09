package problems

func IsValidSudoku(board [][]byte) bool {
	return validRows(board) && validCols(board) && valid3x3(board)
}

func validRows(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		temp := make(map[byte]bool)
		for j := 0; j < len(board[i]); j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}
			_, ok := temp[val]
			if ok {
				return false
			}
			temp[val] = true
		}
	}
	return true
}

func validCols(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		temp := make(map[byte]bool)
		for j := 0; j < len(board[i]); j++ {
			val := board[j][i]
			if val == '.' {
				continue
			}
			_, ok := temp[val]
			if ok {
				return false
			}
			temp[val] = true
		}
	}
	return true
}

func valid3x3(board [][]byte) bool {
	for x := 0; x < 9; x += 3 {
		for y := 0; y < 9; y += 3 {
			temp := make(map[byte]bool)
			for i := x; i < x+3; i++ {
				for j := y; j < y+3; j++ {
					val := board[i][j]
					if val == '.' {
						continue
					}
					_, ok := temp[val]
					if ok {
						return false
					}
					temp[val] = true
				}
			}
		}
	}
	return true
}
