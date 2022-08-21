package problems

func ContainsDuplicate(nums []int) bool {
	cache := make(map[int]int)
	for _, num := range nums {
		_, ok := cache[num]
		if ok {
			return true
		}
		cache[num] = 1
	}
	return false
}
