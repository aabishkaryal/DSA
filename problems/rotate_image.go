package problems

func RotateImage(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		var temp int
		for j := i; j < n-i-1; j++ {
			temp = matrix[i][j]
			// Because i, j <- n-j-1, i.
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = temp
		}
	}
}
