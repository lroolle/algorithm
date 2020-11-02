"+++
title = "0141. Linked List Cycle 图解：环形链表 "
author = ["cheng-xu-yuan-xiao-hao"]
date = 2020-07-23T14:12:55+08:00
tags = ["Leetcode", "Algorithms", "C", "C++", "C#", "Go", "Java", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 图解：环形链表

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [图解：环形链表](https://leetcode-cn.com/problems/linked-list-cycle/solution/tu-jie-huan-xing-lian-biao-by-cheng-xu-yuan-xiao-h/) by [cheng-xu-yuan-xiao-hao](https://leetcode-cn.com/u/cheng-xu-yuan-xiao-hao/)

![图片.png](https://pic.leetcode-cn.com/da42c76fe60de354aee53a22493a86bc54d1d6a495719988484b0b393f37cac0-%E5%9B%BE%E7%89%87.png)
- - 本题题解如下（具体代码见文末）

![image.png](https://pic.leetcode-cn.com/a6a7e241b7eeb369dc931f239e155a29d89c702594f8143bd74332987d940ae6-image.png)


```go
func hasCycle(head *ListNode) bool {  
    if head == nil {
        return false
    }
    fast := head.Next       // 快指针，每次走两步
    for fast != nil && head != nil && fast.Next != nil {
        if fast == head {   // 快慢指针相遇，表示有环
            return true
        }
        fast = fast.Next.Next  
        head = head.Next        // 慢指针，每次走一步
    }
    return false
}
```
```javascript
var hasCycle = function(head) {
    try{
        JSON.stringify(head)
    }catch(e){
        return true
    }
    return false
};
```

我把我写的题解编排成了一本 30w 字的 pdf，包括 140 道高频面试算法题分析，每道题目都配有图解。[直接获取](https://www.geekxh.com/0.0.%E5%AD%A6%E4%B9%A0%E9%A1%BB%E7%9F%A5/03.html?MdXE)

![image.png](https://pic.leetcode-cn.com/df2702dd54d89a89a9b06f6c887300aa87d4b8ac83fe27d9c00395ed64143731-image.png)