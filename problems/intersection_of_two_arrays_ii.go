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

func IntersectSorted(nums1, nums2 []int) []int {
	result := make([]int, 0)
	x, y := 0, 0
	for x < len(nums1) && y < len(nums2) {
		switch {
		case nums1[x] == nums2[y]:
			result = append(result, nums1[x])
			x++
			y++
		case nums1[x] < nums2[y]:
			x++
		case nums2[y] < nums1[x]:
			y++
		}
	}
	return result
}
