"+++
title = "0047. Permutations II 回溯加剪枝 "
author = ["zkk-c"]
date = 2020-08-25T07:37:50+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 回溯加剪枝

> [0047. Permutations II](https://leetcode-cn.com/problems/permutations-ii/)
> [回溯加剪枝](https://leetcode-cn.com/problems/permutations-ii/solution/hui-su-jia-jian-zhi-by-zkk-c-2/) by [zkk-c](https://leetcode-cn.com/u/zkk-c/)

### 解题思路
大体思路同[全排列](https://leetcode-cn.com/problems/permutations/solution/hui-su-jia-jian-zhi-by-zkk-c/)和[组合总和ii](https://leetcode-cn.com/problems/combination-sum-ii/solution/hui-su-jian-zhi-by-zkk-c/)的结合
不同的是要加上
not visited[i - 1]
以[1,1,2]为例，当前层选择第二个1，那当第一个1没有使用过的时候才会跳过

### 代码

```python3
class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        result = []
        nums.sort()
        visited = [False] * len(nums)
        def backtrack(tmp):
            if len(tmp) == len(nums):
                result.append(tmp)
                return 
            
            for i in range(len(nums)):
                if visited[i] or (i > 0 and nums[i] == nums[i - 1] and not visited[i - 1]):
                    continue
                
                visited[i] = True
                backtrack(tmp + [nums[i]])
                visited[i] = False

        backtrack([])
        return result
```