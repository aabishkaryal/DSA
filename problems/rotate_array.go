package problems

func Rotate(nums []int, k int) {
	n := len(nums)
	numLoops := GCD(n, k)
	for i := 0; i < numLoops; i++ {
		store, j := nums[i], (i+k)%n
		for j != i {
			nums[j], store = store, nums[j]
			j = (j + k) % n
		}
		nums[i] = store
	}
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
