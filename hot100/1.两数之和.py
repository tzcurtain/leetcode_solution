#
# @lc app=leetcode.cn id=1 lang=python3
#
# [1] 两数之和
#

# @lc code=start
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        m = {}
        now = 0
        for num in nums:
            if m.get(target - num, None) is not None:
                return [now, m[target-num]]
            m[num] = now
            now += 1


        
# @lc code=end

