"+++
title = "0206. Reverse Linked List 单向链表反转 "
author = ["kimee"]
date = 2020-09-09T15:40:44+08:00
tags = ["Leetcode", "Algorithms", "Swift"]
categories = ["leetcode"]
draft = false
+++

# 单向链表反转

> [0206. Reverse Linked List](https://leetcode-cn.com/problems/reverse-linked-list/)
> [单向链表反转](https://leetcode-cn.com/problems/reverse-linked-list/solution/dan-xiang-lian-biao-fan-zhuan-by-kimee/) by [kimee](https://leetcode-cn.com/u/kimee/)

### 解题思路
特殊处理
当头节点 == nil，或只有1个节点时，直接返回
正常流程
当前节点p，记录当前遍历节点
新的头节点newHead，记录拆链后新链表的头节点
while循环步骤：
1、临时变量tmp = p.next记录拆链前的下一个节点；
2、拆链，将当前p节点的next指向新的头节点（初始为nil）p.next = newHead；
3、更新新的头节点newHead为当前节点p，newHead = p；
4、更新当前遍历节点p为下一个将要遍历的节点，p = tmp；
最后返回新链表头节点 return newHead

### 代码

```swift
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public var val: Int
 *     public var next: ListNode?
 *     public init(_ val: Int) {
 *         self.val = val
 *         self.next = nil
 *     }
 * }
 */
class Solution {
    func reverseList(_ head: ListNode?) -> ListNode? {
        if head == nil || head?.next == nil {
            return head
        }
        var p = head
        var newHead: ListNode? = nil
        while p != nil {
            let tmp = p?.next
            p?.next = newHead
            newHead = p
            p = tmp
        }
        return newHead
    }
}
```