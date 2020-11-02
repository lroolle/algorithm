"+++
title = "0242. Valid Anagram 借助集合set "
author = ["xu-chen-chen-d"]
date = 2020-08-21T11:08:54+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 借助集合set

> [0242. Valid Anagram](https://leetcode-cn.com/problems/valid-anagram/)
> [借助集合set](https://leetcode-cn.com/problems/valid-anagram/solution/jie-zhu-ji-he-set-by-xu-chen-chen-d/) by [xu-chen-chen-d](https://leetcode-cn.com/u/xu-chen-chen-d/)

### 解题思路
此处撰写解题思路
把其中一项用集合set（无重复性质）遍历每个字母的个数，如果每个都相等则为True，只要出现一次不等直接False结束。
### 代码

```python3
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        set1 = set(s)
        for i in set1:
            if s.count(i) != t.count(i):
                return False
        return True
```