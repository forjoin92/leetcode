package leetcode_122

func maxProfit(prices []int) int {
	days := len(prices)
	if days < 2 {
		return 0
	}
	profit := 0
	from, to := 0, 1
	for to < days {
		if prices[to] < prices[to-1] {
			profit = profit + prices[to-1] - prices[from]
			from = to
		}
		to++
	}
	if to > from && prices[to-1] > prices[from] {
		profit = profit + prices[to-1] - prices[from]
	}
	return profit
}
