"+++
title = "0006. ZigZag Conversion Z 字形变换极简解法 "
author = ["zoffer"]
date = 2020-05-24T15:03:35+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "Java", "Python", "Go"]
categories = ["leetcode"]
draft = false
+++

# Z 字形变换极简解法

> [0006. ZigZag Conversion](https://leetcode-cn.com/problems/zigzag-conversion/)
> [Z 字形变换极简解法](https://leetcode-cn.com/problems/zigzag-conversion/solution/ji-jian-jie-fa-by-ijzqardmbd/) by [zoffer](https://leetcode-cn.com/u/zoffer/)

### 解题思路
![v.png](https://pic.leetcode-cn.com/52c6e3914ff4167230ef4d06513f0e3a2075b99988172705e82871ae39b16c77-v.png)
- 以 V 字型为一个循环, 循环周期为 `n = (2 * numRows - 2)` （2倍行数 - 头尾2个）。
- 对于字符串索引值 $\color{red}i$，计算 `x = i % n` 确定在循环周期中的位置。
- 则行号 $\color{blue}y$ = `min(x, n - x)`。

以 `numRows = 4` 为例，则 `n = 6`， 规律如下：

rows |||||
--|--|--|--|--
$\color{blue}0$ | x = $\color{red}0$ % 6 = $\color{blue}0$ ✔️ <br> 6 - x = 6| | | x = $\color{red}6$ % 6 = $\color{blue}0$ ✔️ <br> 6 - x = 6
$\color{blue}1$ | x = $\color{red}1$ % 6 = $\color{blue}1$ ✔️ <br> 6 - x = 5| | x = $\color{red}5$ % 6  = 5 <br> 6 - x = $\color{blue}1$ ✔️ | x = $\color{red}7$ % 6 = $\color{blue}1$ ✔️ <br> 6 - x = 5
$\color{blue}2$ | x = $\color{red}2$ % 6 =  $\color{blue}2$ ✔️ <br> 6 - x = 4| x = $\color{red}4$ % 6  = 4 <br> 6 - x = $\color{blue}2$ ✔️ | | x = $\color{red}8$ % 6 = $\color{blue}2$ ✔️ <br> 6 - x = 4
$\color{blue}3$ | x = $\color{red}3$ % 6 = $\color{blue}3$ ✔️ <br> 6 - x = 3| | | ...

根据规律可拼接出每一行的的字符串，最后合并所有行即可。

### 代码

```javascript
/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function(s, numRows) {
    if (numRows === 1) return s;
    const rows = new Array(numRows).fill("");
    const n = 2 * numRows - 2;
    for(let i = 0; i < s.length; i++) {
        const x = i % n;
        rows[Math.min(x, n - x)] += s[i];
    }
    return rows.join("");
};
```
```python3
class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1: return s
        rows = [""] * numRows
        n = 2 * numRows - 2
        for i, char in enumerate(s):
            x = i % n
            rows[min(x, n - x)] += char
        return "".join(rows)
```
```java
class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1) return s;
        StringBuilder[] rows = new StringBuilder[numRows];
        for(int i = 0; i < numRows; i++) rows[i] = new StringBuilder();
        int n = 2 * numRows - 2;
        for (int i = 0; i < s.length(); i++) {
            int x = i % n;
            rows[Math.min(x, n - x)].append(s.charAt(i));
        }
        return String.join("", rows);
    }
}
```
```go
func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    rows := make([]string, numRows)
    n := 2 * numRows - 2
    for i, char := range s {
        x := i % n
        rows[min(x, n - x)] += string(char)
    }
    return strings.Join(rows, "")
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```