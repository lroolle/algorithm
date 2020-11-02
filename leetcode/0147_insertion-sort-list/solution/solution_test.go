package solution

import (
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"Example 1",
			args{&ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}},
			&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		},
		{
			"Example 2",
			args{&ListNode{Val: -1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 0}}}}}},
			&ListNode{Val: -1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSortList(tt.args.head); !got.Compare(tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
