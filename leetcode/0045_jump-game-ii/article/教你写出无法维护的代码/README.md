"+++
title = "0045. Jump Game II 教你写出无法维护的代码 "
author = ["newyear-ic"]
date = 2020-09-16T11:11:30+08:00
tags = ["Leetcode", "Algorithms", "C"]
categories = ["leetcode"]
draft = false
+++

# 教你写出无法维护的代码

> [0045. Jump Game II](https://leetcode-cn.com/problems/jump-game-ii/)
> [教你写出无法维护的代码](https://leetcode-cn.com/problems/jump-game-ii/solution/jiao-ni-xie-chu-wu-fa-wei-hu-de-dai-ma-by-newyear-/) by [newyear-ic](https://leetcode-cn.com/u/newyear-ic/)

### 解题思路
执行用时：
12 ms, 在所有 C 提交中击败了69.88%的用户
内存消耗：
6.4 MB, 在所有 C 提交中击败了82.66%的用户

### 代码

```c
#define i1Ili1o0Oo0O0Oili1lii1Ii 1
#define ZERO 0
#define ilIli1o0Oo0OoOili1li11Ii1 int
ilIli1o0Oo0OoOili1li11Ii1 jump(ilIli1o0Oo0OoOili1li11Ii1* ilIli1o0Oo0OoOili1li11Iil, ilIli1o0Oo0OoOili1li11Ii1 ilIli1o0Oo0OoOili1lii1Ii){
    ilIli1o0Oo0OoOili1li11Ii1 i1Ili1o0Oo0OoOili1lii1Ii = ZERO;
    for (ilIli1o0Oo0OoOili1li11Ii1 i1Ili1o0Oo0O0Oilil1i1iIi = ZERO, ilIli1o0Oo0OoOili1li11Ii; i1Ili1o0Oo0O0Oilil1i1iIi < ilIli1o0Oo0OoOili1lii1Ii - i1Ili1o0Oo0O0Oili1lii1Ii; i1Ili1o0Oo0O0Oilil1i1iIi = ilIli1o0Oo0OoOili1li11Ii, i1Ili1o0Oo0OoOili1lii1Ii++) {
        if (i1Ili1o0Oo0O0Oilil1i1iIi + ilIli1o0Oo0OoOili1li11Iil[i1Ili1o0Oo0O0Oilil1i1iIi] >= ilIli1o0Oo0OoOili1lii1Ii - i1Ili1o0Oo0O0Oili1lii1Ii) {
            return i1Ili1o0Oo0OoOili1lii1Ii + i1Ili1o0Oo0O0Oili1lii1Ii;
        }
        ilIli1o0Oo0OoOili1li11Ii = i1Ili1o0Oo0O0Oilil1i1iIi + ilIli1o0Oo0OoOili1li11Iil[i1Ili1o0Oo0O0Oilil1i1iIi];
        for (ilIli1o0Oo0OoOili1li11Ii1 i1Ili1o0Oo0O0Oili1li1iIi = i1Ili1o0Oo0O0Oili1lii1Ii; i1Ili1o0Oo0O0Oili1li1iIi < ilIli1o0Oo0OoOili1li11Iil[i1Ili1o0Oo0O0Oilil1i1iIi] && i1Ili1o0Oo0O0Oilil1i1iIi + i1Ili1o0Oo0O0Oili1li1iIi < ilIli1o0Oo0OoOili1lii1Ii; i1Ili1o0Oo0O0Oili1li1iIi++) {
            if (i1Ili1o0Oo0O0Oilil1i1iIi + i1Ili1o0Oo0O0Oili1li1iIi + ilIli1o0Oo0OoOili1li11Iil[i1Ili1o0Oo0O0Oilil1i1iIi + i1Ili1o0Oo0O0Oili1li1iIi] > ilIli1o0Oo0OoOili1li11Ii + ilIli1o0Oo0OoOili1li11Iil[ilIli1o0Oo0OoOili1li11Ii]) {
                ilIli1o0Oo0OoOili1li11Ii = i1Ili1o0Oo0O0Oilil1i1iIi + i1Ili1o0Oo0O0Oili1li1iIi;
            }
        }
    }
    return i1Ili1o0Oo0OoOili1lii1Ii;
}
```