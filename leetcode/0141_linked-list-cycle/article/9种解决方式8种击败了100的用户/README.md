"+++
title = "0141. Linked List Cycle 9种解决方式，8种击败了100%的用户 "
author = ["sdwwld"]
date = 2020-10-09T04:29:37+08:00
tags = ["Leetcode", "Algorithms", "Java", "LinkedList"]
categories = ["leetcode"]
draft = false
+++

# 9种解决方式，8种击败了100%的用户

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [9种解决方式，8种击败了100%的用户](https://leetcode-cn.com/problems/linked-list-cycle/solution/9chong-jie-jue-fang-shi-8chong-ji-bai-liao-100de-y/) by [sdwwld](https://leetcode-cn.com/u/sdwwld/)


### 1，快慢指针解决

判断链表是否有环应该是老生常谈的一个话题了，最简单的一种方式就是快慢指针，**慢指针针每次走一步，快指针每次走两步**，如果相遇就说明有环，如果有一个为空说明没有环。代码比较简单

```
public boolean hasCycle(ListNode head) {
    if (head == null)
        return false;
    //快慢两个指针
    ListNode slow = head;
    ListNode fast = head;
    while (fast != null && fast.next != null) {
        //慢指针每次走一步
        slow = slow.next;
        //快指针每次走两步
        fast = fast.next.next;
        //如果相遇，说明有环，直接返回true
        if (slow == fast)
            return true;
    }
    //否则就是没环
    return false;
}
```


到这里问题好像并没有结束，为什么快慢指针就一定能判断是否有环。我们可以这样来思考一下，**假如有环，那么快慢指针最终都会走到环上**，假如环的长度是m，快慢指针最近的间距是n，如下图中所示

![image.png](https://pic.leetcode-cn.com/1602042204-ekumfo-image.png)
快指针每次走两步，慢指针每次走一步，所以每走一次快慢指针的间距就要缩小一步，在图一中当走n次的时候就会相遇，在图二中当走m-n次的时候就会相遇。

看下运行结果
![image.png](https://pic.leetcode-cn.com/1602042180-PAlkGo-image.png)

<br>

### 2，存放到集合中
这题还可以把节点存放到集合set中，每次存放的时候判断当前节点是否存在，如果存在，说明有环，直接返回true，比较容易理解
```
public boolean hasCycle(ListNode head) {
    Set<ListNode> set = new HashSet<>();
    while (head != null) {
        //如果重复出现说明有环
        if (set.contains(head))
            return true;
        //否则就把当前节点加入到集合中
        set.add(head);
        head = head.next;
    }
    return false;
}
```

<br>

### 3，逐个删除
一个链表从头节点开始一个个删除，**所谓删除就是让他的next指针指向他自己**。如果没有环，从头结点一个个删除，最后肯定会删完，如下图所示

![image.png](https://pic.leetcode-cn.com/1602042271-IRCKvm-image.png)

如果是环形的，那么有两种情况，一种是o型的，一种是6型的。原理都是一样，我们就看一下o型的

![image.png](https://pic.leetcode-cn.com/1602042284-WdrKhg-image.png)

如上图所示，如果删到最后，肯定会出现**head=head.next；**
```
public boolean hasCycle(ListNode head) {
    //如果head为空，或者他的next指向为空，直接返回false
    if (head == null || head.next == null)
        return false;
    //如果出现head.next = head表示有环
    if (head.next == head)
        return true;
    ListNode nextNode = head.next;
    //当前节点的next指向他自己，相当于把它删除了
    head.next = head;
    //然后递归，查看下一个节点
    return hasCycle(nextNode);
}
```
看一下运行结果
![image.png](https://pic.leetcode-cn.com/1602042338-HHNYOF-image.png)

<br>

### 4，先反转在比较
关于链表的反转可以看下[432，剑指 Offer-反转链表的3种方式](https://mp.weixin.qq.com/s?__biz=MzU0ODMyNDk0Mw==&mid=2247488340&idx=1&sn=c3d6adc9f737672aab544931502dda2e&chksm=fb418074cc360962b46cb764068a5818f58bed6a4cd05ef61057823918d95f3a192550f02408&token=723562173&lang=zh_CN#rd)。```如果有环，那么链表反转之后，原来的头结点和反转之后的头结点一定是同一个```

```
    public ListNode reverseList(ListNode head) {
        //新链表
        ListNode newHead = null;
        while (head != null) {
            //先保存访问的节点的下一个节点，保存起来
            //留着下一步访问的
            ListNode temp = head.next;
            //每次访问的原链表节点都会成为新链表的头结点，
            //其实就是把新链表挂到访问的原链表节点的
            //后面就行了
            head.next = newHead;
            //更新新链表
            newHead = head;
            //重新赋值，继续访问
            head = temp;
        }
        //返回新链表
        return newHead;
    }

    public boolean hasCycle(ListNode head) {
        ListNode rev = reverseList(head);
        if (head != null && head.next != null && rev == head) {
            return true;
        }
        return false;
    }
```


看一下运行结果
![image.png](https://pic.leetcode-cn.com/1602207381-UfCJXk-image.png)

<br>

### 5，快指针每次走3步，慢指针每次走1步

```
    public boolean hasCycle(ListNode head) {
        if (head == null)
            return false;
        ListNode slow = head;
        ListNode fast = head;
        while (fast != null && fast.next != null && fast.next.next != null) {
            slow = slow.next;
            fast = fast.next.next.next;
            if (slow == fast)
                return true;
        }
        return false;
    }
```
看下运行结果
![image.png](https://pic.leetcode-cn.com/1602207737-bgSqiX-image.png)

<br>
### 6，快指针每次走4步，慢指针每次走1步

```
    public boolean hasCycle(ListNode head) {
        if (head == null)
            return false;
        ListNode slow = head;
        ListNode fast = head;
        while (fast != null && fast.next != null && fast.next.next != null && fast.next.next.next != null) {
            slow = slow.next;
            fast = fast.next.next.next.next;
            if (slow == fast)
                return true;
        }
        return false;
    }
```
看下运行结果
![image.png](https://pic.leetcode-cn.com/1602208450-nOwHFB-image.png)



<br>
### 7，快指针每次走5步，慢指针每次走1步

```
    public boolean hasCycle(ListNode head) {
        if (head == null)
            return false;
        ListNode slow = head;
        ListNode fast = head;
        while (fast != null && fast.next != null && fast.next.next != null && fast.next.next.next != null && fast.next.next.next.next != null) {
            slow = slow.next;
            fast = fast.next.next.next.next.next;
            if (slow == fast)
                return true;
        }
        return false;
    }
```
看下运行结果
![image.png](https://pic.leetcode-cn.com/1602207826-aWAbkp-image.png)

<br>

### 8，快指针每次走3步，慢指针每次走2步

```
    public boolean hasCycle(ListNode head) {
        if (head == null)
            return false;
        ListNode slow = head;
        ListNode fast = head;
        while (fast != null && fast.next != null && fast.next.next != null) {
            slow = slow.next.next;
            fast = fast.next.next.next;
            if (slow == fast)
                return true;
        }
        return false;
    }
```
看下运行结果
![image.png](https://pic.leetcode-cn.com/1602207876-lOeajQ-image.png)

<br>

### 9，快指针每次走4步，慢指针每次走2步
```
    public boolean hasCycle(ListNode head) {
        if (head == null)
            return false;
        ListNode slow = head;
        ListNode fast = head;
        while (fast != null && fast.next != null && fast.next.next != null && fast.next.next.next != null) {
            slow = slow.next.next;
            fast = fast.next.next.next.next;
            if (slow == fast)
                return true;
        }
        return false;
    }
```
看下运行结果
![image.png](https://pic.leetcode-cn.com/1602207941-buAjda-image.png)


<br>

### 总结
**前面4种方式都很好理解，后面几种，我试了一下如果慢指针每次走m步，快指针每次走n步，只要m≠n，好像都能通过，如果在这样写下去，那么答案就无穷多了，这里就不在继续写下去了。关于这个问题是否永远成立，我也不太清楚，有兴趣的可以证明一下，并且可以在下面写出**




<br>

**如果觉得有用就给个赞吧，你的赞是给我最大的鼓励，也是我写作的最大动力**

![image.png](https://pic.leetcode-cn.com/d56a80459005b444404d2ad6fbaabdabecd2b9ed3cb2cf432e570c315ae2fcf7-image.png)
##### 查看更多答案，可扫码关注我微信公众号“**[数据结构和算法](https://img-blog.csdnimg.cn/20200807155236311.png)**”
