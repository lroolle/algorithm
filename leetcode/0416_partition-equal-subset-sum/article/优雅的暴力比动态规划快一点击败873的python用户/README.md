"+++
title = "0416. Partition Equal Subset Sum 优雅的暴力？比动态规划快一点。击败87.3%的python用户 "
author = ["bai-yi-fei"]
date = 2020-06-17T09:59:10+08:00
tags = ["Leetcode", "Algorithms", "Python3", "哈希", "Python"]
categories = ["leetcode"]
draft = false
+++

# 优雅的暴力？比动态规划快一点。击败87.3%的python用户

> [0416. Partition Equal Subset Sum](https://leetcode-cn.com/problems/partition-equal-subset-sum/)
> [优雅的暴力？比动态规划快一点。击败87.3%的python用户](https://leetcode-cn.com/problems/partition-equal-subset-sum/solution/you-ya-de-bao-li-bi-dong-tai-gui-hua-kuai-yi-dian-/) by [bai-yi-fei](https://leetcode-cn.com/u/bai-yi-fei/)

### 解题思路
看了下动态规划的时间复杂度O(n * sum(nums)).而暴力法的时间复杂度O(2**n)(每一个元素都有取和不取两种情况).优化这么多的原因是当目前元素的和超过sum(nums)/2时就不用考虑啦。加上这个限制，暴力法的时间复杂度也能约束在O(n * sum(nums))以内。
实际上动态规划的时间复杂度比起暴力法还略高一些。因为动态规划每一次循环需要操作sum(nums)/2次，循环n次；而暴力法每次循环需要操作的次数必定小于sum(nums)/2(当元素总和超过sum(nums)/2就舍弃，再加上用哈希表排除重复元素).

### 代码

```python3
class Solution:
    def canPartition(self, nums: List[int]) -> bool:
        n = len(nums)
        target = sum(nums)
        if target % 2:
            return False
        target //= 2
        dic = set() #用于储存当前元素和的可能情况
        dic.add(0)
        for i in range(n):
            dic_tmp = set() #用于存储经过这一步操作，增加的元素和状态
            for j in dic:
                tmp = j + nums[i]
                if tmp == target:
                    return True
                if tmp < target:
                    dic_tmp.add(tmp)
            for j in dic_tmp: #有没有办法直接加在dic里呀？ 感觉这个地方可以优化一下
                dic.add(j)
        return False
```