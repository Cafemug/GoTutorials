import "strings"

func repeatedSubstringPattern(s string) bool {
	size := len(s)
	for i := 1; i < size/2+1; i++ {
		if size%i == 0 && s == strings.Repeat(s[:i], size/i) {
			return true
		}
	}
	return false
}