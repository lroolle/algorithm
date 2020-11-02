"+++
title = "0416. Partition Equal Subset Sum 416.分割等和子集【动态规划、递归】 "
author = ["suukii"]
date = 2020-08-20T14:32:13+08:00
tags = ["Leetcode", "Algorithms", "DynamicProgramming", "Memoization", "Recursion", "JavaScript", "Python"]
categories = ["leetcode"]
draft = false
+++

# 416.分割等和子集【动态规划、递归】

> [0416. Partition Equal Subset Sum](https://leetcode-cn.com/problems/partition-equal-subset-sum/)
> [416.分割等和子集【动态规划、递归】](https://leetcode-cn.com/problems/partition-equal-subset-sum/solution/416fen-ge-deng-he-zi-ji-dong-tai-gui-hua-di-gui-by/) by [suukii](https://leetcode-cn.com/u/suukii/)

## 思路1: DP

**题目理解**

一开始看到题目提到两个子数组，还有点不知道怎么跟 DP 扯上关系。

但其实题目可以换一个说法：**给定数组 nums，是否存在一个子数组，该子数组的和等于 nums 元素和的一半**。

这样就清晰多了。

**解题**

我们还是用一个一维数组 dp 来记录题目的解，dp[i] 表示**是否存在元素和为 i 的子数组**。

对于 nums 中的每个数字 n 来说，都有**选**和**不选**两种选择，选的话问题变成 dp[i - n]，不选的话问题还是 dp[i]，所以：

```
dp[i] = dp[i - n] or dp[i]
```

### 复杂度

- 时间复杂度：$O(len*target)$, len 是 nums 的长度，target 是 nums 元素和的一半。
- 空间复杂度：$O(target)$, target 是 nums 数组元素和的一半。

### 代码

JavaScript Code
```js
/**
 * @param {number[]} nums
 * @return {boolean}
 */
var canPartition = function(nums) {
    const target = nums.reduce((a, b) => a + b) / 2
    if (target % 1 !== 0) return false // nums 和为奇数

    const dp = Array(target + 1).fill(false)
    dp[0] = true
    for (const n of nums) {
        for (let i = target; i >= n; i--) {
            if (dp[target]) return true // 提前返回结果
            dp[i] = dp[i] || dp[i - n]
        }
    }
    return dp[target]
};
```

## 思路2: DFS

将两个子数组 a 和 b 分别初始化为 nums 和 []，然后不断从 a 中取出数字放到 b 中，当两个子数组的和相等时，返回 true。

对于 a 中的每个数字，都有**选**与**不选**两种选择。

### 复杂度

- 时间复杂度：$O(2^n)$，n 为数组 nums 长度，可以看成是二叉树的遍历。
- 空间复杂度：$O(logn)$，n 为数组 nums 长度，递归栈消耗的空间。

### 代码

Python Code
```py
class Solution(object):
    def canPartition(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        @lru_cache
        def dfs(i, sum1, sum2):
            if sum1 == sum2: return True
            if sum2 > sum1 or i >= len(nums): return False
            return dfs(i + 1, sum1 - nums[i], sum2 + nums[i]) or dfs(i + 1, sum1, sum2)

        total = sum(nums)
        if total % 2 == 1: return False
        return dfs(0, total, 0)
```