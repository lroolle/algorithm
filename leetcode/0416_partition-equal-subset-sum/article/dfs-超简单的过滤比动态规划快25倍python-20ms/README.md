"+++
title = "0416. Partition Equal Subset Sum DFS + 超简单的过滤，比动态规划快25倍（python 20ms） "
author = ["PennX"]
date = 2020-07-21T12:24:29+08:00
tags = ["Leetcode", "Algorithms", "Python", "DepthfirstSearch"]
categories = ["leetcode"]
draft = false
+++

# DFS + 超简单的过滤，比动态规划快25倍（python 20ms）

> [0416. Partition Equal Subset Sum](https://leetcode-cn.com/problems/partition-equal-subset-sum/)
> [DFS + 超简单的过滤，比动态规划快25倍（python 20ms）](https://leetcode-cn.com/problems/partition-equal-subset-sum/solution/dfs-chao-jian-dan-de-jian-zhi-xing-neng-ti-gao-25b/) by [PennX](https://leetcode-cn.com/u/pennx/)

![leetcode-416.PNG](https://pic.leetcode-cn.com/b7240d92dec49c4c785ce3b2e0d2a9737e5016f34fbd40e7aad164a1f297cb3f-leetcode-416.PNG)


本题解只过滤一种情况：当数组中的最大值超过数组总和的一半。

解释：相当于数组的最大值超过其他所有数之和，必然无法等和分割。

剩下的就是正常的DFS，只要注意不重复使用数字即可。

（如果用动态规划，执行时间为500ms）

python代码如下
```python
class Solution(object):
    def canPartition(self, nums):
        if not nums: return True
        total = sum(nums)
        if total & 1:  # 和为奇数
            return False
        total = total >> 1  # 除2
        nums.sort(reverse=True)  # 逆排序
        if total < nums[0]:  # 当数组最大值超过总和的一半
            return False
        return self.dfs(nums, total)

    def dfs(self, nums, total):
        if total == 0:
            return True
        if total < 0:
            return False
        for i in range(len(nums)):
            if self.dfs(nums[i+1:], total - nums[i]):  # 除去i及其之前，保证每个数只用一次
                return True
        return False

```
