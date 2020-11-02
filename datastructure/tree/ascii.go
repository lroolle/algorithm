// Build AsciiNode for print Ascii tree

package tree

import (
	"fmt"
	"math"
	"strings"
)

type AsciiNode struct {
	Left, Right *AsciiNode
	Label       string
	LabelLen    int
	//-1=I am left, 0=I am root, 1=right
	ParentDir int
	// length of the edge from this node to its children
	EdgeLen int
	Height  int
}

func (anode *AsciiNode) PrintPreorder() {
	if anode == nil {
		return
	}
	fmt.Printf("%s p: %d e: %d h: %d\n", anode.Label, anode.ParentDir, anode.EdgeLen, anode.Height)
	anode.Left.PrintPreorder()
	anode.Right.PrintPreorder()
}

const MaxHeight = 100
const Gap = 2

var PrintNext int
var Lprofile = make([]int, MaxHeight)
var Rprofile = make([]int, MaxHeight)
var AsciiBuilder strings.Builder

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BuildAsciiTree(t *BinaryNode) (node *AsciiNode) {
	if t == nil {
		return nil
	}
	node = new(AsciiNode)
	node.Left = BuildAsciiTree(t.Left)
	node.Right = BuildAsciiTree(t.Right)

	if node.Left != nil {
		node.Left.ParentDir = -1
	}

	if node.Right != nil {
		node.Right.ParentDir = 1
	}

	node.Label = fmt.Sprintf("%d", t.Val)
	node.LabelLen = len(node.Label)

	return node
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func ComuputeLprofile(node *AsciiNode, x, y int) {
	var isLeft int
	if node == nil {
		return
	}
	isLeft = bool2int(node.ParentDir == -1)
	Lprofile[y] = Min(Lprofile[y], x-((node.LabelLen-isLeft)/2))
	if node.Left != nil {
		for i := 0; i < node.EdgeLen; i++ {
			Lprofile[y+i] = Min(Lprofile[y+i], x-i)
		}
	}
	ComuputeLprofile(node.Left, x-node.EdgeLen-1, y+node.EdgeLen+1)
	ComuputeLprofile(node.Right, x+node.EdgeLen+1, y+node.EdgeLen+1)
}

func ComuputeRprofile(node *AsciiNode, x, y int) {
	var notLeft int
	if node == nil {
		return
	}
	notLeft = bool2int(node.ParentDir != -1)
	Rprofile[y] = Max(Rprofile[y], x+((node.LabelLen-notLeft)/2))
	if node.Right != nil {
		for i := 1; i <= node.EdgeLen; i++ {
			Rprofile[y+i] = Max(Rprofile[y+i], x+i)
		}
	}
	ComuputeRprofile(node.Left, x-node.EdgeLen-1, y+node.EdgeLen+1)
	ComuputeRprofile(node.Right, x+node.EdgeLen+1, y+node.EdgeLen+1)
}

func FillEdgeLen(node *AsciiNode) {
	var h, hmin, delta int
	if node == nil {
		return
	}
	FillEdgeLen(node.Left)
	FillEdgeLen(node.Right)

	if node.Right == nil && node.Left == nil {
		node.EdgeLen = 0
	} else {
		if node.Left != nil {
			for i := 0; i < node.Left.Height; i++ {
				Rprofile[i] = -math.MaxInt32
			}
			ComuputeRprofile(node.Left, 0, 0)
			hmin = node.Left.Height
		} else {
			hmin = 0
		}
		if node.Right != nil {
			for i := 0; i < node.Right.Height; i++ {
				Lprofile[i] = math.MaxInt32
			}
			ComuputeLprofile(node.Right, 0, 0)
			hmin = Min(node.Right.Height, hmin)
		} else {
			hmin = 0
		}

		delta = 4
		for i := 0; i < hmin; i++ {
			delta = Max(delta, Gap+1+Rprofile[i]-Lprofile[i])
		}

		if ((node.Left != nil && node.Left.Height == 1) ||
			(node.Right != nil && node.Right.Height == 1)) && delta > 4 {
			delta--
		}
		node.EdgeLen = ((delta + 1) / 2) - 1
	}

	h = 1
	if node.Left != nil {
		h = Max(node.Left.Height+node.EdgeLen+1, h)
	}
	if node.Right != nil {
		h = Max(node.Right.Height+node.EdgeLen+1, h)
	}
	node.Height = h
}

// Print the given level of tree
func PrintLevel(node *AsciiNode, x, level int) {
	var i, isLeft int
	if node == nil {
		return
	}

	isLeft = bool2int(node.ParentDir == -1)
	if level == 0 {
		for i = 0; i < x-PrintNext-((node.LabelLen-isLeft)/2); i++ {
			fmt.Fprintf(&AsciiBuilder, "%s", " ")
		}
		PrintNext += i
		fmt.Fprintf(&AsciiBuilder, "%s", node.Label)
		PrintNext += node.LabelLen
	} else if node.EdgeLen >= level {
		if node.Left != nil {
			for i = 0; i < x-PrintNext-level; i++ {
				fmt.Fprintf(&AsciiBuilder, "%s", " ")
			}
			PrintNext += i
			fmt.Fprintf(&AsciiBuilder, "%s", "/")
			PrintNext++
		}
		if node.Right != nil {
			for i = 0; i < x-PrintNext+level; i++ {
				fmt.Fprintf(&AsciiBuilder, "%s", " ")
			}
			PrintNext += i
			fmt.Fprintf(&AsciiBuilder, "%s", "\\")
			PrintNext++
		}
	} else {
		PrintLevel(node.Left, x-node.EdgeLen-1, level-node.EdgeLen-1)
		PrintLevel(node.Right, x+node.EdgeLen+1, level-node.EdgeLen-1)
	}
}

// Return a ascii string representation of the tree
func (t *BinaryNode) AsciiBuilder() string {
	var proot *AsciiNode
	var xmin int
	if t == nil {
		return AsciiBuilder.String()
	}

	proot = BuildAsciiTree(t)
	FillEdgeLen(proot)

	for i := 0; i < proot.Height; i++ {
		Lprofile[i] = math.MaxInt32
	}
	ComuputeLprofile(proot, 0, 0)
	xmin = 0
	for i := 0; i < proot.Height; i++ {
		xmin = Min(xmin, Lprofile[i])
	}

	if proot.Height > MaxHeight {
		fmt.Printf("The tree is too fucking high than %d! how high do you want to be?\n",
			MaxHeight)
	}
	for i := 0; i < proot.Height; i++ {
		PrintNext = 0
		PrintLevel(proot, -xmin, i)
		fmt.Fprintf(&AsciiBuilder, "%s", "\n")
	}
	s := AsciiBuilder.String()
	AsciiBuilder.Reset()
	return s
}

// Print the ascii string representation of the tree to stdout
func (t *BinaryNode) PrintAscii() {
	fmt.Println(t.AsciiBuilder())
}
