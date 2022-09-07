package problems

func Intersect(nums1, nums2 []int) []int {
	result := make([]int, 0)
	countNums1 := make(map[int]int)

	for _, num := range nums1 {
		countNums1[num]++
	}

	for _, num := range nums2 {
		count, ok := countNums1[num]
		if ok && count > 0 {
			countNums1[num]--
			result = append(result, num)
		}
	}
	return result
}
