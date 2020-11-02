"+++
title = "0147. Insertion Sort List C++思路清晰实现 "
author = ["swaiwi"]
date = 2020-09-09T07:42:25+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# C++思路清晰实现

> [0147. Insertion Sort List](https://leetcode-cn.com/problems/insertion-sort-list/)
> [C++思路清晰实现](https://leetcode-cn.com/problems/insertion-sort-list/solution/csi-lu-qing-xi-shi-xian-by-swaiwi/) by [swaiwi](https://leetcode-cn.com/u/swaiwi/)

算法：
1.判断链表是否为空，为空直接返回
2.新建排序链表头和尾都指向head: sorted_head=head;sorted_tail=head;
3.从链表第二位开始插入排序(curr=head->next)
    先保存curr-next,next=curr->next
    (1)判断curr->val是否小于排序链表头部值sorted->head->val;
        如果小于，更新排序链表头部curr->next = sorted->head; sorted->head = curr;
    (2)否则判断curr->val是否大于等于排序链表尾部值sorted->tail->val;
        如果大于等于，更新排序链表尾部sorted->tail->next = curr; sorted->tail = curr;
    (3)否则，落入排序链表中间：
        找到排序链表位置p，使得p->val<=curr->val<p->next->val;将curr插入到p和p->next中间
        curr->next = p->next; p->next = curr;
    更新curr，curr=next
4.链表问题记得断尾求生
    sorted_tail->next = nullptr;
5.返回排序链表
    return sorted_head;

```
class Solution {
public:
    ListNode* insertionSortList(ListNode* head) {
        if(!head) return head;
        ListNode* sorted_head = head;
        ListNode* sorted_tail = head;
        ListNode* curr = head->next;
        while(curr){
            // cout<<curr->val<<endl;
            ListNode* next = curr->next;
            if(curr->val<sorted_head->val){
                curr->next = sorted_head;
                sorted_head = curr;
            } else if(curr->val>=sorted_tail->val){
                sorted_tail->next = curr;
                sorted_tail = curr;
            } else {
                ListNode* p = sorted_head;
                while(p!=sorted_tail){
                    if(curr->val>=p->val && curr->val<p->next->val){
                        curr->next = p->next;
                        p->next = curr;
                        break;
                    }
                    p = p->next;
                }
            }
            curr = next;
        }
        sorted_tail->next = nullptr;
        return sorted_head;
    }
};

```
