package twosumiiinputarrayissorted

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
		{"Brute Force - Basic", TwoSumBruteForce, []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"Brute Force - Not Found", TwoSumBruteForce, []int{1, 2, 3}, 10, nil},
		{"Brute Force - Single", TwoSumBruteForce, []int{5}, 5, nil},
		{"Brute Force - Negative Numbers", TwoSumBruteForce, []int{-3, 0, 3, 4}, 0, []int{1, 3}},

		// TwoSumBinarySearch
		{"Binary Search - Basic", TwoSumBinarySearch, []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"Binary Search - Not Found", TwoSumBinarySearch, []int{1, 2, 3}, 6, nil},
		{"Binary Search - Negative Numbers", TwoSumBinarySearch, []int{-4, -1, 1, 2, 5}, 1, []int{1, 5}},

		// TwoSumHashMap
		{"Hash Map - Basic", TwoSumHashMap, []int{2, 7, 11, 15}, 9, []int{2, 1}},
		{"Hash Map - Not Found", TwoSumHashMap, []int{1, 2, 3}, 0, nil},
		{"Hash Map - Negative", TwoSumHashMap, []int{-3, 4, 3, 90}, 0, []int{3, 1}},

		// TwoSumTwoPointers
		{"Two Pointers - Basic", TwoSumTwoPointers, []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"Two Pointers - Not Found", TwoSumTwoPointers, []int{1, 2, 3}, 7, nil},
		{"Two Pointers - Edge", TwoSumTwoPointers, []int{-10, -3, 0, 1, 6}, -3, []int{2, 3}},
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
