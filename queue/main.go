package main

import "fmt"

func main(){
	fila := newQueue(5)
	fmt.Println(fila)
	fila.Append(3)
	fmt.Println(fila.First.Data)
	fmt.Println(fila.First.Next.Data)
}

type Node struct {
	Data int
	Next *Node
}
type Queue struct {
	First *Node
	Size  int
}
func newQueue(firstElem int) Queue{
	return Queue{
		First: &Node{Data: firstElem, Next: nil},
		Size: 1,
	}
}
func (q *Queue) Append(elem int){
	newNode := &Node{Data: elem, Next: nil}
	if q.First == nil {
		q.First = newNode
		return
	}
	current := q.First
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	q.Size++
}