"+++
title = "0145. Binary Tree Postorder Traversal 迭代：一行代码  变换  二叉树的前序，中序，后序遍历 "
author = ["lsyf"]
date = 2020-09-01T11:34:41+08:00
tags = ["Leetcode", "Algorithms", "Go", "DepthfirstSearch", "Tree"]
categories = ["leetcode"]
draft = false
+++

# 迭代：一行代码  变换  二叉树的前序，中序，后序遍历

> [0145. Binary Tree Postorder Traversal](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)
> [迭代：一行代码  变换  二叉树的前序，中序，后序遍历](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/solution/yi-xing-dai-ma-bian-huan-er-cha-shu-de-qian-xu-zho/) by [lsyf](https://leetcode-cn.com/u/lsyf/)

### 解题思路
1. 利用栈的后进先出，分别对每一个节点 压入中，右，左，弹出栈顶获取其值，重复上述

2. 所以每个节点遍历三次
	- 先是被压入
	- 弹出后 又压入 （压入中，右，左节点）
	- 第三次是 当左右节点遍历完成后，再弹出 获取其值
	
3. 例如root,首先被压入；然后弹出 并压入root，root.Right和root.Left；处理完root.Left后，再次弹出root节点获取其值

**注意**：因为弹出是左，右，中的顺序，导致中节点 弹出栈时，有可能再次压入中，右，左节点，所以需要标记已处理防止陷入死循环

### 代码

```golang
func postorderTraversal(root *TreeNode) []int {
	//存储结果
	var res []int
	if root == nil {
		return res
	}
	//通过栈控制遍历顺序
	stack := []*TreeNode{root}
	//用map标记某个节点是否 遍历过
	marks := make(map[*TreeNode]bool)

	var node *TreeNode
	for len(stack) != 0 {
		//pop最后一个节点
		i := len(stack) - 1
		node = stack[i]
		stack = stack[:i]

		//如果节点已经遍历过，则打印
		_, ok := marks[node]
		if ok {
			res = append(res, node.Val)
		} else {
			//否则，入栈该节点和右孩子，左孩子，并标记当前节点已经遍历过(下次不再处理，可直接打印)

			//TODO 更改下行代码的位置 可以转换 前序，中序，后序遍历(当前为后序遍历)
			stack = append(stack, node)

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			//TODO 上述代码放到这里即 中序遍历
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			//TODO 上述代码放到这里即 前序遍历
			marks[node] = true
		}
	}
	return res
}

```