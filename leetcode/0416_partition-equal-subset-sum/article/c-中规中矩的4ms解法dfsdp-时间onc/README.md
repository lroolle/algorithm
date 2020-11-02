"+++
title = "0416. Partition Equal Subset Sum C++ 中规中矩的4ms解法（dfs+dp 时间O(NC)） "
author = ["Gary_coding"]
date = 2020-08-01T12:53:44+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# C++ 中规中矩的4ms解法（dfs+dp 时间O(NC)）

> [0416. Partition Equal Subset Sum](https://leetcode-cn.com/problems/partition-equal-subset-sum/)
> [C++ 中规中矩的4ms解法（dfs+dp 时间O(NC)）](https://leetcode-cn.com/problems/partition-equal-subset-sum/solution/c-zhong-gui-zhong-ju-de-8msjie-fa-dfsjian-zhi-by-g/) by [Gary_coding](https://leetcode-cn.com/u/gary_coding/)

### dfs+剪枝 时间O(N*max(logN,C))
N是数组长度，C是数组元素和的一半
```cpp
class Solution {
    int n;
public:
    bool canPartition(vector<int>& nums) {
        n = nums.size();
        int target = accumulate(nums.begin(), nums.end(), 0); if (target & 1) return false; target >>= 1;
        sort(nums.rbegin(), nums.rend());
        if (nums[0] > target) return false;
        return dfs(nums, 0, target);

    }
    bool dfs(vector<int>& nums, int idx, int target) {
        if (idx >= n) return false;
        if (nums[idx] == target) return true;
        if (nums[idx] > target) return dfs(nums, idx + 1, target);

        return dfs(nums, idx + 1, target - nums[idx]) || dfs(nums, idx + 1, target);
    }
};
```

### dp 时间O(NC)
N是数组长度，C是数组元素和的一半
```cpp
class Solution {
public:
    bool canPartition(vector<int>& nums) {
        int target = accumulate(nums.begin(), nums.end(), 0);
        if (target & 1) return false;
        target >>= 1;

        int n = nums.size();
        vector<bool> dp(target + 1, false);
        dp[0] = true;
        for (int i = 1; i < n; ++i) {
            int cur = nums[i];
            for (int j = target; j >= 1; --j) if (j - cur >= 0) dp[j] = dp[j] | dp[j - cur];
            if (dp[target]) return true;
        }
        return false;
    }
};
```