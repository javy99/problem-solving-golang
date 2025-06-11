package containerwithmostwater

// Time complexity: O(n^2)
// Space complexity: O(1)
func MaxAreaBruteForce(heights []int) int {
	n := len(heights)
	maxArea := 0

	for p1 := 0; p1 < n; p1++ {
		for p2 := p1 + 1; p2 < n; p2++ {
			height := heights[p1]
			if heights[p2] < height {
				height = heights[p2]
			}
			width := p2 - p1
			area := height * width

			if area > maxArea {
				maxArea = area
			}

			// height := min(heights[p1], heights[p2])
			// width := p2 - p1
			// area := height * width
			// maxArea = max(maxArea, area)
		}
	}
	return maxArea
}

// Time complexity: O(n)
// Space complexity: O(1)
func MaxAreaTwoPointers(heights []int) int {
	p1, p2, maxArea := 0, len(heights)-1, 0

	for p1 < p2 {
		height := heights[p1]
		if heights[p2] < height {
			height = heights[p2]
		}
		width := p2 - p1
		area := height * width

		if area > maxArea {
			maxArea = area
		}

		// height := min(heights[p1], heights[p2])
		// width := p2 - p1
		// area := height * width
		// maxArea = max(maxArea, area)

		if heights[p1] < heights[p2] {
			p1++
		} else {
			p2--
		}
	}
	return maxArea
}
