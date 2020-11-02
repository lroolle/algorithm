package main

import (
	"fmt"

	"github.com/lroolle/algorithm/datastructure/stack"
)

func main() {
	s := &stack.Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push("123")
	fmt.Printf("Size: %d\n", s.Size)
	fmt.Printf("Popd: %#v\n", s.Pop())
	fmt.Printf("Popd: %#v\n", s.Pop())
	fmt.Printf("Popd: %#v\n", s.Pop())
	fmt.Printf("Popd: %#v\n", s.Pop())
}
