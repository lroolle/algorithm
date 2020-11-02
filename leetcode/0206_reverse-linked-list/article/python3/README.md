"+++
title = "0206. Reverse Linked List 图解流程 Python3迭代详解 "
author = ["han-han-a-gou"]
date = 2020-06-27T14:08:48+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 图解流程 Python3迭代详解

> [0206. Reverse Linked List](https://leetcode-cn.com/problems/reverse-linked-list/)
> [图解流程 Python3迭代详解](https://leetcode-cn.com/problems/reverse-linked-list/solution/tu-jie-liu-cheng-python3die-dai-xiang-jie-by-han-h/) by [han-han-a-gou](https://leetcode-cn.com/u/han-han-a-gou/)

### 解题思路：

要将链表 `1 -> 2 -> 3 -> 4 -> Null` 反转为  `4 -> 3 -> 2 -> 1 -> Null` ，需要一个 `cur` 指针表示当前遍历到的节点；一个 `pre` 指针表示当前节点的前驱节点；在循环中还需要一个中间变量 `temp` 来保存当前节点的后驱节点。

#### 算法流程：


- 首先 `pre` 指针指向 `Null`，`cur` 指针指向 `head`；

- 当 `cur != Null`，执行循环。
    - 先将 `cur.next` 保存在 `temp` 中防止链表丢失：`temp = cur.next`
    
    - 接着把 `cur.next` 指向前驱节点 `pre`：`cur.next = pre`
    - 然后将 `pre` 往后移一位也就是移到当前 `cur` 的位置：`pre = cur`
    - 最后把 `cur` 也往后移一位也就是 `temp` 的位置：`cur = temp`

- 当 `cur == Null`，结束循环，返回 `pre`。

#### 图解流程：
![image-20200627220535158.png](https://pic.leetcode-cn.com/039a6eff23111dba77d91ed909bbd35539b1185c07664372b129216a7b779b4a-image-20200627220535158.png)


### Python3 代码：
```
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        pre = None
        cur = head
        while cur:
            temp = cur.next   # 先把原来cur.next位置存起来
            cur.next = pre
            pre = cur
            cur = temp
        return pre
```
