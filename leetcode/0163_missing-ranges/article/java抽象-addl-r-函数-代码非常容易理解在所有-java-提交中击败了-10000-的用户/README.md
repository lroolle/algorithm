"+++
title = "0163. Missing Ranges java,抽象 add(l, r) 函数, 代码非常容易理解,在所有 java 提交中击败了 100.00% 的用户 "
author = ["ZZYuting"]
date = 2019-10-29T03:24:52+08:00
tags = ["Leetcode", "Algorithms", ]
categories = ["leetcode"]
draft = false
+++

# java,抽象 add(l, r) 函数, 代码非常容易理解,在所有 java 提交中击败了 100.00% 的用户

> [0163. Missing Ranges](https://leetcode-cn.com/problems/missing-ranges/)
> [java,抽象 add(l, r) 函数, 代码非常容易理解,在所有 java 提交中击败了 100.00% 的用户](https://leetcode-cn.com/problems/missing-ranges/solution/javaeasy-to-understand-chou-xiang-addl-r-han-shu-d/) by [ZZYuting](https://leetcode-cn.com/u/zzyuting/)
给定一个排序的整数数组 nums ，其中元素的范围在 闭区间 [lower, upper] 当中，返回不包含在数组中的缺失区间。

示例：

输入: nums = [0, 1, 3, 50, 75], lower = 0 和 upper = 99,
输出: ["2", "4->49", "51->74", "76->99"]

```java
//reference videos:https://www.youtube.com/watch?v=Cb_67muiXmo
//time:o(n)
/*
testCase:
[0, 1, 3, 50, 75] , [0,99]
add(-1,0), return 
add(0,1), return
add(1,3), return res += "2"
add(3,50), return res += "4->49"
add(50, 75) return res += "51, 74"
add(75,99) return res += "76,99"
*/
class Solution {
    public List<String> findMissingRanges(int[] nums, int lower, int upper) {
      	//防止溢出
        long l = (long) lower;
        long u = (long) upper;
        int n = nums.length;
        List<String> res = new ArrayList<>();
      	//如果长度为0，直接把l，u加入即可
        if(nums == null || n == 0){
            add(res, l-1, u+1);
            return res;
        }
        add(res, l-1, nums[0]);
        for(int i=1; i<n; i++){
            add(res, nums[i-1], nums[i]);
        }
        add(res, nums[n-1], u+1);
        return res;
    }
    private void add(List<String> res, long l, long r){
        if(l == r || l+1 == r){
            return;
        }else if(l+1 == r-1){
            // res.add(String.valueOf(l+1));
            res.add(String.valueOf(l+1));
        }else{
            // StringBuilder sb = new StringBuilder();
            // sb.append(String.valueOf(l+1));
            // sb.append("->");
            // sb.append(String.valueOf(r-1));
            // res.add(sb.toString());
            res.add((l+1) + "->" + (r-1));
        }
    }
}
```

