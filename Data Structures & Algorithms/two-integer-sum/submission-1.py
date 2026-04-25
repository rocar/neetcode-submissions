class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i, n in enumerate(nums):
            for j, s in enumerate(nums[i+1:]):
                if n + s == target:
                    return [i,j+i+1]
        
    