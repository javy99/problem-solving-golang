package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		name     string
		fn       func([]int, int) []int
		nums     []int
		target   int
		expected []int
	}{
		// TwoSumBruteForce
		{"Brute Force - Valid Pair", TwoSumBruteForce, []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Brute Force - No Pair", TwoSumBruteForce, []int{1, 2, 3}, 7, nil},
		{"Brute - Empty", TwoSumBruteForce, []int{}, 0, nil},
		{"Brute - Single Element", TwoSumBruteForce, []int{5}, 5, nil},
		{"Brute - Negative Numbers", TwoSumBruteForce, []int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{"Brute - Duplicates", TwoSumBruteForce, []int{3, 3}, 6, []int{0, 1}},

		// Sorting (Two Pointer)
		{"Sorted - Valid Pair", TwoSumSorted, []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Sorted - No Pair", TwoSumSorted, []int{1, 2, 3}, 7, nil},
		{"Sorted - Empty", TwoSumSorted, []int{}, 0, nil},
		{"Sorted - Single Element", TwoSumSorted, []int{5}, 5, nil},
		{"Sorted - Negative Numbers", TwoSumSorted, []int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{"Sorted - Duplicates", TwoSumSorted, []int{3, 3}, 6, []int{0, 1}},

		// Map Two-Pass
		{"MapTwoPass - Valid Pair", TwoSumMapTwoPass, []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"MapTwoPass - No Pair", TwoSumMapTwoPass, []int{1, 2, 3}, 7, nil},
		{"MapTwoPass - Empty", TwoSumMapTwoPass, []int{}, 0, nil},
		{"MapTwoPass - Single", TwoSumMapTwoPass, []int{5}, 5, nil},
		{"MapTwoPass - Negative", TwoSumMapTwoPass, []int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{"MapTwoPass - Duplicates", TwoSumMapTwoPass, []int{3, 3}, 6, []int{0, 1}},

		// Map One-Pass
		{"MapOnePass - Valid Pair", TwoSumMapOnePass, []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"MapOnePass - No Pair", TwoSumMapOnePass, []int{1, 2, 3}, 7, nil},
		{"MapOnePass - Empty", TwoSumMapOnePass, []int{}, 0, nil},
		{"MapOnePass - Single", TwoSumMapOnePass, []int{5}, 5, nil},
		{"MapOnePass - Negative", TwoSumMapOnePass, []int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{"MapOnePass - Duplicates", TwoSumMapOnePass, []int{3, 3}, 6, []int{0, 1}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.fn(c.nums, c.target)
			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("%s failed: expected %v, got %v", c.name, c.expected, got)
			}
		})
	}
}
