"+++
title = "0147. Insertion Sort List C++  98.69% 三种解法 简洁易懂 "
author = ["mrethan"]
date = 2020-07-02T08:37:33+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# C++  98.69% 三种解法 简洁易懂

> [0147. Insertion Sort List](https://leetcode-cn.com/problems/insertion-sort-list/)
> [C++  98.69% 三种解法 简洁易懂](https://leetcode-cn.com/problems/insertion-sort-list/solution/c-9869-san-chong-jie-fa-jian-ji-yi-dong-by-mrethan/) by [mrethan](https://leetcode-cn.com/u/mrethan/)

方法0：原地排序。
思路如下：
1.跳过已经排序的节点；
2.找到未排序的节点，将其从链表上摘下；
3.从头开始查找合适的插入点，未建立哨兵节点，需处理head更新的情况；

方法1：
方法0的改进版，建立哨兵节点dummy_head，便于从头开始查找合适的插入位置；

方法2：
建立新链表dummy_head，一个一个从原链表摘下节点，按顺序插入到新链表中。

效率：方法1>方法0>方法2
![image.png](https://pic.leetcode-cn.com/75b84f88534cb4b4275953cd2248226dd7f86d7ea55b9f959daf704732e34d23-image.png)
```
class Solution {
public:
    ListNode* insertionSortList0(ListNode* head) {
        if (!head || !head->next) {
            return head;
        }
        
        ListNode* work = NULL;
        ListNode* prev = head;
        ListNode* p = head->next;

        //p为未排序节点
        while (p) {
            //跳过本身已经有序节点
            if (p->val >= prev->val) {
                prev = p;
                p = p->next;
            }
            else {
                //p摘链
                prev->next = p->next;
                work = p;
                p = p->next;
                //从头开始找比p小的节点，插入到它的前面
                if (head->val > work->val) {
                    //插入到头节点之前
                    work->next = head;
                    head = work;
                }
                else {
                    //插入到头节点后
                    ListNode* temp = head;
                    while (temp->next->val <= work->val) {
                        temp = temp->next;
                    }
                    work->next = temp->next;
                    temp->next = work;
                }
            }
        }

        return head;
    }

    //使用哨兵节点，无需特殊处理头结点改变的情况
    ListNode* insertionSortList1(ListNode* head) {
        if (!head || !head->next) {
            return head;
        }
        ListNode* dummy_head = new ListNode(0);
        ListNode* cur = NULL, *prev = NULL, *ans = NULL;
        
        dummy_head->next = head;
        prev = head; //prev指向当前已排序的最大节点
        cur = head->next;

        while (cur) {
            if (cur->val >= prev->val) {
                //跳过已排序部分
                prev = cur;
                cur = cur->next;
                continue;
            }
            //cur节点摘链
            prev->next = cur->next;
            ListNode* p = dummy_head;
            //从头开始查找合适位置插入
            while (p->next->val <= cur->val) {
                p = p->next;
            }
            //找到比cur大的节点为p->next，把cur插入到p节点后面
            cur->next = p->next;
            p->next = cur;
            //继续遍历
            cur = prev->next;
        }
        ans = dummy_head->next;

        delete dummy_head;

        return ans;
    }

    //新建链表，从原链表取节点插入新链表
    ListNode* insertionSortList2(ListNode* head) {
        if (!head || !head->next) {
            return head;
        }

        ListNode* oldNode = NULL;
        ListNode* dummy_head = new ListNode(0);
        ListNode* ans = NULL;

        /* 摘下第一个节点插入新链表 */
        dummy_head->next = head;
        head = head->next;
        dummy_head->next->next = NULL;

        while (head) {
            //从原链表取节点
            oldNode = head;
            head = head->next;
            //断开摘下的节点与原链表的连接
            oldNode->next = NULL;

            //把节点按顺序插入新链表
            ListNode* p = dummy_head->next;
            ListNode* prev = dummy_head;
            while (p != NULL && p->val <= oldNode->val) {
                prev = p;
                p = p->next;
            }
            oldNode->next = prev->next;
            prev->next = oldNode;
        }
        ans = dummy_head->next;

        delete dummy_head;

        return ans;
    }
};
```
