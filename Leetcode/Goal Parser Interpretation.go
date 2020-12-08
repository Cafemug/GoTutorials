import "strings"

func interpret(command string) string {
	replacer := strings.NewReplacer("()", "o", "(al)", "al")
	return replacer.Replace(command)
}