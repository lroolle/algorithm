"+++
title = "0501. Find Mode in Binary Search Tree 真正符合题目要求的O(1)空间复杂度解法 "
author = ["faterazer"]
date = 2020-07-15T10:55:24+08:00
tags = ["Leetcode", "Algorithms", "C++", "Java", "Python3", "Python"]
categories = ["leetcode"]
draft = false
+++

# 真正符合题目要求的O(1)空间复杂度解法

> [0501. Find Mode in Binary Search Tree](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/)
> [真正符合题目要求的O(1)空间复杂度解法](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/solution/zhen-zheng-fu-he-ti-mu-yao-qiu-de-o1kong-jian-fu-z/) by [faterazer](https://leetcode-cn.com/u/faterazer/)

我看了题解区和评论区的很多生成 $O(1)$ 的高赞解法，但实际上这些解法都不是真正 $O(1)$ 的。考虑这样一种最坏情况：`[1,2,3,4,...,n-2,n-1,n-1]`，返回的答案应该只有一个，但中间开辟的额外空间会达到 $n-2$ 个，即在这种最坏情况下，空间复杂度是 $O(n)$，这显然是不符合题目的进阶要求的。

怎么才能做到真正 $O(1)$ 的空间复杂度呢？我认为需要两点：

1. 使用中序遍历。BST 的中序遍历结果是有序的。
2. 遍历两次。第一次确定结果集的大小，第二次得到结果集。

我的解法：

C++版本：

```cpp
class Solution
{
public:
    vector<int> findMode(TreeNode *root)
    {
        vector<int> ret;    // 要返回的结果
        // max_count: 结果元素的出现次数
        // cur_count: 当前元素的出现次数
        // ret_count: 结果 vector 的长度
        int max_count = 0, cur_count = 0, ret_count = 0;
        // pre: 上一个访问的结点
        TreeNode *pre = nullptr;
        inOrder(root, pre, ret, ret_count, max_count, cur_count);
        // 第二次中序遍历
        ret.resize(ret_count);
        cur_count = 0;
        ret_count = 0;
        pre = nullptr;
        inOrder(root, pre, ret, ret_count, max_count, cur_count);
        return ret;
    }
private:
    void inOrder(TreeNode *root, TreeNode *&pre, vector<int> &ret, int &ret_count, int &max_count, int &cur_count)
    {
        if (!root)
            return;
        inOrder(root->left, pre, ret, ret_count, max_count, cur_count);
        // 如果上一个结点为空，或与当前值相同，cur_count + 1
        if (pre && pre->val == root->val)
            cur_count++;
        else    // 否则，cur_count 重新计数
            cur_count = 1;
        if (cur_count > max_count) {
            max_count = cur_count;
            ret_count = 1;
        }
        else if (cur_count == max_count) {
            if (ret.size()) // 注意这里，第一次遍历时，这里的if不会执行
                ret[ret_count] = root->val;
            ret_count++;
        }
        pre = root;
        inOrder(root->right, pre, ret, ret_count, max_count, cur_count);
    }
};
```

Python 版本：

```python
class Solution:
    def __init__(self):
        self.pre = None
        self.ret = []
        self.ret_count, self.max_count, self.cur_count = 0, 0, 0

    def findMode(self, root: TreeNode) -> List[int]:
        self.inOrder(root)
        self.pre = None
        self.ret = [0] * self.ret_count
        self.ret_count, self.cur_count = 0, 0
        self.inOrder(root)
        return self.ret

    def inOrder(self, root: TreeNode) -> None:
        if not root:
            return
        self.inOrder(root.left)
        if self.pre and self.pre.val == root.val:
            self.cur_count += 1
        else:
            self.cur_count = 1
        if self.cur_count > self.max_count:
            self.max_count = self.cur_count
            self.ret_count = 1
        elif self.cur_count == self.max_count:
            if len(self.ret):
                self.ret[self.ret_count] = root.val
            self.ret_count += 1
        self.pre = root
        self.inOrder(root.right)
```

Java 版本：

```java
class Solution {
    private TreeNode pre = null;
    private int[] ret;
    private int retCount = 0;
    private int maxCount = 0;
    private int currCount = 0;

    public int[] findMode(TreeNode root) {
        inOrder(root);
        pre = null;
        ret = new int[retCount];
        retCount = 0;
        currCount = 0;
        inOrder(root);
        return ret;
    }

    private void inOrder(TreeNode root) {
        if (root == null)
            return;
        inOrder(root.left);
        if (pre != null && pre.val == root.val)
            currCount++;
        else
            currCount = 1;
        if (currCount > maxCount) {
            maxCount = currCount;
            retCount = 1;
        }
        else if (currCount == maxCount) {
            if (ret != null)
                ret[retCount] = root.val;
            retCount++;
        }
        pre = root;
        inOrder(root.right);
    }
}
```