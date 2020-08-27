var check = make([]int, 366, 366)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func contain(days []int, day int) bool {
	for _, item := range days {
		if item == day {
			return true
		}
	}
	return false
}

func dp(i int, days []int, costs []int) int {
	if i > 365 {
		return 0
	}
	if check[i] != 0 {
		return check[i]
	}
	var ans int
	if contain(days, i) == true {
		ans = min(min(dp(i+1, days, costs)+costs[0],
			dp(i+7, days, costs)+costs[1]),
			dp(i+30, days, costs)+costs[2])

	} else {
		ans = dp(i+1, days, costs)

	}
	check[i] = ans

	return ans
}

func mincostTickets(days []int, costs []int) int {
	return dp(1, days, costs)
}
