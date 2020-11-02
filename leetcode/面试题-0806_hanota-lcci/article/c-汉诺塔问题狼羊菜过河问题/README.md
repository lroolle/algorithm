"+++
title = "面试题 08.06. Hanota LCCI C++ 汉诺塔问题/狼羊菜过河问题 "
author = ["yizhe-shi"]
date = 2020-09-15T07:24:16+08:00
tags = ["Leetcode", "Algorithms", "cpp"]
categories = ["leetcode"]
draft = false
+++

# C++ 汉诺塔问题/狼羊菜过河问题

> [面试题 08.06. Hanota LCCI](https://leetcode-cn.com/problems/hanota-lcci/)
> [C++ 汉诺塔问题/狼羊菜过河问题](https://leetcode-cn.com/problems/hanota-lcci/solution/c-jian-dan-jie-fa-by-yizhe-shi-14/) by [yizhe-shi](https://leetcode-cn.com/u/yizhe-shi/)

### 娱乐解法
```cpp
class Solution {
public:
    void hanota(vector<int>& A, vector<int>& B, vector<int>& C) {
        swap(A, C);
    }
};
```
### 正经解法
```cpp
class Solution {
public:
    void move(int n, vector<int>& A, vector<int>& B, vector<int>& C){
        if(n == 1){
            C.push_back(A.back());
            A.pop_back();
        } else{
            //n - 1 个盘子通过 C 从 A 移向 B
            move(n - 1, A, C, B);
            //最底下的盘子移向 C
            move(1, A, B, C);
            //n - 1 个盘子通过 A 从 B 移向 C
            move(n - 1, B, A, C);
        }
    }
    void hanota(vector<int>& A, vector<int>& B, vector<int>& C) {
        // swap(A, C);
        int n = A.size();
        //n 个盘子通过 B 从 A 移向 C
        move(n, A, B, C);
    }
};
```
### 打印路径解法
```cpp
class Solution {
public:
    void move(int n, vector<int>& A, vector<int>& B, vector<int>& C, char a, char b, char c){
        if(n == 1){
            cout << a << " -> " << c << endl; //显示出来路径
            C.push_back(A.back());
            A.pop_back();
        } else{
            //n - 1 个盘子通过 C 从 A 移向 B
            move(n - 1, A, C, B, a, c, b);
            //A 中的盘子移向 C
            move(1, A, B, C, a, b, c);
            //n - 1 个盘子通过 A 从 B 移向 C
            move(n - 1, B, A, C, b, a, c);
        }
    }
    void hanota(vector<int>& A, vector<int>& B, vector<int>& C) {
        // swap(A, C);
        int n = A.size();
        //n 个盘子通过 B 从 A 移向 C
        move(n, A, B, C, 'A', 'B', 'C');
    }
};
```
### 扩展
既然来了这么经典的汉诺塔问题了，那不提一下 `狼、羊、菜和农夫过河问题` 就不合适了
大致题意是：农夫需要把狼、羊、菜和自己运到河对岸去，只有农夫能够划船，而且船比较小，除农夫之外每次只能运一种东西，还有一个棘手问题，就是如果没有农夫看着，羊会偷吃菜，狼会吃羊。请考虑一种方法，让农夫能够安全地安排这些东西和他自己过河。
经典吧，这道题应该都或多或少地听说过，感兴趣大家可以去查一下，leetcode没引入这个问题太可惜了。

对于狼或者羊等等他们都有两种状态：过河或者没过河，我们分别用 1 或 0 表示

```C++
#include<iostream>
#include<vector>
#include<string>
#include<queue>
#include<stack>
using namespace std;

struct Node {
    string status;//菜、狼、羊、人的状态，A岸=0,B岸=1，例如初始状态“0000”
    int way;//用那种方式到达此状态,-1是初始状态，0-3分别表示带菜、带狼、带羊、不带
    int father;
    Node(string s,int w,int f) : status(s), way(w), father(f) {}
};

vector<Node> nodes;
string S[10] = { "0000","1000","0100","0010","1100","1111","0111","1011","1101","0011" };
const string S_NULL = "";

string CrossRiver(string status,int way) {
    if (way == 3) {
        status[3] = '1' - status[3] + '0';
    }
    else if (status[way] != status[3]) return S_NULL;//要带的和自己不在一边肯定不能带
    else status[way] = status[3] = '1' - status[3] + '0';
    //如果现在是合法的状态就可以，就是羊和菜、狼和羊不能单独在一起
    for (int i = 0; i < 10; i++) {
        if (status == S[i]) return status;
    }
    return S_NULL;
}

int BFS() {         
    nodes.push_back(Node("0000", -1, -1));
    int idx = 0;
    while (true) {
        Node node = nodes[idx];
        string status = node.status;
        //都到另一边了
        if (status == "1111") return idx;
        //选择三种方式
        for (int i = 0; i < 4; i++) {
            string t = status;
            if ((t = CrossRiver(status, i)) != S_NULL) {
                nodes.push_back(Node(t, i, idx));
            }
        }
        idx ++;
    }
}

void Print(stack<int> s) {
    while (!s.empty()) {
        int way, status;
        way = nodes[s.top()].way;
        status = nodes[s.top()].status[3] - '0';
        s.pop();
        switch (way) {
            case 0: printf("vegetable_"); break;
            case 1: printf("wolf_"); break;
            case 2: printf("sheep_"); break;
            case 3: printf("nothing_"); break;
        }
        if (status == 1) {
            printf("go\n");
        }
        else {
            printf("come\n");
        }
    }
    printf("succeed\n");
}

int main(){
    int ret = BFS();
    stack<int> s;
    s.push(ret);
    //ret 是成功时的最后一个结点
    while (nodes[ret].father != 0) {
        ret = nodes[ret].father;
        s.push(ret);
    }
    Print(s);
    return 0;
}
```
