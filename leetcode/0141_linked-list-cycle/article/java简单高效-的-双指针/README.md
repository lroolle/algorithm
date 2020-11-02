"+++
title = "0141. Linked List Cycle 【Java】简单高效 的 双指针 "
author = ["leetcoder-youzg"]
date = 2020-10-09T00:24:24+08:00
tags = ["Leetcode", "Algorithms", "Java", "TwoPointers", "指针"]
categories = ["leetcode"]
draft = false
+++

# 【Java】简单高效 的 双指针

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [【Java】简单高效 的 双指针](https://leetcode-cn.com/problems/linked-list-cycle/solution/java-jian-dan-gao-xiao-de-shuang-zhi-zhen-by-leetc/) by [leetcoder-youzg](https://leetcode-cn.com/u/leetcoder-youzg/)

## 法1 —— 存在性判别法：
### 解题思路
> 用集合记录每一个遍历到的 节点
> 每次遍历时，先判断该节点是否存在
> 若 **当前节点已存在**，则表明 **存在环**

### 运行结果
![image.png](https://pic.leetcode-cn.com/1602202796-aJCZXb-image.png)

### 代码

```java
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public boolean hasCycle(ListNode head) {
        if (head == null) {
            return false;
        }
        ArrayList<ListNode> memoryList = new ArrayList<>();
        while (head.next != null) {
            if (memoryList.contains(head)) {
                return true;
            }
            memoryList.add(head);
            head = head.next;
        }
        return false;
    }
}
```

---
## 法2 —— 快慢指针(双指针)法：
### 解题思路
> 给定两个指针，快指针从head.next开始，慢指针从head开始
> 快指针每次移动两个单位，慢指针每次移动一个单位
> 若 **快指针和慢指针能相遇**，则表明 **存在环**

### 运行结果
![image.png](https://pic.leetcode-cn.com/1602202936-owkPqE-image.png)

### 代码

```java
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public boolean hasCycle(ListNode head) {
        if (head == null || head.next == null) {
            return false;
        }
        ListNode fastp = head.next;
        ListNode slowp = head;
        while (fastp != null && fastp.next != null) {
            if (fastp == slowp) {
                return true;
            }
            fastp = fastp.next.next;
            slowp = slowp.next;
        }
        return false;
    }
}
```
打卡80天，加油！！！
昨晚将搜索功能做出来了，MP是真的强