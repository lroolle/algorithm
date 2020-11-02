"+++
title = "0230. Kth Smallest Element in a BST 第K小元素的三种找法 "
author = ["lan-se-bei-ban-qiu"]
date = 2020-09-11T02:19:01+08:00
tags = ["Leetcode", "Algorithms", "Java", "Tree", "DFS"]
categories = ["leetcode"]
draft = false
+++

# 第K小元素的三种找法

> [0230. Kth Smallest Element in a BST](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/)
> [第K小元素的三种找法](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/solution/di-kxiao-yuan-su-de-san-chong-zhao-fa-by-lan-se-2/) by [lan-se-bei-ban-qiu](https://leetcode-cn.com/u/lan-se-bei-ban-qiu/)

## 方法一
根据二叉搜索树的中序遍历序列是升序的这一特性，我们可以对树进行中序遍历，得到树中元素的升序序列，最后再返回序列中的第k-1个元素即可。
```java
class Solution {
    List<Integer> res = new ArrayList<>();
    public int kthSmallest(TreeNode root, int k) {
        dfs(root);
        return res.get(k - 1);
    }

    public void dfs(TreeNode root){
        if(root == null)
            return;
        dfs(root.left);
        res.add(root.val);
        dfs(root.right);
    }
}
```

* 时间复杂度: $O(n)$
* 空间复杂度: $O(n)$

## 方法二
因为二叉搜索树左子树的元素都小于根节点，右子树的元素都大于根节点。因此：  
* 如果k不大于root的左子树的节点个数，那么以root为根的二叉搜索树的第k小元素也就是左子树的第k小元素
* 如果k正好比root的左子树的节点个数大一，那么第k小元素就是根节点root本身。
* 否则，二叉搜索树的第k小元素存在于右子树中，也即要寻找右子树的第$k - leftNum - 1$小的元素。（注：$leftNum+1$为左子树加根节点的节点个数）
```java
public int kthSmallest(TreeNode root, int k) {
    //获取左子树的节点数leftNum
    int leftNum = number(root.left);
    if(leftNum >= k)
        return kthSmallest(root.left, k);
    else if(leftNum + 1 == k)
        return root.val;
    else
        return kthSmallest(root.right, k - leftNum - 1);
}

//此函数用于获取以root为根节点的树的节点个数
public int number(TreeNode root){
    if(root == null)
        return 0;
    return number(root.left) + number(root.right) + 1;
}
```

* 时间复杂度: $O(nlogn)$
* 空间复杂度: $O(1)$

## 方法三
方法三可看作是对方法一的一个改进，我们在对二叉搜索树进行中序遍历的时候，不再用额外空间来存储树中节点信息。而是用一个变量count来记录节点当前已遍历到的节点个数，当遍历到第k个节点时(count等于k)，直接返回这个节点的值。
```java
class Solution {
    int count = 0;
    int res;
    int k;
    public int kthSmallest(TreeNode root, int k) {
        this.k = k;
        dfs(root);
        return res;
    }

    public void dfs(TreeNode root){
        if(root == null)
            return;
        dfs(root.left);
        if(++count == k){
            res = root.val;
            return;
        }
        dfs(root.right);
    }
}
```

* 时间复杂度: $O(n)$
* 空间复杂度: $O(1)$

---
若您觉得本题解对您有所帮助，欢迎关注[我的github仓库](https://github.com/wyh317/Leetcode)访问更多leetcode题解，期待和小伙伴们一起交流讨论！