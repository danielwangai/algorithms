class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        
        countS = self._count_letters(s)
        countT = self._count_letters(t)
        for i in s:
            if i in countS.keys() and i in countT.keys():
                if countS[i] != countT[i]:
                    return False
            else:
                return False
        
        return True
        
    def _count_letters(self, s: str):
        count = {}
        for i in s:
            if i not in count:
                count[i] = 1
            else:
                count[i] += 1
        
        return count

sol = Solution()
print(sol.isAnagram("rat", "car"))
# caar, crra