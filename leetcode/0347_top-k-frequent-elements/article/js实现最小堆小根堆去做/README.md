"+++
title = "0347. Top K Frequent Elements JS实现最小堆（小根堆）去做 "
author = ["xiao_ben_zhu"]
date = 2020-09-07T07:35:54+08:00
tags = ["Leetcode", "Algorithms", "JavaScript"]
categories = ["leetcode"]
draft = false
+++

# JS实现最小堆（小根堆）去做

> [0347. Top K Frequent Elements](https://leetcode-cn.com/problems/top-k-frequent-elements/)
> [JS实现最小堆（小根堆）去做](https://leetcode-cn.com/problems/top-k-frequent-elements/solution/jsshi-xian-zui-xiao-dui-xiao-gen-dui-qu-zuo-by-xia/) by [xiao_ben_zhu](https://leetcode-cn.com/u/xiao_ben_zhu/)


#### 思路
实现一个最小堆，维护一个heap数组，长度为 k，存储数字本身，对应的频次是从低到高，当heap数组长度不够k时，就从后push推入数字，然后执行“上浮”，即元素在最小堆中，交换到它合适的位置。当heap数组长度够k时，如果新的数字的频次比堆顶的数字的频次高，我们将它去替换堆顶的数字，然后执行“下沉”，交换到最小堆中合适的位置。

最后，heap数组存的是频次从低到高的 k 个数字，然后我们逆序一下，输出即可。
#### 代码

```js
const topKFrequent = (nums, k) => {
  const freq = {};       // 存储数字出现的频次
  const uniqueNums = []; // 不重复的数字
  for (const num of nums) {
    if (freq[num]) {     // 出现过，频次+1
      freq[num]++;
    } else {             // 没出现过，频次为1
      freq[num] = 1;
      uniqueNums.push(num);
    }
  }

  const heap = [];       // 代表heap的数组

  const swap = (i, j) => { // 交换heap数组的元素
    const t = heap[i];
    heap[i] = heap[j];
    heap[j] = t;
  };

  const bubbleUp = (index) => {
    while (index > 0) {
      const parent = (index - 1) >>> 1;  // 找到父节点在heap数组中的位置
      if (freq[heap[parent]] > freq[heap[index]]) { // 如果父节点数字对应的频率要高于插入的数字的频次
        swap(parent, index); // 交换它们的位置
        index = parent;      // index更新
      } else {               // 满足最小堆的特点，break
        break;
      }
    }
  };

  const bubbleDown = (index) => { // 做“下沉”
    while (2 * index + 1 < heap.length) { // 
      let child = 2 * index + 1;
      if (child + 1 < heap.length && freq[heap[child + 1]] < freq[heap[child]]) { // 左右孩子中取较小的去比较
        child++;
      }
      if (freq[heap[index]] > freq[heap[child]]) {
        swap(index, child); // 交换
        index = child;      // 更新 index
      } else { // 如果满足最小堆的属性，break
        break;
      }
    }
  };

  for (const num of uniqueNums) {
    if (heap.length < k) { // heap数组的长度还不够k
      heap.push(num);      // 推入heap数组
      bubbleUp(heap.length - 1); // 执行上浮，频率小的上浮上去
    } else if (freq[num] > freq[heap[0]]) { // 如果当前数字的频次比堆顶数字的频率要大
      heap[0] = num; // 堆顶的数字要更换
      bubbleDown(0); // 然后要做下沉，下沉到合适的位置
    }
  }
  return heap.sort((a, b) => { // 最终它前面是频次少的，最后返回是要逆序回来
    return freq[b] - freq[a];
  });
};
```
#### 如果有帮助，点个小赞鼓励我继续写下去，如果哪里不对或更好的建议，指出我我我继续修改。