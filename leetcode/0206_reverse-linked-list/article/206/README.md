"+++
title = "0206. Reverse Linked List 206. 反转链表:【双指针法】【递归法】详解！ "
author = ["carlsun-2"]
date = 2020-08-15T10:25:33+08:00
tags = ["Leetcode", "Algorithms", "cpp", "LinkedList"]
categories = ["leetcode"]
draft = false
+++

# 206. 反转链表:【双指针法】【递归法】详解！

> [0206. Reverse Linked List](https://leetcode-cn.com/problems/reverse-linked-list/)
> [206. 反转链表:【双指针法】【递归法】详解！](https://leetcode-cn.com/problems/reverse-linked-list/solution/206-fan-zhuan-lian-biao-shuang-zhi-zhen-fa-di-gui-/) by [carlsun-2](https://leetcode-cn.com/u/carlsun-2/)

> 反转链表的写法很简单，一些同学甚至可以背下来但过一阵就忘了该咋写，主要是因为没有理解真正的翻转过程。

## 思路 

如果再定义一个新的链表，实现链表元素的反转，其实这是对内存空间的浪费。

其实只需要改变链表的next指针的指向，直接将链表反转 ，而不用重新定义一个新的链表，如图所示:

![206_反转链表.png](https://pic.leetcode-cn.com/034c40988d85ede256466be34510e0a0ae1b634d3e9b551259489dcbd769a5f7-206_%E5%8F%8D%E8%BD%AC%E9%93%BE%E8%A1%A8.png)
之前链表的头节点是元素1， 反转之后头结点就是元素5 ，这里并没有添加或者删除节点，仅仅是改表next指针的方向。

那么接下来看一看是如何反转呢？

我们拿有示例中的链表来举例，如动画所示：

![206.翻转链表.mp4](35da887c-65ec-4544-9229-49a2ab42b831)
首先定义一个cur指针，指向头结点，再定义一个pre指针，初始化为null。

然后就要开始反转了，首先要把 cur->next 节点用tmp指针保存一下，也就是保存一下这个节点。

为什么要保存一下这个节点呢，因为接下来要改变 cur->next 的指向了，将cur->next 指向pre ，此时已经反转了第一个节点了。

接下来，就是循环走如下代码逻辑了，继续移动pre和cur指针。

最后，cur 指针已经指向了null，循环结束，链表也反转完毕了。 此时我们return pre指针就可以了，pre指针就指向了新的头结点。

## C++代码

### 双指针法
```
class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        ListNode* temp; // 保存cur的下一个节点
        ListNode* cur = head;
        ListNode* pre = NULL;
        while(cur) {
            temp = cur->next;  // 保存一下 cur的下一个节点，因为接下来要改变cur->next
            cur->next = pre; // 翻转操作
            // 更新pre 和 cur指针
            pre = cur;
            cur = temp;
        }
        return pre;
    }
};
```

### 递归法

递归法相对抽象一些，但是其实和双指针法是一样的逻辑，同样是当cur为空的时候循环结束，不断将cur指向pre的过程。 

关键是初始话的地方，可能有的同学会不理解， 可以看到双指针法中初始化 cur = head，pre = NULL，在递归法中可以从如下代码看出初始化的逻辑也是一样的，只不过写法变了。

具体可以看代码（已经详细注释），双指针法写出来之后，理解如下递归写法就不难了，代码逻辑都是一样的。
```
class Solution {
public:
    ListNode* reverse(ListNode* pre,ListNode* cur){
        if(cur == NULL) return pre;
        ListNode* temp = cur->next;
        cur->next = pre;
        // 可以和双指针法的代码进行对比，如下递归的写法，其实就是做了这两步
        // pre = cur;
        // cur = temp;
        return reverse(cur,temp);
    }
    ListNode* reverseList(ListNode* head) {
        // 和双指针法初始化是一样的逻辑
        // ListNode* cur = head;
        // ListNode* pre = NULL;
        return reverse(NULL, head);
    }

};
```

**我是程序员Carl，先后在腾讯和百度从事技术研发多年，利用工作之余重刷leetcode，更多[精彩算法文章](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzUxNjY5NTYxNA==&action=getalbum&album_id=1485825793120387074&scene=173#wechat_redirect)尽在公众号：[代码随想录](https://img-blog.csdnimg.cn/20200815195519696.png)，关注后，回复「Java」「C++」「python」「简历模板」等等，有我整理多年的学习资料，可以加我[微信](https://img-blog.csdnimg.cn/20200814140330894.png)，备注「简单自我介绍」+「组队刷题」，拉你进入刷题群（无任何广告，纯个人分享），每天一道经典题目分析，我选的每一道题目都不是孤立的，而是由浅入深一脉相承的，如果跟住节奏每篇连续着看，定会融会贯通。本文  [https://github.com/youngyangyang04/leetcode-master](https://github.com/youngyangyang04/leetcode-master ) 已经收录，里面还有leetcode刷题攻略、各个类型经典题目刷题顺序、思维导图，看一看一定会有所收获，如果对你有帮助也给一个star支持一下吧！**
**以下资料希望对你有帮助：**

* [我在B站上讲KMP算法！](https://www.bilibili.com/video/BV1PD4y1o7nd/)
* [leetcode刷题攻略](https://github.com/youngyangyang04/leetcode-master)
* [程序员应该如何写简历（附简历模板）](https://mp.weixin.qq.com/s/PkBpde0PV65dJjj9zZJYtg)
* [一线互联网公司技术面试的流程以及注意事项](https://mp.weixin.qq.com/s/1VMvQ_6HbVpEn85CNilTiw)
* [如何学习C++？ B站](https://www.bilibili.com/video/BV1rK4y1e7ed)
* [手把手带你读C++ primer！B站](https://www.bilibili.com/video/BV1Z5411874t)
* [C++面试&C++学习指南知识点整理](https://github.com/youngyangyang04/TechCPP)
**链表相关题目精选：** 
* [关于链表，你该了解这些！](https://mp.weixin.qq.com/s/ntlZbEdKgnFQKZkSUAOSpQ)
* [链表：听说用虚拟头节点会方便很多？](https://mp.weixin.qq.com/s/slM1CH5Ew9XzK93YOQYSjA)
* [链表：一道题目考察了常见的五个操作！](https://mp.weixin.qq.com/s/Cf95Lc6brKL4g2j8YyF3Mg)
* [链表：听说过两天反转链表又写不出来了？](https://mp.weixin.qq.com/s/pnvVP-0ZM7epB8y3w_Njwg)
* [链表：环找到了，那入口呢？](https://mp.weixin.qq.com/s/_QVP3IkRZWx9zIpQRgajzA)
**如果感觉题解对你有帮助，不要吝啬给一个👍吧！**