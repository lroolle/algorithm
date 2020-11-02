"+++
title = "0739. Daily Temperatures C语言单调栈，97.87%，100% "
author = ["pandab-2"]
date = 2020-06-11T00:24:25+08:00
tags = ["Leetcode", "Algorithms", "C", "Stack"]
categories = ["leetcode"]
draft = false
+++

# C语言单调栈，97.87%，100%

> [0739. Daily Temperatures](https://leetcode-cn.com/problems/daily-temperatures/)
> [C语言单调栈，97.87%，100%](https://leetcode-cn.com/problems/daily-temperatures/solution/cyu-yan-dan-diao-zhan-9787100-by-pandab-2/) by [pandab-2](https://leetcode-cn.com/u/pandab-2/)

### 解题思路
是的，这不就是 [503. 下一个更大元素 II](https://leetcode-cn.com/problems/next-greater-element-ii/)么，单调栈玩儿多了，都是一个套路
1. 题面要求下一个超过当日的温度，肯定得构造一个单调递减栈，当遇到比栈顶元素大的温度，就要考虑出栈咯；
2. 单调栈中，存入的是每日温度的下标，当发现更高的温度时，栈顶元素`monoStack[stackTop]`和当前元素下标`tIter`相减，即为天数差了。

### 代码

```c
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *dailyTemperatures(int *T, int TSize, int *returnSize)
{
    int *monoStack = (int *)malloc(sizeof(int) * TSize);
    memset(monoStack, 0, sizeof(int) * TSize);
    int stackTop = -1;
    int tIter = 0;

    int *res = (int *)malloc(sizeof(int) * TSize);
    memset(res, 0, sizeof(int) * TSize);

    while (tIter < TSize) {
        /* monoStack[stackTop]是栈顶元素在T中的下标，T[monoStack[stackTop]] 才是真正的栈顶温度 */
        while (stackTop != -1 && T[monoStack[stackTop]] < T[tIter]) {
            int r = monoStack[stackTop];
            stackTop--;
            
            res[r] = tIter - r;
        }
        monoStack[++stackTop] = tIter++;        
    }
    *returnSize = TSize;
    return res;
}
```