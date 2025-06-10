package sum

import (
	"reflect"
	"sort"
	"testing"
)

func sortTriplets(triplets [][]int) {
	for _, triplet := range triplets {
		sort.Ints(triplet)
	}
	sort.Slice(triplets, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if triplets[i][k] != triplets[j][k] {
				return triplets[i][k] < triplets[j][k]
			}
		}
		return false
	})
}

func TestThreeSum(t *testing.T) {
	cases := []struct {
		name     string
		fn       func([]int) [][]int
		nums     []int
		expected [][]int
	}{
		// ThreeSumBruteForce
		{"Brute Force - Typical case with multiple triplets", ThreeSumBruteForce, []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"Brute Force - All zeroes", ThreeSumBruteForce, []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"Brute Force - No valid triplet", ThreeSumBruteForce, []int{1, 2, -2, -1}, [][]int{}},
		{"Brute Force - Empty array", ThreeSumBruteForce, []int{}, [][]int{}},
		{"Brute Force - Single element", ThreeSumBruteForce, []int{1}, [][]int{}},
		{"Brute Force - Two elements", ThreeSumBruteForce, []int{1, 2}, [][]int{}},

		// HashMap
		{"Hash Map - Typical case with multiple triplets", ThreeSumHashMap, []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"Hash Map - All zeroes", ThreeSumHashMap, []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"Hash Map - No valid triplet", ThreeSumHashMap, []int{1, 2, -2, -1}, [][]int{}},
		{"Hash Map - Empty array", ThreeSumHashMap, []int{}, [][]int{}},
		{"Hash Map - Single element", ThreeSumHashMap, []int{1}, [][]int{}},
		{"Hash Map - Two elements", ThreeSumHashMap, []int{1, 2}, [][]int{}},

		// TwoPointers
		{"Two Pointers - Typical case with multiple triplets", ThreeSumTwoPointers, []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"Two Pointers - All zeroes", ThreeSumTwoPointers, []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"Two Pointers - No valid triplet", ThreeSumTwoPointers, []int{1, 2, -2, -1}, [][]int{}},
		{"Two Pointers - Empty array", ThreeSumTwoPointers, []int{}, [][]int{}},
		{"Two Pointers - Single element", ThreeSumTwoPointers, []int{1}, [][]int{}},
		{"Two Pointers - Two elements", ThreeSumTwoPointers, []int{1, 2}, [][]int{}},
		{"Two Pointers - Large numbers", ThreeSumTwoPointers, []int{-4, -1, -1, 0, 1, 2}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"Two Pointers - Duplicates with valid triplet", ThreeSumTwoPointers, []int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.fn(c.nums)
			sortTriplets(got)
			sortTriplets(c.expected)
			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("%s failed: expected %v, got %v", c.name, c.expected, got)
			}
		})
	}
}
