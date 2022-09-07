package problems

func Intersect(nums1, nums2 []int) []int {
	result := make([]int, 0)
	for _, num := range nums1 {
		for j, check := range nums2 {
			if num == check {
				result = append(result, num)
				nums2 = removeElement(nums2, j)
				break
			}
		}
	}
	return result
}

func removeElement(arr []int, index int) []int {
	n := len(arr)
	arr[index] = arr[n-1]
	return arr[:n-1]
}
