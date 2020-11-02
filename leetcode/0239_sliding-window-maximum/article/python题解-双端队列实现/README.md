"+++
title = "0239. Sliding Window Maximum python题解--双端队列实现 "
author = ["xiao-xue-66"]
date = 2020-08-23T14:48:59+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# python题解--双端队列实现

> [0239. Sliding Window Maximum](https://leetcode-cn.com/problems/sliding-window-maximum/)
> [python题解--双端队列实现](https://leetcode-cn.com/problems/sliding-window-maximum/solution/pythonti-jie-shuang-duan-dui-lie-shi-xian-by-xiao-/) by [xiao-xue-66](https://leetcode-cn.com/u/xiao-xue-66/)

### 解题思路
![image.png](https://pic.leetcode-cn.com/1598194114-KuXvLn-image.png)
- 双端队列实现，队列里面存储的是下标而不是元素，这点重要

### 代码

```python3
class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        from collections import deque
        dq = deque([])
        res = []

        for i in range(k):
            while dq and nums[dq[0]] < nums[i]:
                dq.popleft()
            dq.appendleft(i)
        res.append(nums[dq[-1]])
        
        j = 0
        for i in range(k,len(nums)):
            while dq and dq[-1] <= j:
                dq.pop()

            while dq and nums[dq[0]] < nums[i]:
                dq.popleft()

            dq.appendleft(i)

            res.append(nums[dq[-1]])
            j += 1

        return res
            
```