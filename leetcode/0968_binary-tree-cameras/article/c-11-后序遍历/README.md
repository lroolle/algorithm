"+++
title = "0968. Binary Tree Cameras c++ 11 后序遍历 "
author = ["zintrulcre"]
date = 2020-09-22T09:53:26+08:00
tags = ["Leetcode", "Algorithms", "cpp", "后序遍历"]
categories = ["leetcode"]
draft = false
+++

# c++ 11 后序遍历

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [c++ 11 后序遍历](https://leetcode-cn.com/problems/binary-tree-cameras/solution/c-11-hou-xu-bian-li-by-zintrulcre/) by [zintrulcre](https://leetcode-cn.com/u/zintrulcre/)

```cpp
class Solution {
public:
    int minCameraCover(TreeNode* root) {
        int res = 0;
        enum class Status : int8_t
        {
            HasCamera,
            Covered,
            NotCovered
        };
        function<Status(TreeNode*)> Traverse = [&](TreeNode* node) -> Status
        {
            if (!node)
                return Status::Covered;
                
            Status left_status = Traverse(node->left);
            Status right_status = Traverse(node->right);

            if (left_status == Status::NotCovered || right_status == Status::NotCovered)
            {
                ++res;
                return Status::HasCamera;
            }
            
            if (left_status == Status::HasCamera || right_status == Status::HasCamera)
            {
                return Status::Covered;
            }
            
            // left_status == Status::Covered && right_status == Status::Covered
            return Status::NotCovered;
        };
        res += (Traverse(root) == Status::NotCovered);
        return res;
    }
};
![image.png](https://pic.leetcode-cn.com/1600768403-DUxljH-image.png)

```