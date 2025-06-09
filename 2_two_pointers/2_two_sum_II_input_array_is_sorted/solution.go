package twosumiiinputarrayissorted

// Time complexity: O(n^2)
// Space complexity: O(1)
func TwoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return nil
}

// Time complexity: O(n * log n)
// Space complexity: O(1)
func TwoSumBinarySearch(nums []int, target int) []int {
	for i := range nums {
		left, right := i+1, len(nums)-1
		needed := target - nums[i]
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] == needed {
				return []int{i + 1, mid + 1}
			} else if nums[mid] < needed {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return nil
}

// Time complexity: O(n)
// Space complexity: O(n)
func TwoSumHashMap(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		if j, ok := seen[target-num]; ok {
			return []int{i + 1, j + 1}
		}
		seen[num] = i
	}
	return nil
}

// Time complexity: O(n)
// Space complexity: O(1)
func TwoSumTwoPointers(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
