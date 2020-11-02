"+++
title = "0206. Reverse Linked List 动画演示+多种解法 206. 反转链表 "
author = ["wang_ni_ma"]
date = 2019-10-22T01:39:33+08:00
tags = ["Leetcode", "Algorithms", "Python", "Java", "LinkedList", "Recursion"]
categories = ["leetcode"]
draft = false
+++

# 动画演示+多种解法 206. 反转链表

> [0206. Reverse Linked List](https://leetcode-cn.com/problems/reverse-linked-list/)
> [动画演示+多种解法 206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/solution/dong-hua-yan-shi-206-fan-zhuan-lian-biao-by-user74/) by [wang_ni_ma](https://leetcode-cn.com/u/wang_ni_ma/)


### 利用外部空间
这种方式很简单，先申请一个动态扩容的数组或者容器，比如 ArrayList 这样的。
然后不断遍历链表，将链表中的元素添加到这个容器中。
再利用容器自身的 API，反转整个容器，这样就达到反转的效果了。
最后同时遍历容器和链表，将链表中的值改为容器中的值。
因为此时容器的值是：
```
5 4 3 2 1
```
链表按这个顺序重新被设置一边，就达到要求啦。   
当然你可以可以再新创建 N 个节点，然后再返回，这样也可以达到目的。   
这种方式很简单，但你在面试中这么做的话，面试官 100% 会追问是否有更优的方式，比如不用外部空间。所以我就不做代码和动画演示了，下面来看看如何用 $O(1)$ 空间复杂度来实现这道题。   
   
### 双指针迭代
我们可以申请两个指针，第一个指针叫 pre，最初是指向 null 的。   
第二个指针 cur 指向 head，然后不断遍历 cur。   
每次迭代到 cur，都将 cur 的 next 指向 pre，然后 pre 和 cur 前进一位。   
都迭代完了(cur 变成 null 了)，pre 就是最后一个节点了。   
动画演示如下：   
![迭代.gif](https://pic.leetcode-cn.com/7d8712af4fbb870537607b1dd95d66c248eb178db4319919c32d9304ee85b602-%E8%BF%AD%E4%BB%A3.gif)

动画演示中其实省略了一个```tmp```变量，这个```tmp```变量会将```cur```的下一个节点保存起来，考虑到一张动画放太多变量会很混乱，所以我就没加了，具体详细执行过程，请点击下面的幻灯片查看。

<![1.JPG](https://pic.leetcode-cn.com/8f03740e93c8a0ecff54a57a25a710ab82574faf99b3efb8ddd2b6eea1d78d49-1.JPG),![2.JPG](https://pic.leetcode-cn.com/aebc5e3ef0a6b942473a7c61f4de22e268cbfbe549ff1a6ad5d3eef7c667f875-2.JPG),![3.JPG](https://pic.leetcode-cn.com/4c160db670c00f7467c1dae3b22ac117b84c0ba00690ca30708fabcdd5ff35b1-3.JPG),![4.JPG](https://pic.leetcode-cn.com/95da97748c8a006b5b827d13154bd2202bed5c2b9d779bea6467473d16fd7d66-4.JPG),![5.JPG](https://pic.leetcode-cn.com/bbc75893b050a8c482178029bd239bd75c75e6a58a4a16fb9270b13044e4ddf3-5.JPG),![6.JPG](https://pic.leetcode-cn.com/9eda023eab4ecc76af90211626f77b695a4874b451084f8e2545717347815006-6.JPG),![7.JPG](https://pic.leetcode-cn.com/1829f246e991d3e27aa3dce9b18ab88282a1bd34a5363969eb5730fe74cf8fbf-7.JPG),![8.JPG](https://pic.leetcode-cn.com/3380338ba4b563c37199b07058d35fde9f9600d3112ef0b074e5986c2f4510e1-8.JPG),![9.JPG](https://pic.leetcode-cn.com/db831df3a7db0fc405a4dc8293bfc92ba6bba6cb1719613a4e8ab7afdf2a53dd-9.JPG),![10.JPG](https://pic.leetcode-cn.com/db7d85400d03237909fc904d52b71cf721dee82af47686756400fdf4d1f07493-10.JPG),![11.JPG](https://pic.leetcode-cn.com/217edaea46af7111109f3dfb4846aad5e307935cec6e48ca42e0059d36792c25-11.JPG),![12.JPG](https://pic.leetcode-cn.com/549d1ed958285f378aea9a0db3a603253c83aff5439a76c0b8c35f2e237b9f4b-12.JPG),![13.JPG](https://pic.leetcode-cn.com/9683a6f2e465086ca1a9bcdb6b4cce9cdcc86da3269e1009867368923c76363d-13.JPG),![14.JPG](https://pic.leetcode-cn.com/9ef5549e7ff4b4748e0ff295ee37d3e172e7b910b397aa8718110b4779a9f689-14.JPG),![15.JPG](https://pic.leetcode-cn.com/074dcf7c17ba256aedbf4ba10d101ce93ed5425a6a1aa6d296a79be7bd65fc30-15.JPG),![16.JPG](https://pic.leetcode-cn.com/f6d200f6b154e8b4e9cf86c26e90c0e71af048fd4dac2f9a35eb8c9dc93bfa15-16.JPG),![17.JPG](https://pic.leetcode-cn.com/3b5504463012d23800d78e50948d1623cf405e32072cd1c630c71636a318d002-17.JPG),![18.JPG](https://pic.leetcode-cn.com/8c474aaee63b49b4cf981637039f65a41d70979a79fec22257576cfbcfb8c4c6-18.JPG),![19.JPG](https://pic.leetcode-cn.com/5eaed6c73d6dec3c3dc7728fda178711ff1a6a558ddf8f2a1946adf22525a44c-19.JPG)>


代码实现：
```Java
class Solution {
	public ListNode reverseList(ListNode head) {
		//申请节点，pre和 cur，pre指向null
		ListNode pre = null;
		ListNode cur = head;
		ListNode tmp = null;
		while(cur!=null) {
			//记录当前节点的下一个节点
			tmp = cur.next;
			//然后将当前节点指向pre
			cur.next = pre;
			//pre和cur节点都前进一位
			pre = cur;
			cur = tmp;
		}
		return pre;
	}
}
```
```Python
class Solution(object):
	def reverseList(self, head):
		"""
		:type head: ListNode
		:rtype: ListNode
		"""
		# 申请两个节点，pre和 cur，pre指向None
		pre = None
		cur = head
		# 遍历链表，while循环里面的内容其实可以写成一行
		# 这里只做演示，就不搞那么骚气的写法了
		while cur:
			# 记录当前节点的下一个节点
			tmp = cur.next
			# 然后将当前节点指向pre
			cur.next = pre
			# pre和cur节点都前进一位
			pre = cur
			cur = tmp
		return pre	
```
   
   		
### 递归解法
这题有个很骚气的递归解法，递归解法很不好理解，这里最好配合代码和动画一起理解。
递归的两个条件：
1. 终止条件是当前节点或者下一个节点==null
2. 在函数内部，改变节点的指向，也就是 head 的下一个节点指向 head 递归函数那句
```
head.next.next = head
```
很不好理解，其实就是 head 的下一个节点指向head。   
递归函数中每次返回的 cur 其实只最后一个节点，在递归函数内部，改变的是当前节点的指向。   
动画演示如下：   
![递归.gif](https://pic.leetcode-cn.com/dacd1bf55dec5c8b38d0904f26e472e2024fc8bee4ea46e3aa676f340ba1eb9d-%E9%80%92%E5%BD%92.gif)
   
   
#### 幻灯片演示 
   
感谢[@zhuuuu-2](/u/zhuuuu-2/)的建议，递归的解法光看动画比较容易理解，但真到了代码层面理解起来可能会有些困难，我补充了下递归调用的详细执行过程。
   
以```1->2->3->4->5```这个链表为例，整个递归调用的执行过程，对应到代码层面(用java做示范)是怎么执行的，以及递归的调用栈都列出来了，请点击下面的幻灯片查看吧。

<![幻灯片1.jpg](https://pic.leetcode-cn.com/c573939fa872edd4da5ac39e703bccdc65c32849a95ba7637dea5cfdfb6eced6-%E5%B9%BB%E7%81%AF%E7%89%871.jpg),![幻灯片2.jpg](https://pic.leetcode-cn.com/bfa449ed16ecfc905f8ef8e9b049ed397b608f30dad84d37b2005e19797b1d49-%E5%B9%BB%E7%81%AF%E7%89%872.jpg),![幻灯片3.jpg](https://pic.leetcode-cn.com/c790acd5bd2ebf285ef2f687d76c9a48f2e32418acdb32909cc76cc069501ac7-%E5%B9%BB%E7%81%AF%E7%89%873.jpg),![幻灯片4.jpg](https://pic.leetcode-cn.com/d078fc5c224964b468210b205838ecac26efde1a0bd6fade6f3a758a87b82068-%E5%B9%BB%E7%81%AF%E7%89%874.jpg),![幻灯片5.jpg](https://pic.leetcode-cn.com/cd9c77c8c7873c4aa39bbb7a185b729151516021cee7009575c258fb82f77383-%E5%B9%BB%E7%81%AF%E7%89%875.jpg),![幻灯片6.jpg](https://pic.leetcode-cn.com/f89c9e095d414b8366dd7d490508ba9c99d6c93953b4429168af262169f18e83-%E5%B9%BB%E7%81%AF%E7%89%876.jpg),![幻灯片7.jpg](https://pic.leetcode-cn.com/485dbe9ad44ab7b05e01c46a1bc1718187a01f8a9fe8331f497e7011f9508b57-%E5%B9%BB%E7%81%AF%E7%89%877.jpg),![幻灯片8.jpg](https://pic.leetcode-cn.com/04550182527e6570d9d04f2eeae330848b83d3ff13b23ea4153410210586dc85-%E5%B9%BB%E7%81%AF%E7%89%878.jpg),![幻灯片9.jpg](https://pic.leetcode-cn.com/b167f5ee08eafab82d680cf782898b649dd2a4453cd122ad690e83e35c387a11-%E5%B9%BB%E7%81%AF%E7%89%879.jpg),![幻灯片10.jpg](https://pic.leetcode-cn.com/950f6edfb553cbeec65ce3c9679dd8d3401fd1837ad9eb16989217bf83f30e58-%E5%B9%BB%E7%81%AF%E7%89%8710.jpg),![幻灯片11.jpg](https://pic.leetcode-cn.com/dcd5bef9b8ff3de98e9533da7da3a8411643bbef9bc5e1a576085b4403197649-%E5%B9%BB%E7%81%AF%E7%89%8711.jpg),![幻灯片12.jpg](https://pic.leetcode-cn.com/1d6aaab3d9a42c20420fb6087e520ea05bd6d0789213f228481e55891b847b1e-%E5%B9%BB%E7%81%AF%E7%89%8712.jpg),![幻灯片13.jpg](https://pic.leetcode-cn.com/74bb9de62d3bb36a9c174d7702c7116974838b46eb96933e2921c054a34fa5d9-%E5%B9%BB%E7%81%AF%E7%89%8713.jpg),![幻灯片14.jpg](https://pic.leetcode-cn.com/e2bea51573eb5a763b5b9c0873887ced8932db95bc6bf4ff5befb363c8960984-%E5%B9%BB%E7%81%AF%E7%89%8714.jpg),![幻灯片15.jpg](https://pic.leetcode-cn.com/95db06ce86448f3265b0f3178baf7372915601e4b367fcde54ff5cf509532bc8-%E5%B9%BB%E7%81%AF%E7%89%8715.jpg),![幻灯片16.jpg](https://pic.leetcode-cn.com/fc96f0786c05a3a2ae0987757b4f569b3b18e59da0b5cb3afb063b9fc0ace069-%E5%B9%BB%E7%81%AF%E7%89%8716.jpg),![幻灯片17.jpg](https://pic.leetcode-cn.com/6a1a4b01a7dd20e3a5a560c482b514c7bd3d9462d8f297d373232229995e6681-%E5%B9%BB%E7%81%AF%E7%89%8717.jpg),![幻灯片18.jpg](https://pic.leetcode-cn.com/1318eda47ab9b0d505b98b5c11503012053a1bbcf754ce0c0145af14de5d41cd-%E5%B9%BB%E7%81%AF%E7%89%8718.jpg),![幻灯片19.jpg](https://pic.leetcode-cn.com/a0669e29700f80938c95faf9a5cc839d316a31b6f6613da88bcdd31636897d96-%E5%B9%BB%E7%81%AF%E7%89%8719.jpg),![幻灯片20.jpg](https://pic.leetcode-cn.com/30db7c04d56193840e53c4dc6c6f095ebe80b2b7a71107ca5ed19da499467914-%E5%B9%BB%E7%81%AF%E7%89%8720.jpg),![幻灯片21.jpg](https://pic.leetcode-cn.com/052984cbdec432f46e771e0c7f214b7303941e4a11a7f59ee53a519e6b400adc-%E5%B9%BB%E7%81%AF%E7%89%8721.jpg),![幻灯片22.jpg](https://pic.leetcode-cn.com/e096fac734909bc4b1ff55d8446332b5d7e67ec238266b0427ad0d69bc50174b-%E5%B9%BB%E7%81%AF%E7%89%8722.jpg),![幻灯片23.jpg](https://pic.leetcode-cn.com/4e0975fd8d898364d24177684967fdf72d9476e64341c375ac2b5df7130b2f0a-%E5%B9%BB%E7%81%AF%E7%89%8723.jpg),![幻灯片24.jpg](https://pic.leetcode-cn.com/5625ed08dbaa84719f0b4b631fa6c4c7f08c1c5cf18646dbab2f65d410e99e8f-%E5%B9%BB%E7%81%AF%E7%89%8724.jpg)>


   
代码实现：
```Java
class Solution {
	public ListNode reverseList(ListNode head) {
		//递归终止条件是当前为空，或者下一个节点为空
		if(head==null || head.next==null) {
			return head;
		}
		//这里的cur就是最后一个节点
		ListNode cur = reverseList(head.next);
		//这里请配合动画演示理解
		//如果链表是 1->2->3->4->5，那么此时的cur就是5
		//而head是4，head的下一个是5，下下一个是空
		//所以head.next.next 就是5->4
		head.next.next = head;
		//防止链表循环，需要将head.next设置为空
		head.next = null;
		//每层递归函数都返回cur，也就是最后一个节点
		return cur;
	}
}
```
```Python
class Solution(object):
	def reverseList(self, head):
		"""
		:type head: ListNode
		:rtype: ListNode
		"""
		# 递归终止条件是当前为空，或者下一个节点为空
		if(head==None or head.next==None):
			return head
		# 这里的cur就是最后一个节点
		cur = self.reverseList(head.next)
		# 这里请配合动画演示理解
		# 如果链表是 1->2->3->4->5，那么此时的cur就是5
		# 而head是4，head的下一个是5，下下一个是空
		# 所以head.next.next 就是5->4
		head.next.next = head
		# 防止链表循环，需要将head.next设置为空
		head.next = None
		# 每层递归函数都返回cur，也就是最后一个节点
		return cur
```
(全文完)
    
**欢迎关注 👉👉👉  [【公众号】](https://pic.leetcode-cn.com/cdc143ca40e2cc2dc86ddb23feadb02464bd44a180d2a2d969a44a368ba13d70-206.jpg) 👈👈👈**   
   
**如果能再点个赞👍👍 就更感激啦💓💓**


