"+++
title = "0347. Top K Frequent Elements python/c++ 自己实现堆，优先队列 "
author = ["xxinjiee"]
date = 2019-08-08T22:10:45+08:00
tags = ["Leetcode", "Algorithms", "Python", "Heap", "HashTable", "优先队列", "Python3", "cpp"]
categories = ["leetcode"]
draft = false
+++

# python/c++ 自己实现堆，优先队列

> [0347. Top K Frequent Elements](https://leetcode-cn.com/problems/top-k-frequent-elements/)
> [python/c++ 自己实现堆，优先队列](https://leetcode-cn.com/problems/top-k-frequent-elements/solution/python-dui-pai-xu-by-xxinjiee/) by [xxinjiee](https://leetcode-cn.com/u/xxinjiee/)

### 解题思路：
更新 0623： 更新为位操作，占位节点版本，简化的代码，方便记忆。

这题是对 **堆，优先队列**  很好的练习，因此有必要自己用python实现研究一下。**堆 处理海量数据的 topK，分位数** 非常合适，**优先队列** 应用在元素优先级排序，比如本题的频率排序非常合适。与基于比较的排序算法 时间复杂度 $O(nlogn)$ 相比，使用 **堆，优先队列** 复杂度可以下降到 $O(nlogk)$，在总体数据规模 `n` 较大，而维护规模 `k` 较小时，时间复杂度优化明显。

**堆，优先队列** 的本质其实就是个完全二叉树，有其下重要性质
ps: 堆 `heap[0]` 插入一个占位节点，此时堆顶为 index 为 1 的位置，可以更方便的运用位操作。
`[1,2,3] -> [0,1,2,3]`
1. 父节点 index 为 `i`
2. 左子节点 index 为 `i << 1`
3. 右子节点 index 为 `i << 1 | 1`
4. 大顶堆中每个父节点大于子节点，小顶堆每个父节点小于子节点
5. 优先队列以优先级为堆的排序依据
因为性质 1，2，3，堆可以用数组直接来表示，不需要通过链表建树。

**堆，优先队列** 有两个重要操作，时间复杂度均是 $O(logk)$。以小顶锥为例：
1. 上浮 sift up: 向堆尾新加入一个元素，堆规模 +1，依次向上与父节点比较，如小于父节点就交换。
2. 下沉 sift down: 从堆顶取出一个元素（堆规模 -1，用于堆排序）或者更新堆中一个元素（本题），依次向下与子节点比较，如大于子节点就交换。

对于 topk 问题：**最大堆求topk小，最小堆求 topk 大。**
- topk小：构建一个 `k` 个数的最大堆，当读取的数小于根节点时，替换根节点，重新塑造最大堆
- topk大：构建一个 `k` 个数的最小堆，当读取的数大于根节点时，替换根节点，重新塑造最小堆

**这一题的总体思路** 总体时间复杂度 **$O(nlogk)$**
- 遍历统计元素出现频率 $O(n)$
- 前k个数构造 **规模为 k+1 的最小堆** minheap， $O(k)$， 注意 +1 是因为占位节点。
- 遍历规模k之外的数据，大于堆顶则入堆，下沉维护规模为k的最小堆 minheap. $O(nlogk)$
- (如需按频率输出，对规模为k的堆进行排序)


```Python
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        def sift_down(arr, root, k):
            """下沉log(k),如果新的根节点>子节点就一直下沉"""
            val = arr[root] # 用类似插入排序的赋值交换
            while root<<1 < k:
                child = root << 1
                # 选取左右孩子中小的与父节点交换
                if child|1 < k and arr[child|1][1] < arr[child][1]:
                    child |= 1
                # 如果子节点<新节点,交换,如果已经有序break
                if arr[child][1] < val[1]:
                    arr[root] = arr[child]
                    root = child
                else:
                    break
            arr[root] = val

        def sift_up(arr, child):
            """上浮log(k),如果新加入的节点<父节点就一直上浮"""
            val = arr[child]
            while child>>1 > 0 and val[1] < arr[child>>1][1]:
                arr[child] = arr[child>>1]
                child >>= 1
            arr[child] = val

        stat = collections.Counter(nums)
        stat = list(stat.items())
        heap = [(0,0)]

        # 构建规模为k+1的堆,新元素加入堆尾,上浮
        for i in range(k):
            heap.append(stat[i])
            sift_up(heap, len(heap)-1) 
        # 维护规模为k+1的堆,如果新元素大于堆顶,入堆,并下沉
        for i in range(k, len(stat)):
            if stat[i][1] > heap[1][1]:
                heap[1] = stat[i]
                sift_down(heap, 1, k+1) 
        return [item[0] for item in heap[1:]]
```
```C++
class Solution {
public:
    void sift_up(vector<vector<int>> &heap, int chlid){
        vector<int> val = heap[chlid];
        while (chlid >> 1 > 0 && val[1] < heap[chlid>>1][1]){
            heap[chlid] = heap[chlid>>1];
            chlid >>= 1;
        heap[chlid] = val;
        }
    }

    void sift_down(vector<vector<int>> &heap, int root, int k){
        vector<int> val = heap[root];
        while (root << 1 < k){
            int chlid = root << 1;
            // 注意这里位运算优先级要加括号
            if ((chlid|1) < k && heap[chlid|1][1] < heap[chlid][1]) chlid |= 1;
            if (heap[chlid][1] < val[1]){
                heap[root] = heap[chlid];
                root = chlid;
            }
            else break;
        }
        heap[root] = val;
    }

    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> stat;
        for (auto &num : nums) stat[num]++;
        vector<vector<int>> vec_stat;
        for (auto &item : stat) vec_stat.push_back({item.first, item.second});

        vector<vector<int>> heap;
        heap.push_back({0, 0});
        for (int i = 0; i < k; i++){
            heap.push_back(vec_stat[i]);
            sift_up(heap, heap.size()-1);
        }

        for (int i = k; i < vec_stat.size(); i++){
            if (vec_stat[i][1] > heap[1][1]){
                heap[1] = vec_stat[i];
                sift_down(heap, 1, k+1);
            }
        }

        vector<int> result;
        for (int i = 1; i < k+1; i++) result.push_back(heap[i][0]);
        return result;
    }
};
```

再附上堆排序(从小到大输出)，注意这里是大顶堆
1. 从后往前非叶子节点下沉，依次向上保证每一个子树都是大顶堆，构造大顶锥
2. 依次把大顶堆根节点与尾部节点交换(不再维护，堆规模 -1)，新根节点下沉。

```Python 
def heapSort(arr):
    def sift_down(arr, root, k):
        val = arr[root]
        while root<<1 < k:
            chlid = root << 1
            if chlid|1 < k and arr[chlid|1] > arr[chlid]:
                chlid |= 1
            if arr[chlid] > val:
                arr[root] = arr[chlid]
                root = chlid
            else:
                break
        arr[root] = val

    arr = [0] + arr
    k = len(arr)
    for i in range((k-1)>>1, 0, -1):
        sift_down(arr, i, k) 
    for i in range(k-1, 0, -1):
        arr[1], arr[i] = arr[i], arr[1]
        sift_down(arr, 1, i)
    return arr[1:]
```
<br>
#### 更多的几个堆的练习
[295. 数据流的中位数](https://leetcode-cn.com/problems/find-median-from-data-stream/)
[215. 数组中的第K个最大元素](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)
[面试题40. 最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)
[347. 前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements)