"+++
title = "0416. Partition Equal Subset Sum 【题解】416.分割等和子集 "
author = ["qingbin"]
date = 2020-10-01T21:46:49+08:00
tags = ["Leetcode", "Algorithms", "Go", "DynamicProgramming"]
categories = ["leetcode"]
draft = false
+++

# 【题解】416.分割等和子集

> [0416. Partition Equal Subset Sum](https://leetcode-cn.com/problems/partition-equal-subset-sum/)
> [【题解】416.分割等和子集](https://leetcode-cn.com/problems/partition-equal-subset-sum/solution/ti-jie-416fen-ge-deng-he-zi-ji-by-qingbin/) by [qingbin](https://leetcode-cn.com/u/qingbin/)

# 初步分析：
    - 若 nums 之和 Sum 为基数，不可能存在两个和相等的子集，返回 False。
    - 若 nums 之和 Sum 为偶数，试图找到和 N = Sum/2 的子集。

假设存在子集: **Si = {A1, A2, ...... Ai-1, Ai}**，其中 **Sum(Si) == N**。考虑退后一步，则必然有：
**Si-1 = {A1, A2, ...... Ai-2, Ai-1}, Sum(Si-1) == N-Ai** 成立。也就是说，**Sum(Si)** 是否等于 **N**，依赖于 **Sum(Si-1)** 是否等于 **N-Ai** 这个子问题的解，属于是否存在可行解的动态规划问题。

# 动态规划分析：
- 对问题状态建模：
根据初步分析，可另第 **1~x** 个连续元素组成的 **nums** 子集，我们将其标记为 **nums[:x]**。则有任意子集 **nums[:x]** 和任意可能的求和 **y** 之间存在两种状态关系，即 **nums[:x]** 中存在或不存在和为 **y** 的子集。我们将这种状态记作：**F(x, y)**，其取值范围为 **{True, False}**

- 最终状态即为问题解：
根据 **F(x, y)** 的定义便可知，问题的最终解等价于求解 **F(length, N)** 的值，即 **nums[:length]** 中是否存在一个和为 **N** 的子集，其中 **length** 为 **nums** 数组的长度。

- 建立状态转移方程
对于 **nums[:x]** 中任意的 **x**，存在两种可能的选择，即（此处 **x** 下标从 **1** 开始，即第 **x** 个元素）：
    1. 若 **x** 属于目标子集，则存在转换关系 **F(x, y) = F(x-1, y-nums[x])**
        即 **nums[:x]** 中是否存在和为 **y** 的子集，取决于 **nums[:x-1]** 中是否存在和为 **y-nums[x]** 的子集。

    2. 若 **x** 不属于目标子集，则存在转换关系 **F(x, y) = F(x-1, y)**
        即 **nums[:x]** 中是否存在和为 **y** 的子集，取决于 **nums[:x-1]** 中是否存在和为 **y** 的子集。
- 状态初始化
    对于任意的 **y**, 存在 **F(0, y) = False**，即 **nums[:0]** 中不可能存在和为 **y** 的子集。
    对于任意的 **x**, 存在 **F(x, 0) = True**，即 **y** 恒等于 **0** 的情况，可视为 **nums** 中各元素均为 0。
- 计算顺序
从 **1...length** 构建 **x**，从 **1...N(Sum/2)** 构建 **y**，建立状态表 **f**。
        
# 实现：
```
func canPartition(nums []int) bool {
    sum, n := 0, len(nums)
    if nums == nil || n <= 1 {
        return false
    }
    for _, v := range nums {
        sum += v
    }
    if sum % 2 == 1 {
        return false
    }
    sum /= 2
    f := make([][]bool, n + 1)
    for x := 0; x < len(f); x++ {
        f[x] = make([]bool, sum+1)
        f[x][0] = true
    }
    for x := 1; x <= n; x++ {
        for y := 1; y <= sum; y++ {
            if y < nums[x-1] {
                f[x][y] = f[x-1][y]
            } else {
                f[x][y] = f[x-1][y] || f[x-1][y-nums[x-1]]
            }
        }
    }
    return f[n][sum]
}
```

# 优化空间复杂度
从代码分析可知，**f[x-1][*]** 仅在构建 **f[x][*]** 时使用一次便被废弃。而最终取值也仅使用数组的最后一维，即 **f[n]**。所以可以考虑复用 **f[0]** 实现空间压缩。又由于计算 **f[x][y]** 的构建依赖 **f[x-1][y-nums[x-1]]**，故倒序计算 **y** 便可复用上一层计算 **x-1** 时的 **y-nums[x-1]**。

**优化后的代码：**
```
func canPartition(nums []int) bool {
    sum, n := 0, len(nums)
    if nums == nil || n <= 1 {
        return false
    }
    for _, v := range nums {
        sum += v
    }
    if sum % 2 == 1 {
        return false
    }
    sum /= 2
    f := make([]bool, sum+1)
    f[0] = true
    for x := 0; x < n; x++ {
        for y := sum; y >= 0; y-- {
            if y >= nums[x] {
                f[y] = f[y] || f[y-nums[x]]
            }
        }
    }
    return f[sum]
}
```
