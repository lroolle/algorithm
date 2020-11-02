"+++
title = "0242. Valid Anagram 242. 有效的字母异位词:【数组在哈希法中的应用】详解！ "
author = ["carlsun-2"]
date = 2020-08-16T07:06:47+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 242. 有效的字母异位词:【数组在哈希法中的应用】详解！

> [0242. Valid Anagram](https://leetcode-cn.com/problems/valid-anagram/)
> [242. 有效的字母异位词:【数组在哈希法中的应用】详解！](https://leetcode-cn.com/problems/valid-anagram/solution/242-you-xiao-de-zi-mu-yi-wei-ci-shu-zu-zai-ha-xi-f/) by [carlsun-2](https://leetcode-cn.com/u/carlsun-2/)

## 思路

先说一个特殊示例，输入：s = “car”, t= “car”，输出应该是什么呢？ 后台判题逻辑返回的依然是true，其实这道题目真正题意是字符串s是否可以排练组合为字符串t。

那么来看一下应该怎么做， 首先是暴力的解法，两层for循环，同时还要记录 字符是否重复出现，很明显时间复杂度是 O(n^2)，暴力的方法这里就不做介绍了，直接看一下有没有更优的方式。

数组其实就是一个简单哈希表，而且这道题目中字符串只有小写字符，那么就可以定义一个数组，来记录字符串s里字符出现的次数。

需要定义一个多大的数组呢，定一个数组叫做record，大小为26 就可以了，初始化为0，因为字符a到字符z的ASCII也是26个连续的数值。

为了方便举例，判断一下字符串s= "aee", t = "eae"。 

操作动画如下：

![242.有效的字母异位词.gif](https://pic.leetcode-cn.com/14a4707c5995f3754c058f68e4241bbe0d3e18c16bc414379daa4e5055b8a995-242.%E6%9C%89%E6%95%88%E7%9A%84%E5%AD%97%E6%AF%8D%E5%BC%82%E4%BD%8D%E8%AF%8D.gif)

定义一个数组叫做record用来上记录字符串s里字符出现的次数。

如何记录呢，需要把字符映射到数组也就是哈希表的索引下表上，那么**因为字符a到字符z的ASCII是26个连续的数值，所以字符a映射为下表0，相应的字符z映射为下表25。**

再遍历 字符串s的时候，**只需要将 s[i] - ‘a’ 所在的元素做+1 操作即可，并不需要记住字符a的ASCII，只要求出一个相对数值就可以了。** 这样就将字符串s中字符出现的次数，统计出来了。
那看一下如何检查 字符串t中是否出现了这些字符，同样在遍历字符串t的时候，对t中出现的字符映射哈希表索引上的数值再做-1的操作。

那么最后检查一下，**record数组如果有的元素不为零0，说明字符串s和t 一定是谁多了字符或者谁少了字符，return false。**

最后如果record数组所有元素都为零0，说明字符串s和t是字母异位词，return true。

时间复杂度为O(n)，空间上因为定义是的一个常量大小的辅助数组，所以空间复杂度为O(1)。

## C++ 代码
```
class Solution {
public:
    bool isAnagram(string s, string t) {
        int record[26] = {0};
        for (int i = 0; i < s.size(); i++) {
            // 并不需要记住字符a的ASCII，只要求出一个相对数值就可以了
            record[s[i] - 'a']++;
        }
        for (int i = 0; i < t.size(); i++) {
            record[t[i] - 'a']--;
        }
        for (int i = 0; i < 26; i++) {
            if (record[i] != 0) {
                // record数组如果有的元素不为零0，说明字符串s和t 一定是谁多了字符或者谁少了字符。
                return false;
            }
        }
        // record数组所有元素都为零0，说明字符串s和t是字母异位词
        return true;
    }
};
```

**我是程序员Carl，先后在腾讯和百度从事技术研发多年，利用工作之余重刷leetcode，更多精彩文章尽在公众号：[代码随想录](https://img-blog.csdnimg.cn/20200815195519696.png)，关注后，回复「Java」「C++」「python」「简历模板」等等，有我整理多年的学习资料，可以加我[微信](https://img-blog.csdnimg.cn/20200814140330894.png)，备注「组队刷题」，拉你进入刷题群，每天一道经典题目分析，我选的每一道题目都不是孤立的，而是由浅入深一脉相承的，如果跟住节奏每篇连续着看，定会融会贯通。本文  [https://github.com/youngyangyang04/leetcode-master](https://github.com/youngyangyang04/leetcode-master ) 已经收录，里面还有leetcode刷题攻略、各个类型经典题目刷题顺序、思维导图，看一看一定会有所收获！**

**以下资料希望对你有帮助：**
* [如何学习C++B站视频](https://www.bilibili.com/video/BV1rK4y1e7ed)
* [手把手带你读C++ primer](https://www.bilibili.com/video/BV1Z5411874t)
* [C++面试&C++学习指南知识点整理](https://github.com/youngyangyang04/TechCPP)
* [程序员应该如何写简历（附简历模板）](https://mp.weixin.qq.com/s/PkBpde0PV65dJjj9zZJYtg)
* [leetcode刷题攻略](https://github.com/youngyangyang04/leetcode-master)

**哈希表经典题目**
* [关于哈希表，你该了解这些！](https://mp.weixin.qq.com/s/g8N6WmoQmsCUw3_BaWxHZA)
* [哈希表：可以拿数组当哈希表来用，但哈希值不要太大](https://mp.weixin.qq.com/s/vM6OszkM6L1Mx2Ralm9Dig)
* [哈希表：哈希值太大了，还是得用set](https://mp.weixin.qq.com/s/N9iqAchXreSVW7zXUS4BVA)
* [哈希表：今天你快乐了么？](https://mp.weixin.qq.com/s/G4Q2Zfpfe706gLK7HpZHpA)
* [哈希表：map等候多时了](https://mp.weixin.qq.com/s/uVAtjOHSeqymV8FeQbliJQ)
* [哈希表：其实需要哈希的地方都能找到map的身影](https://mp.weixin.qq.com/s/Ue8pKKU5hw_m-jPgwlHcbA)
* [哈希表：这道题目我做过？](https://mp.weixin.qq.com/s/sYZIR4dFBrw_lr3eJJnteQ)
* [哈希表：解决了两数之和，那么能解决三数之和么？](https://mp.weixin.qq.com/s/r5cgZFu0tv4grBAexdcd8A)
* [双指针法：一样的道理，能解决四数之和](https://mp.weixin.qq.com/s/nQrcco8AZJV1pAOVjeIU_g)

**如果感觉题解对你有帮助，不要吝啬给一个👍吧！**
