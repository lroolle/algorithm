"+++
title = "0045. Jump Game II 一次循环遍历通过，代码比较简单。 "
author = ["ben-da-xiong"]
date = 2020-08-27T09:01:36+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 一次循环遍历通过，代码比较简单。

> [0045. Jump Game II](https://leetcode-cn.com/problems/jump-game-ii/)
> [一次循环遍历通过，代码比较简单。](https://leetcode-cn.com/problems/jump-game-ii/solution/yi-ci-xun-huan-bian-li-tong-guo-by-ben-da-xiong/) by [ben-da-xiong](https://leetcode-cn.com/u/ben-da-xiong/)

一次循环遍历通过，代码比较简单。
```
class Solution {
public:
    int jump(vector<int> &nums) {
        int curr_max_idx = 0;   // 当前跳跃范围最远坐标
        int next_max_idx = 0;   // 下一个跳跃范围最远坐标
        int res = 0;            // 结果
        int target = nums.size() - 1;  // 目标位置
        if (target == 0) return 0;
        for (int i = 0; i <= curr_max_idx; ++i) {   // 在当前跳跃范围内遍历，获得下次跳跃能到达的最远距离
            next_max_idx = max(next_max_idx, i + nums[i]);
            if (next_max_idx >= target) return res + 1;     // 提前结束
            if (i == curr_max_idx) {    // 更新跳跃范围与跳跃次数
                curr_max_idx = next_max_idx;
                ++res;
            }
        }
        return res;
    }
};
```

