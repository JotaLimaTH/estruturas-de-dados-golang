package main

import "fmt"

func main(){
	var stack Stack = newStack(5)
	fmt.Println(stack.Last)
	stack.Append(2)
	fmt.Println(stack.Last, stack.Last.Previous.Data)
	stack.Remove()
	fmt.Println(stack.Last.Data)
}

type Node struct {
	Data int
	Previous *Node
}
type Stack struct {
	Last *Node
	Size int
}
func newStack(lastElem int) Stack {
	return Stack{
		Last: &Node{Data: lastElem, Previous: nil},
		Size: 1,
	}
}
func (s *Stack) Append(elem int){
	newNode := &Node{Data: elem, Previous: s.Last}
	if s.Last == nil {
		s.Last = newNode
		s.Size = 1
		return
	}
	s.Last = newNode
	s.Size ++
}
func (s *Stack) Remove(){
	s.Last = s.Last.Previous
	s.Size --
}