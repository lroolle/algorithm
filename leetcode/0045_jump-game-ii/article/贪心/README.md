"+++
title = "0045. Jump Game II 贪心 "
author = ["rPSXFVr7e3"]
date = 2020-08-27T08:49:59+08:00
tags = ["Leetcode", "Algorithms", "Greedy", "Python", "Python3"]
categories = ["leetcode"]
draft = false
+++

# 贪心

> [0045. Jump Game II](https://leetcode-cn.com/problems/jump-game-ii/)
> [贪心](https://leetcode-cn.com/problems/jump-game-ii/solution/tan-xin-by-rpsxfvr7e3-2/) by [rPSXFVr7e3](https://leetcode-cn.com/u/rPSXFVr7e3/)

### 解题思路
此处撰写解题思路
1.start<=end时，说明是在end的一步内，则不断去判断这一步内下一步的最大距离
2.当start大于end时，说明是在end的一步外了，则步数要加一，赋予end下一步的最大距离
3.当end能到达或超过最后位置时，则计算出步数

### 代码

```python
class Solution(object):
    def jump(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        count = 0
        start = 0
        end = 0
        next_end = 0
        while end<len(nums)-1:
            if start<=end:
                next_end = max(next_end,start+nums[start])
                start+=1
            else:
                count+=1
                end = next_end
        return count
```