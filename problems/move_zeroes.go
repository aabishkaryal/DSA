package problems

func MoveZeroes(nums []int) {
	zero, nonZero := 0, 0
	for zero < len(nums) && nonZero < len(nums) {
		if nums[zero] == 0 && nums[nonZero] != 0 {
			nums[zero], nums[nonZero] = nums[nonZero], nums[zero]
			zero++
		} else if nums[zero] != 0 {
			zero++
		}
		nonZero++
	}
}
