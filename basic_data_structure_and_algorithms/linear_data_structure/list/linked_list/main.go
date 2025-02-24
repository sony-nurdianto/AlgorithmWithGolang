package main

import "fmt"

type Node struct {
	Property int
	NextNode *Node
}

type LinkedList struct {
	HeadNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	node := &Node{Property: property, NextNode: linkedList.HeadNode}
	linkedList.HeadNode = node
}

func (linkedList *LinkedList) IterateList() {
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		fmt.Println(node.Property)
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	node := linkedList.HeadNode
	if node == nil {
		return nil
	}

	for node.NextNode != nil {
		node = node.NextNode
	}
	return node
}

func (linkedList *LinkedList) AddToEnd(property int) {
	node := &Node{Property: property}

	if linkedList.HeadNode == nil {
		linkedList.HeadNode = node
		return
	}

	lastNode := linkedList.LastNode()
	lastNode.NextNode = node
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.Property == property {
			return node
		}
	}
	return nil
}

func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	nodeWith := linkedList.NodeWithValue(nodeProperty)
	if nodeWith == nil {
		fmt.Printf("Node dengan property %d tidak ditemukan\n", nodeProperty)
		return
	}

	node := &Node{Property: property, NextNode: nodeWith.NextNode}
	nodeWith.NextNode = node
}

func main() {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	linkedList.IterateList()
}
