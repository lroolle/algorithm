"+++
title = "面试题 08.06. Hanota LCCI 一杯茶, 一包烟, 一个easy做半天 "
author = ["shetia"]
date = 2020-07-21T07:31:14+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 一杯茶, 一包烟, 一个easy做半天

> [面试题 08.06. Hanota LCCI](https://leetcode-cn.com/problems/hanota-lcci/)
> [一杯茶, 一包烟, 一个easy做半天](https://leetcode-cn.com/problems/hanota-lcci/solution/yi-bei-cha-yi-bao-yan-yi-ge-easyzuo-ban-tian-by-sh/) by [shetia](https://leetcode-cn.com/u/shetia/)

### 解题思路

还做不出来o(╥﹏╥)o

别说小朋友了, 我心中也有很多问号o(╥﹏╥)o

### 代码

```javascript
/**
 * @param {number[]} A
 * @param {number[]} B
 * @param {number[]} C
 * @return {void} Do not return anything, modify C in-place instead.
 */
var hanota = function(A, B, C) {
  let n = A.length
  let move = function (m, a, b, c) {
    if (m === 1) {          // 当只有一个时直接加到c中
      c.push(a.pop())
    } else {
      move(m - 1, a, c, b)  // 将 a 上的 n - 1 个 通过 c 移到 b
      c.push(a.pop())       // 把 a 中剩下的一个直接放到 c
      move(m - 1, b, a, c)  // 在把 b 上的 n - 1 个 通过 a 放到 c
    }
  }
  move(n, A, B, C)
};
```