package list

import (
	"fmt"
	"testing"
)

func TestSingleListNode_Print(t *testing.T) {
	type fields struct {
		Val  int
		Next *SingleListNode
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"2a", fields{Val: 2, Next: &SingleListNode{Val: "a"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &SingleListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			fmt.Println("SingleListNode Print")
			node.Print()
		})
	}
}

func TestDoubleListNode_Print(t *testing.T) {
	var node3 = &DoubleListNode{Val: 3}
	var node2 = &DoubleListNode{Val: 2}
	var node1 = &DoubleListNode{Val: "a"}
	node3.Next = node2
	node2.Prev = node3
	node2.Next = node1
	node1.Prev = node2

	tests := []struct {
		name string
		node *DoubleListNode
	}{
		{"32a", node1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := tt.node
			fmt.Println("DoubleListDode Print")
			node.Print()
		})
	}
}

func TestSingleListNode_Compare(t *testing.T) {
	tests := []struct {
		name  string
		node1 *SingleListNode
		node2 *SingleListNode
		want  bool
	}{
		{
			"Example 1",
			&SingleListNode{Val: 1, Next: &SingleListNode{Val: 2}},
			&SingleListNode{Val: 2, Next: &SingleListNode{Val: 2}},
			false,
		},
		{
			"Example 2",
			&SingleListNode{Val: 1, Next: &SingleListNode{Val: 2}},
			&SingleListNode{Val: 2},
			false,
		},
		{
			"Example 3",
			&SingleListNode{},
			&SingleListNode{Val: 2},
			false,
		},
		{
			"Example 4",
			&SingleListNode{},
			&SingleListNode{},
			true,
		},
		{
			"Example 5",
			&SingleListNode{Val: 1},
			&SingleListNode{Val: 1},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node1.Compare(tt.node2); got != tt.want {
				t.Errorf("SingleListNode.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleListNode_Insert(t *testing.T) {
	tests := []struct {
		name string
		root *SingleListNode
		args []int
		want *SingleListNode
	}{
		{"Nil 1", &SingleListNode{}, []int{1, 2, 3},
			&SingleListNode{Val: 1, Next: &SingleListNode{Val: 2, Next: &SingleListNode{Val: 3}}},
		},
		{"Example 1", &SingleListNode{Val: 4, Next: &SingleListNode{Val: 5}}, []int{1, 2, 3},
			&SingleListNode{Val: 1, Next: &SingleListNode{Val: 2, Next: &SingleListNode{Val: 3, Next: &SingleListNode{Val: 4, Next: &SingleListNode{Val: 5}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got = tt.root
			for _, val := range tt.args {
				got = got.Insert(val)
			}
			if !got.Compare(tt.want) {
				t.Errorf("%+v: Insert(%v) = %v, want %v", tt.root, tt.args, got.Sprint(), tt.want.Sprint())
			}
		})
	}
}
