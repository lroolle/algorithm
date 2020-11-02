"+++
title = "0025. Reverse Nodes in k-Group （图示）直白思路，使用头插法 "
author = ["zumin"]
date = 2020-09-11T07:26:19+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# （图示）直白思路，使用头插法

> [0025. Reverse Nodes in k-Group](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)
> [（图示）直白思路，使用头插法](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/tu-shi-zhi-bai-si-lu-shi-yong-tou-cha-fa-by-zumin/) by [zumin](https://leetcode-cn.com/u/zumin/)

## 思路
- 对每组链表进行单独翻转，然后建立每组的连接关系
- 首先，需要计算链表的节点个数，并找出翻转后链表的起点和剩余不翻转的第一个节点
- 然后对每组进行翻转。在翻转前需要保存当前翻转组的头节点和上一组的头节点。翻转后让上一翻转组的头节点连接到当前翻转组的最后一个节点
- 最后将翻转后链表的最后一个节点与剩余不翻转的第一个节点相连，然后返回之前记录的链表的起点
- 图示（链表为`1->2->3->4->5`，k=2）
![1.jpg](https://pic.leetcode-cn.com/1599809146-jOuEIC-1.jpg)


## 代码
```
ListNode *reverseKGroup(ListNode *head, int k) {
  if (head == NULL || head->next == NULL || k == 1) {
    return head;
  }

  //计算链表的节点个数，并找出翻转后链表的起点和剩余不翻转的第一个节点
  ListNode *p = head;
  ListNode *start, *end;
  int count = 0;
  while (p) {
    ++count;
    //第count个的节点即为翻转后链表的起点
    if (count == k) {
      start = p;
    }
    p = p->next;
    //记录翻转组的下一个节点，最后就为剩余不翻转的第一个节点
    if (count % k == 0) {
      end = p;
    }
  }

  //计算翻转组的个数
  count /= k;
  //nowHead为当前翻转组的头节点，preHead为上一翻转组的头节点
  ListNode *now = head, *n = head, *h, *nowHead = NULL, *preHead;
  while (count--) {
    preHead = nowHead;
    nowHead = now;
    //翻转
    for (int i = 0; i < k; ++i) {
      n = n->next;
      now->next = h;
      h = now;
      now = n;
    }
    //将上一翻转组的头节点连接到当前翻转组的最后一个节点
    if (preHead != NULL) {
      preHead->next = h;
    }
  }
  //将翻转后链表的最后一个节点与剩余不翻转的第一个节点相连
  nowHead->next = end;

  return start;
}
```
