"+++
title = "0739. Daily Temperatures 【程序员的自我修养】739. Daily Temperatures "
author = ["aver58"]
date = 2019-11-29T01:35:59+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 【程序员的自我修养】739. Daily Temperatures

> [0739. Daily Temperatures](https://leetcode-cn.com/problems/daily-temperatures/)
> [【程序员的自我修养】739. Daily Temperatures](https://leetcode-cn.com/problems/daily-temperatures/solution/cheng-xu-yuan-de-zi-wo-xiu-yang-739-daily-temperat/) by [aver58](https://leetcode-cn.com/u/aver58/)

#### 程序员的自我修养
#### 解题思路

第一个温度值是 23 摄氏度，它要经过 1 天才能等到温度的升高，也就是在第二天的时候，温度升高到 24 摄氏度，所以对应的结果是 1。
接下来，从 25 度到下一次温度的升高需要等待 4 天的时间，那时温度会变为 26 度。
 
#### 思路 1：暴力法 $O(n2)$ 每个去遍历
最直观的做法就是针对每个温度值向后进行依次搜索，找到比当前温度更高的值，这样的计算复杂度就是 $O(n2)$。
 
但是，在这样的搜索过程中，产生了很多重复的对比。
例如，从 25 度开始往后面寻找一个比 25 度更高的温度的过程中，经历了 21 度、19 度和 22 度，而这是一个温度由低到高的过程，
也就是说在这个过程中已经找到了 19 度以及 21 度的答案，它就是 22 度。


#### 思路 2：一个堆栈 stack 算法复杂度是 $O(n)$ 216 ms 98%
可以运用一个堆栈 `stack` 来快速地知道需要经过多少天就能等到温度升高。
从头到尾扫描一遍给定的数组 T，如果当天的温度比堆栈 `stack` 顶端所记录的那天温度还要高，那么就能得到结果。
        
对第一个温度 23 度，堆栈为空，把它的下标压入堆栈；
下一个温度 24 度，高于 23 度高，因此 23 度温度升高只需 1 天时间，把 23 度下标从堆栈里弹出，把 24 度下标压入；
同样，从 24 度只需要 1 天时间升高到 25 度；
21 度低于 25 度，直接把 21 度下标压入堆栈；
19 度低于 21 度，压入堆栈；
22 度高于 19 度，从 19 度升温只需 1 天，从 21 度升温需要 2 天；
由于堆栈里保存的是下标，能很快计算天数；
22 度低于 25 度，意味着尚未找到 25 度之后的升温，直接把 22 度下标压入堆栈顶端；
后面的温度与此同理。
该方法只需要对数组进行一次遍历，每个元素最多被压入和弹出堆栈一次，算法复杂度是 $O(n)$。
![739.gif](https://pic.leetcode-cn.com/7a133e857271e638c04b3a27c1eabc29570e585cc44d7da60eb039459a7f89cd-739.gif)


```
	vector<int> dailyTemperatures(vector<int>& T) {
		stack<int> helpSatck;
		int length = T.size();
		vector<int> res(length);
		for (int i = 0; i < length; i++)
		{
			while (!helpSatck.empty() && T[i] > T[helpSatck.top()])
			{
				res[helpSatck.top()] = i - helpSatck.top();
				helpSatck.pop();
			}
			
			helpSatck.push(i);
		}

		return res;
	}
```

