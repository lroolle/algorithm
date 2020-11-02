"+++
title = "0045. Jump Game II 单（hua）调（dong）队（chuang）列（kou）优化dp "
author = ["RinHoshizora"]
date = 2020-05-04T12:01:49+08:00
tags = ["Leetcode", "Algorithms", "C++", "DynamicProgramming"]
categories = ["leetcode"]
draft = false
+++

# 单（hua）调（dong）队（chuang）列（kou）优化dp

> [0045. Jump Game II](https://leetcode-cn.com/problems/jump-game-ii/)
> [单（hua）调（dong）队（chuang）列（kou）优化dp](https://leetcode-cn.com/problems/jump-game-ii/solution/dan-diao-dui-lie-you-hua-dp-by-zyounan/) by [RinHoshizora](https://leetcode-cn.com/u/rinhoshizora/)

不难写出 $dp$ 方程。设 $dp_i$ 代表到达点 $i,0 \leqslant i < n$ 时需要的最小步数：
$$
dp_{i} = \min_{j < i,\ \ a_{j} \geqslant i - j}\{dp_{j}\} + 1
$$
然而直接枚举是$O(n^2)$的，超时。

注意到这样的 $dp$ 是经典的 [1D/1D](https://wenku.baidu.com/view/681d161ca300a6c30c229f70.html) 形式，直接单调队列优化就好了 ($\texttt{escape}$
```cpp
const int maxn = 1e5 + 5;
class Solution {
    int dp[maxn];
    int q[maxn],l=1,r;  //也可以用 std::deque 代替
public:
    int jump(const vector<int>& a) {
        int n = a.size();
        memset(dp,0x3f,sizeof dp);
        dp[0] = 0,q[++r] = 0;   //我们从第二个开始转移，这样就先把第一个放入队列
        for(int i = 1;i < n;++i) {
            while(l <= r && a[q[l]] + q[l] < i) ++l;    
            if(l <= r) dp[i] = dp[q[l]] + 1;        //转移，队首为最小值
            while(l <= r && dp[i] < dp[q[r]])  --r;     //维护单调队列，保证队列单调递增
            q[++r] = i;         
        }
        return dp[n-1];
    }
};
```
