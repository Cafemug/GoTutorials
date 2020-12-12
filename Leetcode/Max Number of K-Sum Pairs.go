import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	size := len(nums)
	res := 0

	for i, j := 0, size-1; i < j; {
		temp_sum := nums[i] + nums[j]
		if temp_sum == k {
			res++
			i++
			j--
		} else if temp_sum > k {
			j--
		} else {
			i++
		}
	}

	return res
}