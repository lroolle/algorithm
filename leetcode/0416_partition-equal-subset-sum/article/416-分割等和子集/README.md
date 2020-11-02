"+++
title = "0416. Partition Equal Subset Sum 416. 分割等和子集 "
author = ["Geralt_U"]
date = 2020-05-20T10:44:13+08:00
tags = ["Leetcode", "Algorithms", "Java", "DynamicProgramming"]
categories = ["leetcode"]
draft = false
+++

# 416. 分割等和子集

> [0416. Partition Equal Subset Sum](https://leetcode-cn.com/problems/partition-equal-subset-sum/)
> [416. 分割等和子集](https://leetcode-cn.com/problems/partition-equal-subset-sum/solution/416-fen-ge-deng-he-zi-ji-by-ming-zhi-shan-you-m9rf/) by [Geralt_U](https://leetcode-cn.com/u/geralt_u/)

### 解题思路
##### 0-1背包问题

本题是一道0-1背包的动态规划题目。首先在这里介绍一下0-1背包问题。

假设现在有1个背包和N个物品。背包的最大载重为$w$。每个物品有两个属性：重量和价格。第$i$个物品的重量为$w[i]$，价值为$p[i]$。问这个背包最多能装多少价值的东西？

这是典型的动态规划问题。对于上述的物品来说，只有两个状态:**"不放进背包（0）"**和**“放进背包（1）”**。所以这种一般都称为**0-1背包问题**。

按照之前大佬所说的框架第一步要**明确状态与选择**。

- 状态：对背包而言，**背包的重量**，**背包能够选择的物品**

- 选择：对物品而言，选择物品**放进背包**或**不放进背包**

然而，动态规划就是在某种状态下做出最优的选择。

接下来需要定义dp数组。

定义dp数组：**$dp[i][j]$表示在背包载重最大为$j$的情况下，对前$i$个物品进行选择，所能装的最大价值。**

例如：$dp[3][4] = 5$表示在背包载重最大为4的情况下，在前3个物品中进行选择后，所能装的最大价值为5。

回到我们的0-1背包问题中，我们最终想要的到的就是$dp[N][w]$。

base case为:

- $dp[0][\cdots] = 0$，表示在各种背包载重的情况下，在前0个物品中进行选择（也就是没有物品可选），所以最大价值为0.
- $dp[\cdots][0]$，表示在背包载重为0的情况下，对于前任意物品选择，也都放不进去。所以最大价值为0。
对于上述问题来说，有一个基本的框架。

```java
int[][] dp = new int[N+1][w+1];
dp[0][...] = 0;
dp[...][0] = 0;
for(int i = 1; i<=N; i++){
    for(int j = 1; j<=w; j++){
        dp[i][j] = max(不选择把物品i放进背包，选择把物品i放进背包)
    }
}
return dp[N][w];
```

其中关键的点就在于$dp[i][j]$ = max(不选择把物品i放进背包，选择把物品i放进背包).

- 不选择把i放进去：说明当前的$dp[i][j]$只能继承之前的$dp[i-1][j]$的状态，因为没有放里面。
  $$
  dp[i][j] = dp[i-1][j]
  $$

- 选择把i放进去
  $$
  dp[i][j] = dp[i-1][j-w[i-1]] + w[i-1]
  $$
  这里需要强调一下，因为dp中的 i 是从 1 开始的，而w索引是从 0 开始的，所以第 i 个物品的重量应该是 w[i-1]。这里怎么理解呢？把i放进去，肯定价值中肯定就有$w[i-1]$了，那背包里面之前的价值是多少呢？就是在没放i之前背包的最大载重是$j-w[i-1]$，那就是$dp[i-1][j-w[i-1]]$了。将这两部分加在一起就是上式。

那0-1背包动态规划的基本框架就是

```java
int[][] dp = new int[N+1][w+1];
dp[0][...] = 0;
dp[...][0] = 0;
for(int i = 1; i<=N; i++){
    for(int j = 1; j<=w; j++){
        dp[i][j] = max(dp[i-1][j],
                       dp[i-1][j-w[j-1]] + w[i-1]);
    }
}
return dp[N][w];
```
##### 本题题解

解释完0-1背包问题后，回到本题，要求分割成**两个等和的子集**。乍一看好像没有什么动态规划的思路。但我们换一种说法，能不能找出**和为数组总和一半的子集**。这就相当给你一个背包的最大载重量为数组总和的一半，对于N个元素，能不能正好放进去。

这里我们定义dp数组：

**$dp[i][j]$表示 在载重量为$j$的情况下，如果前$i$个元素能够将背包放满，则$dp[i][j] = true$,如果不能够放满，则$dp[i][j] = false$**。

那我们最终要求的就是$dp[n][sum/2]$。

明确这个定义后我们就转换为了动态规划问题，我们可以按照下的流程写了。

- 计算数组的和sum，如果sum为奇数，那肯定不能分成两部分，直接返回false；
- 如果sum为偶数，按照上面的0-1背包框架声明dp数组。
- 初始化base case :$dp[\cdots][0] = true$，相当于当载重量为0的时候，肯定什么东西也不用放，背包肯定默认是满的，因为载重量为0嘛，所以是true；$dp[0][\cdots] = false$,相当于在任一载重量时，什么东西都不放，那肯定背包没有满，所以是false。
- 接下来按照0-1背包框架进行推算。

### 代码

```java
class Solution {
    public boolean canPartition(int[] nums) {
        int sum = 0;
        //计算数组的和sum，如果sum为奇数，那肯定不能分成两部分，直接返回false
        for(int i = 0; i<nums.length; i++) sum += nums[i];
        if(sum % 2 != 0) return false;
        sum = sum/2;
        //初始化base case :dp[...][0] = true，相当于当载重量为0的时候，肯定什么东西也不用放，背包肯定默认是满的，因为载重量为0嘛，所以是true；dp[0][...] = false,相当于在任一载重量时，什么东西都不放，那肯定背包没有满，所以是false
        boolean[][] dp = new boolean[nums.length+1][sum+1];
        for(int i = 0; i<nums.length+1; i++) dp[i][0] = true;
        //这里可以省略，因为java中boolean量默认是false，这里没有注释掉是因为想把逻辑表达清楚。
        for(int i = 0; i<sum+1; i++) dp[0][i] = false;
        
        
        for(int i = 1; i<=nums.length; i++){
            for(int j = 1; j<=sum; j++){
                //如果当前的背包容量比要放的数量都小，那就没法放，只能继承之前的状态
                if(j < nums[i-1]) dp[i][j] = dp[i-1][j];
                else{
                    //放入或者不放入，不管哪种状态，只要能放满就可以
                    dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]];
                }
            }
        }
        return dp[nums.length][sum];
    }
}
```