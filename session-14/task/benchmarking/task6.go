package benchmarking

import "strings"

func ConcatenateWithPlus(strs []string) string {
	result := ""
	for _, s := range strs {
		result += s
	}
	return result
}

func ConcatenateWithJoin(strs []string) string {
	return strings.Join(strs, "")
}
