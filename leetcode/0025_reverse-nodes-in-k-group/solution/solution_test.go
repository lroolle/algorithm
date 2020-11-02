package solution

import (
	"reflect"
	"testing"
)

var example1 = ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
var example2 = ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Example 1", args{&example1, 2}, example1.Next},
		{"Example 2", args{&example2, 3}, example2.Next.Next},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(tt.args.head, tt.args.k)
			got.Print()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
