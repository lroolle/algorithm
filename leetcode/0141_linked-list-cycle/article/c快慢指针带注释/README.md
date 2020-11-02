"+++
title = "0141. Linked List Cycle C++快慢指针（带注释） "
author = ["xi-yu-shi-liu-guang-3"]
date = 2020-02-12T05:07:52+08:00
tags = ["Leetcode", "Algorithms", "C++", "LinkedList"]
categories = ["leetcode"]
draft = false
+++

# C++快慢指针（带注释）

> [0141. Linked List Cycle](https://leetcode-cn.com/problems/linked-list-cycle/)
> [C++快慢指针（带注释）](https://leetcode-cn.com/problems/linked-list-cycle/solution/ckuai-man-zhi-zhen-dai-zhu-shi-by-xi-yu-shi-liu-gu/) by [xi-yu-shi-liu-guang-3](https://leetcode-cn.com/u/xi-yu-shi-liu-guang-3/)

### 解题思路
假如该链表是循环链表，那我们可以定义两个指针，一个每次向前移动两个节点，另一个每次向前移动一个节点。这就和田径比赛是一样的，假如这两个运动员跑的是直道，那快的运动员和慢的运动员在起点位于同一位置，但快的运动员必将先到达终点，期间这两个运动员不会相遇。而如果绕圈跑的话（假设没有米数限制），跑的快的运动员在超过跑的慢的运动员一圈的时候，他们将会相遇，此刻就是循环链表。

### 代码

```cpp
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
bool hasCycle(ListNode* head) 
{
	//两个运动员位于同意起点head
	ListNode* faster{ head };  //快的运动员
	ListNode* slower{ head };  //慢的运动员

	if (head == NULL)  //输入链表为空，必然不是循环链表
		return false;

	while (faster != NULL && faster->next != NULL)
	{
		faster = faster->next->next;  //快的运动员每次跑两步
		slower = slower->next;  //慢的运动员每次跑一步
		if (faster == slower)  //他们在比赛中相遇了
			return true;  //可以断定是环形道，直道不可能相遇
	}
	return false;  //快的运动员到终点了，那就是直道，绕圈跑不会有终点
}
};
```