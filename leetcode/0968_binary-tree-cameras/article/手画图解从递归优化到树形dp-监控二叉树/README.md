"+++
title = "0968. Binary Tree Cameras 「手画图解」从递归优化到树形DP | 监控二叉树 "
author = ["xiao_ben_zhu"]
date = 2020-09-22T01:02:41+08:00
tags = ["Leetcode", "Algorithms", "JavaScript", "DynamicProgramming"]
categories = ["leetcode"]
draft = false
+++

# 「手画图解」从递归优化到树形DP | 监控二叉树

> [0968. Binary Tree Cameras](https://leetcode-cn.com/problems/binary-tree-cameras/)
> [「手画图解」从递归优化到树形DP | 监控二叉树](https://leetcode-cn.com/problems/binary-tree-cameras/solution/shou-hua-tu-jie-cong-di-gui-you-hua-dao-dong-tai-g/) by [xiao_ben_zhu](https://leetcode-cn.com/u/xiao_ben_zhu/)

#### 思路引入
这道题和「打家劫舍III」很像，一样是二叉树。装摄像头难道是防盗？666。

打家劫舍III 规定不能打劫相邻的节点，求打劫节点值的最大收益。对于每个节点，要么打劫，要么不打劫，描述一个子树的最大收益需要两个变量：它的根节点、是否打劫根节点。

#### 思路
对于一个节点，它有什么状态，仅仅是放与不放相机吗？还有：是否被监控到。

两者可以重叠吗？可以的，被监控到的节点，也可以放相机。

有三个变量去描述一个节点的状态：节点本身（代表不同的子树）、是否放了相机、是否被监控到。

我们定义递归函数 minCam，计算：以当前 root 为根节点的子树，需要放置的最少相机数。

求整个树的 minCam，拆分为求子树的 minCam，`当前子树的minCam = 左子树的minCam + 右子树的minCam + 1(不一定+1)`，位于底部的 base case 易得，自上而下地调用，答案自下而上地返回。

递归函数接收的参数，对应上述三个变量：
1. 当前遍历的 root 节点。
2. placeCam：root 处是否放相机。
3. watched：root 是否被父亲或自己监控。
   - 因为递归是父亲调用儿子，对于当前节点，它只知道父亲和自己有没有监控自己，不知道儿子有没有监控自己，所以，watched 只代表是否被父亲或自己监控。

#### 分情况讨论（看图即可）

对于一个子树的根节点，它的状态无非下面三种，我们对应求出：子树不同状态下的 minCam。

![image.png](https://pic.leetcode-cn.com/1600753755-xhouuB-image.png)

1. 当前节点 root 放了相机（当前子树的相机数，保底为1）
    1. 左右儿子都没有放相机，都被父亲监控
        `minCam(root.left, false, true) + minCam(root.right, false, true)`
    2. 左儿子放了相机，被监控，右儿子没有相机，被父亲监控
        `minCam(root.left, true, true) + minCam(root.right, false, true)`
    3. 左儿子没有相机，被父亲监控，右儿子放了相机，被监控
        `minCam(root.left, false, true) + minCam(root.right, true, true)`
2. 当前节点 root 没有相机，但被父亲监控了
    1. 两个儿子都放了相机，被监控
    2. 左儿子放了相机，被监控，右儿子没有相机，没有被父亲和自己监控，但被自己儿子监控
    3. 右儿子放了相机，被监控，左儿子没有相机，没有被父亲和自己监控，但被自己儿子监控
    4. 两个儿子都没有相机，没有被父亲和自己监控，但被自己儿子监控
3. 当前节点 root 没有相机，也没有被父亲监控，是被儿子监控
    1. 两个儿子都放了相机，被监控
    2. 左儿子有相机，被监控，右儿子没有相机，没被父亲和自己监控，被自己儿子监控
    3. 左儿子没有相机，没被父亲和自己监控，被自己儿子监控。右儿子有相机，被监控


#### 递归的入口
整个树的根节点，它被监控，分两种情况：
1. 根节点有相机。
2. 根节点没有相机，且根节点没有父亲，没有被父亲监控，但被儿子监控。

对这两种情况求 minCam，返回结果较小者。
```js
const minCameraCover = (root) => {
  const minCam = (root, placeCam, watched) => { 
    // 省略
  };
  return Math.min(minCam(root, true, true), minCam(root, false, false));
};
```
#### 递归的结束条件
即 base case。当遍历到 null 节点时：

```js
if (root == null) {  // 遍历到null节点
  if (placeCam) {    // 父节点问自己，有相机的minCam，但null节点不可能有相机
    return Infinity; // 返回无穷大，让这个返回值被忽略掉（因为Math.min）
  } else {           // null没有相机，也没有子节点，下面也没有相机，返回0
    return 0;
  }
}
```

#### 执行结果：对是对，遇到这个 case 就超时

![image.png](https://pic.leetcode-cn.com/1600734234-IUIKcx-image.png)

#### 代码
```js
const minCameraCover = (root) => {
  // 以root为根节点的子树，放置的最少相机数
  const minCam = (root, placeCam, watched) => {
    if (root == null) {  // 遍历到null节点
      if (placeCam) {    // 父节点问自己有相机的情况，但臣妾办不到
        return Infinity; // 返回一个无穷大，让这个返回值失效
      } else {
        return 0;
      }
    }
    if (placeCam) {        // root放置相机
      return 1 + Math.min( // root放了相机，相机数有 1 保底
        minCam(root.left, false, true) + minCam(root.right, false, true), 
        minCam(root.left, true, true) + minCam(root.right, false, true), 
        minCam(root.left, false, true) + minCam(root.right, true, true) 
      );  
    } else {
      if (watched) { // root没放相机，但被父亲监控了
        return Math.min(
          minCam(root.left, true, true) + minCam(root.right, true, true),
          minCam(root.left, true, true) + minCam(root.right, false, false), 
          minCam(root.left, false, false) + minCam(root.right, true, true), 
          minCam(root.left, false, false) + minCam(root.right, false, false) 
        );
      } else {      // root没有相机，也没被父亲监控，被儿子监控了
        return Math.min(
          minCam(root.left, true, true) + minCam(root.right, true, true), 
          minCam(root.left, true, true) + minCam(root.right, false, false), 
          minCam(root.left, false, false) + minCam(root.right, true, true) 
        );
      }
    }
  };
  return Math.min(
    minCam(root, true, true),  // 根节点有相机
    minCam(root, false, false) // 根节点没有相机，因为没有父亲，没有被父亲监控，是被儿子监控
  );
};
```
#### 怎么优化？哪里做了重复的计算？
可以看到，我们反复地在调用左右子树，传的参数就有很多重复。

我们看到，对于一个子树，它有三种返回值，即三种状态下的最优解。我们这样，对于每一个子树，即每一次递归调用，都计算三种状态下的最优解 minCam：
1. withCam ：当前子树 root 有相机，情况下的minCam
2. noCamWatchByDad：当前子树 root 没有相机，被父亲监控，情况下的minCam
3. noCamWatchBySon ：当前子树 root 没有相机，被儿子监控，情况下的minCam

相当于`dp[root][0]`,`dp[root][1]`,`dp[root][2]`，放入一个对象，返回出来，返回给父调用。

```js
return { withCam, noCamWatchByDad, noCamWatchBySon };
```
当前以 root 为根的子树，递归调用左右子树，拿到子树的三种最优解，递推算出当前树的三种最优解：withCam、noCamWatchByDad、noCamWatchBySon。
```js
const left = minCam(root.left);
const right = minCam(root.right);
```
这样，一次递归调用，就只执行两次子调用。

#### 优化后代码

```js
const minCameraCover = (root) => {
  const minCam = (root) => {
    if (root == null) {   // base case
      return {
        withCam: Infinity,
        noCamWatchByDad: 0,
        noCamWatchBySon: 0
      };
    }
    const left = minCam(root.left);   // 以左儿子为根的左子树的minCam
    const right = minCam(root.right); // 以右儿子为根的右子树的minCam
    // 下面相当于状态转移方程
    const withCam = 1 + Math.min(     
      left.noCamWatchByDad + right.noCamWatchByDad,
      left.withCam + right.noCamWatchByDad,
      left.noCamWatchByDad + right.withCam
    );

    const noCamWatchByDad = Math.min(
      left.withCam + right.withCam,
      left.withCam + right.noCamWatchBySon,
      left.noCamWatchBySon + right.withCam,
      left.noCamWatchBySon + right.noCamWatchBySon
    );

    const noCamWatchBySon = Math.min(
      left.withCam + right.withCam,
      left.withCam + right.noCamWatchBySon,
      left.noCamWatchBySon + right.withCam
    );

    return { withCam, noCamWatchByDad, noCamWatchBySon };
  };

  const res = minCam(root); // 相当于dp[root]
  return Math.min(res.withCam, res.noCamWatchBySon); // 相当于 dp[root][0]、dp[root][2]
};
```
#### 复盘总结
> **贪心算法我不熟，而且有点像碰运气，有的行得通，有的行不通，不好用数学证明其正确性。**

树形 DP 不像常规 DP 那样在迭代中 “填表”，而是在递归遍历中 “填表”。

这里没有开辟 DP 数组去存中间状态，而是通过子调用将状态返回出去，提供给父调用，其实就是动态规划填表的思想。

因为动态规划：根据过去的状态求出当前状态，按顺序一个个求。这里是沿着一棵树去求解子问题，可以理解为在一棵树上填表。

可以理解为，递归调用栈把中间计算结果暂存了，子调用的结果“交付”给父调用，它就销毁了，并没有加入记忆化。

当然，你也可以开辟一个容器，key 是节点（子树），因此没法用数组，那就用 HashMap，value 是子树的计算结果。

可以，但没必要，因为不需要存储中间计算结果，所以这算是降维优化了。

随着递归的出栈，子调用不断向上返回，子问题（子树）被一个个解决。最后求出大问题：整个树的最小相机数。

#### 感谢阅读，点赞更棒。真诚地在写，话也是我想说的，相信你也感受的到。尽量在用短句，字斟句酌，有时候不能言简意赅一针见血，我尽量努力。欢迎具体建议和反馈。
