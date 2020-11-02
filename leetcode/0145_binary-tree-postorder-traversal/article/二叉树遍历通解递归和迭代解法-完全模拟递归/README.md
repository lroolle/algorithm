"+++
title = "0145. Binary Tree Postorder Traversal 二叉树遍历通解（递归和迭代解法）- 完全模拟递归 "
author = ["long_wotu"]
date = 2020-07-14T14:03:26+08:00
tags = ["Leetcode", "Algorithms", "二叉树", "Java", "Stack", "猿辅导", "DepthfirstSearch", "迭代", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 二叉树遍历通解（递归和迭代解法）- 完全模拟递归

> [0145. Binary Tree Postorder Traversal](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)
> [二叉树遍历通解（递归和迭代解法）- 完全模拟递归](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/solution/bian-li-tong-jie-by-long_wotu/) by [long_wotu](https://leetcode-cn.com/u/long_wotu/)

# 二叉树的遍历
> * 先序遍历
>   1. 判空
>   1. 访问节点
>   2. 左孩子入栈（再次从 1 开始执行）
>   3. 右孩子入栈（再次从 1 开始执行）
>
> * 中序遍历
>   * 判空
>   * 左孩子入栈
>   * 访问节点
>   * 右孩子入栈
>
> * 后序遍历
>   * 判空
>   * 左孩子入栈
>   * 右孩子入栈
>   * 访问节点
>
> **遍历分为两种解法 - 递归和迭代**
## 递归解法
二叉树遍历递归框架
```java 
List<Integer> ret = new ArrayList<Integer>();
public List<Integer> traversal(TreeNode root) {
    dfs(root);
    return ret;
}

public void dfs(TreeNode root) {
    if (root = null) return;
    // 先序遍历 
    dfs(root.left);
    // 中序遍历
    dfs(root.right);   
    // 后序遍历
}
```
递归框架思路（**`函数与方法等同`**）：
1. 如果当前节点存在左孩子，以左孩子为参数的递归函数——`public void dfs(TreeNode root){...}`**入系统栈**，执行**栈顶**方法（直到当前节点没有左孩子) -`dfs(root.left);`
2. 如果当前节点存在右孩子，则以右孩子为参数的递归函数——`public void dfs(TreeNode root){...}`**入系统栈**，执行**栈顶**方法（从第一步操作开始继续执行) - `dfs(root.right);`

基于递归框架的遍历思路：
* 先序遍历
    1. 访问当前节点 
    2. 如果当前节点存在左孩子，以左孩子为参数的递归函数——`public void dfs(TreeNode root){...}`**入系统栈**，执行**栈顶**方法（直到当前节点没有左孩子) -`dfs(root.left);`
    3. 如果当前节点存在右孩子，则以右孩子为参数的递归函数——`public void dfs(TreeNode root){...}`**入系统栈**，执行**栈顶**方法（从第一步操作开始继续执行) - `dfs(root.right);`
* 中序遍历
    1. 如果当前节点存在左孩子，以左孩子为参数的递归函数——`public void dfs(TreeNode root){...}`**入系统栈**，执行**栈顶**方法（直到当前节点没有左孩子) -`dfs(root.left);`
    2. 访问当前节点 
    3. 如果当前节点存在右孩子，则以右孩子为参数的递归函数——`public void dfs(TreeNode root){...}`**入系统栈**，执行**栈顶**方法（从第一步操作开始继续执行) - `dfs(root.right);`
* 后序遍历
    1. 如果当前节点存在左孩子，以左孩子为参数的递归函数——`public void dfs(TreeNode root){...}`**入系统栈**，执行**栈顶**方法（直到当前节点没有左孩子) -`dfs(root.left);`
    2. 如果当前节点存在右孩子，则以右孩子为参数的递归函数——`public void dfs(TreeNode root){...}`**入系统栈**，执行**栈顶**方法（从第一步操作开始继续执行) - `dfs(root.right);`
    3. 访问当前节点 
## 迭代解法
二叉树遍历迭代框架
```java
public List<Integer> traversal(TreeNode root) {
    if (root == null) return new ArrayList<Integer>();
    
    TreeNode node = root;
    List<Integer> ret = new ArrayList<Integer>();
    
    Stack<TreeNode> stack = new Stack<TreeNode>();
    while(node != null || !stack.isEmpty()) {
        while (node != null) {
            stack.push(node);
            // 先序遍历
            node = node.left;
        }
        node = stack.pop();
        // 中序遍历
        node = node.right;
        // 后序遍历
    }
    return ret;
}
```
迭代框架思路：

首先，分析当前节点在**递归框架**中的执行路径
1. 判断当前节点是否为空，若不为空，执行下面操作，否则返回，出系统栈，更新当前节点为栈顶节点。
2. 当前节点是否存在左孩子，若存在，则左孩子入系统栈，执行栈顶函数，更新当前节点为左孩子节点，循环执行这类操作，直到当前节点不存在左孩子为止。
3. 当前节点是否存在右孩子，若存在，则右孩子入系统栈，执行栈顶函数，更新当前节点为右孩子节点。

