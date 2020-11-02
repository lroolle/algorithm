"+++
title = "0226. Invert Binary Tree 前、中、后序遍历、层序遍历 "
author = ["liweiwei1419"]
date = 2019-12-10T17:39:30+08:00
tags = ["Leetcode", "Algorithms", "DepthfirstSearch", "BreadthfirstSearch"]
categories = ["leetcode"]
draft = false
+++

# 前、中、后序遍历、层序遍历

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [前、中、后序遍历、层序遍历](https://leetcode-cn.com/problems/invert-binary-tree/solution/qian-zhong-hou-xu-bian-li-ceng-xu-bian-li-by-liwei/) by [liweiwei1419](https://leetcode-cn.com/u/liweiwei1419/)
思路：只要是遍历的方法都适用。

### 方法一：前序遍历

**参考代码 1**：

```Java
class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }
}

public class Solution {

    public TreeNode invertTree(TreeNode root) {
        if (root == null) {
            return null;
        }

        // 左子树和右子树交换，即使左右子树都空也不影响正确性
        TreeNode temp = root.left;
        root.left = root.right;
        root.right = temp;

        // 递归翻转左右子树
        invertTree(root.left);
        invertTree(root.right);
        return root;
    }
}
```
```Python
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
class Solution:

    def invertTree(self, root: TreeNode) -> TreeNode:
        if root is None:
            return

        # 交换左右子树
        root.left, root.right = root.right, root.left

        self.invertTree(root.left)
        self.invertTree(root.right)
        return root
```

### 方法二：中序遍历

> **注意**：写中序遍历的时候，不能仅仅只是将前序遍历的代码顺序调整一下。

因为在“中序遍历”的时候，左右子树已经交换过了，因此原来写 `invertTree(root.right);` 的地方，应该写作 `invertTree(root.left);`。

**参考代码 2**：

```Java
public class Solution {

    public TreeNode invertTree(TreeNode root) {
        if (root == null) {
            return null;
        }

        invertTree(root.left);

        TreeNode temp = root.left;
        root.left = root.right;
        root.right = temp;

        // 注意：因为左右子树已经交换了，因此这里不能写 invertTree(root.right);
        // 即：这里的 root.left 就是交换之前的 root.right
        invertTree(root.left);
        return root;
    }
}
```
```Python
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
class Solution:

    def invertTree(self, root: TreeNode) -> TreeNode:
        if root is None:
            return
        self.invertTree(root.left)

        # 交换左右子树
        root.left, root.right = root.right, root.left

        # 注意：这里的 root.left 就是交换之前的 root.right
        self.invertTree(root.left)
        return root
```

### 方法三：后序遍历

**参考代码 3**：

```Java
public class Solution {

    public TreeNode invertTree(TreeNode root) {
        if (root == null) {
            return null;
        }

        invertTree(root.left);
        invertTree(root.right);

        TreeNode temp = root.left;
        root.left = root.right;
        root.right = temp;
        return root;
    }
}
```
```Python
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
class Solution:

    def invertTree(self, root: TreeNode) -> TreeNode:
        if root is None:
            return

        self.invertTree(root.left)
        self.invertTree(root.right)

        # 交换左右子树
        root.left, root.right = root.right, root.left
        return root
```

### 方法四：层序遍历

**参考代码 4**：

```Java
import java.util.LinkedList;
import java.util.Queue;

public class Solution {

    public TreeNode invertTree(TreeNode root) {
        // 结点为空的特殊情况要先考虑
        if (root == null) {
            return null;
        }

        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        while (!queue.isEmpty()) {
            TreeNode curNode = queue.poll();
            // 只要其中之一非空，我都交换，并且把非空的结点添加到队列里
            if (curNode.left != null || curNode.right != null) {
                // 先翻转
                TreeNode temp = curNode.left;
                curNode.left = curNode.right;
                curNode.right = temp;
                // 把非空的节点加入队列
                if (curNode.left != null) {
                    queue.offer(curNode.left);
                }
                if (curNode.right != null) {
                    queue.offer(curNode.right);
                }
            }
        }
        return root;
    }
}
```
```Python
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
class Solution:

    def invertTree(self, root: TreeNode) -> TreeNode:
        if root is None:
            return None

        queue = [root]

        while queue:
            cur_node = queue.pop(0)

            if cur_node.left or cur_node.right:
                cur_node.left, cur_node.right = cur_node.right, cur_node.left

                if cur_node.left:
                    queue.append(cur_node.left)
                if cur_node.right:
                    queue.append(cur_node.right)

        return root
```