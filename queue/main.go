package main

import "fmt"

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
		&First: {Data: firstElem, Next: nil}
		Size: 1
	}
}
func (q *Queue) Append(node int){
	
}