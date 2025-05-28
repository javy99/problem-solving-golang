package containsduplicate

import "testing"

func TestContainsDuplicates(t *testing.T) {
	cases := []struct {
		name     string
		fn       func([]int) bool
		input    []int
		expected bool
	}{
		// ContainsDuplicatesBruteForce
		{"Brute Force - Duplicates", ContainsDuplicatesBruteForce, []int{1, 2, 3, 1}, true},
		{"Brute Force - No Duplicates", ContainsDuplicatesBruteForce, []int{1, 2, 3, 4}, false},
		{"Brute Force - Empty Slice", ContainsDuplicatesBruteForce, []int{}, false},
		{"Brute Force - Single Element", ContainsDuplicatesBruteForce, []int{42}, false},

		// ContainsDuplicatesSort
		{"Sort - Duplicates", ContainsDuplicatesSort, []int{1, 2, 3, 3}, true},
		{"Sort - No Duplicates", ContainsDuplicatesSort, []int{1, 2, 3, 4}, false},
		{"Sort - Empty Slice", ContainsDuplicatesSort, []int{}, false},
		{"Sort - Single Element", ContainsDuplicatesSort, []int{42}, false},

		// ContainsDuplicatesMapBool
		{"MapBool - Duplicates", ContainsDuplicatesMapBool, []int{1, 2, 3, 3}, true},
		{"MapBool - No Duplicates", ContainsDuplicatesMapBool, []int{1, 2, 3, 4}, false},
		{"MapBool - Empty Slice", ContainsDuplicatesMapBool, []int{}, false},
		{"MapBool - Single Element", ContainsDuplicatesMapBool, []int{42}, false},

		// ContainsDuplicatesMapStruct
		{"MapStruct - Duplicates", ContainsDuplicatesMapStruct, []int{1, 2, 3, 3}, true},
		{"MapStruct - No Duplicates", ContainsDuplicatesMapStruct, []int{1, 2, 3, 4}, false},
		{"MapStruct - Empty Slice", ContainsDuplicatesMapStruct, []int{}, false},
		{"MapStruct - Single Element", ContainsDuplicatesMapStruct, []int{42}, false},

		// ContainsDuplicatesMapCount
		{"MapCount - Duplicates", ContainsDuplicatesMapCount, []int{1, 2, 3, 3}, true},
		{"MapCount - No Duplicates", ContainsDuplicatesMapCount, []int{1, 2, 3, 4}, false},
		{"MapCount - Empty Slice", ContainsDuplicatesMapCount, []int{}, false},
		{"MapCount - Single Element", ContainsDuplicatesMapCount, []int{42}, false},
	}

	for _, c := range cases {
		result := c.fn(c.input)
		if result != c.expected {
			t.Errorf("%s failed: expected %v, got %v", c.name, c.expected, result)
		}
	}
}
