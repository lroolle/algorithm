"+++
title = "0416. Partition Equal Subset Sum 「手画图解」分割等和子集 | 记忆化递归 思路详解 "
author = ["xiao_ben_zhu"]
date = 2020-10-11T00:50:07+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 「手画图解」分割等和子集 | 记忆化递归 思路详解

> [0416. Partition Equal Subset Sum](https://leetcode-cn.com/problems/partition-equal-subset-sum/)
> [「手画图解」分割等和子集 | 记忆化递归 思路详解](https://leetcode-cn.com/problems/partition-equal-subset-sum/solution/shou-hua-tu-jie-fen-ge-deng-he-zi-ji-dfshui-su-si-/) by [xiao_ben_zhu](https://leetcode-cn.com/u/xiao_ben_zhu/)

#### 思路

题意其实是：给你一个非空数组，和为`sum`，能否找到一个子序列，和为`sum/2`。

- 如果`sum`为奇数，肯定找不到，因为`sum/2`为小数，而数组只包含正整数，子序列的和也为整数。
- 如果`sum`为偶数，有可能找到。

对于每个元素，都存在「选或不选它」去组成该子序列的问题。

想到了 DFS 回溯，穷举出所有的情况，判断是否能找到这样的子序列。


![image.png](https://pic.leetcode-cn.com/1602372398-tqYjTN-image.png)


上图是解的空间树，每次考察一个元素，用索引`i`描述，还有一个变量：现有已累加的`curSum`，它们共同描述一个节点。

递归函数定义：基于已选的元素，和为`curSum`，从第`i`项开始，继续选，能否选出和为`sum/2`的子序列。

每次递归，都有两个选择：
- 选`nums[i]`。基于选它，往下继续选（递归）：`helper(curSum + nums[i], i + 1)`
- 不选`nums[i]`。基于不选它，往下继续选（递归）：`helper(curSum, i + 1)`

递归的终止条件有三种情况：
- 指针越界，考察完所有元素，始终没有返回true，于是返回false，即没有找到满足条件的子序列。
- `curSum > target`，当前的累加和已经爆了，不用继续选数字了，终止递归，返回false。
- `curSum == target`，正好满足条件，不用继续选数字了，终止递归，返回true。

#### 代码（超时）

```js
const canPartition = (nums) => {
  let sum = 0;
  for (const n of nums) { // 求和
    sum += n;
  }
  if (sum % 2 != 0) return false; // 如果 sum 为奇数，直接返回 false
    
  const target = sum / 2; // 目标和
 
  const helper = (curSum, i) => {                 // curSum是当前累加和，i是指针
    if (i > nums.length - 1 || curSum > target) { // 递归的出口
      return false;
    }
    if (curSum == target) {                       // 递归的出口
      return true;
    }
    // 选nums[i]，当前和变为curSum+nums[i]，考察的指针移动一位
    // 不选nums[i]，当前和还是curSum，考察的指针移动一位
    return helper(curSum + nums[i], i + 1) || helper(curSum, i + 1);
  };

  return helper(0, 0); // 递归的入口，当前和为0，指针为0
};
```
#### 遇到特定case，执行结果超时
![image.png](https://pic.leetcode-cn.com/1602364484-esIxTt-image.png)



#### 加入记忆化
描述一个子问题的两个变量是`curSum`和`i`，我们将它合并为 key，存入 hashMap，值为对应的计算结果。

当遇到重复的子问题时，就直接返回 map 中的缓存值，不用进行重复的递归计算。

#### AC代码
```js
const canPartition = (nums) => {
  let sum = 0;
  for (const n of nums) { // 求和
    sum += n;
  }
  if (sum % 2 != 0) return false; // 如果 sum 为奇数，直接返回 false

  const target = sum / 2; // 目标和
  const memo = new Map();

  const helper = (curSum, i) => { 
    const key = curSum + '&' + i;   // 描述一个问题的key
    if (memo.has(key)) {            // 如果memo中有对应的缓存值，直接使用
      return memo.get(key);
    }
    if (i > nums.length - 1 || curSum > target) { 
      return false;
    }
    if (curSum == target) {    
      return true;
    }
    const res = helper(curSum + nums[i], i + 1) || helper(curSum, i + 1);
    memo.set(key, res);  // 计算的结果存入memo
    return res;
  };

  return helper(0, 0); // 递归的入口，当前和为0，指针为0
};
```


#### 动态规划的解法，有空的话会补充上来，从记忆化递归往动态规划上靠，不难的。感谢阅读，点赞更棒。