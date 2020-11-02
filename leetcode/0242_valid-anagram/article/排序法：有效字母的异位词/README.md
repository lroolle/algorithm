"+++
title = "0242. Valid Anagram 排序法：有效字母的异位词 "
author = ["liwuhen-2"]
date = 2020-08-22T09:01:49+08:00
tags = ["Leetcode", "Algorithms", "Python3"]
categories = ["leetcode"]
draft = false
+++

# 排序法：有效字母的异位词

> [0242. Valid Anagram](https://leetcode-cn.com/problems/valid-anagram/)
> [排序法：有效字母的异位词](https://leetcode-cn.com/problems/valid-anagram/solution/pai-xu-fa-you-xiao-zi-mu-de-yi-wei-ci-by-liwuhen-2/) by [liwuhen-2](https://leetcode-cn.com/u/liwuhen-2/)


对两个字符串排序后比较它们是否相等进而判断是否是字母异位词。

```
class Solution(object):
    def isAnagram(self, s, t):
    	if sorted(s) == sorted(t):
    		return True
    	else:
    		return False
```