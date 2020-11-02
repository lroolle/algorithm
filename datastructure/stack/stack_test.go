package stack

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestStack_Push(t *testing.T) {
	type fields struct {
		Head *Node
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"1", fields{&Node{}}, args{1}},
		{"'2'", fields{&Node{}}, args{"2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				Head: tt.fields.Head,
			}
			s.Push(tt.args.data)
			fmt.Println(s.Pop())
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		Mutex sync.Mutex
		Head  *Node
		Size  int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				Mutex: tt.fields.Mutex,
				Head:  tt.fields.Head,
				Size:  tt.fields.Size,
			}
			if got := s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
