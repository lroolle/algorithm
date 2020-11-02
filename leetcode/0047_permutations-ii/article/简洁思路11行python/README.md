"+++
title = "0047. Permutations II 简洁思路+11行Python "
author = ["xiao-yan-gou"]
date = 2020-09-18T04:59:04+08:00
tags = ["Leetcode", "Algorithms", "Python", "Python3", "DepthfirstSearch"]
categories = ["leetcode"]
draft = false
+++

# 简洁思路+11行Python

> [0047. Permutations II](https://leetcode-cn.com/problems/permutations-ii/)
> [简洁思路+11行Python](https://leetcode-cn.com/problems/permutations-ii/solution/jian-ji-si-lu-11xing-python-by-xiao-yan-gou/) by [xiao-yan-gou](https://leetcode-cn.com/u/xiao-yan-gou/)

给你五个数字，那就是有五个位置要填空
依次填空，只要保证当前空格填的数字不重复，那你的整个历史就是不重复的
例子：
[5,5,6,6,8]

填第一个空格时，保证当前空格填的数字不重复，那就是选[5,6,8]中的一个填
于是有
[5]
[6]
[8]

填第二个空格时，如果你的历史是[5],那么你还剩下[5,6,6,8]这几个选项
保证当前空格填的数字不重复，还是选[5,6,8]中的一个填
于是有
[5,5]
[5,6]
[5,8]

填第三个空格时，如果你的历史是[5,6],那么你还剩下[5,6,8]这几个选项
保证当前空格填的数字不重复，还是选[5,6,8]中的一个填
于是有
[5,6,5]
[5,6,6]
[5,6,8]

你是不是在担心，是否会出现[5*,6,5],[5,6,5*]像这样出现两次的情况？（用*区分两个不同的5）
并不会，因为你一直在保证你的历史从来没有重复过。
```python
class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        ret = []

        def search(left,history):
            nonlocal ret
            if not left: #如果没有可以搜的了，说明所有数字用完了
                ret.append(history)

            for i in set(left): #只考虑了当前位置不重复选择，那也就能保证history不重复，所以直接用一个集合来维护
                left.remove(i)
                left.append(i)
                search(left[:-1],history+[i])

        search(nums,[])
        return ret
```
