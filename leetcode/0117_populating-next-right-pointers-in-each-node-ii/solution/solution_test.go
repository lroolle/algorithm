package solution

import (
	"fmt"
	"reflect"
	"testing"
)

var n5 = &Node{Val: 5}
var n7 = &Node{Val: 7}
var n3 = &Node{Val: 3, Right: n7}

var example1 = &Node{
	Val:   1,
	Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: n5},
	Right: n3,
}

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"Example 1", args{example1}, example1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Connect1")
			tt.args.root.Print()
			got := connect(tt.args.root)
			fmt.Println()
			got.Print()
			fmt.Println()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Connect2")
			tt.args.root.Print()
			got := connect2(tt.args.root)
			fmt.Println()
			got.Print()
			fmt.Println()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect2() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Connect3")
			tt.args.root.Print()
			got := connect3(tt.args.root)
			fmt.Println()
			got.Print()
			fmt.Println()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect2() = %v, want %v", got, tt.want)
			}
		})
	}
}
