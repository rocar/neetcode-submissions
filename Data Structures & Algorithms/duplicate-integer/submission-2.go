func hasDuplicate(nums []int) bool {
	hashMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		if hashMap[num] {
			return true
		}
		hashMap[num] = true
	}
	return false
}
