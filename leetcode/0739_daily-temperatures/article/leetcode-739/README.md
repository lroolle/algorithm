"+++
title = "0739. Daily Temperatures  LeetCode 图解 | 739.每日温度 "
author = ["MisterBooo"]
date = 2020-01-10T03:19:48+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

#  LeetCode 图解 | 739.每日温度

> [0739. Daily Temperatures](https://leetcode-cn.com/problems/daily-temperatures/)
> [ LeetCode 图解 | 739.每日温度](https://leetcode-cn.com/problems/daily-temperatures/solution/leetcode-tu-jie-739mei-ri-wen-du-by-misterbooo/) by [MisterBooo](https://leetcode-cn.com/u/misterbooo/)

今天分享的题目来源于 LeetCode 第 739 号问题：每日温度。

#### 题目描述：

根据每日 `气温` 列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 `0` 来代替。

例如，给定一个列表 `temperatures = [73, 74, 75, 71, 69, 72, 76, 73]`，你的输出应该是 `[1, 1, 4, 2, 1, 1, 0, 0]`。

提示：`气温` 列表长度的范围是 `[1, 30000]`。每个气温的值的均为华氏度，都是在 `[30, 100]` 范围内的整数。

#### 题目解析：

这道题目最 “难” 的一个点是题目的理解。

给定列表 `temperatures = [73, 74, 75, 71, 69, 72, 76, 73]`，为啥输出就是 `[1, 1, 4, 2, 1, 1, 0, 0]`  ？

下面来一个个进行解释。

对于输入 73，它需要 **经过一天** 才能等到温度的升高，也就是在第二天的时候，温度升高到 74 ，所以对应的结果是 1。

对于输入 74，它需要 **经过一天** 才能等到温度的升高，也就是在第三天的时候，温度升高到 75 ，所以对应的结果是 1。

对于输入 75，它经过 1 天后发现温度是 71，没有超过它，继续等，一直 **等了四天**，在第七天才等到温度的升高，温度升高到 76 ，所以对应的结果是 4 。

对于输入 71，它经过 1 天后发现温度是 69，没有超过它，继续等，一直 **等了两天**，在第六天才等到温度的升高，温度升高到 72 ，所以对应的结果是 2 。

对于输入 69，它 **经过一天** 后发现温度是 72，已经超过它，所以对应的结果是 1 。

对于输入 72，它 **经过一天** 后发现温度是 76，已经超过它，所以对应的结果是 1 。

对于输入 76，后续 **没有温度** 可以超过它，所以对应的结果是 0 。

对于输入 73，后续 **没有温度** 可以超过它，所以对应的结果是 0 。

好了，理解了题意我们来思考如何求解。

第一个想法就是针对每个温度值 **向后进行依次搜索** ，找到比当前温度更高的值，这是最容易想到的办法。

其原理是从左到右除了最后一个数其他所有的数都遍历一次，最后一个数据对应的结果肯定是 0，就不需要计算。

遍历的时候，每个数都去向后数，直到找到比它大的数，数的次数就是对应输出的值。

**代码如下：**

```Java
public int[] dailyTemperatures(int[] T) {
    int length = T.length;
    int[] result = new int[length];
    for (int i = 0; i < length; i++) {
        int current = T[i];
        if (current < 100) {
            for (int j = i + 1; j < length; j++) {
                if (T[j] > current) {
                    result[i] = j - i;
                    break;
                }
            }
        }
    }
    return result;
}
```

这个题目的标签是 **栈** ，我们考虑一下怎么借助 **栈** 来解决。

不过这个栈有点特殊，它是 **递减栈** ：栈里只有递减元素。

**具体操作如下：**

遍历整个数组，如果栈不空，且当前数字大于栈顶元素，那么如果直接入栈的话就不是 **递减栈** ，所以需要取出栈顶元素，由于当前数字大于栈顶元素的数字，而且一定是第一个大于栈顶元素的数，直接求出下标差就是二者的距离。

继续看新的栈顶元素，直到当前数字小于等于栈顶元素停止，然后将数字入栈，这样就可以一直保持递减栈，且每个数字和第一个大于它的数的距离也可以算出来。

#### 动画理解
![739.m4v](c868335a-e5fd-40e2-913d-d0c8436cb5c8)
#### 复杂度分析

该方法只需要对数组进行一次遍历，每个元素最多被压入和弹出堆栈一次，算法复杂度是 $O(n)$。
#### 代码实现

```c++
class Solution {
public:
    vector<int> dailyTemperatures(vector<int>& temperatures) {
        int n = temperatures.size();
        vector<int> res(n, 0);
        stack<int> st;
        for (int i = 0; i < temperatures.size(); ++i) {
            while (!st.empty() && temperatures[i] > temperatures[st.top()]) {
                auto t = st.top(); st.pop();
                res[t] = i - t;
            }
            st.push(i);
        }
        return res;
    }
};
```

#### 相关题目推荐

利用堆栈，还可以解决如下常见问题：

- 求解算术表达式的结果（LeetCode 224、227、772、770)
- 求解直方图里最大的矩形区域（LeetCode 84）