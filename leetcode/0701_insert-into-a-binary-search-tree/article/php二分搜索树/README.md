"+++
title = "0701. Insert into a Binary Search Tree php二分搜索树 "
author = ["mek1986"]
date = 2020-05-29T12:00:26+08:00
tags = ["Leetcode", "Algorithms", "PHP", "BinarySearchTree"]
categories = ["leetcode"]
draft = false
+++

# php二分搜索树

> [0701. Insert into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)
> [php二分搜索树](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/solution/phper-fen-sou-suo-shu-by-mek1986-5/) by [mek1986](https://leetcode-cn.com/u/mek1986/)

思路：考察的还是数据结构，二分搜素树的插入操作。

### 代码

```php
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($val = 0, $left = null, $right = null) {
 *         $this->val = $val;
 *         $this->left = $left;
 *         $this->right = $right;
 *     }
 * }
 */
class Solution {

    /**
     * @param TreeNode $root
     * @param Integer $val
     * @return TreeNode
     */
    function insertIntoBST($root, $val) {
        return $this->add($root,$val);
    }

    function add($root,$val){
        if($root==null){
            return new TreeNode($val);
        }

        if($root->val>$val){
            $root->left=$this->add($root->left,$val);
        }else if($root->val<$val){
            $root->right=$this->add($root->right,$val);
        }

        return $root;
    }
}
```
![image.png](https://pic.leetcode-cn.com/1dbd73af1f62d9de91256ca2bc0acca8d7e86145c82cee7781b99fb8163d174a-image.png)
