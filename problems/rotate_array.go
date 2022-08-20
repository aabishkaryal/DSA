package problems

func Rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n)
	reverse(nums, 0, k)
	reverse(nums, k, n)
}

// Reverse the array from start(inclusive) to end(exclusive).
func reverse(arr []int, start, end int) {
	mid := (start + end) / 2
	for i := 0; i < (mid - start); i++ {
		arr[start+i], arr[end-i-1] = arr[end-i-1], arr[start+i]
	}
}
