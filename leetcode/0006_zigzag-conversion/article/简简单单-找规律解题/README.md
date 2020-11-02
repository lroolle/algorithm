"+++
title = "0006. ZigZag Conversion 简简单单-找规律解题 "
author = ["lang-shi-4"]
date = 2020-08-31T05:37:07+08:00
tags = ["Leetcode", "Algorithms", "Python"]
categories = ["leetcode"]
draft = false
+++

# 简简单单-找规律解题

> [0006. ZigZag Conversion](https://leetcode-cn.com/problems/zigzag-conversion/)
> [简简单单-找规律解题](https://leetcode-cn.com/problems/zigzag-conversion/solution/jian-jian-dan-dan-zhao-gui-lu-jie-ti-by-lang-shi-4/) by [lang-shi-4](https://leetcode-cn.com/u/lang-shi-4/)

```
class Solution:
    def convert(self, s: str, numRows: int) -> str:
        #base case
        if numRows==1:return str(s)
        res=''
        for irow in range(numRows):
            if irow>0 and irow<numRows-1:
                tmp_index=irow
                ichar_index=irow
                while ichar_index<len(s):
                    res+=s[ichar_index]
                    #关键！！！！
                    #找中间的数z的斜边
                    if (ichar_index+numRows*2-2-(irow)*2)<len(s):
                        #print(ichar_index+numRows*2-2-(irow)*2,len(s))
                        res+=s[ichar_index+numRows*2-2-(irow)*2]
                    ichar_index+=(numRows*2-2)
                    
            else:
                #最上和最下的两行
                tmp_index=irow
                while irow<len(s):
                    res+=s[irow]
                    irow+=(numRows*2-2)
        return res
```
