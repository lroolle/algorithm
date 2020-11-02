"+++
title = "0047. Permutations II 47. 全排序 ii "
author = ["shu-cheng"]
date = 2020-08-27T03:22:59+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 47. 全排序 ii

> [0047. Permutations II](https://leetcode-cn.com/problems/permutations-ii/)
> [47. 全排序 ii](https://leetcode-cn.com/problems/permutations-ii/solution/47-quan-pai-xu-ii-by-shu-cheng/) by [shu-cheng](https://leetcode-cn.com/u/shu-cheng/)

##### 思路

* 一个问题是，通过怎样的方式选择 `tmpPath` 第一个或下一个元素，才能使最后得出的结果中没有重复的项？或者说，在何种情况下不选择 `nums` 的元素，作为 `tmpPath` 的元素？
  * 在递归中，该数组元素此前已经调用。
  * 在不是 `tmpPath` 开头元素（`i>1`）的情况下，元素与前一个相同（`nums[i]=nums[i-1]`），且 `nums[i-1]` 没有被调用过（`!hash[i-1]`）。
  * 具体程序：
    * `if (hash[i] || (i>0 && !hash[i-1] && nums[i]==nums[i-1])) continue;`



##### 思路及程序参考

* [秦时明月：47. 全排序 ii](https://leetcode-cn.com/problems/permutations-ii/solution/47-quan-pai-lie-ii-by-alexer-660/)



```javascript
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permuteUnique = function(nums) {
    let n =nums.length;
    nums = nums.sort((a, b) => {return a - b});
    let res = [];
    let tmpPath = [];
    let hash = {};
    let backTrack = (tmpPath) => {
        if (tmpPath.length == n) {
            res.push(tmpPath);
            return;
        }
        for (let i=0; i<n; i++) {
            if (hash[i] || (i>0 && !hash[i-1] && nums[i]==nums[i-1])) continue;
            hash[i] = true;
            tmpPath.push(nums[i]);
            backTrack(tmpPath.slice());
            hash[i] = false;
            tmpPath.pop();
        }
    }
    backTrack(tmpPath);
    return res;
};
```



##### 复杂度分析

* 时间复杂度：O(n!)。n 为数组 `nums` 的长度。`n!` 也为 `nums` 可以形成的全排序的个数。虽然重复的项不会计入结果，可也还是会有一次判断。