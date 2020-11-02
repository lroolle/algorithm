"+++
title = "0147. Insertion Sort List C++ 95.83% 插入排序 简洁易懂（小白版） "
author = ["Smart_Shelly"]
date = 2020-03-03T15:33:27+08:00
tags = ["Leetcode", "Algorithms", "C++", "LinkedList", "Sort"]
categories = ["leetcode"]
draft = false
+++

# C++ 95.83% 插入排序 简洁易懂（小白版）

> [0147. Insertion Sort List](https://leetcode-cn.com/problems/insertion-sort-list/)
> [C++ 95.83% 插入排序 简洁易懂（小白版）](https://leetcode-cn.com/problems/insertion-sort-list/solution/ccha-ru-pai-xu-xiao-bai-ban-by-tryangel/) by [Smart_Shelly](https://leetcode-cn.com/u/smart_shelly/)

### 解题思路
*(第一次写题解hhh多多包涵~）*
![题解1.1.png](https://pic.leetcode-cn.com/b9f2e2e4b5ec50e56f54fdb4d8f84848efe242b9452d1b8a35ad9090238b0e3e-%E9%A2%98%E8%A7%A31.1.png)
本来以为理解了数组的插入排序会很好写，谁知....下面是我的心路历程：
首先，第一反应是像数组的插入排序一样，找到相应插入位置后将元素逐个后移，但发现这样复杂度过高。
然后，为了降低复杂度，则要减少移动的次数：**改变指针的指向！**

此题利用了**dummyhead伪头指针**，即构造一个指向head的指针，使得可以从head开始遍历。且由于指针在插入排序的过程中，所指向的元素在链表中的位置会发生改变，而dummyhead始终指向第一个元素不变，所以只有返回dummyhead->next才正确。
（在传参的时候，head指向的就是未排序中的第一个元素，且排序后head的指向仍未发生改变即始终指向排序前的第一个元素。但排序过程中，第一个元素的值可能发生改变。而dummyhead->next才会永远指向排序后的第一个元素）

BTW,我所理解的**tail指针**，即为指向已排好元素的最后一个的指针，即此处我所用的prev指针。且由于测试样例中，很多样例都是已经按顺序排好的，故将node和prev进行比较，可*用于提速*！
![题解1.png](https://pic.leetcode-cn.com/dc5194e5b45b3cfe28b31713ff110a70c1924ee77e4d66eed2459bf09701934b-%E9%A2%98%E8%A7%A31.png)
上图的两次提交分别都是使用了改变指针指向的方式。其中第一次（用时较长的）是while循环判断中取了=，（其实是更稳定的算法），而第二次（用时较短的）是while循环没取=的做法。（实际不够稳定）

（注释都是我编写代码的时候踩过的坑555）

### 代码

```cpp
class Solution {
public:
ListNode* insertionSortList(ListNode* head) {
	if (!head || !head->next)
		return head;
	ListNode *dummyhead = new ListNode(-1);//伪头指针
	dummyhead->next = head;
	ListNode *prev = head;
	ListNode *node = head->next;
	while (node)
	{
		if (node->val < prev->val)
		{
			ListNode* temp = dummyhead;//！！temp要等于dummyhead，这样才可以比较第一个元素
			while (temp->next->val < node->val)//！！！这里是temp->next：因为要修改前面的temp的指向
			{
				temp = temp->next;//指针后移
			}
			prev->next = node->next;
			node->next = temp->next;
			temp->next = node;
			node = prev->next;//此时不用改变prev指向！因为prev没有变，只是待排序元素变了位置。
		}
		else
		{
			prev = prev->next;
			node = node->next;
		}
	}
	return dummyhead->next;//!!!不能返回head！！因为后面会改变head所指向内存的位置！
}
};
```

如果题解对你有帮助或启发的话，不妨点个赞~
Thanks for your reading!!