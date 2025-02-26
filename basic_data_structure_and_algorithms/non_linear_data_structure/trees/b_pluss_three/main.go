package main

import (
	"fmt"
	"sort"
)

type BPlusTree struct {
	root  *node
	order int
}

type node struct {
	keys     []int
	children []*node
	isLeaf   bool
	next     *node // For linked list in leaf nodes
}

func NewBPlusTree(order int) *BPlusTree {
	return &BPlusTree{order: order}
}

func (tree *BPlusTree) Insert(key int) {
	if tree.root == nil {
		tree.root = &node{keys: []int{key}, isLeaf: true}
		return
	}

	// Insert helper function
	tree.root = insertIntoNode(tree.root, key, tree.order)
}

func insertIntoNode(n *node, key, order int) *node {
	if n.isLeaf {
		n.keys = append(n.keys, key)
		sort.Ints(n.keys)

		if len(n.keys) > order {
			return splitLeaf(n, order)
		}
		return n
	}

	// Find correct child
	idx := sort.Search(len(n.keys), func(i int) bool { return key < n.keys[i] })
	n.children[idx] = insertIntoNode(n.children[idx], key, order)

	if len(n.children) > order {
		return splitInternal(n, order)
	}
	return n
}

func splitLeaf(n *node, order int) *node {
	mid := order / 2
	newLeaf := &node{keys: append([]int{}, n.keys[mid:]...), isLeaf: true}
	n.keys = n.keys[:mid]

	newLeaf.next = n.next
	n.next = newLeaf

	return &node{keys: []int{newLeaf.keys[0]}, children: []*node{n, newLeaf}}
}

func splitInternal(n *node, order int) *node {
	mid := order / 2
	newNode := &node{keys: append([]int{}, n.keys[mid+1:]...), children: append([]*node{}, n.children[mid+1:]...)}
	n.keys = n.keys[:mid]
	n.children = n.children[:mid+1]

	return &node{keys: []int{n.keys[mid]}, children: []*node{n, newNode}}
}

func (tree *BPlusTree) Print() {
	current := tree.root
	for current != nil && !current.isLeaf {
		current = current.children[0]
	}
	for current != nil {
		fmt.Println(current.keys)
		current = current.next
	}
}

func main() {
	tree := NewBPlusTree(4)
	values := []int{10, 20, 5, 6, 12, 30, 7, 17}

	for _, v := range values {
		tree.Insert(v)
	}

	tree.Print()
}
