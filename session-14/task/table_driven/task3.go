package table_driven

func ReverseString(s string) string {
	runes := []rune(s)
	left := 0
	right := len(runes) - 1
	for left < right {
		temp := runes[left]
		runes[left] = runes[right]
		runes[right] = temp
		left++
		right--
	}
	return string(runes)
}
