package main

import "fmt"

type Node struct {
	Property     int
	NextNode     *Node
	PreviousNode *Node
}

type LinkedList struct {
	HeadNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	node := &Node{
		Property: property,
	}

	if linkedList.HeadNode != nil {
		node.NextNode = linkedList.HeadNode
		linkedList.HeadNode.PreviousNode = node
	}

	linkedList.HeadNode = node
}

func (linkedList *LinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *Node {
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.PreviousNode != nil && node.NextNode != nil {
			if node.PreviousNode.Property == firstProperty && node.NextNode.Property == secondProperty {
				return node
			}
		}
	}
	return nil
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	for node := linkedList.HeadNode; node != nil; node = node.NextNode {
		if node.Property == property {
			return node
		}
	}
	return nil
}

func (linkedList *LinkedList) AddAfterNode(nodeProperty int, property int) {
	nodeWith := linkedList.NodeWithValue(nodeProperty)

	if nodeWith == nil {
		return
	}

	node := &Node{
		Property:     property,
		NextNode:     nodeWith.NextNode,
		PreviousNode: nodeWith,
	}

	if nodeWith.NextNode != nil {
		nodeWith.NextNode.PreviousNode = node
	}

	nodeWith.NextNode = node
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
	node := &Node{
		Property: property,
		NextNode: nil,
	}

	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.NextNode = node
		node.PreviousNode = lastNode
	} else {
		linkedList.HeadNode = node
	}
}

func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	nodeWith := linkedList.NodeWithValue(nodeProperty)
	if nodeWith == nil {
		fmt.Printf("Node dengan property %d tidak ditemukan\n", nodeProperty)
		return
	}

	node := &Node{
		Property:     property,
		NextNode:     nodeWith.NextNode,
		PreviousNode: nodeWith,
	}

	if nodeWith.NextNode != nil {
		nodeWith.NextNode.PreviousNode = node
	}

	nodeWith.NextNode = node
}

func main() {
	var linkedList LinkedList
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)

	fmt.Println(linkedList.HeadNode.Property)

	node := linkedList.NodeBetweenValues(1, 5)
	if node != nil {
		fmt.Println(node.Property)
	} else {
		fmt.Println("Node tid ditemukan")
	}
}
