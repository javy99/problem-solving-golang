package containerwithmostwater

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		name     string
		fn       func([]int) int
		heights  []int
		expected int
	}{
		// MaxAreaBruteForce
		{"BruteForce - Example 1", MaxAreaBruteForce, []int{1, 7, 2, 5, 4, 7, 3, 6}, 36},
		{"BruteForce - Example 2", MaxAreaBruteForce, []int{2, 2, 2}, 4},
		{"BruteForce - Minimal input size 2", MaxAreaBruteForce, []int{1, 1}, 1},
		{"BruteForce - All zeroes", MaxAreaBruteForce, []int{0, 0, 0, 0}, 0},
		{"BruteForce - Increasing heights", MaxAreaBruteForce, []int{1, 2, 3, 4, 5}, 6},
		{"BruteForce - Decreasing heights", MaxAreaBruteForce, []int{5, 4, 3, 2, 1}, 6},

		// MaxAreaTwoPointers
		{"TwoPointers - Example 1", MaxAreaTwoPointers, []int{1, 7, 2, 5, 4, 7, 3, 6}, 36},
		{"TwoPointers - Example 2", MaxAreaTwoPointers, []int{2, 2, 2}, 4},
		{"TwoPointers - Minimal input size 2", MaxAreaTwoPointers, []int{1, 1}, 1},
		{"TwoPointers - All zeroes", MaxAreaTwoPointers, []int{0, 0, 0, 0}, 0},
		{"TwoPointers - Increasing heights", MaxAreaTwoPointers, []int{1, 2, 3, 4, 5}, 6},
		{"TwoPointers - Decreasing heights", MaxAreaTwoPointers, []int{5, 4, 3, 2, 1}, 6},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.fn(c.heights)
			if got != c.expected {
				t.Errorf("%s failed: expected %v, got %v", c.name, c.expected, got)
			}
		})
	}
}
