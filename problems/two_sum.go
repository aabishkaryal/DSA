package problems

import "sort"

func TwoSum(nums []int, target int) []int {
	original := append([]int{}, nums...)
	sort.Ints(nums)
	lhs, rhs := 0, len(nums)-1
	for lhs < rhs {
		switch sum := nums[lhs] + nums[rhs]; {
		case sum == target:
			n1, n2 := nums[lhs], nums[rhs]
			return findIndex(original, n1, n2)
		case sum < target:
			lhs++
		default:
			rhs--
		}
	}
	return []int{}
}

func findIndex(nums []int, n1, n2 int) []int {
	i1, i2 := -1, -1
	for i, num := range nums {
		if i1 < 0 && num == n1 {
			i1 = i
		} else if i2 < 0 && num == n2 {
			i2 = i
		}
		if i1 >= 0 && i2 >= 0 {
			break
		}
	}
	return []int{i1, i2}
}