其次，将其执行路径通过**用户栈**进行模拟（形成先将左孩子入栈，再将右孩子入栈的顺序）
1. 判断当前节点或用户栈是否为空，若不为空，执行下面操作，否则全部流程结束。
2. 判断当前节点是否为空，若为空，则执行下一步，若不为空则将当前节点入用户栈（在递归框架中，隐含了第一个节点——root节点，的入栈操作，也就是第一次调用 `public void dfs(TreeNode root){...}`方法时，进行入系统栈操作），通过`while`循环执行持续的将左孩子入用户栈的操作。
3. 更新当前节点为用户栈出栈节点
4. 更新当前节点为其右孩子节点，进入下一个循环。 
区别：递归是将函数压入**系统栈-操作系统内部数据结构**，栈内存储的是函数，迭代是将节点压入**用户栈-自定义数据结构**，栈内存储的是二叉树节点。

----
1. 先序遍历
```java
public List<Integer> traversal(TreeNode root) {
    if (root == null) return new ArrayList<Integer>();
    
    TreeNode node = root;
    List<Integer> ret = new ArrayList<Integer>();
    
    Stack<TreeNode> stack = new Stack<TreeNode>();
    while(node != null || !stack.isEmpty()) {
        while (node != null) {
            stack.push(node);
            // 先序遍历
            ret.add(node.val);
            node = node.left;
        }
        node = stack.pop();
        node = node.right;          
    }
    return ret;
}
```
先序遍历思路
1. 判断当前节点或用户栈是否为空，若不为空，执行下面操作，否则全部流程结束。
2. 判断当前节点是否为空，若为空，则执行下一步，若不为空则**访问当前节点**并将当前节点入用户栈（在递归框架中，隐含了第一个节点——root节点，的入栈操作，也就是第一次调用 `public void dfs(TreeNode root){...}`方法时，进行入系统栈操作），通过`while`循环执行持续的将左孩子入用户栈的操作。
3. 更新当前节点为用户栈出栈节点
4. 更新当前节点为其右孩子节点，进入下一个循环。 

---

2. 中序遍历
```java
public List<Integer> traversal(TreeNode root) {
    if (root == null) return new ArrayList<Integer>();
    
    TreeNode node = root;
    List<Integer> ret = new ArrayList<Integer>();
    
    Stack<TreeNode> stack = new Stack<TreeNode>();
    while(node != null || !stack.isEmpty()) {
        while (node != null) {
            stack.push(node);
            node = node.left;
        }
        node = stack.pop();
        // 中序遍历
        ret.add(node.val);
        node = node.right;
    }
    return ret;
}
```
中序遍历思路
1. 判断当前节点或用户栈是否为空，若不为空，执行下面操作，否则全部流程结束。
2. 判断当前节点是否为空，若为空，则执行下一步，若不为空则将当前节点入用户栈（在递归框架中，隐含了第一个节点——root节点，的入栈操作，也就是第一次调用 `public void dfs(TreeNode root){...}`方法时，进行入系统栈操作），通过`while`循环执行持续的将左孩子入用户栈的操作。
3. 更新当前节点为用户栈出栈节点
4. **访问当前节点**
5. 更新当前节点为其右孩子节点，进入下一个循环。 

----

3. 后序遍历
```java
public List<Integer> postorderTraversal(TreeNode root) {
        if (root == null) return new ArrayList<Integer>();
        
        TreeNode node = root;
        List<Integer> ret = new ArrayList<Integer>();
        
        Stack<TreeNode> stack = new Stack<TreeNode>();
        while(node != null || !stack.isEmpty()) {
            while (node != null) {
                stack.push(node);
                node = node.left;
            }
            node = stack.pop();
            // 后序遍历
            // 如果没有右孩子或者右孩子被访问过了 {@Alex Zheng 感谢建议哈～}
            if (node.right == null || 
                    (ret.size() != 0 && ret.get(ret.size() - 1).equals(node.right.val)) ) {
                ret.add(node.val);
                node = null;
            }  else {
                stack.push(node);
                node = node.right;
            }
        }
        return ret;
    }
```
后序遍历思路
1. 判断当前节点或用户栈是否为空，若不为空，执行下面操作，否则全部流程结束。
2. 判断当前节点是否为空，若为空，则执行下一步，若不为空则将当前节点入用户栈（在递归框架中，隐含了第一个节点——root节点，的入栈操作，也就是第一次调用 `public void dfs(TreeNode root){...}`方法时，进行入系统栈操作），通过`while`循环执行持续的将左孩子入用户栈的操作。
3. 更新当前节点为用户栈出栈节点
4. 当右孩子为空及右孩子被访问过时，**访问当前节点**，更新当前节点为空，为下一步的出栈作准备
4. 当右孩子不为空并且右孩子未被访问过，则更新当前节点为右孩子 