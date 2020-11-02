"+++
title = "0739. Daily Temperatures 单调栈+从后向前遍历（Python3详细注释） "
author = ["dz-lee"]
date = 2020-06-11T04:04:49+08:00
tags = ["Leetcode", "Algorithms", "Stack", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 单调栈+从后向前遍历（Python3详细注释）

> [0739. Daily Temperatures](https://leetcode-cn.com/problems/daily-temperatures/)
> [单调栈+从后向前遍历（Python3详细注释）](https://leetcode-cn.com/problems/daily-temperatures/solution/dan-diao-zhan-cong-hou-xiang-qian-bian-li-python3x/) by [dz-lee](https://leetcode-cn.com/u/dz-lee/)

### 解题思路
1. 首先想到的解决思路是暴力搜索，复杂度是O(n^2)，不能满足要求。
2. 然后看了官方题解，可以从后向前遍历，每次遍历的时候，维护当前温度的最小下标。因为最高温度有上限，所以时间复杂度是O(nm)
3. 看了官方**单调栈**的解法，太厉害了，**这种解法一定要多多练习、掌握**。时间复杂度是O(n)

### 代码
#### 从后向前遍历解法
```python3
class Solution:
    def dailyTemperatures(self, T: List[int]) -> List[int]:
        d = {}
        ans = [0] * len(T)
        for i in range(len(T) - 1, -1, -1): # 从后向前遍历
            d[T[i]] = i         # 记录当前问题的最小下标
            tmp = [d[t] - i for t in range(T[i] + 1, 101) if t in d]
            ans[i] = (min(tmp) if tmp else 0)
        return ans
```

#### 单调栈解法
```python3
class Solution:
    def dailyTemperatures(self, T: List[int]) -> List[int]:
        # 可以维护一个存储下标的单调栈，从栈底到栈顶的下标对应的温度列表中的温度依次递减。
        # 如果一个下标在单调栈里，则表示尚未找到下一次温度更高的下标。
        ans = [0] * len(T)
        stack = []
        for i in range(len(T)):
            while stack and T[stack[-1]] < T[i]:    # 栈不为空 && 栈顶温度小于当前温度
                ans[stack[-1]] = i - stack[-1]
                stack.pop()
            stack.append(i)
        return ans
```