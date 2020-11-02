"+++
title = "0141. Linked List Cycle 141. 环形链表:【双指针找环】详解 "
author = ["carlsun-2"]
date = 2020-10-09T01:26:59+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 141. 环形链表:【双指针找环】详解

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [141. 环形链表:【双指针找环】详解](https://leetcode-cn.com/problems/linked-list-cycle/solution/141-huan-xing-lian-biao-shuang-zhi-zhen-zhao-huan-/) by [carlsun-2](https://leetcode-cn.com/u/carlsun-2/)

## 思路 

可以使用快慢指针法，  分别定义 fast 和 slow指针，从头结点出发，fast指针每次移动两个节点，slow指针每次移动一个节点，如果 fast 和 slow指针在途中相遇 ，说明这个链表有环。

为什么fast 走两个节点，slow走一个节点，有环的话，一定会在环内相遇呢，而不是永远的错开呢？

首先第一点： **fast指针一定先进入环中，如果fast 指针和slow指针相遇的话，一定是在环中相遇，这是毋庸置疑的。**

那么来看一下，**为什么fast指针和slow指针一定会相遇呢？**

可以画一个环，然后让 fast指针在任意一个节点开始追赶slow指针。 

会发现最终都是这种情况， 如下图：

![142环形链表1.png](https://pic.leetcode-cn.com/1602206805-izTwNO-142%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A81.png)
fast和slow各自再走一步， fast和slow就相遇了

这是因为fast是走两步，slow是走一步，**其实相对于slow来说，fast是一个节点一个节点的靠近slow的**，所以fast一定可以和slow重合。

动画如下：

![142.环形链表.mp4](78e94ea4-906c-405d-86ab-5fea963eb433)
## C++代码如下 

```
class Solution {
public:
    bool hasCycle(ListNode *head) {
        ListNode* fast = head;
        ListNode* slow = head;
        while(fast != NULL && fast->next != NULL) {
            slow = slow->next;
            fast = fast->next->next;
            // 快慢指针相遇，说明有环
            if (slow == fast) return true;
        }
        return false;
    }
};
```
## 扩展 

做完这道题目，可以在做做142.环形链表II，不仅仅要找环，还要找环的入口。

142.环形链表II题解：[链表：环找到了，那入口呢？](https://mp.weixin.qq.com/s/_QVP3IkRZWx9zIpQRgajzA)

----------end----------

**我是程序员Carl，先后在腾讯和百度从事技术研发多年，利用工作之余重刷leetcode，更多[精彩算法文章](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzUxNjY5NTYxNA==&action=getalbum&album_id=1485825793120387074&scene=173#wechat_redirect)尽在：[代码随想录](https://img-blog.csdnimg.cn/20200815195519696.png)，关注后，回复「Java」「C++」「python」「简历模板」等等，有我整理多年的学习资料，可以加我[微信](https://img-blog.csdnimg.cn/20200814140330894.png)，备注「简单自我介绍」+「组队刷题」，拉你进入刷题群（无任何广告，纯个人分享），每天一道经典题目分析，我选的每一道题目都不是孤立的，而是由浅入深一脉相承的，如果跟住节奏每篇连续着看，定会融会贯通。本文  [https://github.com/youngyangyang04/leetcode-master](https://github.com/youngyangyang04/leetcode-master ) 已经收录，里面还有leetcode刷题攻略、各个类型经典题目刷题顺序、思维导图，可以fork到自己仓库，有空看一看一定会有所收获，如果对你有帮助也给一个star支持一下吧！**
**以下资料希望对你有帮助：**

* [学习资料以及我的开源项目](https://github.com/youngyangyang04)
* [如何学习C++？ B站](https://www.bilibili.com/video/BV1rK4y1e7ed)
* [手把手带你读C++ primer！B站](https://www.bilibili.com/video/BV1Z5411874t)
* [帮你把KMP算法学个通透！（理论篇）B站](https://www.bilibili.com/video/BV1PD4y1o7nd/)
* [帮你把KMP算法学个通透！（求next数组代码篇）B站](https://www.bilibili.com/video/BV1M5411j7Xx/)
* [leetcode刷题攻略](https://github.com/youngyangyang04/leetcode-master)
* [程序员应该如何写简历（附简历模板）](https://mp.weixin.qq.com/s/PkBpde0PV65dJjj9zZJYtg)
* [一线互联网公司技术面试的流程以及注意事项](https://mp.weixin.qq.com/s/1VMvQ_6HbVpEn85CNilTiw)
* [C++面试&C++学习指南知识点整理](https://github.com/youngyangyang04/TechCPP)

**双指针经典题目汇总：**

* [数组：就移除个元素很难么？](https://mp.weixin.qq.com/s/wj0T-Xs88_FHJFwayElQlA)
* [字符串：这道题目，使用库函数一行代码搞定](https://mp.weixin.qq.com/s/X02S61WCYiCEhaik6VUpFA)
* [字符串：替换空格](https://mp.weixin.qq.com/s/t0A9C44zgM-RysAQV3GZpg)
* [字符串：花式反转还不够！](https://mp.weixin.qq.com/s/X3qpi2v5RSp08mO-W5Vicw)
* [链表：听说过两天反转链表又写不出来了？](https://mp.weixin.qq.com/s/pnvVP-0ZM7epB8y3w_Njwg)
* [链表：环找到了，那入口呢？](https://mp.weixin.qq.com/s/_QVP3IkRZWx9zIpQRgajzA) 
* [哈希表：解决了两数之和，那么能解决三数之和么？](https://mp.weixin.qq.com/s/r5cgZFu0tv4grBAexdcd8A)
* [双指针法：一样的道理，能解决四数之和](https://mp.weixin.qq.com/s/nQrcco8AZJV1pAOVjeIU_g)
* [双指针法：总结篇！](https://mp.weixin.qq.com/s/_p7grwjISfMh0U65uOyCjA)

**如果感觉题解对你有帮助，不要吝啬给一个👍吧！**
