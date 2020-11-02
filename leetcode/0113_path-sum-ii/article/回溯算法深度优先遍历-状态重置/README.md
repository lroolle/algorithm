"+++
title = "0113. Path Sum II 回溯算法（深度优先遍历 + 状态重置） "
author = ["liweiwei1419"]
date = 2020-04-02T06:56:13+08:00
tags = ["Leetcode", "Algorithms", "Backtracking", "DepthfirstSearch"]
categories = ["leetcode"]
draft = false
+++

# 回溯算法（深度优先遍历 + 状态重置）

> [0113. Path Sum II](https://leetcode-cn.com/problems/path-sum-ii/)
> [回溯算法（深度优先遍历 + 状态重置）](https://leetcode-cn.com/problems/path-sum-ii/solution/hui-su-suan-fa-shen-du-you-xian-bian-li-zhuang-tai/) by [liweiwei1419](https://leetcode-cn.com/u/liweiwei1419/)
这道题写法比较多，但是核心思想还**是「深度优先遍历」（回溯算法）**，以考虑清楚「回溯算法」是「深度优先遍历」的一些细节问题。

读者可以思考：

+ 在 `res.add(new ArrayList<>(path));` 这段代码为什么要套一层 `new ArrayList<>(path)`，不这么做是否可以；
+ 「深度优先遍历」不同的「状态」之间差别很小，因此「状态重置」很方便。「广度优先遍历」就没有这种特点，因此「深度优先遍历」可以成为强大的「回溯搜索」算法；
+ 这里的「状态」是指完成一件事情进行到哪一个阶段，在下面的代码中：`path` 、`sum` 都是状态变量。

**参考代码 1**：（推荐）

注意：不用纠结一些写法特别奇怪的地方，写对其中一个版本就可以了。

```Java
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Deque;
import java.util.List;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }
}
public class Solution {

    public List<List<Integer>> pathSum(TreeNode root, int sum) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) {
            return res;
        }

        // Java 文档中 Stack 类建议使用 Deque 代替 Stack，注意：只使用栈的相关接口
        Deque<Integer> path = new ArrayDeque<>();
        pathSum(root, sum, path, res);
        return res;
    }

    public void pathSum(TreeNode node, int sum, Deque<Integer> path, List<List<Integer>> res) {
        // 打开注释调试
        // System.out.println(path);

        // 递归终止条件
        if (node == null) {
            return;
        }

        // 沿途结点必须选择，这个时候要做两件事：1、sum 减去这个结点的值；2、添加到 path 里
        sum -= node.val;
        path.addLast(node.val);

        if (sum == 0 && node.left == null && node.right == null) {
            // path 全局只有一份，必须做拷贝
            res.add(new ArrayList<>(path));
            // 注意：这里 return 之前必须重置
            path.removeLast();
            return;
        }

        pathSum(node.left, sum, path, res);
        pathSum(node.right, sum, path, res);
        // 递归完成以后，必须重置变量
        path.removeLast();
    }

    public static void main(String[] args) {
        TreeNode treeNode5 = new TreeNode(5);
        TreeNode treeNode4 = new TreeNode(4);
        TreeNode treeNode8 = new TreeNode(8);
        TreeNode treeNode11 = new TreeNode(11);
        TreeNode treeNode13 = new TreeNode(13);
        TreeNode treeNode4_ = new TreeNode(4);
        TreeNode treeNode7 = new TreeNode(7);
        TreeNode treeNode2 = new TreeNode(2);
        TreeNode treeNode5_ = new TreeNode(5);
        TreeNode treeNode1 = new TreeNode(1);

        treeNode5.left = treeNode4;
        treeNode5.right = treeNode8;
        treeNode4.left = treeNode11;
        treeNode11.left = treeNode7;
        treeNode11.right = treeNode2;
        treeNode8.left = treeNode13;
        treeNode8.right = treeNode4_;
        treeNode4_.left = treeNode5_;
        treeNode4_.right = treeNode1;

        Solution solution = new Solution();
        List<List<Integer>> res = solution.pathSum(treeNode5, 22);
        System.out.println(res);
    }
}
```
```Java
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Deque;
import java.util.List;

public class Solution {

    public List<List<Integer>> dfs(TreeNode root, int sum) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) {
            return res;
        }

        // Java 文档中 Stack 类建议使用 Deque 代替 Stack，注意：只使用栈的相关接口
        Deque<Integer> path = new ArrayDeque<>();
        dfs(root, sum, path, res);
        return res;
    }

    private void dfs(TreeNode node, int sum, Deque<Integer> path, List<List<Integer>> res) {
        if (node == null) {
            return;
        }
        if (node.val == sum && node.left == null && node.right == null) {
            path.addLast(node.val);
            res.add(new ArrayList<>(path));
            path.removeLast();
            return;
        }

        path.addLast(node.val);
        dfs(node.left, sum - node.val, path, res);
        dfs(node.right, sum - node.val, path, res);
        path.removeLast();
    }
}
```

第 1 个选项卡代码里 `path` 变量打印输出的结果是：

```
[]
[5]
[5, 4]
[5, 4, 11]
[5, 4, 11, 7]
[5, 4, 11, 7]
[5, 4, 11]
[5, 4]
[5]
[5, 8]
[5, 8, 13]
[5, 8, 13]
[5, 8]
[5, 8, 4]
[5, 8, 4]
[5, 8, 4, 1]
[5, 8, 4, 1]
[[5, 4, 11, 2], [5, 8, 4, 5]]
```
下面再提供一种写法供大家比对，思想是一样的，只是细节不同。

```Java
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Deque;
import java.util.List;

public class Solution {

    public List<List<Integer>> pathSum(TreeNode root, int sum) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) {
            return res;
        }

        Deque<Integer> path = new ArrayDeque<>();
        dfs(root, sum, path, res);
        return res;
    }

    public void dfs(TreeNode node, int sum, Deque<Integer> path, List<List<Integer>> res) {
        if (node == null) {
            return;
        }

        sum -= node.val;
        path.addLast(node.val);

        if (node.left == null && node.right == null && sum == 0) {
            res.add(new ArrayList<>(path));
            return;
        }

        if (node.left != null) {
            dfs(node.left, sum, path, res);
            path.removeLast();
        }

        if (node.right != null) {
            dfs(node.right, sum, path, res);
            path.removeLast();
        }
    }
}
```
测试用例就使用示例：

![image.png](https://pic.leetcode-cn.com/a46657b13c035c0d4873e818c86cd50fad29e1a86b3fd2c4236e15a88a445b79-image.png)
**参考代码 2**：（供比对，不建议这么写）

注意：参考代码 2 的写法只能叫「深度优先遍历」，每一次向下传递，都会复制，复杂度很高，不能叫做「回溯算法」。

**「回溯算法」是全程只使用一个变量去搜索所有可能的情况，在符合条件的时候才做复制和保存**。
```Java
import java.util.ArrayList;
import java.util.List;
public class Solution {

    public List<List<Integer>> pathSum(TreeNode root, int sum) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) {
            return res;
        }

        // 从根结点到叶子结点的路径
        List<Integer> path = new ArrayList<>();
        dfs(root, sum, path, res);
        return res;
    }

    private void dfs(TreeNode root, int sum, List<Integer> path, List<List<Integer>> res) {
        if (root == null) {
            return;
        }

        path.add(root.val);
        sum -= root.val;

        if (root.left == null && root.right == null) {
            if (sum == 0) {
                // 正是因为每一次向下传递的过程中复制整个列表，在叶子结点出直接添加即可
                res.add(path);
                return;
            }
        }

        // 基本数据类型在方法传递过程中的行为是是复制
        // new ArrayList<>() 每一次向下传递的过程中复制整个列表，低效
        dfs(root.left, sum, new ArrayList<>(path), res);
        dfs(root.right, sum, new ArrayList<>(path), res);
        
        // 在递归结束以后无需「状态重置」
    }
}
```

