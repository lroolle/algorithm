package solution

import "fmt"

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (node *Node) Print() {
	if node == nil {
		return
	}
	if node.Next != nil {
		fmt.Printf("%d(->%d) ", node.Val, node.Next.Val)
	} else {
		fmt.Printf("%d ", node.Val)
	}
	node.Left.Print()
	node.Right.Print()
}

func connect(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, last *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root
}

// BFS 层次遍历
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for len(q) > 0 {
		tmp := q
		q = []*Node{}
		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}

type tuple struct {
	node  *Node
	level int
}

func connect3(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*tuple{{node: root, level: 0}}
	var prev *tuple

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.node.Left != nil {
			queue = append(queue, &tuple{cur.node.Left, cur.level + 1})
		}
		if cur.node.Right != nil {
			queue = append(queue, &tuple{cur.node.Right, cur.level + 1})
		}
		if prev != nil && prev.level == cur.level {
			prev.node.Next = cur.node
		}
		prev = cur
	}
	return root
}
