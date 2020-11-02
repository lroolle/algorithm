"+++
title = "面试题 08.06. Hanota LCCI 回溯法 "
author = ["jason-2"]
date = 2020-07-24T02:35:46+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 回溯法

> [面试题 08.06. Hanota LCCI](https://leetcode-cn.com/problems/hanota-lcci/)
> [回溯法](https://leetcode-cn.com/problems/hanota-lcci/solution/hui-su-fa-by-jason-2-15/) by [jason-2](https://leetcode-cn.com/u/jason-2/)

思路：A中有n个盘子。
1. 先将A中顶部的n-1个盘子，通过C，搬到B中。这是个递归子问题。
子问题解决后。A中一个盘子。B中n-1个盘子。C空出。

2. A中剩余1个盘子，将这个盘子，直接搬到C中。
A空出。B中n-1个盘子。C中一个盘子。

4. 再将B中n-1个盘子，通过A，搬到C中。这又是个递归子问题。
A空出，B空出，C中n个盘子。

```
n:一开始A中有n个盘子。
void dfs(A,n,B,C){
    if(A中只有一个盘子){
        将此盘子，从A中直接搬到C中。
        返回
    }
    dfs(A,n-1,C,B);
    将A中剩余的那个盘子，搬到C中；
    dfs(B,n-1,A,C);
}
```
代码
```
class Solution {
public:
    void dfs(vector<int>& A,int n, vector<int>& B, vector<int>& C){
        if(n <= 0) return;
        if(n == 1){
            C.push_back(A.back());A.pop_back();
            return;
        }
        dfs(A,n-1,C,B);
        C.push_back(A.back());A.pop_back();
        dfs(B,n-1,A,C);
    }
    void hanota(vector<int>& A, vector<int>& B, vector<int>& C) {
        dfs(A,A.size(),B,C);
    }
};
```