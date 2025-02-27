**T-Tree** 
adalah struktur data yang digunakan dalam database untuk menyimpan dan mengelola data secara efisien. T-Tree adalah variasi dari **B-Tree** dan **AVL Tree**, yang menggabungkan kelebihan dari kedua struktur data tersebut. T-Tree biasanya digunakan dalam sistem database in-memory karena performanya yang cepat untuk operasi penyisipan, pencarian, dan penghapusan.

Berikut adalah contoh implementasi sederhana **T-Tree** menggunakan Go:

---

### **Implementasi T-Tree di Go**

```go
package main

import (
	"fmt"
)

// Definisi node T-Tree
type TTreeNode struct {
	keys     []int // Array untuk menyimpan key
	children []*TTreeNode // Array untuk menyimpan child node
	isLeaf   bool // Menandakan apakah node adalah leaf
}

// Fungsi untuk membuat node baru
func NewTTreeNode(isLeaf bool) *TTreeNode {
	return &TTreeNode{
		keys:     make([]int, 0),
		children: make([]*TTreeNode, 0),
		isLeaf:   isLeaf,
	}
}

// Definisi T-Tree
type TTree struct {
	root *TTreeNode
	t    int // Derajat minimum T-Tree (minimal jumlah key dalam node)
}

// Fungsi untuk membuat T-Tree baru
func NewTTree(t int) *TTree {
	return &TTree{
		root: NewTTreeNode(true), // Root awalnya adalah leaf
		t:    t,
	}
}

// Fungsi untuk mencari key dalam T-Tree
func (tree *TTree) Search(key int) bool {
	return tree.searchNode(tree.root, key)
}

// Fungsi rekursif untuk mencari key dalam node
func (tree *TTree) searchNode(node *TTreeNode, key int) bool {
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

// Fungsi untuk menyisipkan key ke dalam T-Tree
func (tree *TTree) Insert(key int) {
	root := tree.root

	if len(root.keys) == (2*tree.t)-1 {
		// Jika root penuh, split root
		newRoot := NewTTreeNode(false)
		newRoot.children = append(newRoot.children, root)
		tree.splitChild(newRoot, 0)
		tree.root = newRoot
	}

	tree.insertNonFull(tree.root, key)
}

// Fungsi untuk menyisipkan key ke dalam node yang tidak penuh
func (tree *TTree) insertNonFull(node *TTreeNode, key int) {
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
func (tree *TTree) splitChild(parent *TTreeNode, index int) {
	t := tree.t
	child := parent.children[index]

	// Buat node baru untuk menyimpan separuh key dari child
	newNode := NewTTreeNode(child.isLeaf)
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

// Fungsi untuk mencetak T-Tree (inorder traversal)
func (tree *TTree) Print() {
	tree.printNode(tree.root)
	fmt.Println()
}

// Fungsi rekursif untuk mencetak node
func (tree *TTree) printNode(node *TTreeNode) {
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
	// Buat T-Tree dengan derajat minimum 2
	tree := NewTTree(2)

	// Sisipkan beberapa key
	keys := []int{10, 20, 5, 6, 12, 30, 7, 17}
	for _, key := range keys {
		tree.Insert(key)
	}

	// Cetak T-Tree
	fmt.Println("T-Tree setelah penyisipan:")
	tree.Print()

	// Cari key
	fmt.Println("Cari key 6:", tree.Search(6)) // true
	fmt.Println("Cari key 15:", tree.Search(15)) // false
}
```

---

### **Penjelasan Kode**

1. **Struktur Data**:
   - `TTreeNode`: Merepresentasikan node dalam T-Tree. Setiap node memiliki:
     - `keys`: Array untuk menyimpan key.
     - `children`: Array untuk menyimpan pointer ke child node.
     - `isLeaf`: Menandakan apakah node adalah leaf.
   - `TTree`: Merepresentasikan T-Tree. Memiliki:
     - `root`: Pointer ke root node.
     - `t`: Derajat minimum T-Tree.

2. **Operasi**:
   - **Insert**: Menyisipkan key ke dalam T-Tree. Jika node penuh, dilakukan split.
   - **Search**: Mencari key dalam T-Tree.
   - **Print**: Mencetak semua key dalam T-Tree secara inorder traversal.

3. **Split Child**:
   - Jika node penuh, node tersebut di-split menjadi dua node, dan key tengah dipindahkan ke parent.

4. **Inorder Traversal**:
   - Mencetak key secara terurut (dari terkecil ke terbesar).

---

### **Output Program**

```
T-Tree setelah penyisipan:
5 6 7 10 12 17 20 30 

Cari key 6: true
Cari key 15: false
```

---

### **Kesimpulan**
- **T-Tree** adalah struktur data yang efisien untuk operasi penyisipan, pencarian, dan penghapusan.
- Implementasi di atas adalah contoh sederhana T-Tree dengan derajat minimum `t = 2`.
- Anda dapat menyesuaikan nilai `t` sesuai kebutuhan.

Jika Anda memiliki pertanyaan lebih lanjut, beri tahu saya! 🚀
