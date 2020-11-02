"+++
title = "0416. Partition Equal Subset Sum DFS判断可行性 "
author = ["zongy17"]
date = 2020-08-07T01:40:38+08:00
tags = ["Leetcode", "Algorithms", "C++", "DepthfirstSearch"]
categories = ["leetcode"]
draft = false
+++

# DFS判断可行性

> [0416. Partition Equal Subset Sum](https://leetcode-cn.com/problems/partition-equal-subset-sum/)
> [DFS判断可行性](https://leetcode-cn.com/problems/partition-equal-subset-sum/solution/dfspan-duan-ke-xing-xing-by-zongy17/) by [zongy17](https://leetcode-cn.com/u/zongy17/)

### 解题思路
先降序排序，再从前往后用DFS进行搜索

### 代码

```cpp
bool comp(int & a, int & b){
    return a > b;
}

class Solution {
public:
    bool DFS(vector<int> & nums, int ptr, int curr_sum, int target)
    {
        int remain = target - curr_sum;
        for (int i = ptr; i < nums.size(); i++)
        {
            if (nums[i]==remain)
                return true;
            else if (nums[i] < remain)
            {
                if (DFS(nums, i+1, curr_sum+nums[i], target))
                    return true;
            }
            //nums[i] > remain的情况不用管，直接下一步
        }
        return false;//全部都搜过一遍还是没有return true
    }

    bool canPartition(vector<int>& nums)
    {
        int n = nums.size();
        int tot_sum = 0;
        for (int i = 0; i < n; i++)
        {
            tot_sum += nums[i];
        }
        if (tot_sum % 2 != 0)
            return false;
        int div_sum = tot_sum / 2;

        sort(nums.begin(), nums.end(), comp);//降序排列
        if (nums[0] > div_sum) return false;//第一个就已经过半了
        else {
            if (DFS(nums, 0, 0, div_sum))
                return true;
            else return false;
        }
    }
};
```