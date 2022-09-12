package problems

func TwoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i, num := range nums {
		valueNeeded := target - num
		index, ok := cache[valueNeeded]
		if ok {
			return []int{i, index}
		}
		cache[num] = i
	}
	return []int{}
}
