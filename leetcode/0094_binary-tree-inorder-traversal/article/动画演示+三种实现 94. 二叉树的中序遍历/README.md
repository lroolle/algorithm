"+++
title = "0094. Binary Tree Inorder Traversal 动画演示+三种实现 94. 二叉树的中序遍历 "
author = ["wang_ni_ma"]
date = 2020-02-15T03:14:13+08:00
tags = ["Leetcode", "Algorithms", "Java", "Python", "Tree", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 动画演示+三种实现 94. 二叉树的中序遍历

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [动画演示+三种实现 94. 二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/dong-hua-yan-shi-94-er-cha-shu-de-zhong-xu-bian-li/) by [wang_ni_ma](https://leetcode-cn.com/u/wang_ni_ma/)

## 递归实现
递归遍历太简单了   
- 前序遍历:打印-左-右
- 中序遍历:左-打印-右
- 后序遍历:左-右-打印
   
题目要求的是中序遍历，那就按照 左-打印-右这种顺序遍历树就可以了，递归函数实现   
- 终止条件:当前节点为空时
- 函数内: 递归的调用左节点，打印当前节点，再递归调用右节点
   
时间复杂度:O(n)   
空间复杂度:O(h)，h是树的高度   
代码实现:   
```java
class Solution {
	public List<Integer> inorderTraversal(TreeNode root) {
		List<Integer> res = new ArrayList<Integer>();
		dfs(res,root);
		return res;
	}
	
	void dfs(List<Integer> res, TreeNode root) {
		if(root==null) {
			return;
		}
		//按照 左-打印-右的方式遍历
		dfs(res,root.left);
		res.add(root.val);
		dfs(res,root.right);
	}
}
```
```python
class Solution(object):
	def inorderTraversal(self, root):
		"""
		:type root: TreeNode
		:rtype: List[int]
		"""
		res = []
		def dfs(root):
			if not root:
				return
			# 按照 左-打印-右的方式遍历	
			dfs(root.left)
			res.append(root.val)
			dfs(root.right)
		dfs(root)
		return res
```
   
   
   
## 迭代实现
这题的真正难点在于如何用非递归的方式实现。   
递归实现时，是函数自己调用自己，一层层的嵌套下去，操作系统/虚拟机自动帮我们用**栈**来保存了每个调用的函数，现在我们需要自己模拟这样的调用过程。   
递归的调用过程是这样的：
```
dfs(root.left)
	dfs(root.left)
		dfs(root.left)
			为null返回
		打印节点
		dfs(root.right)
			dfs(root.left)
				dfs(root.left)
				........
```
递归的调用过程是不断往左边走，当左边走不下去了，就打印节点，并转向右边，然后右边继续这个过程。   
我们在迭代实现时，就可以用栈来模拟上面的调用过程。   
![1.gif](https://pic.leetcode-cn.com/47fff35dd3fd640ba60349c78b85242ae8f4b850f06a282cd7e92c91e6eff406-1.gif)
   
时间复杂度:O(n)   
空间复杂度:O(h)，h是树的高度   
代码实现:   
```java
class Solution {
	public List<Integer> inorderTraversal(TreeNode root) {
		List<Integer> res = new ArrayList<Integer>();
		Stack<TreeNode> stack = new Stack<TreeNode>();
		while(stack.size()>0 || root!=null) {
			//不断往左子树方向走，每走一次就将当前节点保存到栈中
			//这是模拟递归的调用
			if(root!=null) {
				stack.add(root);
				root = root.left;
			//当前节点为空，说明左边走到头了，从栈中弹出节点并保存
			//然后转向右边节点，继续上面整个过程
			} else {
				TreeNode tmp = stack.pop();
				res.add(tmp.val);
				root = tmp.right;
			}
		}
		return res;
	}
}
```
```python
class Solution(object):
	def inorderTraversal(self, root):
		"""
		:type root: TreeNode
		:rtype: List[int]
		"""
		res = []
		stack = []
		while stack or root:
			# 不断往左子树方向走，每走一次就将当前节点保存到栈中
			# 这是模拟递归的调用
			if root:
				stack.append(root)
				root = root.left
			# 当前节点为空，说明左边走到头了，从栈中弹出节点并保存
			# 然后转向右边节点，继续上面整个过程
			else:
				tmp = stack.pop()
				res.append(tmp.val)
				root = tmp.right
		return res
```
   
   
   
## 莫里斯遍历
用递归和迭代的方式都使用了辅助的空间，而莫里斯遍历的优点是没有使用任何辅助空间。   
缺点是改变了整个树的结构，强行把一棵二叉树改成一段链表结构。   
![2.jpg](https://pic.leetcode-cn.com/db36653b8b0385101dffc37179c4153450059a54360aa0e15aae0769e3ad6e6f-2.jpg)
我们将黄色区域部分挂到节点5的右子树上，接着再把2和5这两个节点挂到4节点的右边。   
这样整棵树基本上就变改成了一个链表了，之后再不断往右遍历。   
![3.gif](https://pic.leetcode-cn.com/c1b589b5fc7facd1a847c9f5bab407765222ee2d9e1a887a9e5d61cc9e94dfc6-3.gif)
时间复杂度:找到每个前驱节点的复杂度是O(n)，因为n个节点的二叉树有n-1条边，每条边只可能使用2次(一次定位到节点，一次找到前驱节点)，故时间复杂度为O(n)   
空间复杂度:O(1)   
```java
class Solution {
	public List<Integer> inorderTraversal(TreeNode root) {
		List<Integer> res = new ArrayList<Integer>();
		TreeNode pre = null;
		while(root!=null) {
			//如果左节点不为空，就将当前节点连带右子树全部挂到
			//左节点的最右子树下面
			if(root.left!=null) {
				pre = root.left;
				while(pre.right!=null) {
					pre = pre.right;
				}
				pre.right = root;
				//将root指向root的left
				TreeNode tmp = root;
				root = root.left;
				tmp.left = null;
			//左子树为空，则打印这个节点，并向右边遍历	
			} else {
				res.add(root.val);
				root = root.right;
			}
		}
		return res;
	}
}
```
```python
class Solution(object):
	def inorderTraversal(self, root):
		"""
		:type root: TreeNode
		:rtype: List[int]
		"""
		res = []
		pre = None
		while root:
			# 如果左节点不为空，就将当前节点连带右子树全部挂到
			# 左节点的最右子树下面
			if root.left:
				pre = root.left
				while pre.right:
					pre = pre.right
				pre.right = root
				# 将root指向root的left
				tmp = root
				root = root.left
				tmp.left = None
			# 左子树为空，则打印这个节点，并向右边遍历	
			else:
				res.append(root.val)
				root = root.right
		return res
```
(全文完)
   
**欢迎关注我的 👉👉👉  [【公众号】](https://pic.leetcode-cn.com/0cae070d4a5c154f93cb9da69b4a2e82ba655d92b8b95aed82846c390e851f99-94.jpg) 👈👈👈**   
   
**如果能再点个赞👍👍 就更感激啦💓💓**


