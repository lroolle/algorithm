"+++
title = "0163. Missing Ranges 97+100 将lower,upper融合到数组中去，一次遍历 "
author = ["mnm135"]
date = 2020-05-13T15:59:16+08:00
tags = ["Leetcode", "Algorithms", "Python", "Python3", "Array", "数组结构"]
categories = ["leetcode"]
draft = false
+++

# 97+100 将lower,upper融合到数组中去，一次遍历

> [0163. Missing Ranges](https://leetcode-cn.com/problems/missing-ranges/)
> [97+100 将lower,upper融合到数组中去，一次遍历](https://leetcode-cn.com/problems/missing-ranges/solution/97100-jiang-lowerupperrong-he-dao-shu-zu-zhong-qu-/) by [mnm135](https://leetcode-cn.com/u/mnm135/)

### 解题思路
此处撰写解题思路

### 代码

```python
class Solution(object):
    def findMissingRanges(self, nums, lower, upper):
        """
        :type nums: List[int]
        :type lower: int
        :type upper: int
        :rtype: List[str]
        """
        r = []
        nums = [lower - 1] + nums + [upper + 1]
        for i in range(1, len(nums)):
            if nums[i] - nums[i - 1] > 2:
                r.append(str(nums[i - 1] + 1) + '->' + str(nums[i] - 1))
            elif nums[i] - nums[i - 1] == 2:
                r.append(str(nums[i] - 1))
        return r
            

```