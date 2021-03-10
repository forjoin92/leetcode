package leetcode_122

func maxProfit(prices []int) int {
	days := len(prices)
	if days < 2 {
		return 0
	}
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit = profit + prices[i] - prices[i-1]
		}
	}
	return profit
}
