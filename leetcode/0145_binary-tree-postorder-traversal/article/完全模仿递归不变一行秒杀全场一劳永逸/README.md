"+++
title = "0145. Binary Tree Postorder Traversal 完全模仿递归，不变一行。秒杀全场，一劳永逸 "
author = ["sonp"]
date = 2020-04-27T04:40:46+08:00
tags = ["Leetcode", "Algorithms", "C++"]
categories = ["leetcode"]
draft = false
+++

# 完全模仿递归，不变一行。秒杀全场，一劳永逸

> [0145. Binary Tree Postorder Traversal](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)
> [完全模仿递归，不变一行。秒杀全场，一劳永逸](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/solution/mo-fang-di-gui-zhi-bian-yi-xing-by-sonp/) by [sonp](https://leetcode-cn.com/u/sonp/)

#### 解题思路：

如果你只想掌握其中某一种遍历大可去找那些奇技淫巧的题解（一会判断指针一会儿判断栈就问你怕不怕？面试的时候你能想起来？）；如果你想统一掌握三种遍历，并且希望思路清晰，我强烈建议你阅读下去！因为这里介绍的是递归转迭代的思路，而不仅仅是用迭代的形式完成题目。本题解也是双 100% 哦！

[老思路链接](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/liang-ta-lai-liao-ta-dai-zhao-san-xiong-di-lai-lia/)

新思路

1. 递归的本质就是压栈，了解递归本质后就完全可以按照递归的思路来迭代。
2. 怎么压，压什么？压的当然是待执行的内容，后面的语句先进栈，所以进栈顺序就决定了前中后序。

我们需要一个标志区分每个递归调用栈，这里使用 `nullptr` 来表示。
具体直接看注释，可以参考文章最后 “和递归写法的对比”。先序遍历看懂了，中序和后序也就秒懂。
先序遍历
```C++
class Solution {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> res;  //保存结果
        stack<TreeNode*> call;  //调用栈
        if(root!=nullptr) call.push(root);  //首先介入root节点
        while(!call.empty()){
            TreeNode *t = call.top();
            call.pop();  //访问过的节点弹出
            if(t!=nullptr){
                if(t->right) call.push(t->right);  //右节点先压栈，最后处理
                if(t->left) call.push(t->left);
                call.push(t);  //当前节点重新压栈（留着以后处理），因为先序遍历所以最后压栈
                call.push(nullptr);  //在当前节点之前加入一个空节点表示已经访问过了
            }else{  //空节点表示之前已经访问过了，现在需要处理除了递归之外的内容
                res.push_back(call.top()->val);  //call.top()是nullptr之前压栈的一个节点，也就是上面call.push(t)中的那个t
                call.pop();  //处理完了，第二次弹出节点（彻底从栈中移除）
            }
        }
        return res;
    }
};
```

后序遍历
你没看错，只有注释部分改变了顺序，父>右>左。
```C++
class Solution {
public:
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> res;
        stack<TreeNode*> call;
        if(root!=nullptr) call.push(root);
        while(!call.empty()){
            TreeNode *t = call.top();
            call.pop();
            if(t!=nullptr){
                call.push(t);  //在右节点之前重新插入该节点，以便在最后处理（访问值）
                call.push(nullptr); //nullptr跟随t插入，标识已经访问过，还没有被处理
                if(t->right) call.push(t->right);
                if(t->left) call.push(t->left);
            }else{
                res.push_back(call.top()->val);
                call.pop();
            }
        }
        return res;   
    }
};
```
中序遍历
你没看错，只有注释部分改变了顺序，右>父>左。其他和前序遍历 “一模一样”
```C++
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> res;
        stack<TreeNode*> call;
        if(root!=nullptr) call.push(root);
        while(!call.empty()){
            TreeNode *t = call.top();
            call.pop();
            if(t!=nullptr){
                if(t->right) call.push(t->right);
                call.push(t);  //在左节点之前重新插入该节点，以便在左节点之后处理（访问值）
                call.push(nullptr); //nullptr跟随t插入，标识已经访问过，还没有被处理
                if(t->left) call.push(t->left);
            }else{
                res.push_back(call.top()->val);
                call.pop();
            }
        }
        return res;
    }
};
```

 对比中序遍历的递归写法

```C++
void dfs(t){ //进入函数表示“访问过”，将t从栈中弹出

    dfs(t->left);   //因为要访问t->left, 所以我先把函数中下面的信息都存到栈里。
                //依次call.push(t->right), call.push(t)【t第二次入栈】, call.push(nullptr)【标识t二次入栈】, call.push(t->left)。
                //此时t并没有被处理（卖萌）。栈顶是t->left, 所以现在进入t->left的递归中。

    //res.push_back(t->val)
    t.卖萌();   //t->left 处理完了，t->left被彻底弹出栈。
                //此时栈顶是nullptr, 表示t是已经访问过的。那么我现在需要真正的处理t了（即，执行卖萌操作）。
                //卖萌结束后，t 就被彻底弹出栈了。
    

    dfs(t->right); 
}