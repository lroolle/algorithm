"+++
title = "0047. Permutations II 只在全排列程序上加了两行，清爽的程序！ "
author = ["pan-shao-qin"]
date = 2020-09-18T01:56:21+08:00
tags = ["Leetcode", "Algorithms", "cpp", "Backtracking"]
categories = ["leetcode"]
draft = false
+++

# 只在全排列程序上加了两行，清爽的程序！

> [0047. Permutations II](https://leetcode-cn.com/problems/permutations-ii/)
> [只在全排列程序上加了两行，清爽的程序！](https://leetcode-cn.com/problems/permutations-ii/solution/zhi-zai-quan-pai-lie-cheng-xu-shang-jia-liao-liang/) by [pan-shao-qin](https://leetcode-cn.com/u/pan-shao-qin/)

### 解题思路

先写完标准全排列回溯算法程序，关键是用visited数组来记录已经访问的元素。然后在此程序上加了两行：
一行是对num是进行排列：sort(nums.begin(),nums.end());
另外就是这行，避免重复：if (i>0 && nums[i]==nums[i-1] && visited[i-1])  continue;
就搞定了！

### 代码

```cpp
class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        sort(nums.begin(),nums.end());
        int len=nums.size();
        vector<bool> visited(len,false);
        backtrack(nums,visited,0);
        return res;
    }

    vector<int> path;
    vector<vector<int>> res;
    void backtrack(vector<int>& nums,vector<bool>& visited, int deepth ) {
        int len=nums.size();
        if ( deepth == len) {
            res.push_back(path);      
            return;
        }
        for (int i=0;i<len;i++) {
            if (visited[i]) continue;
            if (i>0 && nums[i]==nums[i-1] && visited[i-1])  continue;
            visited[i] = true;
            path.push_back(nums[i]);
            backtrack(nums,visited,deepth+1);
            path.pop_back();
            visited[i] = false;
        }
    }
};
```