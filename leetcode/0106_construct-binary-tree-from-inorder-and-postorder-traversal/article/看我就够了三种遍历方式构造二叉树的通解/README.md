"+++
title = "0106. Construct Binary Tree from Inorder and Postorder Traversal “看我就够了”三种遍历方式构造二叉树的通解 "
author = ["christmas_wang"]
date = 2020-06-16T08:41:43+08:00
tags = ["Leetcode", "Algorithms", "C++", "二叉树", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# “看我就够了”三种遍历方式构造二叉树的通解

> [0106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
> [“看我就够了”三种遍历方式构造二叉树的通解](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/solution/kan-wo-jiu-gou-liao-san-chong-bian-li-fang-shi-g-2/) by [christmas_wang](https://leetcode-cn.com/u/christmas_wang/)

### 写在前面
【先赞后看，养成习惯】
不知道是否有小伙伴和我一样，做了这三道构造二叉树的题目后半懵半懂。不过在看了大佬的解答过程后，我有了一些感受，可以构造出解这三类题目的通用解法。

### 代码相关
这里我先放上我解题的代码，从中间我们可以发现相似之处,在代码中我们可以看到有两个地方不同，这两个地方也是我们解题的关键
#### 不同之处一 寻找当前根节点
这一部分总的来说是在寻找可以将左右子树划分开的根节点
- 前+后
 首先我们可以显然知道当前根节点为`pre[pre_start]`,并且它在后序中的位置为`post_end`，因此这里我们需要找到能区分左右子树的节点。
 我们知道左子树的根节点为`pre[pre_start+1]`，因此只要找到它在后序中的位置就可以分开左右子树（index的含义）
- 前+中
 首先我们可以显然知道当前根节点为`pre[pre_start]`,只用找出它在中序中的位置，就可以把左右子树分开（index的含义）
- 中+后
 首先我们可以显然知道当前根节点为`post[post_end]`，只用找出它在中序中的位置，就可以把左右子树分开（index的含义）


#### 不同之处二 左右遍历范围
这一部分运用了一个技巧是 “两种遍历中，同一子树的节点数目是相同的”
需要说明的是在"前+后","前+中"我们都运用到了“右子树起始位置为左子树终止位置+1”，其实这个也可以运用上面这个技巧来计算出起始位置
- 前+后
  后序遍历中，我们知道 左子树：`[post_start,index]`, 右子树：`[index+1, post_end-1]`
  在前序遍历中，左子树起始位置为`pre_start+1`,左子树个数一共有(index - post_start)个，因此左子树：`[pre_start+1, pre_start+1 + (index - post_start)]`
  右子树起始位置为左子树终止位置+1，终止位置为`pre_end`，因此右子树：`[ pre_start+1 + (index - post_start) + 1, pre_end]`
- 前+中
  中序遍历中，我们知道 左子树：`[inorder_start,index-1]`, 右子树：`[index+1, inorder_end]`
  在前序遍历中，左子树起始位置为`pre_start+1`,左子树一共有(index-1 - inorder_start)个，因此左子树：`[pre_start+1, pre_start+1 + (index-1 - inorder_start)]`
  右子树起始位置为左子树终止位置+1，终止位置为`pre_end`，因此右子树：`[ pre_start+1 + (index-1 - inorder_start) + 1, pre_end]`
- 中+后
  中序遍历中，我们知道 左子树：`[inorder_start,index-1]`, 右子树：`[index+1, inorder_end]`
  在后序遍历中，左子树起始位置为`post_start`，左子树一共有(index-1 - inorder_start)个，因此左子树：`[post_start, post_start + (index-1 - inorder_start)]`
  右子树的终止位置为`post_end - 1`,右子树一共有(inorder_end - (index+1))个,因此右子树:`[post_end - 1 - (inorder_end - (index+1)), post_end - 1]`

<![幻灯片1.PNG](https://pic.leetcode-cn.com/3fdcb62791ffff37a1f04511ef386da4e3aaf1d7417c48d0861d5c5903070ed9-%E5%B9%BB%E7%81%AF%E7%89%871.PNG),![幻灯片2.PNG](https://pic.leetcode-cn.com/da0579653aa42c32b824a0cbabefb4f3a3ed8bf7a761f1d512467649d959fa2a-%E5%B9%BB%E7%81%AF%E7%89%872.PNG),![幻灯片3.PNG](https://pic.leetcode-cn.com/30c31a015a3aa473fbede48ec6044afec542cc5443ba31cb0eb81769a8245088-%E5%B9%BB%E7%81%AF%E7%89%873.PNG),![幻灯片4.PNG](https://pic.leetcode-cn.com/2af484d91fc681d752b70dd20d30f63fb0647e3f55eb1dfd55c49eb2f559f82c-%E5%B9%BB%E7%81%AF%E7%89%874.PNG)>
### 相关题目
[LeetCode889 根据前序和后序遍历构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/)
[LeetCode105 从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)
[LeetCode106 从中序与后序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
### 我的题解
[LeetCode1262 可被三整除的最大和](https://leetcode-cn.com/problems/greatest-sum-divisible-by-three/solution/dong-tai-gui-hua-yu-zhuang-tai-zhuan-yi-by-christm/)
[LeetCode688 “马”在棋盘上的概率](https://leetcode-cn.com/problems/knight-probability-in-chessboard/solution/zhuang-tai-ji-de-zai-ci-ying-yong-by-christmas_wan/)
[LeetCode967 连续差相同的数字](https://leetcode-cn.com/problems/numbers-with-same-consecutive-differences/solution/cun-chu-kong-jian-ke-bian-de-dpshu-zu-by-christmas/)
[LeetCode873 最长的斐波那契子序列的长度](https://leetcode-cn.com/problems/length-of-longest-fibonacci-subsequence/solution/zhuang-tai-ding-yi-hen-shi-zhong-yao-by-christmas_/)
[LeetCode1218 最长定差子序列](https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference/solution/yi-dao-jian-dan-de-dong-tai-gui-hua-de-you-hua-wen/)
[LeetCode523 连续子数组和](https://leetcode-cn.com/problems/continuous-subarray-sum/solution/qian-zhui-he-yu-intmapde-zai-ci-ying-yong-by-chris/)
[LeetCode576 出界的路径数](https://leetcode-cn.com/problems/out-of-boundary-paths/solution/zhuang-tai-ji-du-shi-zhuang-tai-ji-by-christmas_wa/)
[LeetCode1220 统计元音字母序列的数目](https://leetcode-cn.com/problems/count-vowels-permutation/solution/dang-wo-men-zai-tan-dong-tai-gui-hua-de-shi-hou-wo/)