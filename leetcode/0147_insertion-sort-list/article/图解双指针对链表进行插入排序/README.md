"+++
title = "0147. Insertion Sort List 图解：双指针对链表进行插入排序 "
author = ["lulu27753"]
date = 2020-08-08T14:49:38+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "TwoPointers"]
categories = ["leetcode"]
draft = false
+++

# 图解：双指针对链表进行插入排序

> [0147. Insertion Sort List](https://leetcode-cn.com/problems/insertion-sort-list/)
> [图解：双指针对链表进行插入排序](https://leetcode-cn.com/problems/insertion-sort-list/solution/tu-jie-shuang-zhi-zhen-dui-lian-biao-jin-xing-cha-/) by [lulu27753](https://leetcode-cn.com/u/lulu27753/)

### 解题思路
• dummy: 哨兵节点
• head0: 当前已排序列表的最后一个
• pre: 用于遍历当前已排序列表

head0.next 一直通过 pre.next 找到其可以插入的位置插入。

在做指针移动的时候注意不要乱

![1.png](https://pic.leetcode-cn.com/d58cb63822e1d1d3ed44f5f153945ae7ef6d8130869fa2aa308ffb83ba1839db-1.png)
![2.png](https://pic.leetcode-cn.com/2ea903b27e27cb36538f3b44da86aca8352fc8aefba05cf237b76ed5348f09f8-2.png)
### 代码

```javascript
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var insertionSortList = function(head) {
  const dummy = new ListNode(0)
  dummy.next = head
  let head0 = dummy.next

  while (head0 && head0.next) {
    if (head0.next.val >= head0.val) {
      head0 = head0.next
      continue
    }

    let pre = dummy
    while (pre.next.val < head0.next.val) { pre = pre.next }

    let next = head0.next
    head0.next = next.next
    next.next = pre.next
    pre.next = next
  }

  return dummy.next
}
```
欢迎关注“前端之露”加入内推群