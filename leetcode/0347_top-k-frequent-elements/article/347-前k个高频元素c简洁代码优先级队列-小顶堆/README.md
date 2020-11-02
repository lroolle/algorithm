"+++
title = "0347. Top K Frequent Elements 347 前K个高频元素：C++简洁代码（优先级队列-小顶堆） "
author = ["carlsun-2"]
date = 2020-07-28T06:40:19+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 347 前K个高频元素：C++简洁代码（优先级队列-小顶堆）

> [0347. Top K Frequent Elements](https://leetcode-cn.com/problems/top-k-frequent-elements/)
> [347 前K个高频元素：C++简洁代码（优先级队列-小顶堆）](https://leetcode-cn.com/problems/top-k-frequent-elements/solution/347-qian-kge-gao-pin-yuan-su-cjian-ji-dai-ma-you-x/) by [carlsun-2](https://leetcode-cn.com/u/carlsun-2/)

## 思路

这道题目主要涉及到如下三块内容：
1. 要统计元素出现频率
2. 对频率排序
3. 找出前K个高频元素 

首先统计元素出现的频率，这一类的问题可以使用map来进行统计。

然后是对频率进行排序，这里我们可以使用一种 容器适配器就是优先级队列。 为什么不用快排呢， 使用快排我们要向map转换为vector的结构，然后对整个数组进行排序， 而这种场景下，我们其实只需要维护k个有序的序列就可以了，所以使用优先级队列是最优的。 **需要注意的是我们要定一个小顶堆** 因为要统计最大前k个元素，只有小顶堆每次将最小的元素弹出，最后小顶堆里积累的才是前k个最大元素。

最后我们从优先级队列里找出前k个元素，就可以了。

如图所示：
![347.前K个高频元素.png](https://pic.leetcode-cn.com/1599441297-IZiLKo-347.%E5%89%8DK%E4%B8%AA%E9%AB%98%E9%A2%91%E5%85%83%E7%B4%A0.png)
## C++代码
```
// 时间复杂度：O(nlogk)
// 空间复杂度：O(n)
class Solution {
public:
    // 小顶堆
    class mycomparison {
    public:
        bool operator()(const pair<int, int>& lhs, const pair<int, int>& rhs) {
            return lhs.second > rhs.second;
        }
    };
    vector<int> topKFrequent(vector<int>& nums, int k) {
        // 要统计元素出现频率
        unordered_map<int, int> map; // map<nums[i],对应出现的次数>
        for (int i = 0; i < nums.size(); i++) {
            map[nums[i]]++;
        }

        // 对频率排序
        // 定义一个小顶堆，大小为k
        priority_queue<pair<int, int>, vector<pair<int, int>>, mycomparison> pri_que;
        for (unordered_map<int, int>::iterator it = map.begin(); it != map.end(); it++) {
            pri_que.push(*it);
            if (pri_que.size() > k) {
                pri_que.pop();
            }
        }

        // 找出前K个高频元素，因为小顶堆先弹出的是最小的，所以倒叙来数值数组
        vector<int> result(k);
        for (int i = k - 1; i >= 0; i--) {
            result[i] = pri_que.top().first;
            pri_que.pop();
        }
        return result;

    }
};

```

**我是程序员Carl，先后在腾讯和百度从事技术研发多年，利用工作之余重刷leetcode，更多精彩文章尽在公众号：[代码随想录](https://img-blog.csdnimg.cn/20200815195519696.png)，关注后，回复「Java」「C++」「python」「简历模板」等等，有我整理多年的学习资料，可以加我[微信](https://img-blog.csdnimg.cn/20200814140330894.png)，备注「组队刷题」，拉你进入刷题群（无任何广告，纯个人分享），每天一道经典题目分析，我选的每一道题目都不是孤立的，而是由浅入深一脉相承的，如果跟住节奏每篇连续着看，定会融会贯通。本文  [https://github.com/youngyangyang04/leetcode-master](https://github.com/youngyangyang04/leetcode-master ) 已经收录，里面还有leetcode刷题攻略、各个类型经典题目刷题顺序、思维导图，看一看一定会有所收获，如果给你有帮助也给一个star支持一下吧！**
**以下资料希望对你有帮助：**
* [手把手带你读C++ primer！B站](https://www.bilibili.com/video/BV1Z5411874t)
* [C++ primer 第一章，你要知道的知识点还有这些！B站](https://www.bilibili.com/video/BV1Kv41117Ya)
* [C++面试&C++学习指南知识点整理](https://github.com/youngyangyang04/TechCPP)
* [程序员应该如何写简历（附简历模板）](https://mp.weixin.qq.com/s/PkBpde0PV65dJjj9zZJYtg)
* [leetcode刷题攻略](https://github.com/youngyangyang04/leetcode-master)

**如果感觉题解对你有帮助，不要吝啬给一个👍吧！**
