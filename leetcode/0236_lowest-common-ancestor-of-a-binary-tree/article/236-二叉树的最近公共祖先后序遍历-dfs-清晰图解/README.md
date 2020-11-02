"+++
title = "0236. Lowest Common Ancestor of a Binary Tree 236. 二叉树的最近公共祖先（后序遍历 DFS ，清晰图解） "
author = ["jyd"]
date = 2020-05-09T16:12:30+08:00
tags = ["Leetcode", "Algorithms", "Python3", "Java", "二叉树", "DepthfirstSearch", "Recursion", "Python"]
categories = ["leetcode"]
draft = false
+++

# 236. 二叉树的最近公共祖先（后序遍历 DFS ，清晰图解）

> [0236. Lowest Common Ancestor of a Binary Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)
> [236. 二叉树的最近公共祖先（后序遍历 DFS ，清晰图解）](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/236-er-cha-shu-de-zui-jin-gong-gong-zu-xian-hou-xu/) by [jyd](https://leetcode-cn.com/u/jyd/)

#### 解题思路：

**祖先的定义：** 若节点 $p$ 在节点 $root$ 的左（右）子树中，或 $p = root$ ，则称 $root$ 是 $p$ 的祖先。

![Picture1.png](https://pic.leetcode-cn.com/83402bb4c1bba2746effc5607d9654aeb9c3496b4a846d41ce61adb5af02c0f5-Picture1.png){:width=400}

**最近公共祖先的定义：** 设节点 $root$ 为节点 $p, q$ 的某公共祖先，若其左子节点 $root.left$ 和右子节点 $root.right$ 都不是 $p,q$ 的公共祖先，则称 $root$ 是 “最近的公共祖先” 。

根据以上定义，若 $root$ 是 $p, q$ 的 **最近公共祖先** ，则只可能为以下情况之一：

1. $p$ 和 $q$ 在 $root$ 的子树中，且分列 $root$ 的 **异侧**（即分别在左、右子树中）；
2. $p = root$ ，且 $q$ 在 $root$ 的左或右子树中；  
3. $q = root$ ，且 $p$ 在 $root$ 的左或右子树中；  

![Picture2.png](https://pic.leetcode-cn.com/e48705d412500d43fa81c1d8fdd107bb2d0c7dfa12bdc588cd88f481b4b9f7d8-Picture2.png){:width=400}

考虑通过递归对二叉树进行后序遍历，当遇到节点 $p$ 或 $q$ 时返回。从底至顶回溯，当节点 $p, q$ 在节点 $root$ 的异侧时，节点 $root$ 即为最近公共祖先，则向上返回 $root$ 。

##### 递归解析：

1. **终止条件：**
   1. 当越过叶节点，则直接返回 $null$ ；
   2. 当 $root$ 等于 $p, q$ ，则直接返回 $root$ ；
2. **递推工作：**
   1. 开启递归左子节点，返回值记为 $left$ ；
   2. 开启递归右子节点，返回值记为 $right$ ；
3. **返回值：** 根据 $left$ 和 $right$ ，可展开为四种情况；
   1. 当 $left$ 和 $right$ **同时为空** ：说明 $root$ 的左 / 右子树中都不包含 $p,q$ ，返回 $null$ ；
   2. 当 $left$ 和 $right$ **同时不为空** ：说明 $p, q$ 分列在 $root$ 的 **异侧** （分别在 左 / 右子树），因此 $root$ 为最近公共祖先，返回 $root$ ；
   3. 当 $left$ **为空** ，$right$ **不为空** ：$p,q$ 都不在 $root$ 的左子树中，直接返回 $right$ 。具体可分为两种情况：
      1. $p,q$ 其中一个在 $root$ 的 **右子树** 中，此时 $right$ 指向 $p$（假设为 $p$ ）； 
      2. $p,q$ 两节点都在 $root$ 的 **右子树** 中，此时的 $right$ 指向 **最近公共祖先节点** ；
   4. 当 $left$ **不为空** ， $right$ **为空** ：与情况 `3.` 同理；

> 观察发现， 情况 `1.` 可合并至 `3.` 和 `4.` 内，详见文章末尾代码。

<![Picture1.png](https://pic.leetcode-cn.com/c44f8946548954a2513f7d72e20be260c36c157b506749c788afce1e7bd3416c-Picture1.png),![Picture2.png](https://pic.leetcode-cn.com/55f7683ceee27def129a50c9a26305e56b25175dbd3da55983b5848145559354-Picture2.png),![Picture3.png](https://pic.leetcode-cn.com/0c3217c102953090030aec857ef2e1e96672b38c450a90356a3e64f0dfc97af2-Picture3.png),![Picture4.png](https://pic.leetcode-cn.com/f137a75004bf105ae2f9d2987d3d75c0d0cbddfda54126e549b5a3a99b06a6ef-Picture4.png),![Picture5.png](https://pic.leetcode-cn.com/3334d8bc74cad490584a03ca6e6637f4d431626f75ca589710d6a382fe9ab06b-Picture5.png),![Picture6.png](https://pic.leetcode-cn.com/fd6ef030cf8acac250792828c04df471ccab669d4153b49b934bc4cc3517efcf-Picture6.png),![Picture7.png](https://pic.leetcode-cn.com/e03f2505635e77816e12bdfd2ce5c1c4ace3d2dfa2a0e10eefe683e11e88c98b-Picture7.png),![Picture8.png](https://pic.leetcode-cn.com/a9cf21e0a271c74af5ab00e39da09d485de8a3dabbfaa6d6cd2a2a1c7f60d2a8-Picture8.png),![Picture9.png](https://pic.leetcode-cn.com/6540e7106efa4461cb19c21a682e9b7c9bd33367d6c5a8bd97982cc7bcec9ec3-Picture9.png),![Picture10.png](https://pic.leetcode-cn.com/d249c4379aee12e4a5bce4f20c4ad8b709ff35e358cfdb3255ed6d9c6dc4ae2c-Picture10.png),![Picture11.png](https://pic.leetcode-cn.com/fa87c46c8e8360cf3ab5ea852161ab5f9e4a2ca1b5ff6cd9e9ba0897dcbd455a-Picture11.png),![Picture12.png](https://pic.leetcode-cn.com/91287818231c969dcc8c69ac9c79197a9b29085d120b18b5230dda77d092f6d6-Picture12.png),![Picture13.png](https://pic.leetcode-cn.com/bf71136a0329cf5d48933cb7dc1d8c1bd0cec96c1ad13bcca97cea2f58d14fb5-Picture13.png),![Picture14.png](https://pic.leetcode-cn.com/68aced35d03027033c2552b35d477d700e38e7c01f8c1fa76b4cf8b0a1858d30-Picture14.png),![Picture15.png](https://pic.leetcode-cn.com/ddf6279a32122924d3608ad7fcdf3f518091da353a1961c0e0e1afaf51d09623-Picture15.png),![Picture16.png](https://pic.leetcode-cn.com/b4777e4a6ff72ed49356e20a0a897fd866bb3e4dfdb5e0c8fc1dc0f918029237-Picture16.png),![Picture17.png](https://pic.leetcode-cn.com/df510a1fe4750116a935e61ef63ad30a5092bfefe38845497ff9431b3656a793-Picture17.png),![Picture18.png](https://pic.leetcode-cn.com/0724b87055c4bc4d744ab64775e6eefa348777c0ea0b07a00ff917773f4b494e-Picture18.png)>

##### 复杂度分析：

- **时间复杂度 $O(N)$ ：** 其中 $N$ 为二叉树节点数；最差情况下，需要递归遍历树的所有节点。
- **空间复杂度 $O(N)$ ：** 最差情况下，递归深度达到 $N$ ，系统使用 $O(N)$ 大小的额外空间。

##### 代码：

```python
class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if not root or root == p or root == q: return root
        left = self.lowestCommonAncestor(root.left, p, q)
        right = self.lowestCommonAncestor(root.right, p, q)
        if not left: return right
        if not right: return left
        return root
```

```java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if(root == null || root == p || root == q) return root;
        TreeNode left = lowestCommonAncestor(root.left, p, q);
        TreeNode right = lowestCommonAncestor(root.right, p, q);
        if(left == null) return right;
        if(right == null) return left;
        return root;
    }
}
```

情况 `1.` , `2.` , `3.` , `4.` 的展开写法如下。

```python
class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        if not root or root == p or root == q: return root
        left = self.lowestCommonAncestor(root.left, p, q)
        right = self.lowestCommonAncestor(root.right, p, q)
        if not left and not right: return # 1.
        if not left: return right # 3.
        if not right: return left # 4.
        return root # 2. if left and right:
```

```java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if(root == null || root == p || root == q) return root;
        TreeNode left = lowestCommonAncestor(root.left, p, q);
        TreeNode right = lowestCommonAncestor(root.right, p, q);
        if(left == null && right == null) return null; // 1.
        if(left == null) return right; // 3.
        if(right == null) return left; // 4.
        return root; // 2. if(left != null and right != null)
    }
}
```