"+++
title = "0141. Linked List Cycle Javascript 【环形链表】的三种解法 "
author = ["piao-z"]
date = 2020-07-21T12:01:25+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# Javascript 【环形链表】的三种解法

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [Javascript 【环形链表】的三种解法](https://leetcode-cn.com/problems/linked-list-cycle/solution/javascript-huan-xing-lian-biao-de-san-chong-jie-fa/) by [piao-z](https://leetcode-cn.com/u/piao-z/)

### 节点标记法

最先想到的基础方法，不需要任何前置知识：

我们只需要遍历这个链表，在遍历过的时候给节点打一个tag，这样，如果遍历到循环结束，那该链表自然不是环形链表。

反之，如果我们遍历的时候遇到了tag，则说明我们之前也遇到过它，所以我们已经进入了环中！故`return true`

```javascript
var hasCycle = function(head) {
    // 链表长度小于2时直接return
    if(!head || !head.next) return false
    while(head) {
        if(head.tag) return true
        head.tag = true
        head = head.next
    }
    return false
};
```

- 执行用时：96 ms, 在所有 JavaScript 提交中击败了18.67%的用户
- 内存消耗：38.3 MB, 在所有 JavaScript 提交中击败了33.33%的用户

### hashMap法

在遍历链表时存储遍历过的节点，如果Map中该节点存在，则说明链表中有环

这种解法和节点标记法异曲同工。

```javascript
var hasCycle = (head) => {
    if(!head || !head.next) return false
    let map = new Map()
    while (head) {
        if (map.has(head)) return true
        map.set(head, true)
        head = head.next
    }
    return false
}
```

- 执行用时：80 ms, 在所有 JavaScript 提交中击败了71.12%的用户
- 内存消耗：38.6 MB, 在所有 JavaScript 提交中击败了8.33%的用户

### 快慢指针

定义两个指针一快一慢，如果链表有环，则快慢两指针早晚会相遇，这样就能判断出链表中有环存在。

![](https://pic.leetcode-cn.com/83ff2a14a40255d326f6165892f1958f674b405a98e0bb5f705cc341bb5d0d5a-file_1595332885128)

```javascript
var hasCycle = function(head) {
    if(!head || !head.next) return false
    let slow = head
    let fast = head.next
    while(slow != fast){
        if(!fast || !fast.next) return false
        fast = fast.next.next
        slow = slow.next
    }
    return true
};
```

- 执行用时：76 ms, 在所有 JavaScript 提交中击败了85.77%的用户
- 内存消耗：38.3 MB, 在所有 JavaScript 提交中击败了33.33%的用户

首先定义两个变量（指针）开始循环，快指针的速度是慢指针的两倍，所以当他们相遇时，则链表中存在环，或者快指针走到终点，说明链表无环。

### 分门别类系统学习算法

欢迎各位关注我的公众号：[一个歪卜](https://cdn.byeguo.cn/gzh/logo.jpeg)

在这里，你可以跟着我每周做一道算法题，我会按照分类、系统的让大家认识各种前端算法的不同题型，做到面试时可以信手拈来～

同时，我也会不定时更新前端相关的常见面试题型解法，项目实战等内容，和大家一起学习进步😊