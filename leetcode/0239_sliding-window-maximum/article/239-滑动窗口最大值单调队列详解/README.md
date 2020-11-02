"+++
title = "0239. Sliding Window Maximum 239. 滑动窗口最大值:【单调队列】详解！ "
author = ["carlsun-2"]
date = 2020-08-12T12:12:11+08:00
tags = ["Leetcode", "Algorithms", "cpp", "单调队列"]
categories = ["leetcode"]
draft = false
+++

# 239. 滑动窗口最大值:【单调队列】详解！

> [0239. Sliding Window Maximum](https://leetcode-cn.com/problems/sliding-window-maximum/)
> [239. 滑动窗口最大值:【单调队列】详解！](https://leetcode-cn.com/problems/sliding-window-maximum/solution/239-hua-dong-chuang-kou-zui-da-zhi-dan-diao-dui-3/) by [carlsun-2](https://leetcode-cn.com/u/carlsun-2/)

## 思路

这是使用单调队列的经典题目。

暴力方法，遍历一遍的过程中每次从窗口中在找到最大的数值，这样很明显是O(n * k)的算法。

有的同学可能会想用一个大顶堆也就是优先级队列来存放这个窗口里的k个数字，这样就可以知道最大的最大值是多少了， **但是问题是这个窗口是移动的，而大顶堆每次只能弹出最大值，我们无法移除其他数值，这样就造成大顶堆维护的不是滑动窗口里面的数值了。所以不能用大顶堆。**

**使用单调队列，即单调递减或单调递增的队列。它不是一个独立的数据结构，需要使用其他数据结构来实现单调队列**，例如： deque，deque是双向队列，可以选择 从头部或者尾部pop，同样也可以从头部或者尾部push。

不要以为实现的单调队列就是 对窗口里面的数进行排序，如果排序的话，那和优先级队列又有什么区别了呢。

使用deque实现的单调队列如下：（代码详细注释）

```
class MyQueue { //单调队列（从大到小）
public:
    deque<int> que; // 使用deque来实现单调队列
    // 每次弹出的时候，比较当前要弹出的数值是否等于队列前端的数值，如果相等在pop数据，当然也要判断队列当前是否为空。
    void pop(int value) {
        if (!que.empty() && value == que.front()) {
            que.pop_front();
        }
    }
    // 如果push的数值大于后端的数值，那么就将队列后端的数值弹出，直到push的数值小于等于 队列后端的数值为止。 
    // 然后再将数值push到队列后端，这样就保持了队列里的数值是单调从大到小的了。
    void push(int value) {
        while (!que.empty() && value > que.back()) {
            que.pop_back();
        }
        que.push_back(value);

    }
    // 查询当前队列里的最大值 直接返回队列前端也就是front就可以了。
    int front() {
        return que.front();
    }
};
```

动画解释如下：
![0239.滑动窗口最大值.mp4](7a4c6732-5867-42d0-8ea9-2fb8880b28ae)
这样我们就用deque实现了一个单调队列，接下来解决滑动窗口最大值的问题就很简单了。 

详情看代码吧，已经简洁。

## C++代码

```
class Solution {
public:
    class MyQueue { //单调队列（从大到小）
    public:
        deque<int> que; // 使用deque来实现单调队列
        void pop(int value) {
            if (!que.empty() && value == que.front()) {
                que.pop_front();
            }
        }
        void push(int value) {
            while (!que.empty() && value > que.back()) {
                que.pop_back();
            }
            que.push_back(value);

        }
        int front() {
            return que.front();
        }
    };
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        MyQueue que;
        vector<int> result;
        for (int i = 0; i < k; i++) { // 先将前k的元素放进队列
            que.push(nums[i]);
        }
        result.push_back(que.front()); // result 记录前k的元素的最大值
        for (int i = k; i < nums.size(); i++) {
            que.pop(nums[i - k]); // 模拟滑动窗口的移动
            que.push(nums[i]); // 模拟滑动窗口的移动
            result.push_back(que.front()); // 记录对应的最大值
        }
        return result;
    }
};
```
来看一下时间复杂度，时间复杂度： O(n)，
有的同学可能想了，在队里中 push元素的过程中，还有pop操作呢，感觉不是纯粹了O(n)。

其实，大家可以自己观察一下单调队列的实现，nums 中的每个元素最多也就被 push_back 和 pop_back 一次，没有任何多余操作，所以整体的复杂度还是 O(n)。

空间复杂度因为我们定义一个辅助队列，所以是O(k)。

--------end--------

**我是程序员Carl，先后在腾讯和百度从事技术研发多年，利用工作之余重刷leetcode，更多[精彩算法文章](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzUxNjY5NTYxNA==&action=getalbum&album_id=1485825793120387074&scene=173#wechat_redirect)尽在公众号：[代码随想录](https://img-blog.csdnimg.cn/20200815195519696.png)，关注后，回复「Java」「C++」「python」「简历模板」等等，有我整理多年的学习资料，可以加我[微信](https://img-blog.csdnimg.cn/20200814140330894.png)，备注「简单自我介绍」+「组队刷题」，拉你进入刷题群（无任何广告，纯个人分享），每天一道经典题目分析，我选的每一道题目都不是孤立的，而是由浅入深一脉相承的，如果跟住节奏每篇连续着看，定会融会贯通。本文  [https://github.com/youngyangyang04/leetcode-master](https://github.com/youngyangyang04/leetcode-master ) 已经收录，里面还有leetcode刷题攻略、各个类型经典题目刷题顺序、思维导图，看一看一定会有所收获，如果对你有帮助也给一个star支持一下吧！**
**以下资料希望对你有帮助：**

* [我在B站上讲KMP算法！](https://www.bilibili.com/video/BV1PD4y1o7nd/)
* [如何学习C++？ B站](https://www.bilibili.com/video/BV1rK4y1e7ed)
* [手把手带你读C++ primer！B站](https://www.bilibili.com/video/BV1Z5411874t)
* [leetcode刷题攻略](https://github.com/youngyangyang04/leetcode-master)
* [程序员应该如何写简历（附简历模板）](https://mp.weixin.qq.com/s/PkBpde0PV65dJjj9zZJYtg)
* [一线互联网公司技术面试的流程以及注意事项](https://mp.weixin.qq.com/s/1VMvQ_6HbVpEn85CNilTiw)
* [C++面试&C++学习指南知识点整理](https://github.com/youngyangyang04/TechCPP)

**如果感觉题解对你有帮助，不要吝啬给一个👍吧！**

