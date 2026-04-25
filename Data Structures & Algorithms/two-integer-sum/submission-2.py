class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen = {}
        for index, n in enumerate(nums):
            diff = target - n
            if diff in seen.keys():
                return [seen[diff], index]
            seen[n] = index

        
    