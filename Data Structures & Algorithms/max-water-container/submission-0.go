func maxArea(heights []int) int {
	left, right := 0, len(heights) - 1
	maxWater := 0

	for left < right {
		width := right - left
		currentHeight := min(heights[left], heights[right])
		currentWater := width * currentHeight

		maxWater = max(maxWater, currentWater)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a> b {
		return a
	}
	return b
}
