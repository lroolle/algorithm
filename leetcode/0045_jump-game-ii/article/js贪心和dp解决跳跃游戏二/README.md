"+++
title = "0045. Jump Game II js贪心和DP解决跳跃游戏二 "
author = ["50-hui"]
date = 2020-09-03T03:18:37+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# js贪心和DP解决跳跃游戏二

> [0045. Jump Game II](https://leetcode-cn.com/problems/jump-game-ii/)
> [js贪心和DP解决跳跃游戏二](https://leetcode-cn.com/problems/jump-game-ii/solution/jstan-xin-he-dpjie-jue-tiao-yue-you-xi-er-by-50-hu/) by [50-hui](https://leetcode-cn.com/u/50-hui/)

### 解题思路
    题目和跳跃游戏的第一版本差不多；
    思路使用map方法，将nums里面的值变为能跳到的最大的值,
        定义一个与nums等长数组flag保存跳到下标的最小步数，
        while循环1
            循环2本步能到达的索引
                将flag对应的索引值赋值为本步的步数step
                找出下步能到达的最大索引并赋值给max
            结束 循环二

            判断flag[nums.length]是否被赋值了，
                赋值说明就是能到达终点，步数为flag[nums.length]的值
                有值跳出循环一或者直接返回flag[nums.length];

            判断maxIndex 是否大于等于max 
                大于的话直接跳出循环，表示下一步不能向前面跳了，返回0
            
            将步数step+1;
            将至少下一部到达的最小索引即(maxIndex+1)赋值给minIndex(下一步最小)，
            将max赋值给maxIndex（下一步最大）;

    返回flag[nums.length]; 表示能到该索引的最小距离
### 代码

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var jump = function(nums) {
    let lastlength = nums.length;
    if(lastlength<=1) return 0;
    nums = nums.map((item,i)=>item+i);
    flag = new Array(lastlength).fill(0);
    let max = 0;
    let minIndex = 0;
    let maxIndex = 0;
    let step = 0;
    let count = 0; 
    while(true){
        for(let i = minIndex;i<=maxIndex;i++){
                flag[i] = step;
                max = nums[i]>max?nums[i]:max;
        }
        if(flag[lastlength-1]){
            break;
        }
        // max = Math.max(...nums.slice(minIndex,maxIndex+1))
        step++;
        minIndex = maxIndex+1;
        // console.log(nums,max,maxIndex);
        if(maxIndex>=max){
           return 0;
        }
         maxIndex = max;
    }
    // console.log(flag[lastlength]);
    return flag[lastlength-1];
};
```