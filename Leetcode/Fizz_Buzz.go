import "strconv"

func fizzBuzz(n int) []string {
	count := 1
	var datas []string
	for count <= n {
		var na3 = count % 3
		var na5 = count % 5
		if na3 == 0 && na5 == 0 {
			datas = append(datas, "FizzBuzz")
		} else if na5 == 0 {
			datas = append(datas, "Buzz")
		} else if na3 == 0 {
			datas = append(datas, "Fizz")
		} else {
			datas = append(datas, strconv.Itoa(count))

		}
		count++
	}
	return datas
}
