"+++
title = "0347. Top K Frequent Elements 非标准解法  手工创建LFU统计频次  但性能很高 "
author = ["reikiaura"]
date = 2020-09-06T21:19:35+08:00
tags = ["Leetcode", "Algorithms", "Java", "HashTable", "LinkedList"]
categories = ["leetcode"]
draft = false
+++

# 非标准解法  手工创建LFU统计频次  但性能很高

> [0347. Top K Frequent Elements](https://leetcode-cn.com/problems/top-k-frequent-elements/)
> [非标准解法  手工创建LFU统计频次  但性能很高](https://leetcode-cn.com/problems/top-k-frequent-elements/solution/fei-biao-zhun-jie-fa-shou-gong-chuang-jian-lfutong/) by [reikiaura](https://leetcode-cn.com/u/reikiaura/)

### 基于 LFU 统计频次 ###
本解法非此题考查的主要目的, 仅供参考
使用一种性能较高的 LFU 写法: DoublyLinkedList + HashMap
题目参考: Leetcode 460-LFU缓存
对输入数组进行扫描, 添加至 LFU, 添加操作每步均为 O(1) 整体复杂度为 O(N)



### 提交结果截图 ###
![image.png](https://pic.leetcode-cn.com/1599427399-onkAtb-image.png)



### 代码实现 ###
```
class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        LFU lfu = new LFU();
        for (int val : nums) {
            lfu.put(val);
        }
        return lfu.getTopKFreqent(k);
    }

    private class LFU {
        HashMap<Integer, Node> cache;
        DoublyLinkedList head;
        DoublyLinkedList tail;

        public LFU() {
            cache = new HashMap<>();
            head = new DoublyLinkedList(0);
            tail = new DoublyLinkedList(0);
            head.next = tail;
            tail.prev = head;
        }

        public void put(int key) {
            Node node = cache.get(key);
            if (node == null) {
                node = new Node(key);
                DoublyLinkedList prevList = tail.prev;
                if (prevList.freq == node.freq) {
                    prevList.addToTail(node);
                } else {
                    DoublyLinkedList newList = new DoublyLinkedList(node.freq);
                    newList.addToTail(node);
                    addList(newList, prevList);
                }
            } else {
                node.freq++;
                DoublyLinkedList curList = node.curList;
                DoublyLinkedList prevList = curList.prev;
                curList.remove(node);
                if (curList.head.next == curList.tail) {
                    removeList(curList);
                }
                if (prevList.freq == node.freq) {
                    prevList.addToTail(node);
                } else {
                    DoublyLinkedList newList = new DoublyLinkedList(node.freq);
                    newList.addToTail(node);
                    addList(newList, prevList);
                }
            }
            cache.put(key, node);
        }

        public int[] getTopKFreqent(int k) {
            int[] res = new int[k];
            int cur = 0;
            Node curNode = head.next.head.next;
            while (cur < k) {
                // System.out.println(cur + ", " + curNode.key + ", " + curNode.freq + ", "+ curNode.next.key + ", " + curNode.curList.freq);
                res[cur++] = curNode.key;
                curNode = curNode.next;
                if (curNode == curNode.curList.tail) {
                    curNode = curNode.curList.next.head.next;
                }
            }
            return res;
        }

        private void addList(DoublyLinkedList newList, DoublyLinkedList prevList) {
            prevList.next.prev = newList;
            newList.next = prevList.next;
            prevList.next = newList;
            newList.prev = prevList;
        }

        private void removeList(DoublyLinkedList curList) {
            curList.next.prev = curList.prev;
            curList.prev.next = curList.next;
            curList.next = null;
            curList.prev = null;
        }
    }

    private class DoublyLinkedList {
        DoublyLinkedList prev;
        DoublyLinkedList next;
        Node head;
        Node tail;
        int freq;

        public DoublyLinkedList(int freq) {
            this.freq = freq;
            head = new Node(0);
            tail = new Node(0);
            head.next = tail;
            tail.prev = head;
            // from head/tail can get curList
            head.curList = this;
            tail.curList = this;
        }

        public void addToTail(Node node) {
            tail.prev.next = node;
            node.prev = tail.prev;
            node.next = tail;
            tail.prev = node;
            node.curList = this;
        }

        public void remove(Node node) {
            node.prev.next = node.next;
            node.next.prev = node.prev;
            node.next = null;
            node.prev = null;
            node.curList = null;
        }
    }

    private class Node {
        int key;
        int freq;
        Node prev;
        Node next;
        DoublyLinkedList curList;

        public Node(int key) {
            this.key = key;
            this.freq = 1;
        }
    }
}
```
