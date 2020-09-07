func largestOverlap(A [][]int, B [][]int) int {
	size := len(A[0])
	result := 0
	resulta := 0
	resultb := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			tempa := 0
			tempb := 0
			for x := 0; x < size; x++ {
				for y := 0; y < size; y++ {
					an := 0
					bn := 0
					if y-j >= 0 && x-i >= 0 {
						an = A[x-i][y-j]
						bn = B[x-i][y-j]
					}
					if an == B[x][y] && B[x][y] == 1 {
						tempb++
					}
					if bn == A[x][y] && A[x][y] == 1 {
						tempa++
					}

					if resulta < tempa {
						resulta = tempa
					}
				}
				if resultb < tempb {
					resultb = tempb
				}
			}
		}
	}
	if resulta < resultb {
		result = resultb
	} else {
		result = resulta
	}
	return result
}