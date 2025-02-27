package main

import (
	"fmt"
)

// Definisi node B-Tree
type BTreeNode struct {
	keys     []int
	children []*BTreeNode
	isLeaf   bool
}

// Fungsi untuk membuat node baru
func NewBTreeNode(isLeaf bool) *BTreeNode {
	return &BTreeNode{
		keys:     make([]int, 0),
		children: make([]*BTreeNode, 0),
		isLeaf:   isLeaf,
	}
}

// Definisi B-Tree
type BTree struct {
	root *BTreeNode
	t    int // Derajat minimum B-Tree
}

// Fungsi untuk membuat B-Tree baru
func NewBTree(t int) *BTree {
	return &BTree{
		root: NewBTreeNode(true), // Root awalnya adalah leaf
		t:    t,
	}
}

// Fungsi untuk mencari key dalam B-Tree
func (tree *BTree) Search(key int) bool {
	return tree.searchNode(tree.root, key)
}

// Fungsi rekursif untuk mencari key dalam node
func (tree *BTree) searchNode(node *BTreeNode, key int) bool {
	i := 0
	for i < len(node.keys) && key > node.keys[i] {
		i++
	}

	if i < len(node.keys) && key == node.keys[i] {
		return true // Key ditemukan
	}

	if node.isLeaf {
		return false // Key tidak ditemukan
	}

	return tree.searchNode(node.children[i], key) // Cari di child node
}

// Fungsi untuk menyisipkan key ke dalam B-Tree
func (tree *BTree) Insert(key int) {
	root := tree.root

	if len(root.keys) == (2*tree.t)-1 {
		// Jika root penuh, split root
		newRoot := NewBTreeNode(false)
		newRoot.children = append(newRoot.children, root)
		tree.splitChild(newRoot, 0)
		tree.root = newRoot
	}

	tree.insertNonFull(tree.root, key)
}

// Fungsi untuk menyisipkan key ke dalam node yang tidak penuh
func (tree *BTree) insertNonFull(node *BTreeNode, key int) {
	i := len(node.keys) - 1

	if node.isLeaf {
		// Jika node adalah leaf, sisipkan key ke dalam array keys
		node.keys = append(node.keys, 0)
		for i >= 0 && key < node.keys[i] {
			node.keys[i+1] = node.keys[i]
			i--
		}
		node.keys[i+1] = key
	} else {
		// Jika node bukan leaf, cari child yang tepat
		for i >= 0 && key < node.keys[i] {
			i--
		}
		i++

		if len(node.children[i].keys) == (2*tree.t)-1 {
			// Jika child penuh, split child
			tree.splitChild(node, i)
			if key > node.keys[i] {
				i++
			}
		}

		tree.insertNonFull(node.children[i], key)
	}
}

// Fungsi untuk split child node
func (tree *BTree) splitChild(parent *BTreeNode, index int) {
	t := tree.t
	child := parent.children[index]

	// Buat node baru untuk menyimpan separuh key dari child
	newNode := NewBTreeNode(child.isLeaf)
	newNode.keys = append(newNode.keys, child.keys[t:]...)

	if !child.isLeaf {
		// Jika child bukan leaf, pindahkan separuh children ke newNode
		newNode.children = append(newNode.children, child.children[t:]...)
		child.children = child.children[:t]
	}

	// Pindahkan key tengah ke parent
	parent.keys = append(parent.keys, 0)
	for i := len(parent.keys) - 1; i > index; i-- {
		parent.keys[i] = parent.keys[i-1]
	}
	parent.keys[index] = child.keys[t-1]

	// Potong key child
	child.keys = child.keys[:t-1]

	// Sisipkan newNode ke parent
	parent.children = append(parent.children, nil)
	for i := len(parent.children) - 1; i > index+1; i-- {
		parent.children[i] = parent.children[i-1]
	}
	parent.children[index+1] = newNode
}

// Fungsi untuk mencetak B-Tree (inorder traversal)
func (tree *BTree) Print() {
	tree.printNode(tree.root)
	fmt.Println()
}

// Fungsi rekursif untuk mencetak node
func (tree *BTree) printNode(node *BTreeNode) {
	if node == nil {
		return
	}

	for i := 0; i < len(node.keys); i++ {
		if !node.isLeaf {
			tree.printNode(node.children[i])
		}
		fmt.Print(node.keys[i], " ")
	}

	if !node.isLeaf {
		tree.printNode(node.children[len(node.children)-1])
	}
}

func main() {
	// Buat B-Tree dengan derajat minimum 2
	tree := NewBTree(2)

	// Sisipkan beberapa key
	keys := []int{10, 20, 5, 6, 12, 30, 7, 17}
	for _, key := range keys {
		tree.Insert(key)
	}

	// Cetak B-Tree
	fmt.Println("B-Tree setelah penyisipan:")
	tree.Print()

	// Cari key
	fmt.Println("Cari key 6:", tree.Search(6))   // true
	fmt.Println("Cari key 15:", tree.Search(15)) // false
}
