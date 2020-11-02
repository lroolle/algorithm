"+++
title = "0106. Construct Binary Tree from Inorder and Postorder Traversal 两种实现+动画演示 106. 从中序与后序遍历序列构造二叉树 "
author = ["wang_ni_ma"]
date = 2020-07-08T08:57:24+08:00
tags = ["Leetcode", "Algorithms", "Recursion", "Python", "Java"]
categories = ["leetcode"]
draft = false
+++

# 两种实现+动画演示 106. 从中序与后序遍历序列构造二叉树

> [0106. Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
> [两种实现+动画演示 106. 从中序与后序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/solution/liang-chong-shi-xian-dong-hua-yan-shi-106-cong-zho/) by [wang_ni_ma](https://leetcode-cn.com/u/wang_ni_ma/)


## 解法一
这道题跟【[105.从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solution/dong-hua-yan-shi-105-cong-qian-xu-yu-zhong-xu-bian/)】很像，105题是用前序遍历后的数组+中序遍历后的数组，来构造一个棵二叉树。   
本题则是中序+后序，但是思路都是类似的   
首先，回忆一下后序遍历:
```
dfs(root.left)
dfs(root.right)
print root
```
后序遍历的特点是，根节点始终是最后一个被打印的   
所以根据输入的后序遍历数组的最后一位，就可以确定出根节点了   
![1.jpg](https://pic.leetcode-cn.com/5312c3f53e4dcefa64ad5261d62b4524fa6296112e597914b47f212564fe7161-1.jpg)
因为输入数组的元素不重复，所以拿后序数组的最后一个元素```1```，去中序数组中查找，就可以确定中序数组中```1```的下标。   
     
我们根据后序数组的最后一个元素```1```，找到了对应的中序数组```1```的下标**index**。
因为中序遍历的特点是根节点把数组分成了左右两部分，所以根据index，就可以确定出左子树部分，以及右子树部分 
![2.jpg](https://pic.leetcode-cn.com/9379211cc016c5a56d770f4ce2134faee14aa214d6b500ce787ef5c0dc328dcf-2.jpg)


上图中，我们根据```1```的下标index，将中序数组分成两半。   
后序遍历和中序遍历一样，都是先遍历左子树部分，所以中序数组和后序数组的左子树部分长度是一样的，也就是对应上图中绿色的部分。   
同样黄色部分就对应了右子树部分，由于后序遍历是根节点在最后、中序遍历是根节点在中间，所以它们拆分数组的规则会有不同。

具体的拆分规则如下:(假设数组最后一个元素的下标是END_INDEX)   
中序数组左半边:```[0:index)```   
中序数组右半边:```[index+1:END_INDEX]```  
后序数组的左半边:```[0:index)```   
后序数组的右半边:```[index:END_INDEX-1]```  
   
由此我们可以得到这么一个递归函数
```
root.left = 递归函数(左子树部分)
root.right = 递归函数(右子树部分)
reutrn root
```
   
   
动态演示如下:
![3.gif](https://pic.leetcode-cn.com/9ec30348c10f0ab6b770e258bf00d5c54598b2a3bb8ddde641aebf78f2d2c68f-3.gif)




代码实现:
```java
class Solution {
    public TreeNode buildTree(int[] inorder, int[] postorder) {
        if(inorder==null || postorder==null) {
            return null;
        }
        return helper(inorder,postorder);
    }

    private TreeNode helper(int[] in, int[] post) {
        if(in.length==0) {
            return null;
        }
        //根据后序数组的最后一个元素，创建根节点
        TreeNode root = new TreeNode(post[post.length-1]);
        //在中序数组中查找值等于【后序数组最后一个元素】的下标
        for(int i=0;i<in.length;++i) {
            if(in[i]==post[post.length-1]) {
                //确定这个下标i后，将中序数组分成两部分，后序数组分成两部分
                int[] inLeft = Arrays.copyOfRange(in,0,i);
                int[] inRight = Arrays.copyOfRange(in,i+1,in.length);
                int[] postLeft = Arrays.copyOfRange(post,0,i);
                int[] postRight = Arrays.copyOfRange(post,i,post.length-1);
                //递归处理中序数组左边，后序数组左边
                root.left = helper(inLeft,postLeft);
                //递归处理中序数组右边，后序数组右边
                root.right = helper(inRight,postRight);
                break;
            }
        }
        return root;
    }
}
```
```python
class Solution(object):
    def buildTree(self, inorder, postorder):
        if not (inorder and postorder):
            return None
        def helper(inor,post):
            if not post:
                return None
            # 根据后序数组的最后一个元素，创建根节点    
            root = TreeNode(post[-1])
            # 在中序数组中查找值等于【后序数组最后一个元素】的下标
            idx = inor.index(post[-1])
            # 确定这个下标i后，将中序数组分成两部分，后序数组分成两部
            # 递归处理中序数组左边，后序数组左边
            # 递归处理中序数组右边，后序数组右边
            root.left = helper(inor[:idx],post[:idx])
            root.right = helper(inor[idx+1:],post[idx:-1])
            return root
        return helper(inorder,postorder)
```

## 解法二
解法一中，我们要确定中序数组的下标index，然后根据这个index将中序、后序数组各分成两半，也就是新生成4个数组，之后再递归的处理这些新数组。   
   
这样的效率是属于O(N^2)了，虽然也能通过，但效率不高。主要是每次查找都是线性查找，另外每次都要重新创建新数组，我们可以将这部分优化一下。   
因为是用**后序数组的最后一个元素**去中序数组中查找，所以我们可以先将中序数组的下标和值保存到**哈希表**中，这样就省去了线性查找了。      
之后递归处理的时候，我们不再新创建数组，而是传入数组下标，这样就不用频繁创建新数组了。   
这样的话时间复杂度可以降低到0(N)   
![4.gif](https://pic.leetcode-cn.com/efab9abe389b3797ed105a16abe9ce3a488b1650f2662219f447e589bfab8203-4.gif)


解法二的递归函数跟解法一有点小小不同     
解法二是先处理右子树，再处理左子树
```python
root = TreeNode(val)
root.right = helper(处理右子树)
root.left = helper(处理左子树)
return root
```
从上图中我们可以发现，每次处理后序数组时，都是拿最后一个，然后post_index--，相当于是逆序遍历后序数组，如果先处理左子树下标就会越界，所以先处理右子树。   


代码实现 :
```java
class Solution {
    private int postIdx = 0;
    public TreeNode buildTree(int[] inorder, int[] postorder) {
        if(inorder==null || postorder==null) {
            return null;
        }
        //将中序数组的下标、值保存到map中省去解法一中线性查找
        Map<Integer,Integer> map = new HashMap<Integer,Integer>();
        for(int i=0;i<inorder.length;++i) {
            map.put(inorder[i],i);
        }        
        this.postIdx = postorder.length-1;
		return helper(map,inorder,postorder,0,inorder.length-1);
    }
	    
    private TreeNode helper(Map<Integer,Integer> map, int[] inor, int[] post, int left, int right) {
        if(left>right) {
            return null;
        }
        //从后序数组中拿最后一个元素，根据这个元素去map中找到中序数组对应的index
        //然后递归的处理右边[index+1,right]，递归处理左边[left,index-1]
        int val = post[this.postIdx];
        --this.postIdx;
        TreeNode root = new TreeNode(val);
        int index = map.get(val);
        root.right = helper(map,inor,post,index+1,right);
        root.left = helper(map,inor,post,left,index-1);
        return root;
    }
}
```
```python
class Solution(object):
    def buildTree(self, inorder, postorder):
        if not (inorder and postorder):
            return None
        # 将中序数组的下标、值保存到map中省去解法一中线性查找
        d = {val:idx for idx,val in enumerate(inorder)}
        self.post_idx = len(postorder)-1
        def dfs(left,right):
            if left>right:
                return None
            # 从后序数组中拿最后一个元素，根据这个元素去map中找到中序数组对应的index
            # 然后递归的处理右边[index+1,right]，递归处理左边[left,index-1]
            val = postorder[self.post_idx]
            self.post_idx -= 1
            root = TreeNode(val)
            index = d[val]
            root.right = dfs(index+1,right)
            root.left = dfs(left,index-1)
            return root
        return dfs(0,len(inorder)-1)
```

关于前序、中序、后序遍历序列构造二叉树一共有三道题
| 题目 | 题解 | 难度等级 
|---| --- | --- |
|[从前序与中序遍历序列构造二叉树 ](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/) | [题解链接](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solution/dong-hua-yan-shi-105-cong-qian-xu-yu-zhong-xu-bian/) | 中等 |
|[从中序与后序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)| [题解链接](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/solution/liang-chong-shi-xian-dong-hua-yan-shi-106-cong-zho/)| 中鞥 |
|[根据前序和后序遍历构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/)|[题解链接](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/solution/tu-jie-889-gen-ju-qian-xu-he-hou-xu-bian-li-gou-2/)| 中等 |


**欢迎关注 👉👉👉  [【公众号】](https://share.weiyun.com/0KDRUnfK) 👈👈👈**   

**如果能再点个赞👍👍 就更感激啦💓💓**






















