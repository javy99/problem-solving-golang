package sum

import (
	"sort"
)

// Time complexity: O(n^3)
// Space complexity: O(n)
func ThreeSumBruteForce(nums []int) [][]int {
	n := len(nums)
	result := map[[3]int]bool{}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					triplet := []int{nums[i], nums[j], nums[k]}
					sort.Ints(triplet) // to avoid duplicate triplets in different orders
					var key [3]int
					copy(key[:], triplet)
					result[key] = true
				}
			}
		}
	}

	// convert map to slice - initialize empty slice to avoid nil
	output := make([][]int, 0)
	for k := range result {
		output = append(output, []int{k[0], k[1], k[2]})
	}
	return output
}

// Time complexity: O(n^2)
// Space complexity: O(n)
func ThreeSumHashMap(nums []int) [][]int {
	n := len(nums)
	result := map[[3]int]bool{}

	for i := 0; i < n; i++ {
		target := -nums[i]
		seen := map[int]bool{}
		for j := i + 1; j < n; j++ {
			complement := target - nums[j]
			if seen[complement] {
				triplet := []int{nums[i], nums[j], complement}
				sort.Ints(triplet)
				var key [3]int
				copy(key[:], triplet)
				result[key] = true
			}
			seen[nums[j]] = true
		}
	}

	// convert map to slice - initialize empty slice to avoid nil
	output := make([][]int, 0)
	for k := range result {
		output = append(output, []int{k[0], k[1], k[2]})
	}
	return output
}

// Time complexity: O(n^2)
// Space complexity: O(1)
func ThreeSumTwoPointers(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0) // Initialize empty slice to avoid nil

	for i := 0; i < n-2; i++ {
		// Skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
