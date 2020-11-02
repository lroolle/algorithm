"+++
title = "0235. Lowest Common Ancestor of a Binary Search Tree 二叉搜索树的最近公共祖先（3种解决方式） "
author = ["sdwwld"]
date = 2020-09-27T01:25:29+08:00
tags = ["Leetcode", "Algorithms", "Java", "Tree", "BinarySearchTree"]
categories = ["leetcode"]
draft = false
+++

# 二叉搜索树的最近公共祖先（3种解决方式）

> [0235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)
> [二叉搜索树的最近公共祖先（3种解决方式）](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/solution/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-3c/) by [sdwwld](https://leetcode-cn.com/u/sdwwld/)


### 1，非递归解决

这题让求二叉搜索树的最近公共祖先，而二叉搜索树的特点就是```左子树的所有节点都小于当前节点，右子树的所有节点都大于当前节点，并且每棵子树都具有上述特点```，所以这题就好办了，从更节点开始遍历

- **如果两个节点值都小于根节点，说明他们都在根节点的左子树上，我们往左子树上找**
- **如果两个节点值都大于根节点，说明他们都在根节点的右子树上，我们往右子树上找**
- **如果一个节点值大于根节点，一个节点值小于根节点，说明他们他们一个在根节点的左子树上一个在根节点的右子树上，那么根节点就是他们的最近公共祖先节点。**

画个图看一下，比如要找0和5的最近公共祖先节点，如下图所示

![image.png](https://pic.leetcode-cn.com/1601168445-ZQWJxT-image.png)

```
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        //如果根节点和p,q的差相乘是正数，说明这两个差值要么都是正数要么都是负数，也就是说
        //他们肯定都位于根节点的同一侧，就继续往下找
        while ((root.val - p.val) * (root.val - q.val) > 0)
            root = p.val < root.val ? root.left : root.right;
        //如果相乘的结果是负数，说明p和q位于根节点的两侧，如果等于0，说明至少有一个就是根节点
        return root;
    }
```

看一下运行结果
![image.png](https://pic.leetcode-cn.com/1601169281-EXQktK-image.png)

<br>

### 2，递归解决

也可把它改为递归的方式
```
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        //如果小于等于0，说明p和q位于root的两侧，直接返回即可
        if ((root.val - p.val) * (root.val - q.val) <= 0)
            return root;
        //否则，p和q位于root的同一侧，就继续往下找
        return lowestCommonAncestor(p.val < root.val ? root.left : root.right, p, q);
    }
```
如果嫌代码行数太多，那就一行解决
```
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        return (root.val - p.val) * (root.val - q.val) <= 0 ? root : lowestCommonAncestor(p.val < root.val ? root.left : root.right, p, q);
    }
```



<br>

### 3，参照[236. 二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)
这题和第236题不一样的地方，在于236题不是二叉搜索树，所以第236题不能使用上面的方法解决，但这题可以使用第236题的解法，这里也可以来看下第236题之前写过的题解[236. 二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/javadai-ma-di-gui-he-fei-di-gui-tu-wen-xiang-jie-b/)

```
    public TreeNode lowestCommonAncestor(TreeNode cur, TreeNode p, TreeNode q) {
        if (cur == null || cur == p || cur == q)
            return cur;
        TreeNode left = lowestCommonAncestor(cur.left, p, q);
        TreeNode right = lowestCommonAncestor(cur.right, p, q);
        //如果left为空，说明这两个节点在cur结点的右子树上，我们只需要返回右子树查找的结果即可
        if (left == null)
            return right;
        //同上
        if (right == null)
            return left;
        //如果left和right都不为空，说明这两个节点一个在cur的左子树上一个在cur的右子树上，
        //我们只需要返回cur结点即可。
        return cur;
    }
```


<br>

**如果觉得有用就给个赞吧，你的赞是给我最大的鼓励，也是我写作的最大动力**

![image.png](https://pic.leetcode-cn.com/d56a80459005b444404d2ad6fbaabdabecd2b9ed3cb2cf432e570c315ae2fcf7-image.png)
##### 查看更多答案，可扫码关注我微信公众号“**[数据结构和算法](https://img-blog.csdnimg.cn/20200807155236311.png)**”



