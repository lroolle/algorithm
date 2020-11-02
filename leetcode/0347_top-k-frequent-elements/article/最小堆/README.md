"+++
title = "0347. Top K Frequent Elements 最小堆 "
author = ["elevenxx"]
date = 2020-08-21T06:36:03+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Heap", "Python"]
categories = ["leetcode"]
draft = false
+++

# 最小堆

> [0347. Top K Frequent Elements](https://leetcode-cn.com/problems/top-k-frequent-elements/)
> [最小堆](https://leetcode-cn.com/problems/top-k-frequent-elements/solution/zui-xiao-dui-by-elevenxx/) by [elevenxx](https://leetcode-cn.com/u/elevenxx/)

解一：
使用 Python 自带容器类型模块中的计数器及其函数
```
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        return [e[0] for e in collections.Counter(nums).most_common(k)]
```

解二：
使用计数器之后构建最小堆
堆的元素可以是元组类型
因为求前 K 个高频元素，python 默认最小堆，则将频次取负再入堆

```python3
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        dic = collections.Counter(nums)
        heap, ans = [], []
        for i in dic:
            heapq.heappush(heap, (-dic[i], i))
        for _ in range(k):
            ans.append(heapq.heappop(heap)[1])
        return ans
```

