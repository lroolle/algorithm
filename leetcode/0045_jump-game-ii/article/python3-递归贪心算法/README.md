"+++
title = "0045. Jump Game II python3 递归+贪心算法 "
author = ["ting-ting-28"]
date = 2020-08-24T03:26:56+08:00
tags = ["Leetcode", "Algorithms", "Python", "Python3", "Greedy", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# python3 递归+贪心算法

> [0045. Jump Game II](https://leetcode-cn.com/problems/jump-game-ii/)
> [python3 递归+贪心算法](https://leetcode-cn.com/problems/jump-game-ii/solution/python3-di-gui-tan-xin-suan-fa-by-ting-ting-28/) by [ting-ting-28](https://leetcode-cn.com/u/ting-ting-28/)

# 贪心算法+递归
每一步都取得最优就可以了，不过需要牵连到下一次，所以计算最大值是需要计算下一步，最后递归。

```python3
class Solution:
    def jump(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return 0
        self.ret = 0
        self.nums = nums
        self.n = len(nums)
        self._jump(0)
        return self.ret

    def _jump(self, inx: int):
        if inx >= self.n-1:
            return
        tmp_list = []
        for i in range(inx+1, self.nums[inx]+inx+1):
            if i >= len(self.nums):
                self.ret += 1
                return
            tmp_list.append([i+self.nums[i], i])
        self.ret += 1
        if not tmp_list:
            return
        if len(self.nums)-1 == tmp_list[-1][-1]:
            return
        t = max(tmp_list, key=lambda x:x[0])[1]
        return self._jump(t)
```