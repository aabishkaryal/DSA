package problems

func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		// if today's price is lower than tomorrow's
		// buy today and sell tomorrow making a profit
		// repeat for the next day until the last.
		if prices[i] < prices[i+1] {
			profit += (prices[i+1] - prices[i])
		}
	}
	return profit
}
