"+++
title = "0025. Reverse Nodes in k-Group 从递归到迭代：由浅入深…… "
author = ["fan-xing-man-tian"]
date = 2020-08-22T08:15:08+08:00
tags = ["Leetcode", "Algorithms", "Python3"]
categories = ["leetcode"]
draft = false
+++

# 从递归到迭代：由浅入深……

> [0025. Reverse Nodes in k-Group](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)
> [从递归到迭代：由浅入深……](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/cong-di-gui-dao-die-dai-huan-mei-ru-men-bian-yi-fa/) by [fan-xing-man-tian](https://leetcode-cn.com/u/fan-xing-man-tian/)

##### 前提

以下所有 Python 代码的解法，都基于你做过 [206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/) 并熟知下面的表述方式：

```py
def reverseList(self, head: ListNode) -> ListNode:
    pre, cur = None, head
    while cur:
        cur.next, pre, cur = pre, cur, cur.next
    return pre
```
若看不懂 Python 的多重赋值，请移步我写的博文 [链表反转问题 1：整体反转](https://blog.csdn.net/weixin_43932942/article/details/107921785)，那里有详细的说明。

##### 递归

若你做过 [24. 两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/) 并了解递归做法，可能会更容易理解后面的代码：
```py
def swapPairs(self, head: ListNode) -> ListNode:
    if not head or not head.next:
        return head
    a, b = head, head.next
    b.next, a.next = a, self.swapPairs(b.next)
    return b
```


更准确的叫法是：递归 + 迭代。

```py
def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
    cur = head
    for _ in range(k):
        if not cur:
            return head
        cur = cur.next
 
    pre, cur = None, head
    for _ in range(k):
        cur.next, pre, cur = pre, cur, cur.next

    head.next = self.reverseKGroup(cur, k)
    return pre
```
前半部分判断剩余元素是否达到 $k$ 个：若没达到，就直接返回结果；若达到，继续下面的运算。
下面部分是 k 个元素的反转，用到了迭代。
最后将反转好的 k 个元素与剩余递归结束后的元素进行拼接，返回结果。

##### 迭代

由于本题要求：
> 你的算法只能使用常数的额外空间。

所以必须考虑迭代解法。

递归本质是调用内置栈完成的，可以考虑使用栈完成，然而我不会这种~~~

我的设想是：反转 k 个元素，然后一直执行这个操作。反转前要考虑剩余元素的个数，看是否需要反转。

代码细节已经注明：

```py
def reverseKGroup(self, head: ListNode, k: int) -> ListNode:

    # 反转 k 个元素，返回反转后的头、尾节点。
    def revrese(head, k):
        pre, cur = None, head
        for _ in range(k):
            cur.next, pre, cur = pre, cur, cur.next
        return pre, head

    pre = dummy = ListNode(0)
    dummy.next = head

    while pre.next:
        
        # 判断剩余节点个数是否有 k 个，若没有直接返回。
        # 另一个目的是将 cur 定位到 k+1 个节点的位置，方便后续的拼接。    
        cur = pre.next
        for _ in range(k):
            if not cur:
                return dummy.next
            cur = cur.next

        pre.next, pre = revrese(pre.next, k)  # 将 pre 连接已经反转好的元素，顺便移动 pre 的位置。
        pre.next = cur  # 拼接剩余元素，否则无法继续迭代。

    return dummy.next
```


