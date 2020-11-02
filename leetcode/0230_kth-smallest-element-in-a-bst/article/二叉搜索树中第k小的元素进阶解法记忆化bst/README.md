"+++
title = "0230. Kth Smallest Element in a BST 二叉搜索树中第k小的元素（进阶解法）：记忆化BST "
author = ["newpp"]
date = 2020-06-28T15:39:12+08:00
tags = ["Leetcode", "Algorithms", "cpp", "BinarySearchTree"]
categories = ["leetcode"]
draft = false
+++

# 二叉搜索树中第k小的元素（进阶解法）：记忆化BST

> [0230. Kth Smallest Element in a BST](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/)
> [二叉搜索树中第k小的元素（进阶解法）：记忆化BST](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/solution/er-cha-sou-suo-shu-zhong-di-kxiao-de-yuan-su-jin-j/) by [newpp](https://leetcode-cn.com/u/newpp/)

对于解决单次kthSmallest问题，使用迭代或者递归都是非常直观的方法，而本题的进阶版本提出了一个问题：**如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？**
#### 思路
对于一棵BST，我们都很熟悉其查找/插入/删除元素的平均时间复杂度是$O(h)$（h为树的高度），那么如果我们在不断变化的树中频繁查找第k小的元素时，我们能否实现平均时间复杂度是$O(h)$的查找呢？这里我们需要引入一种新的BST结点，它除了拥有结点值和左右子结点指针等变量外，还新增了另外两个变量`leftchilds`和`rightchilds`，分别记录了该结点左子树和右子树中结点的数量。

#### 记忆化BST的维护
```
//该结点的定义
struct BstMemoNode{
    int val;
    BstNode* left;
    BstNode* right;
    int leftchinds;
    int rightchilds;
    BstMemoNode(int x):val(x),left(NULL),right(NULL),leftchilds(0),rightchilds(0){}
};
//由于本题中已经定义好了BST的结点，也可使用两个map维护
//unordered_map<BstMemoNode*,int> leftchilds;
//unordered_map<BstMemoNode*,int> rightchilds; 
```
在二叉树的插入/删除的过程中，我们都可以很方便地维护这两个变量。
```
void insertNode(BstMemoNode*& root,int key){
    if(root == NULL){
        root = new BstMemoNode(key);
        return;
    }
    if(key < root -> val){
        root -> leftchilds++;
    //  leftchilds[root]++;
        insertNode(root -> left,key);
    }
    if(key > root -> val){
        root -> rightchilds++;
    //  rightchilds[root]++;
        insertNode(root -> right,key);
    }
}

void deleteNode(BstMemoNode*& root,int key){
    if(key == root -> val){
        delete root;
        return;
    }
    if(key < root -> val){
        root -> leftchilds--;
    //  leftchilds[root]--;
        deleteNode(root -> left,key);
    }
    if(key > root -> val) {
        root -> rightchilds--;
    //  rightchilds[root]--;
        deleteNode(root-> right, key);
    }
}
```
#### 记忆化BST查找第k小的元素
假设我们已经建立好一棵记忆化的二叉搜索树（或者将一棵普通的BST记忆化后），要如何找到树中第K小的元素呢？我们可以利用结点的`leftchilds`和`rightchilds`这两个变量实现一个简单的二分查找：

1. 初始化工作指针`cur`和变量`rank`，其中`cur`一开始指向根结点，`rank`是`cur`当前指向结点的排名（即第`rank`小）。因为根结点的左子树元素都是比根结点小的，故`rank`的初始化为`rank = root -> leftchilds + 1`。

2. 当`k != rank`的时候自顶向下查找:
- 当前的`k`小于`rank`，要找的结点在左子树中，`cur`指向`cur -> left`，**因为上一个根结点和当前根结点的右子树结点都排在当前结点后面**，当前结点的排名相当于前进了`cur -> rightchilds + 1`个位置，所以新的排名：
`rank = rank - cur -> rightchilds - 1`。

- 当前的`k`大于`rank`，要找的结点在右子树中，`cur`指向`cur -> right`，**因为上一个根结点和当前根结点的左子树结点都排在当前结点前面**，当前结点的排名相当于后退了`cur -> leftchilds + 1`个位置，所以新的排名：
`rank = rank + cur -> leftchilds + 1`。

以在下面二叉树中搜索第4小的元素为例

<![幻灯片1.png](https://pic.leetcode-cn.com/7431525d97af91ea94fd6d3e94eafbdac2f6850706b19feb89b25ba49bcad02b-%E5%B9%BB%E7%81%AF%E7%89%871.png),![幻灯片2.png](https://pic.leetcode-cn.com/e82c48263f02663d74141f57427297e3e0b52cc914f997220718b628ce0347b1-%E5%B9%BB%E7%81%AF%E7%89%872.png),![幻灯片3.png](https://pic.leetcode-cn.com/95c665ec9e54b322b9f36b34563c96be127683806d810831559f42869c8d5cf9-%E5%B9%BB%E7%81%AF%E7%89%873.png),![幻灯片4.png](https://pic.leetcode-cn.com/22432a28bbfa7834f64b5ef3a6801591dba00ae70f080211cefebbbed2b1c827-%E5%B9%BB%E7%81%AF%E7%89%874.png)>

代码实现为：
```
int kthSmallest(BstMemoNode* root,int k){
    BstMemoNode* cur = root;
    int rank = root -> leftchilds + 1;
//  int rank = leftchilds[cur] + 1;
    while(k != rank){
        if(k < rank){
            cur = cur -> left;
            rank -= cur -> rightchilds + 1;
        //  rank -= rightchilds[cur] + 1;
        }else{
            cur = cur -> right;
            rank += cur -> leftchilds + 1;
        //  rank += leftchilds[cur] + 1;
        }
    }
    return cur -> val;
}
```

#### 解题代码
```
class Solution {
private:
    unordered_map<TreeNode*,int> leftchilds;
    unordered_map<TreeNode*,int> rightchilds;
public:
    int myKthSmallest(TreeNode* root,int k){
        TreeNode* cur = root;
        int rank = leftchilds[cur] + 1;
        while(k != rank){
            if(k < rank){
                cur = cur -> left;
                rank -= rightchilds[cur] + 1;
            }else{
                cur = cur -> right;
                rank += leftchilds[cur] + 1;
            }
        }
        return cur -> val;
    }
    
    int memoTree(TreeNode* root){
        if(!root){
            return 0;
        }
        leftchilds[root] = memoTree(root -> left);
        rightchilds[root] = memoTree(root -> right);
        return leftchilds[root] + rightchilds[root] + 1;
    }

    int kthSmallest(TreeNode* root, int k) {
        if(!root){
            return 0;
        }
        memoTree(root);
        return myKthSmallest(root,k);      
    }
    
};
```
- 时间复杂度：第一次需要处理整一棵树，时间复杂度为$O(n)$，后续即使经过多次插入/删除，只要维护好每个结点的leftchilds和rightchilds变量，则每次查找第k小元素的平均时间复杂度为$O(h)$。

- 空间复杂度：$O(n)$需要两个map维护各结点的左右子树结点数。
