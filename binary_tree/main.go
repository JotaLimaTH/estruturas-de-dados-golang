package main

import "fmt"

type Node struct {
	Data int
	Left *Node
	Right *Node
	Parent *Node
}

type BinaryTree struct {
	Root *Node
}