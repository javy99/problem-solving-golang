package topkfrequentelements

import (
	"reflect"
	"sort"
	"testing"
)

func sortUnorderedResult(result []int) []int {
	sort.Ints(result)
	return result
}

func containsAll(target, actual []int) bool {
	set := make(map[int]bool)
	for _, v := range target {
		set[v] = true
	}
	for _, v := range actual {
		if !set[v] {
			return false
		}
	}
	return true
}

func TestTopKFrequent(t *testing.T) {
	cases := []struct {
		name       string
		fn         func([]int, int) []int
		nums       []int
		k          int
		expected   []int
		anyOfThese bool // if true, result can contain any subset of expected
	}{
		// TopKFrequentSort
		{"Sort - Basic", TopKFrequentSort, []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}, false},
		{"Sort - Single", TopKFrequentSort, []int{1}, 1, []int{1}, false},
		{"Sort - All Same Freq", TopKFrequentSort, []int{4, 4, 6, 6, 7, 7}, 2, []int{4, 6, 7}, true},

		// TopKFrequentHeap
		{"Heap - Basic", TopKFrequentHeap, []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}, false},
		{"Heap - Single", TopKFrequentHeap, []int{1}, 1, []int{1}, false},
		{"Heap - All Same Freq", TopKFrequentHeap, []int{4, 4, 6, 6, 7, 7}, 2, []int{4, 6, 7}, true},

		// TopKFrequentBucket
		{"Bucket - Basic", TopKFrequentBucket, []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}, false},
		{"Bucket - Single", TopKFrequentBucket, []int{1}, 1, []int{1}, false},
		{"Bucket - All Same Freq", TopKFrequentBucket, []int{4, 4, 6, 6, 7, 7}, 2, []int{4, 6, 7}, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.fn(c.nums, c.k)
			sortedGot := sortUnorderedResult(got)
			if c.anyOfThese {
				if len(got) != c.k || !containsAll(c.expected, got) {
					t.Errorf("%s failed: expected any %d of %v, got %v", c.name, c.k, c.expected, sortedGot)
				}
			} else {
				if !reflect.DeepEqual(sortedGot, sortUnorderedResult(c.expected)) {
					t.Errorf("%s failed: expected %v, got %v", c.name, c.expected, sortedGot)
				}
			}
		})
	}
}
