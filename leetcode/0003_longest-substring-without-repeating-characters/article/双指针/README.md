"+++
title = "0003. Longest Substring Without Repeating Characters 双指针 "
author = ["rPSXFVr7e3"]
date = 2020-09-08T07:52:08+08:00
tags = ["Leetcode", "Algorithms", "Python", "Python3", "TwoPointers"]
categories = ["leetcode"]
draft = false
+++

# 双指针

> [0003. Longest Substring Without Repeating Characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)
> [双指针](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/shuang-zhi-zhen-by-rpsxfvr7e3/) by [rPSXFVr7e3](https://leetcode-cn.com/u/rPSXFVr7e3/)

### 解题思路
此处撰写解题思路
1.两个指针，一个为不重复子串左端，一个为不重复子串右端
2.右端即为循环遍历整个字符串
3.左端需要考虑情况 
    3.1 第一个数，即first为-1
    3.2 不在字典中的数，first不变
    3.3 在字典中的数，注意到前一个first即为s[i-1]的first值，故必须大于等于，但若在这个first值之后出现s[i]值，那么first为dic[s[i]]
4.不断计算子串长度，并自我比较，返回最大值
### 代码

```python
class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        if not s:
            return 0
        
        first = -1
        dic = {}
        ans = 0
        for i in range(len(s)):
            if s[i] not in dic:
                dic[s[i]]=i 
                ans = max(ans,i-first)
            else:
                first = max(dic[s[i]],first)
                dic[s[i]]=i
                ans = max(ans,i-first)
        return ans
        #
        #if s[i] in dic:
        #    first = max(dic[s[i]],first)
        #dic[s[i]]=i
        #ans = max(ans,i-first)
        #
```