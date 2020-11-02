"+++
title = "0739. Daily Temperatures Python3, 单调栈，739 "
author = ["lionKing_njuer"]
date = 2020-08-26T14:08:10+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# Python3, 单调栈，739

> [0739. Daily Temperatures](https://leetcode-cn.com/problems/daily-temperatures/)
> [Python3, 单调栈，739](https://leetcode-cn.com/problems/daily-temperatures/solution/python3-dan-diao-zhan-739-by-littlelion_njuer/) by [lionKing_njuer](https://leetcode-cn.com/u/lionking_njuer/)

### 解题思路
维护一个单调栈；
当栈不为空，且温度大于栈顶温度，弹出栈顶温度，且将差值添加到res.
栈内元素为温度的index。

### 代码

```python3
class Solution:
    def dailyTemperatures(self, T):
        res = [0] * len(T)
        stack = []
        for index, temperature in enumerate(T):
            while stack and temperature > T[stack[-1]]:
                cur = stack.pop()
                res[cur] = index - cur
            stack.append(index)
        return res

```