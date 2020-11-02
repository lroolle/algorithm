"+++
title = "面试题 08.06. Hanota LCCI 为什么要那么麻烦?[滑稽] "
author = ["justfun"]
date = 2020-04-07T12:27:05+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 为什么要那么麻烦?[滑稽]

> [面试题 08.06. Hanota LCCI](https://leetcode-cn.com/problems/hanota-lcci/)
> [为什么要那么麻烦?[滑稽]](https://leetcode-cn.com/problems/hanota-lcci/solution/wei-shi-yao-yao-na-yao-ma-fan-hua-ji-by-justfun/) by [justfun](https://leetcode-cn.com/u/justfun/)

```cpp
class Solution {
public:
    void hanota(vector<int>& A, vector<int>& B, vector<int>& C) {
        swap(A, C);
    }
};
```