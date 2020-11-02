"+++
title = "0226. Invert Binary Tree 递归，BFS和DFS多种方式解决（全部击败了100%的用户） "
author = ["sdwwld"]
date = 2020-09-16T01:28:57+08:00
tags = ["Leetcode", "Algorithms", "Java", "BreadthfirstSearch", "DepthfirstSearch", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 递归，BFS和DFS多种方式解决（全部击败了100%的用户）

> [0226. Invert Binary Tree](https://leetcode-cn.com/problems/invert-binary-tree/)
> [递归，BFS和DFS多种方式解决（全部击败了100%的用户）](https://leetcode-cn.com/problems/invert-binary-tree/solution/di-gui-bfshe-dfsduo-chong-fang-shi-jie-jue-quan-bu/) by [sdwwld](https://leetcode-cn.com/u/sdwwld/)


### 1，递归方式解决
反转二叉树，可以先交换根节点的两个子节点，然后通过同样的方式在交换根节点的子节点的两个子节点……
![image.png](https://pic.leetcode-cn.com/1600217755-etthJR-image.png)
代码比较简单
```
    public TreeNode invertTree(TreeNode root) {
        //递归的边界条件判断
        if (root == null)
            return null;
        //先交换子节点
        TreeNode left = root.left;
        root.left = root.right;
        root.right = left;
        //递归调用
        invertTree(root.left);
        invertTree(root.right);
        return root;
    }
```
看下运行结果，效率还是挺高的
![image.png](https://pic.leetcode-cn.com/1600218112-JXIYjR-image.png)
递归有两种方式，一种是先交换，再递归调用。还一种是先递归调用，再交换。两种方式都可以实现
```
    public TreeNode invertTree(TreeNode root) {
        //递归的边界条件判断
        if (root == null)
            return null;
        //先递归
        TreeNode left = invertTree(root.left);
        TreeNode right = invertTree(root.right);
        //最后在交换
        root.left = right;
        root.right = left;
        return root;
    }
```
看下运行结果
![image.png](https://pic.leetcode-cn.com/1600218310-bbwbBJ-image.png)

### 2，BFS解决
如果对树的遍历比较熟悉的话，我们只要遍历树的所有节点，然后把他们的左右子节点相互交换即可，如果这样写，那么答案就比较多了。这里来看下二叉树的BFS解法，二叉树的BFS就是一层一层的遍历，下面这样
![image.png](https://pic.leetcode-cn.com/1600218566-IHCDSL-image.png)
二叉树的BFS代码如下
```
    public void levelOrder(TreeNode tree) {
        if (tree == null)
            return;
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(tree);//相当于把数据加入到队列尾部
        while (!queue.isEmpty()) {
            //poll方法相当于移除队列头部的元素
            TreeNode node = queue.poll();
            System.out.println(node.val);
            if (node.left != null)
                queue.add(node.left);
            if (node.right != null)
                queue.add(node.right);
        }
    }
```
我们就参照这种方式来写下，每次遍历节点的时候都要交换子节点，所以代码很容易写
```
    public TreeNode invertTree(TreeNode root) {
        if (root == null)
            return root;
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);//相当于把数据加入到队列尾部
        while (!queue.isEmpty()) {
            //poll方法相当于移除队列头部的元素
            TreeNode node = queue.poll();
            //先交换子节点
            TreeNode left = node.left;
            node.left = node.right;
            node.right = left;

            if (node.left != null)
                queue.add(node.left);
            if (node.right != null)
                queue.add(node.right);
        }
        return root;
    }
```
看一下运行结果
![image.png](https://pic.leetcode-cn.com/1600218902-YCDFWW-image.png)

### 3，DFS解决
上面说了只要能遍历二叉树的所有节点，然后交换子节点，就能完成这道题，二叉树还有一种方式是DFS遍历，他的代码如下
```
public static void treeDFS(TreeNode root) {
    Stack<TreeNode> stack = new Stack<>();
    stack.add(root);
    while (!stack.empty()) {
        TreeNode node = stack.pop();
        System.out.println(node.val);
        if (node.right != null) {
            stack.push(node.right);
        }
        if (node.left != null) {
            stack.push(node.left);
        }
    }
}
```
我们来参照这种方式写下
```
    public TreeNode invertTree(TreeNode root) {
        if (root == null)
            return root;
        Stack<TreeNode> stack = new Stack<>();
        stack.add(root);
        while (!stack.empty()) {
            TreeNode node = stack.pop();
            //先交换子节点
            TreeNode left = node.left;
            node.left = node.right;
            node.right = left;
            if (node.right != null) {
                stack.push(node.right);
            }
            if (node.left != null) {
                stack.push(node.left);
            }
        }
        return root;
    }
```
看下运行结果
![image.png](https://pic.leetcode-cn.com/1600219119-nKKcgs-image.png)

### 总结：
这道题其实很简答，只要能遍历二叉树的所有节点都可以轻松完成，我们还可以使用二叉树的中序遍历来解决，二叉树的中序遍历代码如下
```
public static void inOrderTraversal(TreeNode tree) {
    Stack<TreeNode> stack = new Stack<>();
    while (tree != null || !stack.isEmpty()) {
        while (tree != null) {
            stack.push(tree);
            tree = tree.left;
        }
        if (!stack.isEmpty()) {
            tree = stack.pop();
            System.out.println(tree.val);
            tree = tree.right;
        }
    }
}
```
修改一下
```
    public TreeNode invertTree(TreeNode root) {
        if (root == null)
            return root;
        Stack<TreeNode> stack = new Stack<>();
        TreeNode node = root;
        while (node != null || !stack.isEmpty()) {
            while (node != null) {
                stack.push(node);
                node = node.left;
            }
            if (!stack.isEmpty()) {
                node = stack.pop();
                //先交换子节点
                TreeNode left = node.left;
                node.left = node.right;
                node.right = left;
                //注意，这里是交换之后的，所以要修改
                node = node.left;
            }
        }
        return root;
    }
```
看下运行结果
![image.png](https://pic.leetcode-cn.com/1600219635-ZAEBSC-image.png)


<br>


**如果觉得有用就给个赞吧，你的赞是给我最大的鼓励，也是我写作的最大动力**

![image.png](https://pic.leetcode-cn.com/d56a80459005b444404d2ad6fbaabdabecd2b9ed3cb2cf432e570c315ae2fcf7-image.png)
##### 查看更多答案，可扫码关注我微信公众号“**[数据结构和算法](https://img-blog.csdnimg.cn/20200807155236311.png)**”
