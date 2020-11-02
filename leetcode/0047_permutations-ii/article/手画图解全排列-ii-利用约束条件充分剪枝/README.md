"+++
title = "0047. Permutations II 「手画图解」全排列 II | 利用约束条件充分剪枝 "
author = ["xiao_ben_zhu"]
date = 2020-09-18T00:27:27+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "Backtracking"]
categories = ["leetcode"]
draft = false
+++

# 「手画图解」全排列 II | 利用约束条件充分剪枝

> [0047. Permutations II](https://leetcode-cn.com/problems/permutations-ii/)
> [「手画图解」全排列 II | 利用约束条件充分剪枝](https://leetcode-cn.com/problems/permutations-ii/solution/shou-hua-tu-jie-li-yong-yue-shu-tiao-jian-chong-fe/) by [xiao_ben_zhu](https://leetcode-cn.com/u/xiao_ben_zhu/)

天气转凉了，今年入秋有点早。
#### 思路
这道题容易含糊的是，约束条件：
1. 一个数字不能重复地被选，见下图的红色圆圈。
2. 不能产生重复的排列。重复的排列是怎么产生的？见下图的红色三角。
   - 比如`[1,1,2]`，先选第一个`1`和先选第二个`1`，往后产生的排列是一样的。因此选第一个数时，不用考虑第二个`1`。
   - 再比如，已选了`2`，现在，选第一个`1`，和选第二个`1`，往后产生的排列也是一样的。
   - 它们都是 “同一层” 的选择出现重复，或者说，当前的可选选项有重复。
   
![image.png](https://pic.leetcode-cn.com/1600389760-qBZfUD-image.png)
   
重复的选项要忽略掉，为了方便在迭代中识别出重复，先对 nums 的数字升序排列，使相同的数字相邻。

#### 充分地剪枝
对应上面第一点，我们使用一个 used 数组记录使用过的数字，使用过了就不再使用：
```js
if (used[i]) {
  continue;
}
```
对应上面第二点，经过了排序后，如果当前的选项`nums[i]`，与同一层的上一个选项`nums[i - 1]`相同，且`nums[i - 1]`有意义（索引 >= 0），且没有被使用过，那就跳过该选项。
因为`nums[i - 1]`如果被使用过，它会被修剪掉，不是一个选项了，即便它和`nums[i]`重复，`nums[i]`还是可以选的。
```js
if (nums[i - 1] == nums[i] && i - 1 >= 0 && !used[i - 1]) {
  continue;
}
```
比如，`[1,1,2]`第一次选了第一个`1`，第二次可以选第二个`1`，即便它和它前一个`1`相同，因为前一个`1`被使用过了，它在本轮已经不是一个选项了，所以，第二轮中第二个`1`是可选的。
#### 为什么要回溯？
因为不是找到一个排列就完事，我们要找出所有的满足题意的排列。

经过了充分的剪枝，当数字选够了，就可以直接加入解集，结束递归。

此时，我们要撤销当前最新做的选择，回到选择前的状态，去选别的选项，切入别的搜索分支。

掉头，把路走全了，这样才能在解的空间树中，回溯出所有的解。

#### 代码

```javascript
const permuteUnique = (nums) => {
  const res = [];
  const len = nums.length;
  const used = new Array(len);
  nums.sort((a, b) => a - b); // 升序排序

  const helper = (path) => {
    if (path.length == len) { // 个数选够了
      res.push(path.slice()); // path的拷贝 加入解集
      return;                 // 结束当前递归 结束当前分支
    }

    for (let i = 0; i < len; i++) { // 枚举出所有的选择
      if (nums[i - 1] == nums[i] && i - 1 >= 0 && !used[i - 1]) { // 避免产生重复的排列
        continue;
      }
      if (used[i]) {      // 这个数使用过了，跳过。
        continue;
      }
      path.push(nums[i]); // make a choice
      used[i] = true;     // 记录路径上做过的选择
      helper(path);       // explore，基于它继续选，递归
      path.pop();         // undo the choice
      used[i] = false;    // 也要撤销一下对它的记录
    }
  };

  helper([]);
  return res;
};
```
#### 复盘总结
本题难点在于正确地“去重”。

也许我们会想到，先生成所有的排列，再用 hashSet 之类的做去重，是可行，但没有剪枝，会做很多无效的搜索，充分的剪枝可以保证生成出的排列就是满足题意的。

也许你会想到，用库函数类似 includes 去判断重复，但这样每次都花 $O(n)$ 时间，应该用 used 数组，空间换时间。

刚开始时，我忽略了`!used[i - 1]`，画图才发现，这个选择已经被我用另一个约束条件修剪掉了，它不会和同一层的隔壁选项发生冲突了。

画个图会清晰很多，对于剪枝来说。
#### 如果有帮助，点个赞鼓励我继续写下去画下去，如果哪里不对或哪里不够好，欢迎提出来一起完善。