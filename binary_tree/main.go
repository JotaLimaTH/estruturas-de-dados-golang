package main

import "fmt"

func main() {
	binaryTree := NewBinaryTree(5)
	binaryTree.Insert(5, &binaryTree.Root)
	binaryTree.Insert(2, &binaryTree.Root)
	binaryTree.Insert(7, &binaryTree.Root)
	binaryTree.Insert(4, &binaryTree.Root)
	fmt.Println(binaryTree.Root.Data, binaryTree.Root.Left.Data, binaryTree.Root.Right.Data)
}

type Node struct {
	Data int
	Left *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}
func NewBinaryTree(elem int) BinaryTree {
	return BinaryTree{
		Root: &Node{Data: elem},
	}
}
func (b *BinaryTree) Insert(elem int, root **Node) {
	if *root == nil {
		*root = &Node{Data: elem}
		return
	}
	
	if elem < (*root).Data {
		b.Insert(elem, &(*root).Left)
	} else {
		b.Insert(elem, &(*root).Right)
	}
}