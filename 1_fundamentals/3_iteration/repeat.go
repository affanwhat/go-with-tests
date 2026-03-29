package iteration

import "strings"

// Repeat a string for n times
func Repeat(s string, n int) string {
	var repeated strings.Builder
	for range n {
		repeated.WriteString(s)
	}

	return repeated.String()
}