"+++
title = "0163. Missing Ranges Java 用一个左边界指针 主要考察数据溢出 "
author = ["yangqinkuan"]
date = 2020-01-22T16:32:11+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# Java 用一个左边界指针 主要考察数据溢出

> [0163. Missing Ranges](https://leetcode-cn.com/problems/missing-ranges/)
> [Java 用一个左边界指针 主要考察数据溢出](https://leetcode-cn.com/problems/missing-ranges/solution/java-yong-yi-ge-zuo-bian-jie-zhi-zhen-zhu-yao-kao-/) by [yangqinkuan](https://leetcode-cn.com/u/yangqinkuan/)

```
class Solution {
    public List<String> findMissingRanges(int[] nums, int lower, int upper) {
        List<String> res = new ArrayList<>();
        // 记录左边界指针
        long left = lower;
        for(int i=0;i<nums.length;i++){
            if(left+1==nums[i]){
                res.add(String.valueOf(left));
            }else if(left+1<nums[i]){
                res.add(String.valueOf(left)+"->"+String.valueOf(nums[i]-1));
            }
            left = (long)nums[i]+1;
        }
        // 最后与upper比较
        if(left==upper) {
            res.add(String.valueOf(left));
        }else if(left<upper){
            res.add(String.valueOf(left)+"->"+String.valueOf(upper));
        }
        return res;
    }
}
```
