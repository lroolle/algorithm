"+++
title = "0117. Populating Next Right Pointers in Each Node II 0ms无脑简洁带中文注释Golang解 "
author = ["ukcuf"]
date = 2020-08-05T13:24:13+08:00
tags = ["Leetcode", "Algorithms", "Go"]
categories = ["leetcode"]
draft = false
+++

# 0ms无脑简洁带中文注释Golang解

> [0117. Populating Next Right Pointers in Each Node II](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/)
> [0ms无脑简洁带中文注释Golang解](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/solution/0mswu-nao-jian-ji-dai-zhong-wen-zhu-shi-golangjie-/) by [ukcuf](https://leetcode-cn.com/u/ukcuf/)

### 解题思路

![image.png](https://pic.leetcode-cn.com/f83b2996e3689b17bbaa3e2eb8a66b752ab72d82c7ce26fff4af580be6a54431-image.png)

假设当前层所有节点的Next指针都已经连接完毕，那么通过Next来迭代当前层的每个节点，依次处理当前节点的下一层Left和Right节点即可。
### 代码

```golang
func connect(root *Node) *Node {
	n := root
	var prev, nextLevel *Node  // prev是被连接这层的上一次成功连接的节点，nextLevel是这层的第一个节点。
	for n != nil {             // 来迭代吧，一直到节点世界的尽头。。。。。
		prev, nextLevel = nil, nil   // 下一层目前还没有成功连接Next的节点，也找到下一层的第一个节点
		for n != nil {               // 循环到本层没节点为止
			connectNode(n.Left, &prev, &nextLevel)  // 尝试连接本层当前节点的Left child节点
			connectNode(n.Right, &prev, &nextLevel) // 尝试连接本层当前节点的Right child节点
			n = n.Next        // 当前节点的children节点处理完，处理下一个节点
		}
		n = nextLevel         // 这层的children都连接完了，那么该通过相同套路迭代刚刚连接完Next的这层了
	}
	return root
}

func connectNode(n *Node, prev, nextLevel **Node) {
	switch {
	case n == nil:      // 这个节点是nil，直接返回
		return
	case *nextLevel == nil:  
		*nextLevel = n  // 这个节点是这层第一个不是nil的节点，也就是首节点
		*prev = n       // 因此也是第一个我们见到的活着的非nil节点
		return
	}
	(*prev).Next = n   // 之前这层已经连接到了prev节点，那么prev节点的Next指向这个节点吧 
	*prev = n          // 更新已经被连接完的这个节点为prev节点
}
```