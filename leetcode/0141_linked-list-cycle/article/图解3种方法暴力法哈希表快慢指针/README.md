"+++
title = "0141. Linked List Cycle 「图解」3种方法：暴力法、哈希表、快慢指针 "
author = ["xiao_ben_zhu"]
date = 2020-05-25T11:34:03+08:00
tags = ["Leetcode", "Algorithms", "TwoPointers", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# 「图解」3种方法：暴力法、哈希表、快慢指针

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [「图解」3种方法：暴力法、哈希表、快慢指针](https://leetcode-cn.com/problems/linked-list-cycle/solution/lei-si-xiao-xue-de-zhui-ji-wen-ti-kuai-man-zhi-zhe/) by [xiao_ben_zhu](https://leetcode-cn.com/u/xiao_ben_zhu/)


#### 暴力法
- 每遍历一个节点，就从头遍历之前所有节点
 ![微信截图_20200525174038.png](https://pic.leetcode-cn.com/6cf3006161dddb53205efb6ef79d5904aae8c1fc8e20335e97f179266c2b954b-%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20200525174038.png)
-  比如遍历到 3，就从头访问 1 和 2，没有发现之前有节点 3，继续往下遍历。
- 当第二次遍历到 3 时，又从头访问之前的节点，发现之前遍历过节点 3，说明链表有环
- 时间复杂度为 O(n^2)，空间复杂度为 O(1)

```js
var hasCycle = (head) => {
  let cur = head;     // cur指针
  let step1 = 0;      // step1统计cur指针走的步数
  while (cur) {
    step1++
    let cur2 = head;   // cur2指针 从头开始
    let step2 = 0;     // step2统计cur2指针走的步数
    while (cur2) {
      step2++;
      if (cur == cur2) {      // cur和cur2指针指向相同的元素
        if (step1 == step2) { // cur2指针和cur走的步数一样，走到了cur这里
          break;              // 退出内层循环
        } else {              // cur2指针和cur相遇，但步数不一样
          return true;        // 说明链表有环，cur比cur2正好多走一个环
        }
      }
      cur2 = cur2.next; // cur2一次走一步
    }
    cur = cur.next;     // cur一次走一步
  }
  return false;
};
```
#### 方法2：借助哈希表
- 哈希表存储曾经遍历过的节点，遍历每一个节点，都查看哈希表是否存在当前节点，如果存在，则说明链表有环
- 如果不存在，则存入哈希表，并继续遍历下一节点
- 时间复杂度为 O(n)，空间复杂度为 O(n)

```js
var hasCycle = (head) => {
  let map = new Map()
  while (head) {
    if (map.has(head)) return true
    map.set(head, true)
    head = head.next
  }
  return false
}
```

#### 快慢指针法
- 快指针和慢指针，初始都指向头节点
- 慢指针每次走一步，快指针每次走两步，不断比较它们指向的节点的值
- 如果节点值相同，说明有环。如果不同，继续循环。
- 类似“追及问题”
  - 两个人在环形跑道上赛跑，同一个起点出发，一个跑得快一个跑得慢，在某一时刻，跑得快的必定会追上跑得慢的，只要是跑道是环形的，不是环形就肯定追不上。

```js
var hasCycle = (head) => {
  let fastP = head;
  let slowP = head;
  while (fastP) {                         // 快指针指向真实节点
    if (fastP.next == null) return false; // 如果下一个为null，说明没有环
    slowP = slowP.next;                   // 慢的走一步
    fastP = fastP.next.next;              // 快的走两步
    if (slowP == fastP) return true;      // 快慢指针相遇，有环
  }
  return false;                           // fastP指向null了，也始终不相遇
}
```
#### 如果有帮助，不妨点个赞鼓励我继续写下去，如果哪里写得不对不好，指出我我继续修改。**
