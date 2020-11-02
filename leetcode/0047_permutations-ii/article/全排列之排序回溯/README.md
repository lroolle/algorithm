"+++
title = "0047. Permutations II 全排列之排序回溯 "
author = ["shetia"]
date = 2020-09-18T01:22:51+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "Backtracking", "TypeScript"]
categories = ["leetcode"]
draft = false
+++

# 全排列之排序回溯

> [0047. Permutations II](https://leetcode-cn.com/problems/permutations-ii/)
> [全排列之排序回溯](https://leetcode-cn.com/problems/permutations-ii/solution/quan-pai-lie-zhi-pai-xu-hui-su-by-shetia/) by [shetia](https://leetcode-cn.com/u/shetia/)

### 解题思路

用个数组记录已经使用过的数

排序, 用来判断去重

### 代码

```typescript
function permuteUnique( nums: number[] ): number[][] {
  let res: number[][] = []
  let n: number = nums.length
  let cache: boolean[] = Array(n).fill( false )
  nums.sort( (a, b) => a - b )
  dfs( [] )
  return res

  function dfs ( arr: number[] ): void{
    if ( arr.length === n ) {
      res.push( [...arr] )
      return
    }
    for (let i = 0; i < n; i++) {
      if( cache[i] || (i > 0 && nums[i] === nums[i - 1] && !cache[i - 1]) ) continue
      cache[i] = true
      arr.push( nums[i] )
      dfs( arr )
      cache[i] = false
      arr.pop()
    }
  }
};
```