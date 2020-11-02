"+++
title = "0117. Populating Next Right Pointers in Each Node II 纯C，利用当前层的 next 指针构造下层的 next，【117.填充每个节点的下一个右侧节点II】【思路清晰，代码易读】 "
author = ["r0vHWU5AdJ"]
date = 2020-05-28T06:36:03+08:00
tags = ["Leetcode", "Algorithms", "C"]
categories = ["leetcode"]
draft = false
+++

# 纯C，利用当前层的 next 指针构造下层的 next，【117.填充每个节点的下一个右侧节点II】【思路清晰，代码易读】

> [0117. Populating Next Right Pointers in Each Node II](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/)
> [纯C，利用当前层的 next 指针构造下层的 next，【117.填充每个节点的下一个右侧节点II】【思路清晰，代码易读】](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/solution/chun-cli-yong-dang-qian-ceng-de-next-zhi-zhen-go-2/) by [r0vHWU5AdJ](https://leetcode-cn.com/u/r0vhwu5adj/)

### 解题思路
方法二：利用当前层的 next 指针构造下层的 next
1,pHeadNode 指向每层的第一个节点 用于判断是否还有下一层
2,pCurNode 用于遍历当前层所有节点
3,pLastNode 指向当前节点同一层左侧节点

方法一：二叉树的层序遍历
1,构造节点队列，实现向队列中增加元素，和从队列中取出元素
2,利用队列先进先出，实现二叉树的层序遍历
### 代码

```c
/**
 * Definition for a Node.
 * struct Node {
 *     int val;
 *     struct Node *left;
 *     struct Node *right;
 *     struct Node *next;
 * };
 */

//方法二：利用当前层的 next 指针构造下层的 next
//1,pHeadNode 指向每层的第一个节点 用于判断是否还有下一层
//2,pCurNode 用于遍历当前层所有节点
//3,pLastNode 指向当前节点同一层左侧节点

struct Node* connect(struct Node* root) {
    struct Node*    pHeadNode   = NULL;     //指向每层的第一个节点
    struct Node*    pCurNode    = NULL;     //遍历每一层的所有节点
    struct Node*    pLastNode   = NULL;     //指向当前节点同一层左侧节点

    if(NULL == root) return NULL;
    pHeadNode = root;

    //1,二叉树层级循环，当 NULL==pHeadNode 时表示二叉树的最后一层了
    while(NULL != pHeadNode) 
    {
        pCurNode = pHeadNode;
        pHeadNode = NULL;
        pLastNode = NULL;

        //2,遍历同一层的所有节点
        while(NULL != pCurNode)
        {
            if((NULL != pCurNode->left) && (NULL != pCurNode->right))
            {
                //当前节点左右支都存在的情况处理
                if(NULL == pHeadNode) pHeadNode = pCurNode->left;
                
                if(NULL != pLastNode) pLastNode->next = pCurNode->left;
                pCurNode->left->next = pCurNode->right;
                pLastNode = pCurNode->right;
            }
            else if(NULL != pCurNode->left)
            {
                //当前节点只有左支处理
                if(NULL == pHeadNode) pHeadNode = pCurNode->left;

                if(NULL != pLastNode) pLastNode->next = pCurNode->left;
                pLastNode = pCurNode->left;
            }
            else if(NULL != pCurNode->right)
            {
                //当前节点只有右支处理
                if(NULL == pHeadNode) pHeadNode = pCurNode->right;

                if(NULL != pLastNode) pLastNode->next = pCurNode->right;
                pLastNode = pCurNode->right;
            }

            //当前层下一个节点
            pCurNode = pCurNode->next;
        }
    }

    return root;
}
/*
//方法一：二叉树的层序遍历
//1,构造节点队列，实现向队列中增加元素，和从队列中取出元素
//2,利用队列先进先出，实现二叉树的层序遍历

//声明队列节点结构
struct QueueNode {
    struct Node* pTreeNode;     //队列元素：二叉树节点指针
    struct NodeQueue* pNext;    //队列元素：下一个节点指针
};

//声明队列结构
struct NodeQueue {
    int     iNum;                   //队列元素个数
    struct QueueNode*   pHead;      //队列头指针
    struct QueueNode*   pTail;      //队列尾指针
};

//函数一：向队列中增加元素
bool Push_Queue(struct NodeQueue* pQueue, struct Node* pNode){
    struct QueueNode*   pQueueNode = NULL;

    if(NULL == pQueue) return false;

    pQueueNode = (struct QueueNode*)malloc(sizeof(struct QueueNode));
    pQueueNode->pTreeNode = pNode;
    pQueueNode->pNext = NULL;

    if(0 == pQueue->iNum)
    {
        pQueue->pHead = pQueueNode;
        pQueue->pTail = pQueueNode;
        pQueue->iNum += 1;
    }
    else
    {
        pQueue->pTail->pNext = pQueueNode;
        pQueue->pTail = pQueueNode;
        pQueue->iNum += 1;
    }

    return true;
}

//函数二：从队列中取出元素
struct Node* Pop_Queue(struct NodeQueue* pQueue){
    struct Node*        pRet    = NULL;
    struct QueueNode*   pTmp    = NULL;

    if((NULL == pQueue) || (0 == pQueue->iNum)) return NULL;

    pRet = pQueue->pHead->pTreeNode;

    pQueue->iNum -= 1;
    if(0 == pQueue->iNum)
    {
        free(pQueue->pHead);
        pQueue->pHead = NULL;
        pQueue->pTail = NULL;
    }
    else
    {
        pTmp = pQueue->pHead->pNext;
        free(pQueue->pHead);
        pQueue->pHead = pTmp;
    }
    
    return pRet;
}

struct Node* connect(struct Node* root) {
    struct Node*        pCurNode    = NULL;
    struct Node*        pLastNode   = NULL;
	struct NodeQueue    strNodeQueue;

    if(NULL == root) return NULL;
    memset(&strNodeQueue, 0x00, sizeof(struct NodeQueue));

    //1,将根节点压入队列，并加入 NULL 作为层的分隔
    Push_Queue(&strNodeQueue, root);
    Push_Queue(&strNodeQueue, NULL);

    while(strNodeQueue.iNum != 0)
    {
        //2,从队列中取出节点，层序遍历
        pCurNode = Pop_Queue(&strNodeQueue);

        if(NULL == pCurNode)
        {
            //3,一层结束，需要将 NULL 压入队列作为每层的分隔
            if(strNodeQueue.iNum != 0)
            {
                Push_Queue(&strNodeQueue, NULL);
                pLastNode = NULL;
            }
        }
        else
        {
            //4,将每个节点的左右支分别压入队列
            if(NULL != pCurNode->left)
            {
                Push_Queue(&strNodeQueue, pCurNode->left);
            }
            if(NULL != pCurNode->right)
            {
                Push_Queue(&strNodeQueue, pCurNode->right);
            }

            //5,每个节点的右侧节点指针处理
            if(NULL != pLastNode)
            {
                pLastNode->next = pCurNode;
            }
            pLastNode = pCurNode;
        }
    }
    return root;
}
*/
```