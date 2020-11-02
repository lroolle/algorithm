"+++
title = "0047. Permutations II C++ (4~8ms 6行) 这题有那么复杂吗？ "
author = ["chronosphere"]
date = 2020-06-23T01:33:23+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# C++ (4~8ms 6行) 这题有那么复杂吗？

> [0047. Permutations II](https://leetcode-cn.com/problems/permutations-ii/)
> [C++ (4~8ms 6行) 这题有那么复杂吗？](https://leetcode-cn.com/problems/permutations-ii/solution/c-48ms-6xing-zhe-ti-you-na-yao-fu-za-ma-by-chronos/) by [chronosphere](https://leetcode-cn.com/u/chronosphere/)

怎么都在考虑去重。。。

`next_permutaion`是C++库函数，其实现参考[我在31题的解答](https://leetcode-cn.com/problems/next-permutation/solution/c-4ms-5xing-by-chronosphere/)
```cpp
class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        vector<vector<int>> result;
        sort(nums.begin(), nums.end());
        result.push_back(nums);
        while(next_permutation(nums.begin(), nums.end())) {
            result.push_back(nums);
        }
        return result;
    }
};
```

如果想自己实现`next_permutation`也很简单，根据[我在31题的解答](https://leetcode-cn.com/problems/next-permutation/solution/c-4ms-5xing-by-chronosphere/)，多加一个函数即可。
```cpp
class Solution {
public:
    bool nextPermutation(vector<int>& nums) {
        auto i = is_sorted_until(nums.rbegin(), nums.rend()); // 找到末尾的一个降序段[s]及其前一个元素i
        bool has_next = i != nums.rend();
        if(has_next) {
            iter_swap(i, upper_bound(nums.rbegin(), i, *i));  // 找到[s]中比i大的数中最小的
            reverse(nums.rbegin(), i);                        // 序列反转
        }
        return has_next;
    }
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        vector<vector<int>> result;
        sort(nums.begin(), nums.end());
        result.push_back(nums);
        while(nextPermutation(nums)) {
            result.push_back(nums);
        }
        return result;
    }
};
```