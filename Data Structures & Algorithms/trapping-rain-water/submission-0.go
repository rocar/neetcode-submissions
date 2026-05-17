func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }

    l, r := 0, len(height)-1
    leftMax, rightMax := height[l], height[r]
    res := 0

    for l < r {
        // We move the pointer pointing to the smaller maximum height
        if leftMax < rightMax {
            l++
            // Update leftMax to the current height (or keep the old max)
            leftMax = max(leftMax, height[l])
            // Water trapped is the difference between max wall and current floor
            res += leftMax - height[l]
        } else {
            r--
            // Update rightMax
            rightMax = max(rightMax, height[r])
            // Water trapped is the difference
            res += rightMax - height[r]
        }
    }

    return res
}

// Helper function to handle integer max
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}