package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		name     string
		fn       func(string, string) bool
		s, t     string
		expected bool
	}{
		// IsAnagramSort
		{"Sort - Anagram", IsAnagramSort, "listen", "silent", true},
		{"Sort - Not Anagram", IsAnagramSort, "hello", "world", false},
		{"Sort - Empty Strings", IsAnagramSort, "", "", true},
		{"Sort - Different Lengths", IsAnagramSort, "a", "ab", false},
		{"Sort - Unicode Anagrams", IsAnagramSort, "école", "éocle", true},

		// IsAnagramMap
		{"Map - Anagram", IsAnagramMap, "listen", "silent", true},
		{"Map - Not Anagram", IsAnagramMap, "hello", "world", false},
		{"Map - Empty Strings", IsAnagramMap, "", "", true},
		{"Map - Different Lengths", IsAnagramMap, "a", "ab", false},
		{"Map - Unicode Anagrams", IsAnagramMap, "école", "éocle", true},

		// IsAnagramArray
		{"Array - Anagram", IsAnagramArray, "listen", "silent", true},
		{"Array - Not Anagram", IsAnagramArray, "hello", "world", false},
		{"Array - Empty Strings", IsAnagramArray, "", "", true},
		{"Array - Different Lengths", IsAnagramArray, "a", "ab", false},
		{"Array - Unicode Anagrams", IsAnagramArray, "école", "éocle", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.fn(c.s, c.t)
			if got != c.expected {
				t.Errorf("%s failed: expected %v, got %v", c.name, c.expected, got)

			}
		})
	}
}
