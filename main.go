package main

import "math/rand"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	node := &Node{
		Value: value,
	}
	return node
}

func (n *Node) PrintSequentialNode() {
	if n.Left != nil {
		n.Left.PrintSequentialNode()
	}
	println(n.Value)
	if n.Right != nil {
		n.Right.PrintSequentialNode()
	}
}

type Tree struct {
	Root *Node
}

func NewTree() *Tree {
	tree := &Tree{}
	return tree
}

func (t *Tree) Insert(value int) {
	if t.Root == nil {
		t.Root = NewNode(value)
		return
	}

	aux := t.Root

	for aux != nil {
		if value <= aux.Value {
			prox := aux.Left
			if prox == nil {
				aux.Left = NewNode(value)
			}
			aux = prox
		} else {
			prox := aux.Right
			if prox == nil {
				aux.Right = NewNode(value)
			}
			aux = prox
		}
	}
}

func (t *Tree) PrintTree() {
	if t.Root != nil {
		t.Root.PrintSequentialNode()
	}
}

func main() {
	tree := NewTree()
	for i := 0; i < 100; i++ {
		tree.Insert(rand.Intn(100))
	}
	tree.PrintTree()
}
