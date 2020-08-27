func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func mincostTickets(days []int, costs []int) int {
	var dp = make([]int, 366, 366)
	max_day := days[len(days)-1]
	for d, i := 1, 0; d <= max_day; d++ {
		if d == days[i] {
			dp[d] = min(min(dp[d-1]+costs[0], dp[max(0, d-7)]+costs[1]), dp[max(0, d-30)]+costs[2])
			i++

		} else {
			dp[d] = dp[d-1]
		}

	}
	return dp[max_day]
}
