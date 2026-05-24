package array

func maxProfit_naive(price []int, start, end int) int {
	if start >= end {
		return 0
	}

	profit := 0

	for i := start; i < end; i++ {
		for j := i + 1; j <= end; j++ {
			if price[j] > price[i] {
				curr_profit := price[j] - price[i] + maxProfit(price, start, i-1) + maxProfit(price, j + 1, end)

				profit = max(profit, curr_profit)
			}
		}
	}

	return profit
}

func maxProfit(price []int, start, end int) int {
	_ = start
	_ = end

	profit := 0

	for i := 1; i < len(price); i++ {
		if price[i] > price[i - 1] {
			profit += (price[i] - price[i-1])
		}
	}

	return profit
}
