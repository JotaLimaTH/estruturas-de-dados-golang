package main

import "fmt"

func main(){
	var stack Stack = NewStack(5)
	fmt.Println(stack.Last)
	stack.Append(2)
	stack.Append(3)
	stack.Append(4)
	stack.Append(124)
	stack.PrintStack()
}

type Node struct {
	Data int
	Previous *Node
}
type Stack struct {
	Last *Node
	Size int
}
func NewStack(lastElem int) Stack {
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
func (s *Stack) PrintStack(){
	current := s.Last
	for current != nil {
		fmt.Println(current.Data)
		current = current.Previous
	}
}