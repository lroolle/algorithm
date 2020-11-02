"+++
title = "0163. Missing Ranges 163. 缺失的区间 "
author = ["klb"]
date = 2020-09-04T03:01:26+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 163. 缺失的区间

> [0163. Missing Ranges](https://leetcode-cn.com/problems/missing-ranges/)
> [163. 缺失的区间](https://leetcode-cn.com/problems/missing-ranges/solution/163-que-shi-de-qu-jian-by-klb/) by [klb](https://leetcode-cn.com/u/klb/)

### 解题思路

从题意可知， nums 肯定在 [lower, upper] 闭区间内。

遍历 nums 数组，定义一个变量 pre 用来表示遍历到的数字的前一个数字。

pre 初始值为 lower - 1.

每遍历一个数字：

1、如果 nums[i] - pre == 2，说明 pre 和 nums[i] 中间缺失了一个数字，这个缺失的数字就是 pre + 1，也可以用 nums[i] - 1 表示。

2、如果 nums[i] - pre >= 3，说明 pre 和 nums[i] 之间缺失了两个以上的数字，缺失的数字的范围是 pre + 1 到 nums[i] - 1。

3、每处理完一个 nums[i]，更新 pre 为 nums[i]。即，对于 nums[i+1] 来说，它的前一个数字就是 nums[i]。

4、遍历完毕 nums[] 数组，此时 pre == nums[nums.length - 1]，

4.1 如果 upper - pre == 1，说明 nums[] 数组最后一个数字尚未达到边界，还差一个，这个数字为 pre + 1，也可以用 upper - 1 表示；

4.2 如果 upper - pre >= 2，说明 nums[] 数组最后一个数字尚未到达边界，还差两个以上的数字，区间为 pre + 1 到 upper。

### 代码

```java
class Solution {
    public List<String> findMissingRanges(int[] nums, int lower, int upper) {
        List<String> ans = new ArrayList<>();
        long pre = (long) lower - 1;    // 防止 lower 就是 int 最小值，减一就直接溢出了
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] - pre == 2) {
                ans.add(String.valueOf(pre + 1));
            } else if (nums[i] - pre >= 3) {
                ans.add((pre + 1) + "->" + (nums[i] - 1));
            }
            pre = nums[i]; // 'int' to 'long'
        }
        if (upper - pre == 1) {
            ans.add(String.valueOf(pre + 1));
        } else if (upper - pre >= 2) {
            ans.add((pre + 1) + "->" + upper);
        }
        return ans;
    }
}
```