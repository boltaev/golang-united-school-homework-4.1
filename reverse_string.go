package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		output += string(runes[i])
	}
	return output
}
