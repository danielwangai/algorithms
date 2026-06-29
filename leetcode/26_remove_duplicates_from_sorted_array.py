from typing import List

class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        write = 1
        for i in range(1, len(nums)):
            if nums[i] != nums[i-1]:
                nums[write] = nums[i]
                write += 1
        return write

    """
    def _removeDuplicates_not_in_place(self, nums: List[int]) -> int:
        len_nums = len(nums)
        count = 0
        for i in range(1, len(nums)-1):
            if nums[i] == nums[i-1]:
                # delete num at position i
                # append -1 at end of list
                # nums.insert(len_nums - 1, nums.pop(i))
                del nums[i]
                count += 1
            else:
                continue
        return len_nums - count
    """

nums = [1,1,2]
lcp = Solution()
print(lcp.removeDuplicates(nums))