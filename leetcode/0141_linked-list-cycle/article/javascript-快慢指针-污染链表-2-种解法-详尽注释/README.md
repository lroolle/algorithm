"+++
title = "0141. Linked List Cycle JavaScript 快慢指针 + 污染链表 2 种解法 详尽注释 "
author = ["jsliang"]
date = 2020-10-09T02:31:13+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "指针", "快慢指针"]
categories = ["leetcode"]
draft = false
+++

# JavaScript 快慢指针 + 污染链表 2 种解法 详尽注释

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [JavaScript 快慢指针 + 污染链表 2 种解法 详尽注释](https://leetcode-cn.com/problems/linked-list-cycle/solution/javascript-kuai-man-zhi-zhen-wu-ran-lian-biao-2-ch/) by [jsliang](https://leetcode-cn.com/u/jsliang/)

JavaScript 做一个环形链表：

```js
// 设置空对象
let tempLink = {};

// 设置循环链表
let cycleLink = {
  val: 2,
  next: {
    val: 0,
    next: {
      val: 4,
      next: tempLink, // 最后指向空对象
    },
  },
};

// 空对象指向循环链表，使链表变成真的环
tempLink.next = cycleLink;

// 设置其他节点，接着循环链表
const link = {
  val: 3,
  next: cycleLink,
};
// 3 -> 2 -> 0 -> 4 -> 2 -> 0 -> 4 -> cycle

console.log(hasCycle(link));
```

* 快慢指针

慢指针一次走一步，快指针一次走两步，快指针追上慢指针，说明有环

举例：400m 操场，小明一次跑 2 米，小梁一次跑 1 米，小梁跑一圈（400m）后小明（800m）追上小梁

```js
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @return {boolean}
 */
const hasCycle = (head) => {

  // 至少 2 个节点才能构成一个环
  if (!head || !head.next) {
    return false;
  }

  // 设置快慢指针
  let slow = head;
  let fast = head.next;

  // 如果快指针一直没有追上慢指针
  while (slow !== fast) {
    // 如果没有环，则快指针会抵达终点
    if (!fast || !fast.next) {
      return false;
    }
    slow = slow.next;
    fast = fast.next.next;
  }

  // 如果有环，那么快指针会追上慢指针
  return true;
};
```

* 人见人爱解法

```js
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @return {boolean}
 */
const hasCycle = (head) => {
  // 1. 如果有链表
  while (head) {
    // 2. 每经过一个节点，将它渲染成 jsliang
    if (head.val === 'jsliang') {
      // 3. 如果下次找到了自己，证明有环
      return true;
    } else {
      head.val = 'jsliang';
    }
    // 4. 一直往链表尾部走
    head = head.next;
  }

  // 5. 如果没有重复，那么返回 false
  return false;
};
```

每次路过节点，都将它设置为 **jsliang**，这样如果重复了，就知道了，毕竟链表节点不可能设置为 **jsliang**。

就是造成的污染大了点，数据都被污染了。
