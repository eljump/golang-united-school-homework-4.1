package reverse_string

func ReverseString(input string) (output string) {
	sliceInput := []rune(input)

	for i, j := 0, len(sliceInput)-1; i < j; i, j = i+1, j-1 {
		sliceInput[i], sliceInput[j] = sliceInput[j], sliceInput[i]
	}
	output = string(sliceInput)
	return output
}
