from typing import List

class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        """
        all strings have to have a length of at least 1
        - find shortest string
        - loop while i < shortest string length
            - get character at index i for each string
            - if all match, increment by 1
            - if any doesn't match break from loop
        """
        common = 0
        shortest_string = self._shortest_string(strs)
        for i in range(len(shortest_string)):
            for j in range(len(strs)):
                if strs[j][i] != shortest_string[i]:
                    return shortest_string[:i]
        return shortest_string
    
    def _shortest_string(self, strs: List[str]) -> str:
        shortest = strs[0]
        for i in range(1, len(strs)):
            if len(strs[i]) < len(shortest):
                shortest = strs[i]
        return shortest


strs: List[str] = ["flower","flow","flight"]
lcp = Solution()
print(lcp.longestCommonPrefix(strs))