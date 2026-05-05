func groupAnagrams(strs []string) [][]string {
    output := make(map[[26]int][]string)

    for _, str := range strs {
        count := [26]int{}
        for _, c := range str {
            count[c-'a']++
        }
        output[count] = append(output[count], str)
    }

    result := make([][]string, 0, len(output))
    for _, group := range output {
        result = append(result, group)
    }

    return result
}
