"+++
title = "0739. Daily Temperatures 【手绘图解】单调栈思路的形成+模板套路 "
author = ["xiao_ben_zhu"]
date = 2020-06-10T23:29:55+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "Stack"]
categories = ["leetcode"]
draft = false
+++

# 【手绘图解】单调栈思路的形成+模板套路

> [0739. Daily Temperatures](https://leetcode-cn.com/problems/daily-temperatures/)
> [【手绘图解】单调栈思路的形成+模板套路](https://leetcode-cn.com/problems/daily-temperatures/solution/shou-hui-ti-jie-fang-da-guan-cha-dan-diao-zhan-si-/) by [xiao_ben_zhu](https://leetcode-cn.com/u/xiao_ben_zhu/)


### 暴力法
- 题意：每个元素找到它右边第一个比它大的元素的位置，求它们的距离
- $i$ 指向当前元素，$j$ 扫描右边的区间，直到找到比当前元素大的元素，记录 $j-i$
- 然后 $i$ 考察下一位，$j$ 重复上述过程。有两重循环，时间复杂度 $O(n^2)$

![image.png](https://pic.leetcode-cn.com/ac538e728d21576111f4bef239320c015f17462588bc6322ca2b43822dc007cc-image.png)

```
const dailyTemperatures = (T) => {
  const res = new Array(T.length).fill(0)
  for (let i = 0; i < T.length; i++) {
    for (let j = i + 1; j < T.length; j++) {
      if (T[j] > T[i]) {
        res[i] = j - i
        break
      }
    }
  }
  return res
}
```
### 单调栈 解法
- 当前元素关注的是它**右边**的元素
- 我选择从右开始遍历，需要考察的右边元素由少到多
### 开辟空间，换取时间
- 遍历一次数组是少不了的，关键是每个元素寻找目标元素时做了很多重复的遍历
- 我们用一个数据结构去存储右边的项，想将无需重复比较的元素从该数据结构中剔除，且不会再进来了，避免重复遍历
### 什么剔除，什么留下
- $T[i]$ 目标是找到第一个大项，小项被 pass 掉，该被剔除
- 因为比 $T[i]$ 还小的元素，肯定不会是 $T[i-1]$ 想找的大项
- 大项留下，$T[i]$ 也进入到这个“数据结构”中，供 $T[i-1]$ 寻找
### 保持单调，为了更快比较
- 如果“数据结构”中的项无序，则新来的项无法快速地比较大小，找到大项
- 如果从小到大排好，则很容易剔除小项，遇到大项
### 为什么是栈
- 小项 “出” 了，当前项要 “进”，依然要维持排列的单调性
- 所以当前项要从小项 “出” 的地方 “进”
- 只在一端进出——所以是栈
 
![image.png](https://pic.leetcode-cn.com/0051296ee260c55479a5dc139fedadc99d0ee32f30a7d27d216c4c4dcc51c7ad-image.png)

### 具体操作
- 如果当前元素比栈顶大，则让小项逐个出栈，直到当前元素比栈顶小，停止出栈
- 此时的栈顶元素就是当前项右边的第一个比自己大的元素索引，计算距离
- 当前项入栈
- 所以说，当前元素还是要入栈的，只是在入栈前要维护好栈的单调性


![image.png](https://pic.leetcode-cn.com/a24107cefeff7239068268099db90671254c2d357857232f19dc21bdaace5774-image.png)
### 复杂度分析
- 一次遍历：O(n)；每个元素都出栈入栈各一次：线性时间的复杂度。综合：O(n)
- 空间复杂度：O(n)

```js 
const dailyTemperatures = (T) => {
  const res = new Array(T.length).fill(0)
  const stack = []
  for (let i = T.length - 1; i >= 0; i--) {
    while (stack.length && T[i] >= T[stack[stack.length - 1]]) {
      stack.pop()
    }
    if (stack.length) {
      res[i] = stack[stack.length - 1] - i
    }
    stack.push(i)
  }
  return res
}
```
### 总结部分
- 单调递增栈：从 栈底 到 栈顶 递增，栈顶大
- 单调递减栈：从 栈底 到 栈顶 递减，栈顶小

### 什么时候用单调栈
- 通常是一维数组，要寻找任一元素右边（左边）第一个比自己大（小）的元素
- 且要求 O(n) 的时间复杂度
### 模板套路
- **单调递增栈会剔除波峰，留下波谷；单调递减栈会剔除波谷，留下波峰**
![image.png](https://pic.leetcode-cn.com/0051296ee260c55479a5dc139fedadc99d0ee32f30a7d27d216c4c4dcc51c7ad-image.png)

    - 当前项向左找第一个比自己大的位置 —— 从左向右维护一个单调递减栈
    - 当前项向左找第一个比自己小的位置 —— 从左向右维护一个单调递增栈
    - 当前项向右找第一个比自己大的位置 —— 从右向左维护一个单调递减栈
    - 当前项向右找第一个比自己小的位置 —— 从右向左维护一个单调递增栈
- 用伪代码描述：
    ``` 
    insert x
    while !stack.empty() && stack.top()<x
        stack.pop()
    stack.push(x)
    ```


#### 如果有帮助，点个赞鼓励鼓励我继续坚持写下去，如果哪里不对 不好，指出我我继续修改。