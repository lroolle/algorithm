"+++
title = "0141. Linked List Cycle C++简洁代码（快慢指针/set）： "
author = ["OrangeMan"]
date = 2020-09-04T17:02:18+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# C++简洁代码（快慢指针/set）：

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [C++简洁代码（快慢指针/set）：](https://leetcode-cn.com/problems/linked-list-cycle/solution/cjian-ji-dai-ma-duo-chong-fang-fa-by-orangeman-11/) by [OrangeMan](https://leetcode-cn.com/u/orangeman/)

### 解题思路
解题思路:
1. 方法一：**快慢指针**
- 快指针去追赶慢指针
- 若是快指针遇到`NULL`，则返回`false`
- 若是快指针等于慢指针（表示追上了，也就是被套圈了），则返回`true`
2. 方法二：**set**
- 遍历链表，若遇到NULL，则退出循环，返回`false`
- 若`set.count(cur) == 1`，表示已有此节点，则返回`true`
### 代码

```快慢指针
class Solution {
public:
    bool hasCycle(ListNode *head) {
        ListNode* slow = head;
        ListNode* fast = head;
        while(fast && fast->next) {
            fast = fast->next->next;
            slow = slow->next;
            if(fast == slow) return true;
        }
        return false;
    }
};
```
```set
class Solution {
public:
    bool hasCycle(ListNode *head) {
        unordered_set<ListNode*> set;
        ListNode* cur = head;
        while(cur) {
            if(set.count(cur) == 1) return true;
            set.insert(cur);
            cur = cur->next;
        }
        return false;
    }
};
```
