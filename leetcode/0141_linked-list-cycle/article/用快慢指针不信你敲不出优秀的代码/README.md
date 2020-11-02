"+++
title = "0141. Linked List Cycle 用快慢指针不信你敲不出优秀的代码 "
author = ["xmblgt"]
date = 2020-09-20T09:42:10+08:00
tags = ["Leetcode", "Algorithms", "快慢指针", "TwoPointers"]
categories = ["leetcode"]
draft = false
+++

# 用快慢指针不信你敲不出优秀的代码

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [用快慢指针不信你敲不出优秀的代码](https://leetcode-cn.com/problems/linked-list-cycle/solution/yong-kuai-man-zhi-zhen-bu-xin-ni-qiao-bu-chu-you-x/) by [xmblgt](https://leetcode-cn.com/u/xmblgt/)

**分析：**
快慢指针一前一后的遍历，只要存在环，迟早会相遇；
**思路：**
1. 特判： `为空`或者`仅有一个节点`的链表肯定没有环，排除！；
2. 初始化：
    - 慢指针`slow`指向第一个节点；
    - 快指针`fast`指向`slow`后面的节点；
3. 只要`slow`和`fast`不相遇，就进行循环遍历，**每次进行如下操作**：
    - 判断`fast`是否指向空或者最后一个节点，为什么最后一个节点也要，因为`fast`每次走两步，如果当前为最后一个节点的话，那么就办法前行了。
    - `slow`走一步，`fast`走两步（其实只要两者步伐不一致就可以了）。

```java
public class Solution {
    public boolean hasCycle(ListNode head) {
        //特判
        if(head == null || head.next == null)  return false;
        //初始化
        ListNode slow = head;
        ListNode fast = head.next;
        //循环
        while(slow != fast){
            if(fast == null || fast.next == null){
                return false;
            }
            //更新
            slow = slow.next;
            fast = fast.next.next;
        }
        return true; //能跳出循环走到这一步就说明，存在slow==fast，也就是存在环形了
    }
}
```
![微信图片_20200920111058.jpg](https://pic.leetcode-cn.com/1600594923-dnVqWE-%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20200920111058.jpg)

