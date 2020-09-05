func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func partitionLabels(S string) []int {
	alpha := make(map[rune]int)

	for idx, c := range S {
		alpha[c] = idx
	}

	start := 0
	end := 0
	var ans []int
	for idx, _ := range S {
		end = max(end, alpha[rune(S[idx])])
		if idx == end {
			ans = append(ans, idx-start+1)
			start = idx + 1
		}
	}
	return ans
}