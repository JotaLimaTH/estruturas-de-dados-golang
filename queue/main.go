package main

import "fmt"

func main(){
	fila := NewQueue(5)
	fila.Append(3)
	fila.Append(1)
	fila.Append(120)
	fila.Append(-321)
	fila.PrintQueue()
}

type Node struct {
	Data int
	Next *Node
}
type Queue struct {
	First *Node
	Size  int
}
func NewQueue(firstElem int) Queue{
	return Queue{
		First: &Node{Data: firstElem, Next: nil},
		Size: 1,
	}
}
func (q *Queue) Append(elem int){
	newNode := &Node{Data: elem, Next: nil}
	if q.First == nil {
		q.First = newNode
		q.Size = 1
		return
	}
	current := q.First
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	q.Size++
}
func (q *Queue) Remove(){
	q.First = q.First.Next
	q.Size --
}
func (q *Queue) PrintQueue(){
	current := q.First
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}