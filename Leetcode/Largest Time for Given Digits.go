import "strconv"

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func makeStr(A []int) string {
	return strconv.Itoa(A[0]) + strconv.Itoa(A[1]) + strconv.Itoa(A[2]) + strconv.Itoa(A[3])
}

func largestTimeFromDigits(A []int) string {
	per := permutations(A)
	flag := false
	result := "0000"
	for _, item := range per {
		times := makeStr(item)
		if times < "2400" && times[2:] < "60" {
			if result < times {
				result = times
			}
			flag = true
		}
	}
	if flag == false {
		return ""
	}
	return result[:2] + ":" + result[2:]
}