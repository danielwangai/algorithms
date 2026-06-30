from typing import List

class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        """
        delete in place
        write starts at 0
        loop through nums list
            compare number at current index(i) with val(what we want to delete)
                if not same, override number at index write with number at index i
                increment write by 1
        """
        write = 0 # index containting number we want to "delete"
        for i in range(len(nums)):
            if nums[i] != val: # value to remove
                nums[write] = nums[i]
                write += 1
        # print(nums)
        return write

nums = [3,2,2,3]
lcp = Solution()
print(lcp.removeElement(nums, 3))