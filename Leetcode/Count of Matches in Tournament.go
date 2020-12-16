func numberOfMatches(n int) int {
	res := 0
	for n > 1 {
		rem_num := n % 2
		div_num := n / 2
		n = rem_num + div_num
		res += div_num
	}
	return res

}