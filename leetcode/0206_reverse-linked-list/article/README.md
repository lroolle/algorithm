"+++
title = "0206. Reverse Linked List 逐步图解递归 迭代 "
author = ["sucongcjs"]
date = 2020-08-03T06:14:32+08:00
tags = ["Leetcode", "Algorithms", "Recursion", "迭代", "C++"]
categories = ["leetcode"]
draft = false
+++

# 逐步图解递归 迭代

> [0206. Reverse Linked List](https://leetcode-cn.com/problems/reverse-linked-list/)
> [逐步图解递归 迭代](https://leetcode-cn.com/problems/reverse-linked-list/solution/zhu-bu-tu-jie-di-gui-die-dai-by-sucongcjs/) by [sucongcjs](https://leetcode-cn.com/u/sucongcjs/)

# 整条链表反转
## 非递归解法
非递归只需要将链表节点一个个添加到一条新的链表
- `cur` 指向待反转的链表的头一个节点
- `prev` 指向新链表的最后一个节点


![image.png](https://pic.leetcode-cn.com/7a160b91f4c60a0d7eca56cdd4f8d930b83f3e043c2a333537a53c778bfb1367-image.png)
---


![image.png](https://pic.leetcode-cn.com/629a7e857b15925fcb467a6dc843f5d9d50015866c28222e6ca17bf245b79b94-image.png)
---

![image.png](https://pic.leetcode-cn.com/9b42e83869a4baf11f1bf7cb8c07f6d1b07ef4adf528493321266f927dae40b9-image.png)


如此操作直到 `cur==NULL`

代码如下
```cpp
ListNode* reverseList(ListNode* head) {
    ListNode* prev = NULL;  // 新的链表的最后
    ListNode* cur = head;  // 当前待添加到新链表的节点
    while(cur != NULL){
        ListNode* tmp = cur->next;  // 保存cur->next
        cur->next = prev;
        prev = cur;
        cur = tmp;
    }
    return prev;
}
```

## 递归解法
递归的思路就是先递归到底, 找到最后一个节点, 然后从最后一个节点开始, 把箭头方向掉转

![image.png](https://pic.leetcode-cn.com/f776160e91490609ec4d589207ebddbf02f210f17b83c875b89d2e4269c64b06-image.png)


递归结束的条件
```cpp
if(head == NULL || head->next == NULL) 
    return head;
```
得到最后一个节点之后, 返回给上一层递归, `p`指向原链表的最后一个节点, 现在要作为头节点, 之后的`p`都不需要改动, 不断返回给上一层递归

![image.png](https://pic.leetcode-cn.com/48f2b1894560d7b564e3e65382c13a156be070ac1f4a04d4d75c84f5f7930cf2-image.png)


改变`head`的`next`的指向, 让它指向自己
```cpp
head->next->next = head;
```
再把`head`的`next`指向空, 作为新链表的最后一个节点
```cpp
head->next = NULL;
```

![image.png](https://pic.leetcode-cn.com/a2360c321eb4717ce0139038476faea6ebb2602efde0935825a150e7f7efa810-image.png)

递归结束, 最后返回的还是`p`(头节点)

代码如下
```cpp
ListNode* reverseList(ListNode* head) {
    if(head == NULL || head->next == NULL) 
        return head;

    ListNode* p = reverseList(head->next);
    head->next->next = head;
    head->next = NULL;
    return p;
}
```

[链表的各种反转, 欢迎来访!](http://www.sucong.top/archives/%E5%90%83%E9%80%8F%E9%93%BE%E8%A1%A8%E5%8F%8D%E8%BD%AC)