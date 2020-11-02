package tree

import (
	"fmt"
	"testing"
)

func TestBinaryNode_Print(t *testing.T) {
	t.Run("Print", func(t *testing.T) {
		root := &BinaryNode{Val: 1}
		root.Left = &BinaryNode{Val: 2}
		root.Left.Left = &BinaryNode{Val: 4}
		root.Left.Right = &BinaryNode{Val: 5}
		root.Left.Right.Left = &BinaryNode{Val: 8}
		root.Left.Right.Right = &BinaryNode{Val: 9}
		root.Right = &BinaryNode{Val: 3}
		root.Right.Left = &BinaryNode{Val: 6}
		root.Right.Right = &BinaryNode{Val: 7}

		fmt.Println("Recursive PrintPreorder: ")
		root.PrintPreorder()
		fmt.Println()
		fmt.Println("Iteration PrintPreorder: ")
		root.IterPrintPreorder()
		fmt.Println()
		fmt.Println("Recursive PrintInorder")
		root.PrintInorder()
		fmt.Println()
		fmt.Println("Iteration Printinorder")
		root.IterPrintInorder()
		fmt.Println()
		fmt.Println("Morris Inorder")
		root.MorrisInorder()
		fmt.Println()
		fmt.Println("Recursive Postorder")
		root.PrintPostorder()
		fmt.Println()
		// fmt.Println("Recursive Postorder")
		// root.IterPrintPostorder()
		fmt.Println()
	})
}

func TestBinaryNode_Insert(t *testing.T) {
	type fields struct {
		Val   int
		Left  *BinaryNode
		Right *BinaryNode
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &BinaryNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			node.Insert(tt.args.val)
		})
	}
}

func TestBinaryNode_PrintPreorder(t *testing.T) {
	type fields struct {
		Val   int
		Left  *BinaryNode
		Right *BinaryNode
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &BinaryNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			node.PrintPreorder()
		})
	}
}

func TestBinaryNode_IterPrintPreorder(t *testing.T) {
	type fields struct {
		Val   int
		Left  *BinaryNode
		Right *BinaryNode
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &BinaryNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			node.IterPrintPreorder()
		})
	}
}

func TestBinaryNode_PrintInorder(t *testing.T) {
	type fields struct {
		Val   int
		Left  *BinaryNode
		Right *BinaryNode
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &BinaryNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			node.PrintInorder()
		})
	}
}

func TestBinaryNode_IterPrintInorder(t *testing.T) {
	type fields struct {
		Val   int
		Left  *BinaryNode
		Right *BinaryNode
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &BinaryNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			node.IterPrintInorder()
		})
	}
}

func TestBinaryNode_MorrisInorder(t *testing.T) {
	type fields struct {
		Val   int
		Left  *BinaryNode
		Right *BinaryNode
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &BinaryNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			node.MorrisInorder()
		})
	}
}

func TestBinaryNode_PrintPostorder(t *testing.T) {
	type fields struct {
		Val   int
		Left  *BinaryNode
		Right *BinaryNode
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &BinaryNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			node.PrintPostorder()
		})
	}
}

func TestBinaryNode_IterPrintPostorder(t *testing.T) {
	type fields struct {
		Val   int
		Left  *BinaryNode
		Right *BinaryNode
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &BinaryNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			node.IterPrintPostorder()
		})
	}
}

func TestBinaryNode_Compare(t *testing.T) {
	var nil1 *BinaryNode
	var nil2 *BinaryNode
	tests := []struct {
		name         string
		root1, root2 *BinaryNode
		want         bool
	}{
		{name: "Nil Tree", root1: nil1, root2: nil2, want: true},
		{name: "Tree 1",
			root1: &BinaryNode{Val: 1, Left: &BinaryNode{Val: 2}},
			root2: &BinaryNode{Val: 1, Left: &BinaryNode{Val: 2}},
			want:  true,
		},
		{name: "Tree 2",
			root1: &BinaryNode{Val: 1, Left: &BinaryNode{Val: 2}},
			root2: &BinaryNode{Val: 1, Left: &BinaryNode{Val: 3}},
			want:  false,
		},
		{name: "Tree 3",
			root1: &BinaryNode{Val: 1, Left: &BinaryNode{Val: 2}},
			root2: &BinaryNode{Val: 1, Right: &BinaryNode{Val: 2}},
			want:  false,
		},
		{name: "Tree 4",
			root1: &BinaryNode{Val: 1, Right: &BinaryNode{Val: 3}},
			root2: &BinaryNode{Val: 1, Right: &BinaryNode{Val: 4}},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root1.Compare(tt.root2); got != tt.want {
				t.Errorf("BinaryNode.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilder(t *testing.T) {
	var null bool
	var nilNode *BinaryNode
	type args struct {
		in0 []interface{}
	}
	tests := []struct {
		name string
		args args
		want *BinaryNode
	}{
		{"Tree Blank", args{[]interface{}{}}, nilNode},
		{"Tree 1",
			args{[]interface{}{5, 4, 7, 3, null, 2, null, -1, null, 9}},
			&BinaryNode{Val: 5,
				Left:  &BinaryNode{Val: 4, Left: &BinaryNode{Val: 3, Left: &BinaryNode{Val: -1}}},
				Right: &BinaryNode{Val: 7, Left: &BinaryNode{Val: 2, Left: &BinaryNode{Val: 9}}},
			},
		},
		{"Tree 2",
			args{[]interface{}{1, 2, 3}},
			&BinaryNode{Val: 1, Left: &BinaryNode{Val: 2}, Right: &BinaryNode{Val: 3}},
		},
		{"Tree 3",
			args{[]interface{}{1, null, 3}},
			&BinaryNode{Val: 1, Right: &BinaryNode{Val: 3}},
		},
		{"Tree 4",
			args{[]interface{}{1, null, 2, 3}},
			&BinaryNode{Val: 1, Right: &BinaryNode{Val: 2, Left: &BinaryNode{Val: 3}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
			root := TestCaseBuilder(tt.args.in0)
			root.PrintAscii()
			if !root.Compare(tt.want) {
				t.Errorf("Builder(%v):\n%v want:\n%v", tt.args.in0, root.AsciiBuilder(), tt.want.AsciiBuilder())
			}
		})
	}
}
