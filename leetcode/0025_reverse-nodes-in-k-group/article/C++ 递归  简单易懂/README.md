"+++
title = "0025. Reverse Nodes in k-Group C++ 递归  简单易懂 "
author = ["yu-fan-u"]
date = 2020-06-06T11:41:31+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# C++ 递归  简单易懂

> [0025. Reverse Nodes in k-Group](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)
> [C++ 递归  简单易懂](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/c-di-gui-jian-dan-yi-dong-by-yu-fan-u/) by [yu-fan-u](https://leetcode-cn.com/u/yu-fan-u/)
```
class Solution {
public:
    ListNode* reverseKGroup(ListNode* head, int k) {
        if (head == NULL) return NULL;
        ListNode *a = head;
        ListNode *b = head;
        for (int i = 0; i < k; i++) {
            if (b == NULL) return head;
            b = b->next;
        }

        ListNode *newNode = reverseOperator(a,b);
        a->next = reverseKGroup(b,k);
        return newNode;
    }

    ListNode* reverseOperator(ListNode* n,ListNode *b) {
        ListNode *pre, *cur, *nxt;
        pre = NULL; cur = n; nxt = n;
        while (cur != b) {
            nxt = cur->next;
            cur->next = pre;
            pre = cur;
            cur = nxt;
        }
        return pre;
    }
};
```
