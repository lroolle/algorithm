"+++
title = "0617. Merge Two Binary Trees 动画演示 递归+迭代 617.合并二叉树 "
author = ["wang_ni_ma"]
date = 2020-01-31T12:22:59+08:00
tags = ["Leetcode", "Algorithms", "Java", "Python", "Tree"]
categories = ["leetcode"]
draft = false
+++

# 动画演示 递归+迭代 617.合并二叉树

> [0617. Merge Two Binary Trees](https://leetcode-cn.com/problems/merge-two-binary-trees/)
> [动画演示 递归+迭代 617.合并二叉树](https://leetcode-cn.com/problems/merge-two-binary-trees/solution/dong-hua-yan-shi-di-gui-die-dai-617he-bing-er-cha-/) by [wang_ni_ma](https://leetcode-cn.com/u/wang_ni_ma/)

## 递归实现
如果没有头绪的话，可以将这两颗树想象成是两个数组：   
```
1 3 2 5
2 1 3 4 7
```
合并两个数组很直观，将数组2的值合并到数组1中，再返回数组1就可以了。   
对于二叉树来说，如果我们像遍历数组那样，挨个遍历两颗二叉树中的每个节点，再把他们相加，那问题就比较容易解决了。  
   
遍历二叉树很简单，用**前序**遍历就可以了，再依次把访问到的节点值相加，因为题目没有说不能改变树的值和结构，我们不用再创建新的节点了，直接将树2合并到树1上再返回就可以了。   
需要注意：这两颗树并不是长得完全一样，有的树可能有左节点，但有的树没有。 
对于这种情况，我们统一的都把他们挂到树1 上面就可以了，对于上面例子中的两颗树，合并起来的结果如下：
```
	     3
	    / \
	   4   5
	  / \   \ 
	 5   4   7
```
相当于树1少了一条腿，而树2有这条腿，那就把树2的拷贝过来。   
总结下递归的条件：
1. 终止条件：树1的节点为null，或者树2的节点为null   
2. 递归函数内：将两个树的节点相加后，再赋给树1的节点。再递归的执行两个树的左节点，递归执行两个树的右节点   
    
动画演示如下：   
![recursion.gif](https://pic.leetcode-cn.com/23fbf9388a4193475a7606a6390729f575e3329e0a810d2047682f701d3ddd1f-recursion.gif)
时间复杂度：O(N)   
空间复杂度：O(h)，h是树的高度      
代码实现：     
```java
class Solution {
	public TreeNode mergeTrees(TreeNode t1, TreeNode t2) {
		if(t1==null || t2==null) {
			return t1==null? t2 : t1;
		}
		return dfs(t1,t2);
	}
	
	TreeNode dfs(TreeNode r1, TreeNode r2) {
		// 如果 r1和r2中，只要有一个是null，函数就直接返回
		if(r1==null || r2==null) {
			return r1==null? r2 : r1;
		}
		//让r1的值 等于  r1和r2的值累加，再递归的计算两颗树的左节点、右节点
		r1.val += r2.val;
		r1.left = dfs(r1.left,r2.left);
		r1.right = dfs(r1.right,r2.right);
		return r1;
	}
}
```
```python
class Solution(object):
	def mergeTrees(self, t1, t2):
		"""
		:type t1: TreeNode
		:type t2: TreeNode
		:rtype: TreeNode
		"""		
		def dfs(r1,r2):
			# 如果 r1和r2中，只要有一个是null，函数就直接返回
			if not (r1 and r2):
				return r1 if r1 else r2
			# 让r1的值 等于  r1和r2的值累加
			# 再递归的计算两颗树的左节点、右节点
			r1.val += r2.val
			r1.left = dfs(r1.left,r2.left)
			r1.right = dfs(r1.right,r2.right)
			return r1
		return dfs(t1,t2)
```
   
   
   
## 迭代实现
迭代实现用的是广度优先算法，广度优先就需要额外的数据结构来辅助了，我们可以借助栈或者队列来完成。   
只要两颗树的左节点都不为null，就把将他们放入队列中；同理只要两棵树的右节点都不为null了，也将他们放入队列中。     
然后我们不断的从队列中取出节点，把他们相加。   
如果出现  树1的left节点为null，树2的left不为null，直接将树2的left赋给树1就可以了；同理如果树1的right节点为null，树2的不为null，将树2的right节点赋给树1。   
动画图如下：     
![iterator.gif](https://pic.leetcode-cn.com/e252bdefa83701034a5c0551b960e6537650d42fd5acfdadcd58a417a985fe37-iterator.gif)

时间复杂度:O(N)   
空间复杂度:O(N)，对于满二叉树时，要保存所有的叶子节点，即N/2个节点。
   
更多二叉树相关题解请点击 [这里](https://mp.weixin.qq.com/mp/appmsgalbum?action=getalbum&album_id=1377493411015819266&__biz=MzI2NDE1MzY3Mw==#wechat_redirect) 查看   
   
      
代码实现：   
```java
class Solution {
	public TreeNode mergeTrees(TreeNode t1, TreeNode t2) {
	//如果 t1和t2中，只要有一个是null，函数就直接返回
		if(t1==null || t2==null) {
			return t1==null? t2 : t1;
		}
		java.util.LinkedList<TreeNode> queue = new java.util.LinkedList<TreeNode>();
		queue.add(t1);
		queue.add(t2);
		while(queue.size()>0) {
			TreeNode r1 = queue.remove();
			TreeNode r2 = queue.remove();
			r1.val += r2.val;
			//如果r1和r2的左子树都不为空，就放到队列中
			//如果r1的左子树为空，就把r2的左子树挂到r1的左子树上
			if(r1.left!=null && r2.left!=null){
				queue.add(r1.left);
				queue.add(r2.left);
			}
			else if(r1.left==null) {
				r1.left = r2.left;
			}
			//对于右子树也是一样的
			if(r1.right!=null && r2.right!=null) {
				queue.add(r1.right);
				queue.add(r2.right);
			}
			else if(r1.right==null) {
				r1.right = r2.right;
			}
		}
		return t1;
	}
}
```
```python
class Solution(object):
	def mergeTrees(self, t1, t2):
		"""
		:type t1: TreeNode
		:type t2: TreeNode
		:rtype: TreeNode
		"""	
	# 如果 t1和t2中，只要有一个是null，函数就直接返回
		if not (t1 and t2):
			return t2 if not t1 else t1
		queue = [(t1,t2)]
		while queue:
			r1,r2 = queue.pop(0)
			r1.val += r2.val
			# 如果r1和r2的左子树都不为空，就放到队列中
			# 如果r1的左子树为空，就把r2的左子树挂到r1的左子树上
			if r1.left and r2.left:
				queue.append((r1.left,r2.left))
			elif not r1.left:
				r1.left = r2.left
			# 对于右子树也是一样的
			if r1.right and r2.right:
				queue.append((r1.right,r2.right))
			elif not r1.right:
				r1.right = r2.right
		return t1
```
(全文完)	

**欢迎关注 👉👉👉  [【公众号】](https://pic.leetcode-cn.com/60594481855242ec2d1e373f1c706b4c649c336f3dcea1621e34c593e437c59f-617.jpg) 👈👈👈**   
    
**如果能再点个赞👍👍 就更感激啦💓💓** 
