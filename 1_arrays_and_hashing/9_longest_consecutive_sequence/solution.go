package longestconsecutivesequence

import (
	"sort"
)

// Time complexity: O(n^2)
// Space complexity: O(1)
func LongestConsecutiveBruteForce1(nums []int) int {
	maxLen := 0

	for _, num := range nums {
		current := num
		length := 1

		for {
			found := false
			for _, n := range nums {
				if n == current+1 {
					found = true
					current++
					length++
					break
				}
			}
			if !found {
				break
			}
		}

		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}

// Time complexity: O(n^2)
// Space complexity: O(n)
func LongestConsecutiveBruteForce2(nums []int) int {
	res := 0
	store := make(map[int]struct{})
	for _, num := range nums {
		store[num] = struct{}{}
	}

	for _, num := range nums {
		streak, curr := 0, num
		for _, ok := store[curr]; ok; _, ok = store[curr] {
			streak++
			curr++
		}
		if streak > res {
			res = streak
		}
	}
	return res
}

// Time complexity: O(n log n) due to sorting
// Space complexity: O(1) or O(n) depending on sorting algorithm implementation
func LongestConsecutiveSort(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	maxLen := 1
	currLen := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue // skip duplicates
		}
		if nums[i] == nums[i-1]+1 {
			currLen++
		} else {
			currLen = 1
		}
		if currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}

// Time complexity: O(n)
// Space complexity: O(n)
func LongestConsecutiveHashSet(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	maxLen := 0
	for num := range set {
		// Only check starts of sequences
		if !set[num-1] {
			current := num
			length := 1

			for set[current+1] {
				current++
				length++
			}
			if length > maxLen {
				maxLen = length
			}
		}
	}
	return maxLen
}

// Time complexity: O(n)
// Space complexity: O(n)
func LongestConsecutiveHashMap(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lookup := make(map[int]int)
	maxLen := 0

	for _, num := range nums {
		if _, exists := lookup[num]; !exists {
			left := lookup[num-1]
			right := lookup[num+1]
			length := left + right + 1

			lookup[num] = length
			lookup[num-left] = length
			lookup[num+right] = length

			if length > maxLen {
				maxLen = length
			}
		}
	}
	return maxLen
}
