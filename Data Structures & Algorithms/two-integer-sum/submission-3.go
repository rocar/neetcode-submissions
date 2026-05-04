func twoSum(nums []int, target int) []int {
    diffs := make(map[int]int)
    diff := 0
    for i, num := range nums {
        diff = target - num
        if _, ok := diffs[num]; ok {
            return []int{diffs[num], i}
        }
        diffs[diff] = i
    }
    return []int{}
}
