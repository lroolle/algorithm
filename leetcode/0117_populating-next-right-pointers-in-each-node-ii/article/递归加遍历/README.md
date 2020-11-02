"+++
title = "0117. Populating Next Right Pointers in Each Node II 递归加遍历 "
author = ["gin-34"]
date = 2020-08-04T16:34:05+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 递归加遍历

> [0117. Populating Next Right Pointers in Each Node II](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/)
> [递归加遍历](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/solution/di-gui-jia-bian-li-by-gin-34/) by [gin-34](https://leetcode-cn.com/u/gin-34/)

![image.png](https://pic.leetcode-cn.com/f088d48ca130267e335727cf9a88cde4ce00b1d22dbccc296cbfa3840423dfe2-image.png)
和上一题一样，左子树的下一个节点是右子节点，但是需要考虑右子节点为空的情况；
1. 先说不为空的情况，左子节点直接指向右子节点即可，但是需要先把右子节点的next指针链接上，不然之后会存在最右节点没法连接的case；
右子节点需要遍历根节点的next指针，直到next为空或者next的左子节点或者右子节点不为空（从左到右）；
2. 如果右子节点为空，此时左子节点的next一定是根节点中某个next节点的左子节点或者右子节点或者空，因此遍历即可，由于右子节点为空，就不需要进行处理；

```
/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/

class Solution {
public:
    Node* connect(Node* root) {
        connect(root, nullptr);
        return root;
    }

    void connect(Node* root, Node* next) {
        if (root == nullptr) {
            return;
        }
        root->next = next;
        Node* p = nullptr;
        if (root->right != nullptr) {
            p = root->next;
            while (p != nullptr) {
                if (p->left != nullptr) {
                    p = p->left;
                    break;
                }
                if (p->right != nullptr) {
                    p = p->right;
                    break;
                }
                p = p->next;
            }
            connect(root->right, p);
            connect(root->left, root->right);
        } else {
            p = root->next;
            while (p != nullptr) {
                if (p->left != nullptr) {
                    p = p->left;
                    break;
                }
                if (p->right != nullptr) {
                    p = p->right;
                    break;
                }
                p = p->next;
            }
            connect(root->left, p);
        }
    }
};
```
