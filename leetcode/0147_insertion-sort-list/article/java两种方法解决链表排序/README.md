"+++
title = "0147. Insertion Sort List Java，两种方法解决链表排序 "
author = ["whinight"]
date = 2020-08-29T09:42:41+08:00
tags = ["Leetcode", "Algorithms", "Java", "LinkedList"]
categories = ["leetcode"]
draft = false
+++

# Java，两种方法解决链表排序

> [0147. Insertion Sort List](https://leetcode-cn.com/problems/insertion-sort-list/)
> [Java，两种方法解决链表排序](https://leetcode-cn.com/problems/insertion-sort-list/solution/javaliang-chong-fang-fa-jie-jue-lian-biao-pai-xu-b/) by [whinight](https://leetcode-cn.com/u/whinight/)

### 方法一：将链表排序转化为数组排序
1.创建一个list，遍历链表，将链表中每个节点的值存到list当中；
2.将list排序；
3.从头遍历链表，将每个节点的值依次替换为list中的数据。
### 代码

```java
class Solution {
    public ListNode insertionSortList(ListNode head) {
        List<Integer> list=new ArrayList<>();
        ListNode p=head;
        while (p!=null){
            list.add(p.val);
            p=p.next;
        }
        Collections.sort(list);
        p=head;
        int i=0;
        while (p!=null){
            p.val=list.get(i++);
            p=p.next;
        }
        return head;
    }
}
```
### 方法二：链表插入排序
1.创建一个新的头结点dummy，便于链表的插入（这样原先的头结点head也有前驱节点dummy，不再特殊);
2.遍历链表，如果前后俩个节点已经有序，则不需要操作；
3.否则，从头找到当前节点cur需要插入的位置，插入进去

### 代码

```java
class Solution {
   public ListNode insertionSortList(ListNode head) {
        if (head==null ||head.next==null)
            return head;
        ListNode pre=head,cur=head.next;           //使用前驱节点pre便于后续节点的删除操作
        ListNode dummy=new ListNode(0);         //建立一个头结点，便于链表的插入
        dummy.next=head;
        while (cur!=null){
            if (pre.val<cur.val){                   //前后节点已经有序，无需重排
                pre=cur;
                cur=cur.next;
            }
            else {
                ListNode p=dummy;
                while (p.next!=cur &&p.next.val<cur.val)
                    p=p.next;
                pre.next=cur.next;         //删除当前节点
                cur.next=p.next;          //将当前节点连接到对应位置
                p.next=cur;
                cur=pre.next;
            }
        }
        return dummy.next;
    }

}
```