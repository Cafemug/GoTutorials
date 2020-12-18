func minPartitions(n string) int {
	res := 1
	for _, char := range n {
		num := int(char) - 48
		if res < num {
			res = num
		}
	}

	return res
}