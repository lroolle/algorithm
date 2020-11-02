package solution

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	var example1 = ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	var tail = &example1
	for tail.Next != nil {
		tail = tail.Next
	}
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Example 1", args{&example1}, tail},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.args.head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseListRecursive(t *testing.T) {
	var example1 = ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	var tail = &example1
	for tail.Next != nil {
		tail = tail.Next
	}
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Example 1", args{&example1}, tail},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseListRecursive(tt.args.head)
			got.Print()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
