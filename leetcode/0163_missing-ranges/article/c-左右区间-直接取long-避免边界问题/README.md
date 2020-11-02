"+++
title = "0163. Missing Ranges C++ 左右区间  直接取long 避免边界问题 "
author = ["xing-chen-da-hai-36"]
date = 2020-08-12T02:44:52+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# C++ 左右区间  直接取long 避免边界问题

> [0163. Missing Ranges](https://leetcode-cn.com/problems/missing-ranges/)
> [C++ 左右区间  直接取long 避免边界问题](https://leetcode-cn.com/problems/missing-ranges/solution/c-zuo-you-qu-jian-zhi-jie-qu-long-bi-mian-bian-jie/) by [xing-chen-da-hai-36](https://leetcode-cn.com/u/xing-chen-da-hai-36/)

### 解题思路
此处撰写解题思路

### 代码

```cpp
class Solution {
public:
    vector<string> findMissingRanges(vector<int>& nums, int lower, int upper) {
        vector<string> ans;
        long l = (long)lower - 1;
        for(int iter : nums){
            if(iter > upper)
                break;
            if(iter - l == 2){
                ans.push_back(to_string(l + 1));
            }
            else if(iter - l > 2){
                ans.push_back(to_string(l + 1) + "->" + to_string(iter - 1));
            }
            l = iter;
        }
        if(upper - l == 1){
             ans.push_back(to_string(upper));
        }
        else if(upper - l > 1){
            ans.push_back(to_string(l + 1) + "->" + to_string(upper));
        }
        return ans;
    }
};
```