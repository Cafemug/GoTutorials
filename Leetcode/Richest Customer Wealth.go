func maximumWealth(accounts [][]int) int {
	result := 0
	for _, val := range accounts {
		temp_sum := 0
		for _, nested_val := range val {
			temp_sum += nested_val
		}
		if result <= temp_sum {
			result = temp_sum
		}
	}
	return result
}