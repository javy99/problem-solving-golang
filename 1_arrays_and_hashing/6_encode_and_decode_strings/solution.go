package encodeanddecodestrings

import (
	"strconv"
	"strings"
)

// Time complexity: O(n), where n is the total length of all strings in the input slice.
// Space complexity: O(n)
func Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteString("#")
		sb.WriteString(str)
	}
	return sb.String()
}

// Time complexity: O(n), where n is the total length of the encoded string.
// Space complexity: O(n)
func Decode(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	var result []string
	i := 0
	for i < len(s) {
		j := i
		for s[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(s[i:j])
		j++
		str := s[j : j+length]
		result = append(result, str)
		i = j + length
	}
	return result
}
