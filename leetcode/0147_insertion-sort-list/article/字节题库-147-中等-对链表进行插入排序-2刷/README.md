"+++
title = "0147. Insertion Sort List 字节题库 - #147 - 中等 - 对链表进行插入排序 - 2刷 "
author = ["superkakayong"]
date = 2020-08-05T02:16:31+08:00
tags = ["Leetcode", "Algorithms", "C++", "字节跳动", "LinkedList", "Sort"]
categories = ["leetcode"]
draft = false
+++

# 字节题库 - #147 - 中等 - 对链表进行插入排序 - 2刷

> [0147. Insertion Sort List](https://leetcode-cn.com/problems/insertion-sort-list/)
> [字节题库 - #147 - 中等 - 对链表进行插入排序 - 2刷](https://leetcode-cn.com/problems/insertion-sort-list/solution/zi-jie-ti-ku-147-zhong-deng-dui-lian-biao-jin-xing/) by [superkakayong](https://leetcode-cn.com/u/superkakayong/)

### 先放C++代码，详细思路及注释都已写好在代码中

```cpp
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* insertionSortList(ListNode* head) {
        if(!head || !head->next) return head;
        ListNode *dummy = new ListNode(10086); // 创建一个虚拟头结点，赋值10086～
        dummy -> next = head; // 让虚拟头结点指向head
        ListNode *tail = head, *sort = head -> next; // tail指向已排好序部分的尾结点；sort是下一个待排序的结点
        while(sort)
        {
            if(sort->val < tail->val)
            {
                ListNode *pos = dummy; // 定义一个pos指针，用于每次从头（dummy处）寻找该插入sort的结点的位置。比如1->3，sort是2，那pos就在1处。
                while(pos->next->val < sort->val) pos = pos ->next; // 将pos定位到已排好序的且比sort小的部分中的最大的那个结点。
                tail -> next = sort -> next; // 断链。比如dummy->4->2->1->3，tail是4，sort是2，则让4->1，把2拎出来。
                sort -> next = pos -> next; // pos此时在dummy处（因为4大于2），则让 2->4。
                pos -> next = sort; // 上一步相当于断掉了dummy->4，现在重新接起来：dummy->2->4->1->3。
                sort = tail -> next; // 让sort重新回到下一个待排序的结点处
            }
            else
            {
                tail = tail -> next;
                sort = sort -> next;
            }
        }
        return dummy -> next;
    }
};
```
### 执行结果截图
![image.png](https://pic.leetcode-cn.com/018a186ae8041d690b5860538ebc9f56fa3fe2b5b7a76835bad26771123b80ef-image.png)
