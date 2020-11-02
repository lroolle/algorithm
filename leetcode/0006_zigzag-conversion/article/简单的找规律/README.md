"+++
title = "0006. ZigZag Conversion 简单的找规律 "
author = ["chenghua-3"]
date = 2020-09-12T04:17:23+08:00
tags = ["Leetcode", "Algorithms", "C"]
categories = ["leetcode"]
draft = false
+++

# 简单的找规律

> [0006. ZigZag Conversion](https://leetcode-cn.com/problems/zigzag-conversion/)
> [简单的找规律](https://leetcode-cn.com/problems/zigzag-conversion/solution/jian-dan-de-zhao-gui-lu-by-chenghua-3/) by [chenghua-3](https://leetcode-cn.com/u/chenghua-3/)

### 解题思路
![Z字型转换.png](https://pic.leetcode-cn.com/1599884734-vAwEKd-file_1599884733593)
### 代码
```c
char* convert(char* s, int numRows)
{
    int num=strlen(s),count = -1;char* ret= (char*)malloc((1+num)*sizeof(s[1]));
    if(numRows==1 || numRows>num)  return s;
    for (int cir = 0; cir < numRows; cir++)
    {
        int fir, sec; sec = 2 * cir; fir = (2 * numRows - 2) - sec;
        ret[++count] = s[cir];
        for (int c = cir; c < num;)
        {
            c = c + fir;
            if (c < num && cir != (numRows - 1))    ret[++count] = s[c];
            c = c + sec;
            if (c < num && cir != 0)    ret[++count] = s[c];
        }
    }
    ret[num] = '\0';
    return ret;
}
```