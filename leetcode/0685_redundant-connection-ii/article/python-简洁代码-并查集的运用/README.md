"+++
title = "0685. Redundant Connection II python 简洁代码 并查集的运用 "
author = ["yyk"]
date = 2020-04-20T18:19:50+08:00
tags = ["Leetcode", "Algorithms", "Python"]
categories = ["leetcode"]
draft = false
+++

# python 简洁代码 并查集的运用

> [0685. Redundant Connection II](https://leetcode-cn.com/problems/redundant-connection-ii/)
> [python 简洁代码 并查集的运用](https://leetcode-cn.com/problems/redundant-connection-ii/solution/python-jian-ji-dai-ma-bing-cha-ji-de-yun-yong-by-y/) by [yyk](https://leetcode-cn.com/u/yyk/)

### 解题思路

分为两种情况:

1. 都是度为1， 则找出构成环的最后一条边
2. 有度为2的两条边(A->B, C->B)，则删除的边一定是在其中
   先不将C->B加入并查集中，若不能构成环，则C->B是需要删除的点边，反之，则A->B是删除的边(去掉C->B还能构成环，则C->B一定不是要删除的边)

### 代码

```python
# coding:utf-8

class UnionFind(object):

    def __init__(self, n):
        self.pre = range(n)

    def find(self, x):
        if x != self.pre[x]:
            self.pre[x] = self.find(self.pre[x])
        return self.pre[x]

    def union(self, x1, x2):
        root1 = self.find(x1)
        root2 = self.find(x2)
        if root1 != root2:
            self.pre[root2] = root1
            return False

        return True

class Solution(object):
    def findRedundantDirectedConnection(self, edges):
        """
        :type edges: List[List[int]]
        :rtype: List[int]
        """
        n = len(edges)
        uf = UnionFind(n + 1)

        last = []
        parent = {}
        candidates = []
        for st, ed in edges:
            if ed in parent:
                candidates.append([parent[ed], ed])
                candidates.append([st, ed])
            else:
                parent[ed] = st
                if uf.union(st, ed):
                    last = [st, ed]

        if not candidates:
            return last

        return candidates[0] if last else candidates[1]
```