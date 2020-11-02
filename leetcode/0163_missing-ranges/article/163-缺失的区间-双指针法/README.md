"+++
title = "0163. Missing Ranges 163. 缺失的区间 (双指针法) "
author = ["jyd"]
date = 2019-06-22T07:39:58+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Java", "Python"]
categories = ["leetcode"]
draft = false
+++

# 163. 缺失的区间 (双指针法)

> [0163. Missing Ranges](https://leetcode-cn.com/problems/missing-ranges/)
> [163. 缺失的区间 (双指针法)](https://leetcode-cn.com/problems/missing-ranges/solution/missing-ranges-shuang-zhi-zhen-fa-by-jyd/) by [jyd](https://leetcode-cn.com/u/jyd/)

- 使用双指针`low`、`num`，遍历`nums`添加对应范围即可；
- 需要先向`nums`尾部添加`upper + 1`。

```python
class Solution:
    def findMissingRanges(self, nums: [int], lower: int, upper: int) -> [str]:
        res = []
        low = lower - 1
        nums.append(upper + 1)
        for num in nums:
            dif = num - low
            if dif == 2: res.append(str(low+1))
            elif dif > 2: res.append(str(low+1) + "->" + str(num-1))
            low = num
        return res
```
```java
class Solution {
    public List<String> findMissingRanges(int[] nums, int lower, int upper) {
        List<String> res = new ArrayList<>();
        long pre = (long)lower - 1; // prevent 'int' overflow
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] - pre == 2) res.add(String.valueOf(pre + 1));
            else if (nums[i] - pre > 2) res.add((pre + 1) + "->" + (nums[i] - 1));
            pre = nums[i]; // 'int' to 'long'
        }
        if (upper - pre == 1) res.add(String.valueOf(pre + 1));
        else if (upper - pre > 1) res.add((pre + 1) + "->" + upper);
        return res;
    }
}
```