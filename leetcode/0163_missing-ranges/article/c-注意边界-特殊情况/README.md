"+++
title = "0163. Missing Ranges C++ 注意边界 && 特殊情况 "
author = ["duo-la-quan-er-meng"]
date = 2020-05-01T09:18:09+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# C++ 注意边界 && 特殊情况

> [0163. Missing Ranges](https://leetcode-cn.com/problems/missing-ranges/)
> [C++ 注意边界 && 特殊情况](https://leetcode-cn.com/problems/missing-ranges/solution/czhu-yi-bian-jie-by-chi-la-quan-er-meng/) by [duo-la-quan-er-meng](https://leetcode-cn.com/u/duo-la-quan-er-meng/)

1）注意int范围
2）把lower - 1、upper + 1加到数组中，减少特殊情况处理

```
class Solution {
public:
    vector<string> findMissingRanges(vector<int>& nums, int lower, int upper) {
        int64_t left = (int64_t)lower - 1;
        char buf[1024] = {0};
        vector<string> result;
        vector<int64_t> nums_copy;

        nums_copy.emplace_back((int64_t)lower - 1);
        nums_copy.insert(nums_copy.end(), nums.begin(), nums.end());
        nums_copy.emplace_back((int64_t)upper + 1);

        for (int i = 0; i < nums_copy.size(); ++i) {
            auto&& delta = nums_copy[i] - left;
            if (delta > 1) {
                if (delta == 2) sprintf(buf, "%lld", left + 1);
                else sprintf(buf, "%lld->%lld", left + 1, nums_copy[i] - 1);
                result.emplace_back(buf);
            }
            left = nums_copy[i];
        }
        return result;
    }
};
```
