package stack

import (
	"sync"

	"github.com/lroolle/algorithm/datastructure/list"
)

type Node = list.SingleListNode

type Stack struct {
	sync.Mutex
	Head *Node
	Size int
}

func (s *Stack) Push(data interface{}) {
	s.Lock()
	defer s.Unlock()

	node := new(Node)
	node.Val = data

	temp := s.Head
	node.Next = temp

	s.Head = node
	s.Size++

}

func (s *Stack) Pop() interface{} {
	if s.Head == nil {
		return nil
	}

	s.Lock()
	defer s.Unlock()

	r := s.Head.Val
	s.Head = s.Head.Next
	s.Size--

	return r
}
