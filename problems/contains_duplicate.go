package problems

import "sort"

func ContainsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sort.Ints(nums)
	for i := range nums[1:] {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
