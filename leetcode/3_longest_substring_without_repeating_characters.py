class Solution:
    # def lengthOfLongestSubstring(self, s: str) -> int:
    #     """
    #     if len(s) == 0:
    #         return 0
    #     count = 1
    #     log = {}# store first character at position count(1) before loop begins
    #     prev = first character
    #     longest_found = []
    #     loop through s from i = 1:
    #         check if s[i] is in log:
    #             longest_found.push(count)
    #             count = 1
    #             # clear the log
    #             log[s[i]] = s[i]
    #             # continue
    #         if s[i] != prev:
    #             count +=1
    #             log[str[i]] = s[i]
    #         prev = s[i]
        
    #     return max(longest_count)
        
    #     return count
    #     """

    #     if len(s) == 0:
    #         return 0
        
    #     count = 1
    #     log = {}
    #     prev = s[0]
    #     log[prev] = prev
    #     longest_count = []

    #     for i in range(1, len(s)):
    #         if s[i] in log:
    #             longest_count.append(count)
    #             count = 1
    #             log.clear()
    #             log[s[i]] = s[i]
    #             prev = s[i]
    #             continue
    #         if s[i] != prev:
    #             count += 1
    #             log[s[i]] = s[i]
    #             prev = s[i]
        
    #     if len(longest_count) > 0:
    #         return max(longest_count)
        
    #     return len(log.keys())

    def lengthOfLongestSubstring(self, s: str) -> int:
        log = set([])
        longest_str = 0
        left = 0
        right = 0
        while right < len(s):
            # current character in string s
            c = s[right]
            # check if c in set
            if c in log:
                log.remove(s[left])
                left += 1
            else:
                log.add(c)
                longest_str = max(longest_str, len(log))
                right += 1
        
        return longest_str

sol = Solution()
s = "aab"
print(sol.lengthOfLongestSubstring(s))
