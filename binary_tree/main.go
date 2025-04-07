package main

import "fmt"

func main() {
	binaryTree := NewBinaryTree(5)
	binaryTree.Insert(5, &binaryTree.Root)
	binaryTree.Insert(2, &binaryTree.Root)
	binaryTree.Insert(7, &binaryTree.Root)
	binaryTree.Insert(4, &binaryTree.Root)
	binaryTree.Insert(20, &binaryTree.Root)
	binaryTree.Insert(15, &binaryTree.Root)
	binaryTree.Insert(28, &binaryTree.Root)
	fmt.Println(binaryTree.Root.Data, binaryTree.Root.Left.Data, binaryTree.Root.Right.Data)
	fmt.Println("In Order Traversal: ", binaryTree.InOrderTraversal(binaryTree.Root))
	fmt.Println("Pre Order Traversal: ", binaryTree.PreOrderTraversal(binaryTree.Root))
	fmt.Println("Post Order Traversal: ", binaryTree.PostOrderTraversal(binaryTree.Root))
	binaryTree.Delete(binaryTree.Root, 2)
	fmt.Println("Post Order Traversal: ", binaryTree.PostOrderTraversal(binaryTree.Root))
	binaryTree.Delete(binaryTree.Root, 20)
	fmt.Println("Post Order Traversal: ", binaryTree.PostOrderTraversal(binaryTree.Root))
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
func (b *BinaryTree) InOrderTraversal(root *Node) []int { //Isso vai retornar um array com os dados dos elementos da árvore
	if root == nil {
		return []int{}
	}
	var result []int
	left := b.InOrderTraversal(root.Left)
	right := b.InOrderTraversal(root.Right)

	result = append(result, left...)
	result = append(result, root.Data)
	result = append(result, right...)

	return result
}
func (b *BinaryTree) PreOrderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	left := b.PreOrderTraversal(root.Left)
	right := b.PreOrderTraversal(root.Right)

	result = append(result, root.Data)
	result = append(result, left...)
	result = append(result, right...)

	return result
}
func (b *BinaryTree) PostOrderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	left := b.PostOrderTraversal(root.Left)
	right := b.PostOrderTraversal(root.Right)

	result = append(result, left...)
	result = append(result, right...)
	result = append(result, root.Data)

	return result
}
func (b *BinaryTree) Search(root *Node, key int) Node{
	if root == nil || key == root.Data {
		return *root
	}
	if key < root.Data {
		return b.Search(root.Left, key)
	}
	return b.Search(root.Right, key)
}
func (b *BinaryTree) Delete(root *Node, key int) *Node{
	if root == nil {
		return root
	}

	if key < root.Data {
		root.Left = b.Delete(root.Left, key)
	} else if key > root.Data {
		root.Right = b.Delete(root.Right, key)
	} else {
		// Deleção de folha
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// Deleção de node com só um filho
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		if root.Right == nil && root.Left != nil {
			return root.Left
		}

		if root.Left != nil && root.Right != nil {
			minNode := b.FindMin(root.Right)
			root.Data = minNode.Data
			root.Right = b.Delete(root.Right, minNode.Data)
		}
	}

	return root
}
func (b *BinaryTree) FindMin(root *Node) *Node {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
func (b *BinaryTree) GetHeight(root *Node) int {
 	leftHeight := 1
	rightHeight := 1
	pointerLeft := root
	pointerRight := root
	for pointerLeft.Left != nil {
		leftHeight ++
		pointerLeft = pointerLeft.Left
	}
	for pointerRight.Right != nil {
		rightHeight ++
		pointerRight = pointerLeft.Right
	}
	if leftHeight > rightHeight {
		return leftHeight
	}
	return rightHeight
}