package problems

func SingleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	accumulator := 0
	for _, num := range nums {
		accumulator ^= num
	}
	return accumulator
}
