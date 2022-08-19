package problems

func RemoveDuplicates(nums []int) int {
	unique := 0
	for j := 1; j < len(nums); j++ {
		if nums[unique] != nums[j] {
			unique++
			nums[unique] = nums[j]
		}
	}
	return unique + 1
}
