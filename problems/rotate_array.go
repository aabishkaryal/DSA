package problems

func Rotate(nums []int, k int) {
	n := len(nums)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		j := (i + k) % n
		array[j] = nums[i]
	}
	copy(nums, array)
}
