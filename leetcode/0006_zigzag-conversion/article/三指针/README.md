"+++
title = "0006. ZigZag Conversion 三指针 "
author = ["yi-wen-statistics"]
date = 2020-08-25T09:48:09+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 三指针

> [0006. ZigZag Conversion](https://leetcode-cn.com/problems/zigzag-conversion/)
> [三指针](https://leetcode-cn.com/problems/zigzag-conversion/solution/san-zhi-zhen-by-yi-wen-statistics-3/) by [yi-wen-statistics](https://leetcode-cn.com/u/yi-wen-statistics/)

## 方法1——构建二维数组
如代码

### 代码

```python3
class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s

        # 构建Z型字矩阵
        multi = len(s) // (numRows*2-2) 
        residual = len(s) % (numRows*2-2)
        columns = multi*(numRows-1) if residual == 0 else (multi+1)*(numRows-1)
        words = [['' for _ in range(columns)] for _ in range(numRows)]
        i, j, k = 0, 0, 0 # 初始化三指针

        # Z型字填充
        while i < len(s):
            count = numRows*2-2
            while i < len(s) and count > numRows-2:
                words[j][k] = s[i]
                j += 1
                i += 1
                count -= 1
            if numRows == 2:
                j -= 2
                k += 1
            else:
                j -= 1
                while i < len(s) and count > 0:
                    j -= 1
                    k += 1
                    words[j][k] = s[i]
                    count -= 1
                    i += 1
                j -= 1
                k += 1

        # 生成结果
        res = ''
        for m in range(numRows):
            for n in range(columns):
                res += words[m][n]

        return res

```

### 复杂度分析
时间复杂度：O(N)，前面常数项很大，或者干脆是O(N^2)，这个思路很乱，仅仅能解出来并不优雅
空间复杂度：O(N)

## 方法2——构建一维数组
方法2的本质和方法1一样，但是方法2可以避免遍历很多空格

### 代码
代码还可以精简，但编码水平有限，只能用双指针凑活儿一下了

```
class Solution:
    def convert(self, s: str, numRows: int) -> str:
        # 构建Z型字矩阵
        res = ['' for _ in range(numRows)]
        i, j = 0, 0

        # Z型字填充
        while i < len(s):
            while i < len(s) and j < numRows:
                res[j] += s[i]
                j += 1
                i += 1
            if numRows <= 2:
                j -= numRows
            else:
                j -= 1
                while i < len(s) and j > 1:
                    j -= 1
                    res[j] += s[i]
                    i += 1
                j -= 1

        return ''.join(res)
```

### 复杂度分析
时间复杂度：O(N)
空间复杂度：O(N)
