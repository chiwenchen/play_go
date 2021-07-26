package main

func maxProfit(prices []int) int {
	lowest := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < lowest {
			lowest = price
		}
		profit := price - lowest
		if profit > maxProfit {
			maxProfit = profit
		}

	}
	return maxProfit
}
