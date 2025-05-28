package validanagram

import (
	"sort"
	"strings"
)

// Time complexity: O(n * log n)
// Space complexity : O(n)
func IsAnagramSort(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sArr := strings.Split(s, "")
	tArr := strings.Split(t, "")

	sort.Strings(sArr)
	sort.Strings(tArr)

	return strings.Join(sArr, "") == strings.Join(tArr, "")
}

// Time complexity: O(n)
// Space complexity: O(n)
func IsAnagramMap(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	for _, r := range t {
		counts[r]--
		if counts[r] < 0 {
			return false
		}
	}
	return true
}

// Time complexity: O(n)
// Space complexity: O(1)
func IsAnagramArray(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var counts [26]int

	for i := 0; i < len(s); i++ {
		// we ensure the character is between 'a' and 'z'
		if s[i] < 'a' || s[i] > 'z' || t[i] < 'a' || t[i] > 'z' {
			return false
		}

		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}
	return true
}
