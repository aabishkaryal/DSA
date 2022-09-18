package problems

func RotateImage(matrix [][]int) {
	n := len(matrix)
	temp := make([][]int, n)
	for i := range matrix {
		temp[i] = make([]int, n)
		for j := range matrix[i] {
			temp[i][j] = matrix[n-1-j][i]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = temp[i][j]
		}
	}
}
