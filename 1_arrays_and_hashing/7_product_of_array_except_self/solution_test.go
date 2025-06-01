package productofarrayexceptself

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	cases := []struct {
		name     string
		fn       func([]int) []int
		nums     []int
		expected []int
	}{
		// ProductExceptSelfBruteForce
		{"Brute Force - Basic", ProductExceptSelfBruteForce, []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"Brute Force - One Zero", ProductExceptSelfBruteForce, []int{1, 2, 0, 4}, []int{0, 0, 8, 0}},
		{"Brute Force - Two Zeros", ProductExceptSelfBruteForce, []int{0, 2, 0, 4}, []int{0, 0, 0, 0}},
		{"Brute Force - Single", ProductExceptSelfBruteForce, []int{10}, []int{1}},
		{"Brute Force - Same Elements", ProductExceptSelfBruteForce, []int{3, 3, 3}, []int{9, 9, 9}},

		// ProductExceptSelfDivision
		{"Division - Basic", ProductExceptSelfDivision, []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"Division - One Zero", ProductExceptSelfDivision, []int{1, 2, 0, 4}, []int{0, 0, 8, 0}},
		{"Division - Two Zeros", ProductExceptSelfDivision, []int{0, 2, 0, 4}, []int{0, 0, 0, 0}},
		{"Division - Single", ProductExceptSelfDivision, []int{10}, []int{1}},
		{"Division - Same Elements", ProductExceptSelfDivision, []int{3, 3, 3}, []int{9, 9, 9}},

		// ProductExceptSelfPrefixSuffix
		{"Prefix & Suffix - Basic", ProductExceptSelfPrefixSuffix, []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"Prefix & Suffix - One Zero", ProductExceptSelfPrefixSuffix, []int{1, 2, 0, 4}, []int{0, 0, 8, 0}},
		{"Prefix & Suffix - Two Zeros", ProductExceptSelfPrefixSuffix, []int{0, 2, 0, 4}, []int{0, 0, 0, 0}},
		{"Prefix & Suffix - Single", ProductExceptSelfPrefixSuffix, []int{10}, []int{1}},
		{"Prefix & Suffix - Same Elements", ProductExceptSelfPrefixSuffix, []int{3, 3, 3}, []int{9, 9, 9}},

		// ProductExceptSelfPrefixSuffixOptimal
		{"Prefix & Suffix Optimal - Basic", ProductExceptSelfPrefixSuffixOptimal, []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"Prefix & Suffix Optimal - One Zero", ProductExceptSelfPrefixSuffixOptimal, []int{1, 2, 0, 4}, []int{0, 0, 8, 0}},
		{"Prefix & Suffix Optimal - Two Zeros", ProductExceptSelfPrefixSuffixOptimal, []int{0, 2, 0, 4}, []int{0, 0, 0, 0}},
		{"Prefix & Suffix Optimal - Single", ProductExceptSelfPrefixSuffixOptimal, []int{10}, []int{1}},
		{"Prefix & Suffix Optimal - Same Elements", ProductExceptSelfPrefixSuffixOptimal, []int{3, 3, 3}, []int{9, 9, 9}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.fn(c.nums)
			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("%s failed: expected %v, got %v", c.name, c.expected, got)
			}
		})
	}
}
