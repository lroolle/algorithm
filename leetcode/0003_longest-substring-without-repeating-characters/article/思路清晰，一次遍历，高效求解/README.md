"+++
title = "0003. Longest Substring Without Repeating Characters 思路清晰，一次遍历，高效求解 "
author = ["Joe_99"]
date = 2020-08-23T01:44:07+08:00
tags = ["Leetcode", "Algorithms", "Python3"]
categories = ["leetcode"]
draft = false
+++

# 思路清晰，一次遍历，高效求解

> [0003. Longest Substring Without Repeating Characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)
> [思路清晰，一次遍历，高效求解](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/si-lu-qing-xi-yi-ci-bian-li-gao-xiao-qiu-jie-by-jo/) by [Joe_99](https://leetcode-cn.com/u/joe_99/)

![1.png](https://pic.leetcode-cn.com/1598146801-YIsXUt-1.png)
``` python 
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        # 使用一个辅助变量来暂时存储匹配的子串
        ans = ''
        tep = ''
        for i in s:
            # 遍历，若不重复则记录该字符
            if i not in tep:
                tep += i
            # 如果遇到了已经存在的字符，则找到该字符所在位置，删除该字符，并保留该位置之后的子串，并把当前字符加入到最后，完成更新
            else:
                tep = tep[tep.index(i)+1:]
                tep += i   
            # 如果是当前最长的，就替换掉之前存储的最长子串
            if len(tep) > len(ans): 
                    ans = tep 
        return len(ans)
```
