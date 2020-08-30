func rand10() int {
	index := 41
	for index > 40 {
		row := rand7()
		col := rand7()
		index = (row-1)*7 + col
	}
	return 1 + index%10
}