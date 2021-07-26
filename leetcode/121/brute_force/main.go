package main

// Time Limit Exceeded
func maxProfit(prices []int) int {
	highestProfit := 0

	days := len(prices)
	for start := 0; start < days; start++ {
		for end := 0; end < days; end++ {
			if end <= start {
				continue
			}
			profit := prices[end] - prices[start]
			if profit > highestProfit {
				highestProfit = profit
			}
		}
	}

	return highestProfit
}
