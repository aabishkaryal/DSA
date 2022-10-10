package problems

import "math"

func ReverseInt(x int) int {
	y := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		y = (y * 10) + digit
	}
	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}
	return y
}
