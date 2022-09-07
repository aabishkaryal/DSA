package problems

func Intersect(nums1, nums2 []int) []int {
	result := make([]int, 0)
	countNums1 := make(map[int]int)
	countNums2 := make(map[int]int)

	for _, num := range nums1 {
		countNums1[num]++
	}

	for _, num := range nums2 {
		countNums2[num]++
	}

	for num, count := range countNums1 {
		repeatCount := min(count, countNums2[num])
		for i := 0; i < repeatCount; i++ {
			result = append(result, num)
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
