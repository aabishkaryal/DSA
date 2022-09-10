package problems

func PlusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		plusOne := digits[i] + carry
		carry = plusOne / 10
		digits[i] = plusOne % 10
		if carry == 0 {
			break
		}
	}
	if carry != 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
