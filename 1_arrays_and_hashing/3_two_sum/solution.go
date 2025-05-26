package twosum

import "sort"

// Time complexity: O(n^2)
// Space complexity: O(1)
func TwoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

type Pair struct {
	val int
	idx int
}

// Time complexity: O(n log n)
// Space complexity: O(n)
func TwoSumSort(nums []int, target int) []int {
	pairs := make([]Pair, len(nums))
	for i, v := range nums {
		pairs[i] = Pair{val: v, idx: i}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val < pairs[j].val
	})

	left, right := 0, len(pairs)-1
	for left < right {
		sum := pairs[left].val + pairs[right].val
		if sum == target {
			return []int{pairs[left].idx, pairs[right].idx}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

// Time complexity: O(n)
// Space complexity: O(n)
func TwoSumMapTwoPass(nums []int, target int) []int {
	// First pass: build map
	m := make(map[int]int) // value -> index
	for i, num := range nums {
		m[num] = i
	}

	// Second pass: check complements
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}

// Time complexity: O(n)
// Space complexity: O(n)
func TwoSumMapOnePass(nums []int, target int) []int {
	m := make(map[int]int) // value -> index
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}
