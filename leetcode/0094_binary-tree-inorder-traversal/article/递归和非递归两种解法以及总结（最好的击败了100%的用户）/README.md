"+++
title = "0094. Binary Tree Inorder Traversal 递归和非递归两种解法以及总结（最好的击败了100%的用户） "
author = ["sdwwld"]
date = 2020-09-14T02:04:52+08:00
tags = ["Leetcode", "Algorithms", "Java", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 递归和非递归两种解法以及总结（最好的击败了100%的用户）

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [递归和非递归两种解法以及总结（最好的击败了100%的用户）](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/di-gui-he-fei-di-gui-liang-chong-jie-fa-by-sdwwld/) by [sdwwld](https://leetcode-cn.com/u/sdwwld/)


### 1，递归解法
二叉树的前序，中序和后续遍历，使用递归方式是最简单的。中序遍历的顺序是```左子树→根节点→右子树```，就像下面图中这样
![image.png](https://pic.leetcode-cn.com/1600045856-CrdyQS-image.png)
代码也比较简单
```
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> list = new ArrayList();
        helper(list, root);
        return list;
    }

    public void helper(List<Integer> list, TreeNode root) {
        //终止条件
        if (root == null)
            return;
        //遍历当前节点的左子树
        helper(list, root.left);
        //把当前节点加入到集合中
        list.add(root.val);
        //遍历当前节点的右子树
        helper(list, root.right);
    }
```
看下运行结果，可以看到运行时间很快，但内存消耗很差
![image.png](https://pic.leetcode-cn.com/1600045939-VnKbSN-image.png)

上面这种递归比较容易理解，我们还可以换种形式，先递归找到左子树的所有节点把他们加入到集合中，再把根节点加入到集合中，最后再递归找到右子树的所有节点并把他们加入到集合中
```
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        //边界条件判断，如果为空就返回一个空的集合
        if (root == null)
            return res;
        //通过递归把左子树的结点加入到集合中
        res.addAll(inorderTraversal(root.left));
        //把根节点的值加入到集合中
        res.add(root.val);
        //通过递归把右子树的结点加入到集合中
        res.addAll(inorderTraversal(root.right));
        return res;
    }
```
看一下运行结果
![image.png](https://pic.leetcode-cn.com/1600057004-TiAcrI-image.png)



### 2，非递归解法
非递归解法比较绕，就是从根节点开始，沿着他的左子节点一直走下去，一直到某一个节点没有左子节点，然后把这个节点加入到集合中，接着访问他的右子节点，这个时候``这个右子节点我们可以把它按照访问根节点的方式进行访问```，也是一种访问他的左子节点……，来画个简单的图看下
![image.png](https://pic.leetcode-cn.com/1600049007-pWtItP-image.png)
![image.png](https://pic.leetcode-cn.com/1600049020-DvBpmb-image.png)
![image.png](https://pic.leetcode-cn.com/1600049029-BIYRxy-image.png)

```
    public List<Integer> inorderTraversal(TreeNode root) {
        //list存储结果的
        List<Integer> list = new ArrayList<>();
        //栈存储结点的
        Stack<TreeNode> stack = new Stack<>();
        while (root != null || !stack.empty()) {
            //找当前节点的左子节点，一直找下去，直到为空为止
            while (root != null) {
                stack.push(root);
                root = root.left;
            }
            //出栈，这时候root就是最左子节点
            root = stack.pop();
            //然后把最左子节点加入到集合中
            list.add(root.val);
            //最后再访问他的右子节点
            root = root.right;
        }
        return list;
    }
```
看下运行结果，明显不如递归的好
![image.png](https://pic.leetcode-cn.com/1600049088-SyeuYz-image.png)


## 3，总结
有的可能对二叉树的前中后3种遍历方式容易搞混。**二叉树的前序，中序和后续都是根据根节点来判断的**，先遍历根节点就是先序遍历，最后遍历根节点就是后续遍历，那么根节点在中间的时候遍历就是中序遍历了。至于左右子树哪个先遍历，记住一点，这3种方式无论哪种，都是先遍历左子树在遍历右子树，也就是说在同一节点，右子树永远都不可能比左子树先遍历

**前序遍历：根节点→左子树→右子树（根节点在前）**
**中序遍历：左子树→根节点→右子树（根节点在中间）**
**后序遍历：左子树→右子树→根节点（根节点在后）**

**1）前序遍历递归代码**
```
    public static void preOrder(TreeNode tree) {
        if (tree == null)
            return;
        System.out.printf(tree.val + "");
        preOrder(tree.left);
        preOrder(tree.right);
    }
```
**2）前序遍历非递归代码**
```
    public static void preOrder(TreeNode tree) {
        if (tree == null)
            return;
        Stack<TreeNode> q1 = new Stack<>();
        q1.push(tree);//压栈
        while (!q1.empty()) {
            TreeNode t1 = q1.pop();//出栈
            System.out.println(t1.val);
            if (t1.right != null) {
                q1.push(t1.right);
            }
            if (t1.left != null) {
                q1.push(t1.left);
            }
        }
    }
```

**3）中序遍历递归代码**
```
    public static void inOrderTraversal(TreeNode node) {
        if (node == null)
            return;
        inOrderTraversal(node.left);
        System.out.println(node.val);
        inOrderTraversal(node.right);
    }
```

**4）中序遍历非递归代码**
```
    public void inOrderTraversal(TreeNode tree) {
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

**5）后序遍历递归代码**
```
    public void postOrder(TreeNode tree) {
        if (tree == null)
            return;
        postOrder(tree.left);
        postOrder(tree.right);
        System.out.println(tree.val);
    }
```

**6）后序遍历非递归代码**
```
    public static void postOrder(TreeNode tree) {
        if (tree == null)
            return;
        Stack<TreeNode> s1 = new Stack<>();
        Stack<TreeNode> s2 = new Stack<>();
        s1.push(tree);
        while (!s1.isEmpty()) {
            tree = s1.pop();
            s2.push(tree);
            if (tree.left != null) {
                s1.push(tree.left);
            }
            if (tree.right != null) {
                s1.push(tree.right);
            }
        }
        while (!s2.isEmpty()) {
            System.out.print(s2.pop().val + " ");
        }
    }
```

**7）BFS递归代码**
```
    public void levelOrder(TreeNode tree) {
        if (tree == null)
            return;
        LinkedList<TreeNode> list = new LinkedList<>();//链表，这里我们可以把它看做队列
        list.add(tree);//相当于把数据加入到队列尾部
        while (!list.isEmpty()) {
            TreeNode node = list.poll();//poll方法相当于移除队列头部的元素
            System.out.println(node.val);
            if (node.left != null)
                list.add(node.left);
            if (node.right != null)
                list.add(node.right);
        }
    }
```

**8）BFS非递归代码**
```
    public void levelOrder(TreeNode tree) {
        int depth = depth(tree);
        for (int level = 0; level < depth; level++) {
            printLevel(tree, level);
        }
    }

    private static int depth(TreeNode tree) {
        if (tree == null)
            return 0;
        int leftDepth = depth(tree.left);
        int rightDepth = depth(tree.right);
        return Math.max(leftDepth, rightDepth) + 1;
    }


    private static void printLevel(TreeNode tree, int level) {
        if (tree == null)
            return;
        if (level == 0) {
            System.out.print(" " + tree.val);
        } else {
            printLevel(tree.left, level - 1);
            printLevel(tree.right, level - 1);
        }
    }
```

**9）DFS参照前序遍历**

### 4，定义:
**结点的度**：一个结点含有的子结点的个数称为该结点的度；

**叶结点或终端结点**：度为0的结点称为叶结点；

**非终端结点或分支结点**：度不为0的结点；

**双亲结点或父结点**：若一个结点含有子结点，则这个结点称为其子结点的父结点；

**孩子结点或子结点**：一个结点含有的子树的根结点称为该结点的子结点；

**兄弟结点**：具有相同父结点的结点互称为兄弟结点；

**树的度**：一棵树中，最大的结点的度称为树的度；

**结点的层次**：从根开始定义起，根为第1层，根的子结点为第2层，以此类推；

**树的高度或深度**：树中结点的最大层次；

**堂兄弟结点**：双亲在同一层的结点互为堂兄弟；

**结点的祖先**：从根到该结点所经分支上的所有结点；

**子孙**：以某结点为根的子树中任一结点都称为该结点的子孙。

**森林**：由m（m>=0）棵互不相交的树的集合称为森林；

**无序树**：树中任意节点的子结点之间没有顺序关系，这种树称为无序树,也称为自由树;

**有序树**：树中任意节点的子结点之间有顺序关系，这种树称为有序树；

**二叉树**：每个节点最多含有两个子树的树称为二叉树；

**完全二叉树**：若设二叉树的深度为h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树

**满二叉树**：除最后一层无任何子节点外，每一层上的所有结点都有两个子结点的二叉树。

**哈夫曼树**：带权路径最短的二叉树称为哈夫曼树或最优二叉树；

<br/>

**如果觉得有用就给个赞吧，你的赞是给我最大的鼓励，也是我写作的最大动力**

![image.png](https://pic.leetcode-cn.com/d56a80459005b444404d2ad6fbaabdabecd2b9ed3cb2cf432e570c315ae2fcf7-image.png)
##### 查看更多答案，可扫码关注我微信公众号“**[数据结构和算法](https://img-blog.csdnimg.cn/20200807155236311.png)**”