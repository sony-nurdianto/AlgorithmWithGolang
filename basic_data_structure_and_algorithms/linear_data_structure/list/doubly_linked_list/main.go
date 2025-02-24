package main

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

func main() {
}
