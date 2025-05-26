package groupanagrams

import (
	"sort"
	"strings"
)

// Time complexity: O(n * k log k), where n is the number of strings and k is the maximum length of a string.
// Space complexity: O(n * k)
func GroupAnagramsSort(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	anagramMap := make(map[string][]string)

	// for _, str := range strs {
	// 	// Convert to a slice of runes to sort
	// 	runes := []rune(str)
	// 	sort.Slice(runes, func(i, j int) bool {
	// 		return runes[i] < runes[j]
	// 	})
	// 	key := string(runes)
	// 	anagramMap[key] = append(anagramMap[key], str)
	// }

	for _, str := range strs {
		sChars := strings.Split(str, "")
		sort.Strings(sChars)
		sortedStr := strings.Join(sChars, "")

		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	result := make([][]string, 0, len(anagramMap)) // result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

// Time complexity: O(n * k)
// Space complexity: O(n * k)
func GroupAnagramsCount(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	anagramMap := make(map[[26]int][]string)

	for _, str := range strs {
		var key [26]int
		for _, char := range str {
			if char >= 'a' && char <= 'z' {
				key[char-'a']++
			}
		}
		anagramMap[key] = append(anagramMap[key], str)
	}

	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}
