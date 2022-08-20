package problems

func Rotate(nums []int, k int) {
	n := len(nums)
	store := nums[0]
	for i := 0; i < k; i++ {
		for j := 1; j < n; j++ {
			nums[j], store = store, nums[j]
		}
		nums[0] = store
	}
}
