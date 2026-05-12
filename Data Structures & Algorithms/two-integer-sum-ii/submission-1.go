func twoSum(numbers []int, target int) []int {
	expecting := make(map[int]int)
	
	for i, num := range numbers {
		if index1, ok := expecting[num]; ok {
			return []int{index1, i+1}
		} else {
			expecting[target-num] = i+1
		}
	}
	return []int{}
}
