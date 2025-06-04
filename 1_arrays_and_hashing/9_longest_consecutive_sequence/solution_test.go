package longestconsecutivesequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	cases := []struct {
		name     string
		fn       func([]int) int
		nums     []int
		expected int
	}{
		// LongestConsecutiveBruteForce1
		{"Brute Force - Example 1", LongestConsecutiveBruteForce1, []int{2, 20, 4, 10, 3, 4, 5}, 4},
		{"Brute Force - Example 2", LongestConsecutiveBruteForce1, []int{0, 3, 2, 5, 4, 6, 1, 1}, 7},
		{"Brute Force - Empty", LongestConsecutiveBruteForce1, []int{}, 0},
		{"Brute Force - Single Element", LongestConsecutiveBruteForce1, []int{42}, 1},
		{"Brute Force - All Same", LongestConsecutiveBruteForce1, []int{7, 7, 7}, 1},

		// LongestConsecutiveBruteForce2
		{"Brute Force - Example 1", LongestConsecutiveBruteForce2, []int{2, 20, 4, 10, 3, 4, 5}, 4},
		{"Brute Force - Example 2", LongestConsecutiveBruteForce2, []int{0, 3, 2, 5, 4, 6, 1, 1}, 7},
		{"Brute Force - Empty", LongestConsecutiveBruteForce2, []int{}, 0},
		{"Brute Force - Single Element", LongestConsecutiveBruteForce2, []int{42}, 1},
		{"Brute Force - All Same", LongestConsecutiveBruteForce2, []int{7, 7, 7}, 1},

		// LongestConsecutiveSort
		{"Sort - Example 1", LongestConsecutiveSort, []int{2, 20, 4, 10, 3, 4, 5}, 4},
		{"Sort - Example 2", LongestConsecutiveSort, []int{0, 3, 2, 5, 4, 6, 1, 1}, 7},
		{"Sort - Empty", LongestConsecutiveSort, []int{}, 0},
		{"Sort - Single Element", LongestConsecutiveSort, []int{42}, 1},
		{"Sort - All Same", LongestConsecutiveSort, []int{7, 7, 7}, 1},

		// LongestConsecutiveHashSet
		{"Hash Set - Example 1", LongestConsecutiveHashSet, []int{2, 20, 4, 10, 3, 4, 5}, 4},
		{"Hash Set - Example 2", LongestConsecutiveHashSet, []int{0, 3, 2, 5, 4, 6, 1, 1}, 7},
		{"Hash Set - Empty", LongestConsecutiveHashSet, []int{}, 0},
		{"Hash Set - Single Element", LongestConsecutiveHashSet, []int{42}, 1},
		{"Hash Set - All Same", LongestConsecutiveHashSet, []int{7, 7, 7}, 1},

		// LongestConsecutiveHashMap
		{"Hash Map - Example 1", LongestConsecutiveHashMap, []int{2, 20, 4, 10, 3, 4, 5}, 4},
		{"Hash Map - Example 2", LongestConsecutiveHashMap, []int{0, 3, 2, 5, 4, 6, 1, 1}, 7},
		{"Hash Map - Empty", LongestConsecutiveHashMap, []int{}, 0},
		{"Hash Map - Single Element", LongestConsecutiveHashMap, []int{42}, 1},
		{"Hash Map - All Same", LongestConsecutiveHashMap, []int{7, 7, 7}, 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.fn(c.nums)
			if got != c.expected {
				t.Errorf("%s failed: expected %d, got %d", c.name, c.expected, got)
			}
		})
	}
}
