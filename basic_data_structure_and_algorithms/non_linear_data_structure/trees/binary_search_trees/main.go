package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	Key       int
	Value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

func (tree *BinarySearchTree) InsertElement(key int, value int) *TreeNode {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	treeNode := &TreeNode{Key: key, Value: value}

	if tree.rootNode == nil {
		tree.rootNode = treeNode
		return treeNode
	}

	current := tree.rootNode
	for {
		if key == current.Key {
			current.Value = value
			return current // Return node yang diupdate
		} else if key < current.Key {
			if current.LeftNode == nil {
				current.LeftNode = treeNode
				return treeNode
			}
			current = current.LeftNode
		} else {
			if current.RightNode == nil {
				current.RightNode = treeNode
				return treeNode
			}
			current = current.RightNode
		}
	}
}

func inOrderTraverse(treeNode *TreeNode, fn func(int, int)) {
	if treeNode != nil {
		inOrderTraverse(treeNode.LeftNode, fn)
		fn(treeNode.Key, treeNode.Value)
		inOrderTraverse(treeNode.RightNode, fn)
	}
}

func (tree *BinarySearchTree) InOrderTraverseTree(fn func(int, int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverse(tree.rootNode, fn)
}

func preOrderTraverseTree(treeNode *TreeNode, fn func(int, int)) {
	if treeNode != nil {
		fn(treeNode.Key, treeNode.Value)
		preOrderTraverseTree(treeNode.LeftNode, fn)
		preOrderTraverseTree(treeNode.RightNode, fn)
	}
}

func (tree *BinarySearchTree) PreOrderTraverseTree(fn func(int, int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	preOrderTraverseTree(tree.rootNode, fn)
}

func postOrderTraverseTree(treeNode *TreeNode, fn func(int, int)) {
	if treeNode != nil {
		postOrderTraverseTree(treeNode.LeftNode, fn)
		postOrderTraverseTree(treeNode.RightNode, fn)
		fn(treeNode.Key, treeNode.Value) // Pemrosesan setelah subtree
	}
}

func (tree *BinarySearchTree) PostOrderTraverseTree(fn func(int, int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	postOrderTraverseTree(tree.rootNode, fn)
}

func (tree *BinarySearchTree) MinNode() (int, int, bool) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	if tree.rootNode == nil {
		return 0, 0, false
	}

	treeNode := tree.rootNode
	for treeNode.LeftNode != nil {
		treeNode = treeNode.LeftNode
	}
	return treeNode.Key, treeNode.Value, true
}

func (tree *BinarySearchTree) MaxNode() (int, int, bool) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	if tree.rootNode == nil {
		return 0, 0, false
	}

	node := tree.rootNode
	for node.RightNode != nil {
		node = node.RightNode
	}
	return node.Key, node.Value, true
}

func searchNode(treeNode *TreeNode, key int) bool {
	for treeNode != nil {
		if key == treeNode.Key {
			return true
		}
		if key < treeNode.Key {
			treeNode = treeNode.LeftNode
		} else {
			treeNode = treeNode.RightNode
		}
	}
	return false
}

func (tree *BinarySearchTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNode(tree.rootNode, key)
}

func removeNode(treeNode *TreeNode, key int) *TreeNode {
	if treeNode == nil {
		return nil
	}

	if key < treeNode.Key {
		treeNode.LeftNode = removeNode(treeNode.LeftNode, key)
		return treeNode
	}
	if key > treeNode.Key {
		treeNode.RightNode = removeNode(treeNode.RightNode, key)
		return treeNode
	}

	// Node ditemukan, hapus
	if treeNode.LeftNode == nil {
		return treeNode.RightNode
	}
	if treeNode.RightNode == nil {
		return treeNode.LeftNode
	}

	// Cari "leftmost node" di subtree kanan
	leftmostrightNodeParent := treeNode
	leftmostrightNode := treeNode.RightNode
	for leftmostrightNode.LeftNode != nil {
		leftmostrightNodeParent = leftmostrightNode
		leftmostrightNode = leftmostrightNode.LeftNode
	}

	// Ganti nilai node yang dihapus dengan leftmostrightNode
	treeNode.Key, treeNode.Value = leftmostrightNode.Key, leftmostrightNode.Value

	// Hapus leftmostrightNode dari tempat lama
	if leftmostrightNodeParent == treeNode {
		treeNode.RightNode = leftmostrightNode.RightNode
	} else {
		leftmostrightNodeParent.LeftNode = leftmostrightNode.RightNode
	}

	return treeNode
}

func (tree *BinarySearchTree) RemoveNode(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	tree.rootNode = removeNode(tree.rootNode, key) // Update root node jika berubah
}

func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += " "
		}
		format += "---[ "
		level++
		stringify(treeNode.LeftNode, level)
		fmt.Printf(format+"%d\n", treeNode.Key)
		stringify(treeNode.RightNode, level)
	}
}

func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(tree.rootNode, 0)
	fmt.Println("------------------------------------------------")
}

func main() {
	tree := &BinarySearchTree{}

	// Insert nodes
	tree.InsertElement(10, 100)
	tree.InsertElement(5, 50)
	tree.InsertElement(15, 150)
	tree.InsertElement(3, 30)
	tree.InsertElement(7, 70)
	tree.InsertElement(12, 120)
	tree.InsertElement(18, 180)

	// Print Tree
	fmt.Println("Binary Search Tree:")
	tree.String()

	// In-Order Traversal
	fmt.Println("In-Order Traversal:")
	tree.InOrderTraverseTree(func(key, val int) {
		fmt.Printf("Key: %d, Value: %d\n", key, val)
	})

	// Pre-Order Traversal
	fmt.Println("Pre-Order Traversal:")
	tree.PreOrderTraverseTree(func(key, val int) {
		fmt.Printf("Key: %d, Value: %d\n", key, val)
	})

	// Post-Order Traversal
	fmt.Println("Post-Order Traversal:")
	tree.PostOrderTraverseTree(func(key, val int) {
		fmt.Printf("Key: %d, Value: %d\n", key, val)
	})

	// Cari nilai minimum dan maksimum
	if key, val, ok := tree.MinNode(); ok {
		fmt.Printf("Minimum Node: Key: %d, Value: %d\n", key, val)
	}

	if key, val, ok := tree.MaxNode(); ok {
		fmt.Printf("Maximum Node: Key: %d, Value: %d\n", key, val)
	}

	// Cari node tertentu
	searchKey := 7
	fmt.Printf("Apakah key %d ada di tree? %v\n", searchKey, tree.SearchNode(searchKey))

	// Hapus node dan tampilkan hasil
	fmt.Println("\nMenghapus node 5...")
	tree.RemoveNode(5)
	tree.String()

	fmt.Println("\nMenghapus node 10 (root)...")
	tree.RemoveNode(10)
	tree.String()
}
