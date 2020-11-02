"+++
title = "0094. Binary Tree Inorder Traversal 二叉树的中序遍历，新手从递归中找思路 "
author = ["zui-weng-jiu-xian"]
date = 2020-09-14T03:04:16+08:00
tags = ["Leetcode", "Algorithms", "Java", "Stack", "Tree"]
categories = ["leetcode"]
draft = false
+++

# 二叉树的中序遍历，新手从递归中找思路

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [二叉树的中序遍历，新手从递归中找思路](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/er-cha-shu-de-zhong-xu-bian-li-xin-shou-cong-di-gu/) by [zui-weng-jiu-xian](https://leetcode-cn.com/u/zui-weng-jiu-xian/)

## 思路
从递归版本获取灵感 (从大佬题解中抄袭)，如果是递归，前中后遍历无非就是访问根节点的顺序不一样
```java
public void pre_in_post_order(TreeNode root) {
    if (root == null) { return; }
    // visit(root); // 在这访问是前序遍历
    pre_in_post_order(root.left);
    // visit(root); // 在这访问是前序遍历
    pre_in_post_order(root.right);
    // visit(root); // 在这访问是前序遍历
}
```
用栈模拟递归实现迭代，我觉得要注意以下几点：

1. 递归中，`root`是怎么变化的。

2. 什么时候访问根节点。

以中序遍历为例子：
`root`的变化首先一直是`root = root.left`直至为`null` (注释1)；
然后回头，这时访问根节点 (注释2)；
访问后检查它的右子树，对右子树采用以上的相同方法 (注释3)。
```java
public void inorder(TreeNode root) {
    if (root == null) { return; }   
    pre_in_post_order(root.left);   // 1
    visit(root);                    // 2
    pre_in_post_order(root.right);  // 3

}
```

按照这个顺序，用栈模拟递归
使用双端队列`Deque`来实现栈`stack`，可以获得更高的效率。
限制只能队首进，队首出，就是栈了。
限制只能队首进，队尾出，就是普通队列了。
## 中序遍历
```java
class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> ans = new LinkedList<>();         // 结果
        Deque<TreeNode> deque = new ArrayDeque<>();     // dequestack
        while (root != null || !deque.isEmpty()) {      // 栈非空   
            while (root != null) {      // 左子树
                deque.addFirst(root);   // 入栈，为了能回头
                root = root.left;
            }
            root = deque.pollFirst();   // 回头
            ans.add(root.val);          // 这时访问根节点
            root = root.right;          // 对右子树采用以上相同做法
        }
        return ans;     // 送它出道
    }
}
```
## 算法分析
设`n`表示二叉树的节点个数

1. 遍历访问所有节点一次，时间复杂度为$O(n)$.

2. `ans`的空间不算在内，`deque`的空间为$O(n)$，空间复杂度为$O(n)$.

因为初次接触，顺便把前序遍历和后序遍历一起写了

## 前序遍历
只是改变一下访问根节点的时机。
```java
class Solution {
    public List<Integer> preorderTraversal(TreeNode root) {
        List<Integer> ans = new LinkedList<>();
        Deque<TreeNode> deque = new ArrayDeque<>();
        while (root != null || !deque.isEmpty()) {
            while (root != null) {          // 左子树
                ans.add(root.val);          // 访问根节点
                deque.offerFirst(root);     
                root = root.left;
            }
            root = deque.pollFirst();
            root = root.right;
        }
        return ans;
    }
}
```

## 后序遍历

因为后序遍历的顺序是：左右根，从上面的代码可以看到，如果进行`root = deque.pollFirst();`时还不访问根节点，以后都将失去访问机会，因为后序遍历访问完左子树，还要去访问右子树，才能访问根。
所以，后序遍历需要做一个巧妙的变换，颠倒访问顺序为：根右左，这样，只需要对前序遍历的代码做一丢丢改动，最后来一次`reverse`即可。
```java
class Solution {
    public List<Integer> postorderTraversal(TreeNode root) {
        List<Integer> ans = new ArrayList<>();
        Deque<TreeNode> deque = new ArrayDeque<>();
        while (root != null || !deque.isEmpty()) {
            while (root != null) {      // 右子树
                ans.add(root.val);      // 这时访问根节点
                deque.offerFirst(root);
                root = root.right;      // 右
            }
            root = deque.pollFirst();   
            root = root.left;           // 对左子树采用以上相同做法
        }
        Collections.reverse(ans);       // 反转
        return ans;
    }
}
```

栈模拟都是需要$O(n)$空间的，如果想得到$O(1)$空间，参考官方给出的`Morris`算法。
[Morris算法](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/er-cha-shu-de-zhong-xu-bian-li-by-leetcode-solutio/)

另外还发现了一个大佬的颜色标记法，非常简洁，推荐大伙儿看。
[颜色标记法](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/yan-se-biao-ji-fa-yi-chong-tong-yong-qie-jian-ming/)

### 如果本文对你有帮助，可以给个大拇指呀！
### 如果你有什么建议或疑问，可以评论留言呀！



