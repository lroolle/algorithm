"+++
title = "0025. Reverse Nodes in k-Group Java简单递归，清晰明了 "
author = ["winnndy"]
date = 2020-08-21T07:40:36+08:00
tags = ["Leetcode", "Algorithms", "Java"]
categories = ["leetcode"]
draft = false
+++

# Java简单递归，清晰明了

> [0025. Reverse Nodes in k-Group](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)
> [Java简单递归，清晰明了](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/javajian-dan-di-gui-qing-xi-ming-liao-by-winnndy/) by [winnndy](https://leetcode-cn.com/u/winnndy/)

### 解题思路
与反转链表类似，不同在于反转链表时head.next = null, 而在这里head.next指向的是下一组k个节点反转后的头节点。

### 代码

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode next = head;
        int i = 0;
        //找到下一组k个节点现在的头节点
        for(i = 0; i < k; i++){
            if(next == null)
                break;
            next = next.next;
        }
        //节点数不足k个直接返回head
        if(i < k)
            return head;
        ListNode cur = head.next;
        ListNode pre = head;
        ListNode temp;
        //递归，将head.next指向下一组k个节点反转后的头节点
        head.next = reverseKGroup(next, k);
        while(cur != next){
            temp = cur.next;
            cur.next = pre;
            pre = cur;
            cur = temp;
        }
        //返回反转后的头节点
        return pre;
    }
}
```