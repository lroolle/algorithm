"+++
title = "0239. Sliding Window Maximum 双端队列记录滑动窗口中元素的索引 "
author = ["elevenxx"]
date = 2020-08-21T07:22:54+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Queue", "SlidingWindow", "Python"]
categories = ["leetcode"]
draft = false
+++

# 双端队列记录滑动窗口中元素的索引

> [0239. Sliding Window Maximum](https://leetcode-cn.com/problems/sliding-window-maximum/)
> [双端队列记录滑动窗口中元素的索引](https://leetcode-cn.com/problems/sliding-window-maximum/solution/shuang-duan-dui-lie-ji-lu-hua-dong-chuang-kou-zhon/) by [elevenxx](https://leetcode-cn.com/u/elevenxx/)

双端队列记录滑动窗口中元素的索引，队列左边界记录当前滑动窗口中最大元素的索引
遍历数组，当前元素为 num，索引为 i
1. 当队列非空，左边界出界时（滑动窗口向右移导致的），更新左边界
2. 当队列非空，将队列中索引对应的元素值比 num 小的移除
3. 更新队列
4. 当索引 i 大于 k-1，更新输出结果
```python3
class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        deque = collections.deque()
        ans = []
        for i, num in enumerate(nums):
            if deque and deque[0] <= i-k:
                deque.popleft()
            while deque and nums[deque[-1]] < num:
                deque.pop()
            deque.append(i)
            if i >= k-1:
                ans.append(nums[deque[0]])
        return ans
```

