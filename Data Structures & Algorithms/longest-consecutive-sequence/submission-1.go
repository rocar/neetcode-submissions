func longestConsecutive(nums []int) int {
	numSet := make(map[int]int)
	for _, num := range nums {
		numSet[num] = 1
	}


	for n, _ := range numSet {
		if _, exist := numSet[n - 1]; !exist {
			next := n + 1
			for i := 0; i < len(numSet); i++ {
				if _, ok := numSet[next]; ok {
					numSet[n]++
					next++
				}
			}
		}
	}

	max := 0
	for _, v := range numSet {
		if v > max {
			max = v
		}
	}
	return max
}
