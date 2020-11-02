"+++
title = "0416. Partition Equal Subset Sum JAVA DFS+剪枝 1ms 99.31% "
author = ["lava-4"]
date = 2020-01-31T17:13:51+08:00
tags = ["Leetcode", "Algorithms", ]
categories = ["leetcode"]
draft = false
+++

# JAVA DFS+剪枝 1ms 99.31%

> [0416. Partition Equal Subset Sum](https://leetcode-cn.com/problems/partition-equal-subset-sum/)
> [JAVA DFS+剪枝 1ms 99.31%](https://leetcode-cn.com/problems/partition-equal-subset-sum/solution/java-dfsjian-zhi-1ms-9931-by-lava-4/) by [lava-4](https://leetcode-cn.com/u/lava-4/)

把问题抽象为：
是否存在集合S的一个子集X，使得SUM(X)=SUM(S)/2?
S大小N不大于200，因此搜索+剪枝是完全可以AC的（偷懒没往DP上想hhhhhhhh）

策略：
对于S中的每一个元素，对于子集X，都有接受和丢弃两种操作，且这两种操作是互补且等价的
（可以想象有两个子集，接受是放到第一个子集中，丢弃是放到第二个子集里）。
1、当任意一个子集装满SUM(S)/2后，即成功找到一个解。
2、当任意一个子集超过SUM(S)/2后，则在此种分支下不可能找到一个可行解，剪枝。

代码：
```
class Solution {
    public boolean canPartition(int[] nums) {
        //涉及到剪枝的问题，先排个序
        Arrays.sort(nums);
        int sum = 0;
        //算出SUM(S)
        for (int n : nums){
            sum += n;
        }
        //奇数肯定不行
        if ((sum & 1) == 1){
            return false;
        }
        sum >>= 1;
        //搜索
        return canPartition(nums, nums.length-1, sum, sum);
    }

    //DFS idx为当前元素，had为待接受上限，pass为待丢弃上限
    private boolean canPartition(int[] nums, int idx, int had, int pass){
        //找到可行解
        if (had == 0 || pass == 0){
            return true;
        }
        //剪枝
        else if (had < 0 || pass < 0){
            return false;
        }
        //继续搜索
        else{
            return canPartition(nums, idx-1, had-nums[idx], pass) || canPartition(nums, idx-1, had, pass-nums[idx]);
        }
    }
}
```
