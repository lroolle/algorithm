package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildAsciiTree(t *testing.T) {
	type args struct {
		t *BinaryNode
	}
	tests := []struct {
		name     string
		args     args
		wantNode *AsciiNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNode := BuildAsciiTree(tt.args.t); !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("BuildAsciiTree() = %v, want %v", gotNode, tt.wantNode)
			}
		})
	}
}

func Test_bool2int(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bool2int(tt.args.b); got != tt.want {
				t.Errorf("bool2int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComuputeLprofile(t *testing.T) {
	type args struct {
		node *AsciiNode
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ComuputeLprofile(tt.args.node, tt.args.x, tt.args.y)
		})
	}
}

func TestComuputeRprofile(t *testing.T) {
	type args struct {
		node *AsciiNode
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ComuputeRprofile(tt.args.node, tt.args.x, tt.args.y)
		})
	}
}

func TestFillEdgeLen(t *testing.T) {
	type args struct {
		node *AsciiNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FillEdgeLen(tt.args.node)
		})
	}
}

func TestPrintLevel(t *testing.T) {
	type args struct {
		node  *AsciiNode
		x     int
		level int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintLevel(tt.args.node, tt.args.x, tt.args.level)
		})
	}
}

var tree1 = &BinaryNode{Val: 2, Left: &BinaryNode{Val: 1}, Right: &BinaryNode{Val: 3}}
var tree2 = &BinaryNode{
	Val:   1111,
	Left:  &BinaryNode{Val: 11, Left: &BinaryNode{Val: 1}, Right: &BinaryNode{Val: 111}},
	Right: &BinaryNode{Val: 111111, Left: &BinaryNode{Val: 11111}, Right: &BinaryNode{Val: 1111111}},
}
var tree3 = &BinaryNode{Val: 2, Right: &BinaryNode{Val: 3}}

func TestBinaryNode_PrintAscii(t *testing.T) {
	tests := []struct {
		name string
		root *BinaryNode
	}{
		{"Tree1", tree1},
		{"Tree2", tree2},
		{"Tree2", tree3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
			tt.root.PrintAscii()
		})
	}
}
