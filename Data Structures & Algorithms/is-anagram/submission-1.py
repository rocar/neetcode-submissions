from collections import defaultdict
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        ds = {}
        dt = {}
        for c in s:
            if c not in ds:
                ds[c] = 1
            else:
                ds[c] = ds[c] + 1
        for c in t:
            if c not in dt:
                dt[c] = 1
            else:
                dt[c] = dt[c] + 1

        return ds == dt