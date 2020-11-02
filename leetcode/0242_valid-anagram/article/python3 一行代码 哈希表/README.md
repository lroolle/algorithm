"+++
title = "0242. Valid Anagram python3 一行代码 哈希表 "
author = ["ting-ting-28"]
date = 2020-07-18T08:44:37+08:00
tags = ["Leetcode", "Algorithms", "Python3", "HashTable", "Python"]
categories = ["leetcode"]
draft = false
+++

# python3 一行代码 哈希表

> [0242. Valid Anagram](https://leetcode-cn.com/problems/valid-anagram/)
> [python3 一行代码 哈希表](https://leetcode-cn.com/problems/valid-anagram/solution/python3-yi-xing-dai-ma-ha-xi-biao-by-ting-ting-28/) by [ting-ting-28](https://leetcode-cn.com/u/ting-ting-28/)

### 解题思路
利用内置模块判断。

### 代码

```python3
import collections
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        return collections.Counter(s) == collections.Counter(t)
```