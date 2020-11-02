"+++
title = "0416. Partition Equal Subset Sum 剪枝就可以，容易理解还高效 -- 极其适合初学者理解 "
author = ["yi-shun-niu-xing"]
date = 2020-07-11T14:06:21+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# 剪枝就可以，容易理解还高效 -- 极其适合初学者理解

> [0416. Partition Equal Subset Sum](https://leetcode-cn.com/problems/partition-equal-subset-sum/)
> [剪枝就可以，容易理解还高效 -- 极其适合初学者理解](https://leetcode-cn.com/problems/partition-equal-subset-sum/solution/jian-zhi-jiu-ke-yi-rong-yi-li-jie-huan-gao-xiao-ji/) by [yi-shun-niu-xing](https://leetcode-cn.com/u/yi-shun-niu-xing/)

看了很多题解，大家好想普遍都把这个问题列为动态规划的经典例题。

动态规划的思想其实就是空间换时间，通过将子问题的解记录下来，从而共同解答复杂的大问题。

然而动态规划中很难让人理解的就是为什么。。
- 为什么要用这个动态转移方程
- 为什么要用二维数组记录
- 二维数组记录的内容到底是什么

动态规划对于初学者来说很不友好。而对于已经很有经验的大佬来说，看到题就知道动态转移方程怎么来的。

但其实，动态规划的理解应该从暴力求解开始，通过空间换取时间的思想，从而避免解决重复计算的问题（例如斐波那契数的计算）。所以理解暴力求解，再从暴力求解中寻求优化方法，才是至关重要的。

所以本文题解想要给大家一个从暴力求解，到剪枝的优化算法。整个算法甚至更快于一些动态规划算法，并且超过98.88%的提交
![image.png](https://pic.leetcode-cn.com/8489bb5fece00affbac741d87a17f91bc4dd995b318a102b1da0d7fd4f205100-image.png)

## 暴力求解（超时）
```java
class Solution {
    public boolean find(int[] nums, int target, int index){
        if( target == 0) return true;
        for(int i=index; i<nums.length; i++){
            if(target-nums[i]<0) return false;
            if(find(nums, target-nums[i], i+1)) return true;
        }
        return false;
    }
    public boolean canPartition(int[] nums) {
        int total=0;
        for(int num: nums) total+=num;
        if(total %2 !=0) return false;
        if(total==0) return true;
        return find(nums, total/2, 0);
    }
}
```

## 剪枝
区别微乎其微，仅仅加入两行 --> 14行 和 5行。 直接从超时的代码，变为超越98.88%的提交，是不是很神奇呢？
```java
class Solution {
    public boolean find(int[] nums, int target, int index){
        if( target == 0) return true;
        for(int i=index; i<nums.length; i++){
            if(i>index && nums[i]==nums[i-1]) continue;
            if(target-nums[i]<0) return false;
            if(find(nums, target-nums[i], i+1)) return true;
        }
        return false;
    }
    public boolean canPartition(int[] nums) {
        int total=0;
        for(int num: nums) total+=num;
        Arrays.sort(nums);
        if(total %2 !=0) return false;
        if(total==0) return true;
        return find(nums, total/2, 0);
    }
}
```
