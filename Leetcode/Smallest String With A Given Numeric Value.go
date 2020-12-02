import "strings"

func getSmallestString(n int, k int) string {
	start_num := k - n
	var temp_list []string = make([]string, n)
	alphabet := [27]string{"", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for n > 0 {
		if start_num+1 > 26 {
			temp_list[n-1] = alphabet[26]

			start_num = start_num - 25
		} else if start_num == 0 {
			temp_list[n-1] = alphabet[1]

		} else {
			temp_list[n-1] = alphabet[start_num+1]
			start_num = 0
		}

		n--
	}

	return strings.Join(temp_list, "")

}