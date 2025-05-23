package containsduplicate

import "sort"

// Time complexity: O(n^2)
// Space complexity: O(1)
func ContainsDuplicatesBruteForce(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// Time complexity: O(n log n)
// Space complexity: O(1)
func ContainsDuplicatesSorted(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// Time complexity: O(n)
// Space complexity: O(n)
func ContainsDuplicatesMapBool(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

// Time complexity: O(n)
// Space complexity: O(n)
func ContainsDuplicatesMapStruct(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}

// Time complexity: O(n)
// Space complexity: O(n)
func ContainsDuplicatesMapCount(nums []int) bool {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
		if counts[num] > 1 {
			return true
		}
	}
	return false
}
