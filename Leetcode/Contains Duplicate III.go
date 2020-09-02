import "math"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	size := len(nums)
	i := 0
	for i < size {
		n := 1
		for n <= k {
			if i+n >= size {
				break
			}
			if int(math.Abs(float64(nums[i+n]-nums[i]))) <= t {
				return true
			}
			n++
		}
		i++
	}
	return false

}