"+++
title = "0094. Binary Tree Inorder Traversal 彻底吃透二叉树的前中后序递归法和迭代法！！ "
author = ["carlsun-2"]
date = 2020-09-14T09:22:35+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# 彻底吃透二叉树的前中后序递归法和迭代法！！

> [0094. Binary Tree Inorder Traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)
> [彻底吃透二叉树的前中后序递归法和迭代法！！](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/che-di-chi-tou-er-cha-shu-de-qian-zhong-hou-xu-d-2/) by [carlsun-2](https://leetcode-cn.com/u/carlsun-2/)

这篇文章，**彻底讲清楚应该如何写递归，并给出了前中后序三种不同的迭代法，然后分析迭代法的代码风格为什么没有统一，最后给出统一的前中后序迭代法的代码，帮大家彻底吃透二叉树的深度优先遍历。**
* 二叉树深度优先遍历
    * 前序遍历： [0144.二叉树的前序遍历](https://github.com/youngyangyang04/leetcode/blob/master/problems/0144.二叉树的前序遍历.md)
    * 后序遍历： [0145.二叉树的后序遍历](https://github.com/youngyangyang04/leetcode/blob/master/problems/0145.二叉树的后序遍历.md)
    * 中序遍历： [0094.二叉树的中序遍历](https://github.com/youngyangyang04/leetcode/blob/master/problems/0094.二叉树的中序遍历.md)
* 二叉树广度优先遍历 
    * 层序遍历：[0102.二叉树的层序遍历](https://github.com/youngyangyang04/leetcode/blob/master/problems/0102.二叉树的层序遍历.md)

这几道题目建议大家都做一下，本题解先只写二叉树深度优先遍历，二叉树广度优先遍历请看题解[0102.二叉树的层序遍历](https://github.com/youngyangyang04/leetcode/blob/master/problems/0102.二叉树的层序遍历.md)

这里想帮大家一下，明确一下二叉树的遍历规则：

![image.png](https://pic.leetcode-cn.com/0005d6f797d3281bbe2be08effd0f8fa991dc8126aef754929af34edf650626a-image.png)
以上述中，前中后序遍历顺序如下：

* 前序遍历（中左右）：5 4 1 2 6 7 8
* 中序遍历（左中右）：1 4 2 5 7 6 8
* 后序遍历（左右中）：1 2 4 7 8 6 5

# 递归法

接下来我们来好好谈一谈递归，为什么很多同学看递归算法都是“一看就会，一写就废”。主要是对递归不成体系，没有方法论，每次写递归算法 ，都是靠玄学来写代码，代码能不能编过都靠运气。

这里帮助大家确定下来递归算法的三个要素。每次写递归，都按照这三要素来写，可以保证大家写出正确的递归算法！
1. **确定递归函数的参数和返回值：**
确定哪些参数是递归的过程中需要处理的，那么就在递归函数里加上这个参数， 并且还要明确每次递归的返回值是什么进而确定递归函数的返回类型。

2. **确定终止条件：** 
写完了递归算法,  运行的时候，经常会遇到栈溢出的错误，就是没写终止条件或者终止条件写的不对，操作系统也是用一个栈的结构来保存每一层递归的信息，如果递归没有终止，操作系统的内存栈必然就会溢出。

3. **确定单层递归的逻辑：**
确定每一层递归需要处理的信息。在这里也就会重复调用自己来实现递归的过程。

好了，我们确认了递归的三要素，接下来就来练练手：

**以下以前序遍历为例：**

1. **确定递归函数的参数和返回值**：因为要打印出前序遍历节点的数值，所以参数里需要传入vector在放节点的数值，除了这一点就不需要在处理什么数据了也不需要有返回值，所以递归函数返回类型就是void，代码如下：

```
void traversal(TreeNode* cur, vector<int>& vec)
```

2. **确定终止条件**：在递归的过程中，如何算是递归结束了呢，当然是当前遍历的节点是空了，那么本层递归就要要结束了，所以如果当前遍历的这个节点是空，就直接return，代码如下： 

```
if (cur == NULL) return;
```

3. 确定单层递归的逻辑：前序遍历是中左右的循序，所以在单层递归的逻辑，是要先取中节点的数值，代码如下：

```
vec.push_back(cur->val);    // 中
traversal(cur->left, vec);  // 左
traversal(cur->right, vec); // 右
```

单层递归的逻辑就是按照中左右的顺序来处理的，这样二叉树的前序遍历，基本就写完了，在看一下完整代码：

前序遍历：

```
class Solution {
public:
    void traversal(TreeNode* cur, vector<int>& vec) {
        if (cur == NULL) return;
        vec.push_back(cur->val);    // 中
        traversal(cur->left, vec);  // 左
        traversal(cur->right, vec); // 右
    }
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> result;
        traversal(root, result);
        return result;
    }
};
```

中序遍历：

```
    void traversal(TreeNode* cur, vector<int>& vec) {
        if (cur == NULL) return;
        traversal(cur->left, vec);  // 左
        vec.push_back(cur->val);    // 中
        traversal(cur->right, vec); // 右
    }
```

后序遍历： 

```
    void traversal(TreeNode* cur, vector<int>& vec) {
        if (cur == NULL) return;
        traversal(cur->left, vec);  // 左
        traversal(cur->right, vec); // 右
        vec.push_back(cur->val);    // 中
    }
```

# 迭代法 

实践过的同学，应该会发现使用迭代法实现先中后序遍历，很难写出统一的代码，不像是递归法，实现了其中的一种遍历方式，其他两种只要稍稍改一下节点顺序就可以了。而迭代法，貌似需要每一种遍历都要写出不同风格的代码。

那么接下来我先带大家看一看其中的根本原因，其实是可以针对三种遍历方式，使用迭代法可以写出统一风格的代码。

前序遍历（迭代法）不难写出如下代码:

```
class Solution {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        stack<TreeNode*> st;
        vector<int> result;
        st.push(root);
        while (!st.empty()) {
            TreeNode* node = st.top();
            st.pop();
            if (node != NULL) result.push_back(node->val);
            else continue;
            st.push(node->right);
            st.push(node->left);
        }
        return result;
    }
};
```

这时会发现貌似使用迭代法写出先序遍历并不难，确实不难，但难却难在，我们再用迭代法写中序遍历的时候，发现套路又不一样了，目前的这个逻辑无法直接应用到中序遍历上。

## 前后中遍历迭代法不统一的写法

为了解释清楚，我说明一下 刚刚在迭代的过程中，其实我们有两个操作，**一个是处理：将元素放进result数组中，一个是访问：遍历节点。**

分析一下为什么刚刚写的前序遍历的代码，不能和中序遍历通用呢，因为前序遍历的顺序是中左右，要先访问的元素是中间节点，要处理的元素也是中间节点，所以刚刚才能写出相对简洁的代码，**因为要访问的元素和要处理的元素顺序是一致的，都是中间节点。**

那么再看看中序遍历，中序遍历是左中右，先访问的是二叉树顶部的节点，然后一层一层向下访问，直到到达树左面的最底部，再开始处理节点（也就是在把节点的数值放进result数组中），这就造成了**处理顺序和访问顺序是不一致的。**

那么在使用迭代法写中序遍历，就需要借用指针的遍历来帮助访问节点，栈则用来处理节点上的元素。

**中序遍历，可以写出如下代码：**

```
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> result;
        stack<TreeNode*> st;
        TreeNode* cur = root;
        while (cur != NULL || !st.empty()) {
            if (cur != NULL) {
                st.push(cur);
                cur = cur->left;
            } else {
                cur = st.top();
                st.pop();
                result.push_back(cur->val);
                cur = cur->right;
            }
        }
        return result;
    }
};

```

![image.png](https://pic.leetcode-cn.com/4f49ee1ccbd7b2d641a740d63a68f69146dc0bcb6dd5c0471e4289730d902352-image.png)
**所以后序遍历只需要前序遍历的代码稍作修改就可以了，代码如下：**

```
class Solution {
public:

    vector<int> postorderTraversal(TreeNode* root) {
        stack<TreeNode*> st;
        vector<int> result;
        st.push(root);
        while (!st.empty()) {
            TreeNode* node = st.top();
            st.pop();
            if (node != NULL) result.push_back(node->val);
            else continue;
            st.push(node->left); // 相对于前序遍历，这更改一下入栈顺序
            st.push(node->right);
        }
        reverse(result.begin(), result.end()); // 将结果反转之后就是左右中的顺序了
        return result;
    }
};

```

此时我们实现了前后中遍历的三种迭代法，**是不是发现迭代法实现的先中后序，其实风格也不是那么统一，除了先序和后序，有关联，中序完全就是另一个风格了，一会用栈遍历，一会又用指针来遍历。**

**重头戏来了，接下来介绍一下统一写法。**

## 前后中遍历迭代法统一的写法

我们以中序遍历为例，之前说使用栈的话，**无法同时解决处理过程和访问过程不一致的情况**，那我们就将访问的节点放入栈中，把要处理的节点也放入栈中但是要做标记，标记就是要处理的节点放入栈之后，紧接着放入一个空指针作为标记。

中序遍历代码如下：
```
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> result;
        stack<TreeNode*> st;
        if (root != NULL) st.push(root);
        while (!st.empty()) {
            TreeNode* node = st.top();
            if (node != NULL) {
                st.pop(); // 将该节点弹出，避免重复操作，下面再将右中左节点添加到栈中
                if (node->right) st.push(node->right); // 添加右节点

                st.push(node); // 添加中节点
                st.push(NULL); // 中节点访问过，但是还没有处理，需要做一下标记。

                if (node->left) st.push(node->left); // 添加左节点
            } else {
                st.pop(); // 将空节点弹出
                node = st.top(); // 重新取出栈中元素
                st.pop();
                result.push_back(node->val); // 加入到数组中
            }
        }
        return result;
    }
};
```

看代码有点抽象我们来看一下动画：

![中序遍历迭代（统一写法）.mp4](62c6ae4a-9046-428a-ac1f-8a66be274401)
前序遍历代码如下： 

```
class Solution {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> result;
        stack<TreeNode*> st;
        if (root != NULL) st.push(root);
        while (!st.empty()) {
            TreeNode* node = st.top();
            if (node != NULL) {
                st.pop();
                if (node->right) st.push(node->right);  // 右
                if (node->left) st.push(node->left);    // 左
                st.push(node);                          // 中
                st.push(NULL);
            } else {
                st.pop();
                node = st.top();
                st.pop();
                result.push_back(node->val);
            }
        }
        return result;
    }
};
```

后续遍历代码如下：

```
class Solution {
public:
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> result;
        stack<TreeNode*> st;
        if (root != NULL) st.push(root);
        while (!st.empty()) {
            TreeNode* node = st.top();
            if (node != NULL) {
                st.pop();
                st.push(node);                          // 中
                st.push(NULL);

                if (node->right) st.push(node->right);  // 右
                if (node->left) st.push(node->left);    // 左

            } else {
                st.pop();
                node = st.top();
                st.pop();
                result.push_back(node->val);
            }
        }
        return result;
    }
};
```

**我是程序员Carl，先后在腾讯和百度从事技术研发多年，利用工作之余重刷leetcode，更多[精彩算法文章](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzUxNjY5NTYxNA==&action=getalbum&album_id=1485825793120387074&scene=173#wechat_redirect)尽在公众号：[代码随想录](https://img-blog.csdnimg.cn/20200815195519696.png)，关注后，回复「Java」「C++」「python」「简历模板」等等，有我整理多年的学习资料，可以加我[微信](https://img-blog.csdnimg.cn/20200814140330894.png)，备注「简单自我介绍」+「组队刷题」，拉你进入刷题群（无任何广告，纯个人分享），每天一道经典题目分析，我选的每一道题目都不是孤立的，而是由浅入深一脉相承的，如果跟住节奏每篇连续着看，定会融会贯通。本文  [https://github.com/youngyangyang04/leetcode-master](https://github.com/youngyangyang04/leetcode-master ) 已经收录，里面还有leetcode刷题攻略、各个类型经典题目刷题顺序、思维导图，看一看一定会有所收获，如果对你有帮助也给一个star支持一下吧！**
**以下资料希望对你有帮助：**

* [我在B站上讲KMP算法！](https://www.bilibili.com/video/BV1PD4y1o7nd/)
* [leetcode刷题攻略](https://github.com/youngyangyang04/leetcode-master)
* [程序员应该如何写简历（附简历模板）](https://mp.weixin.qq.com/s/PkBpde0PV65dJjj9zZJYtg)
* [一线互联网公司技术面试的流程以及注意事项](https://mp.weixin.qq.com/s/1VMvQ_6HbVpEn85CNilTiw)
* [如何学习C++？ B站](https://www.bilibili.com/video/BV1rK4y1e7ed)
* [手把手带你读C++ primer！B站](https://www.bilibili.com/video/BV1Z5411874t)
* [C++面试&C++学习指南知识点整理](https://github.com/youngyangyang04/TechCPP)

**如果感觉题解对你有帮助，不要吝啬给一个👍吧！**
