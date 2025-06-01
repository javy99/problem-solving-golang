package productofarrayexceptself

// Time complexity: O(n^2), where n is the length of the input array.
// Space complexity: O(n), for the output array.
func ProductExceptSelfBruteForce(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		product := 1
		for j := 0; j < n; j++ {
			if i != j {
				product *= nums[j]
			}
		}
		result[i] = product
	}
	return result
}

// Time complexity: O(n), where n is the length of the input array.
// Space complexity: O(1), for the output array (ignoring the output size).
func ProductExceptSelfDivision(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	totalProduct := 1
	zeroCount := 0

	for _, num := range nums {
		if num != 0 {
			totalProduct *= num
		} else {
			zeroCount++
		}
	}

	for i, num := range nums {
		if zeroCount > 1 {
			result[i] = 0
		} else if zeroCount == 1 {
			if num == 0 {
				result[i] = totalProduct
			} else {
				result[i] = 0
			}
		} else {
			result[i] = totalProduct / num
		}
	}
	return result
}

// Time complexity: O(n), where n is the length of the input array.
// Space complexity: O(n), for the output array.
func ProductExceptSelfPrefixSuffix(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n)
	suffix := make([]int, n)
	result := make([]int, n)

	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	suffix[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		result[i] = prefix[i] * suffix[i]
	}
	return result
}

// Time complexity: O(n), where n is the length of the input array.
// Space complexity: O(1), for the output array (ignoring the output size).
func ProductExceptSelfPrefixSuffixOptimal(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// Prefix pass
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// Suffix pass
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}

	return res
}
