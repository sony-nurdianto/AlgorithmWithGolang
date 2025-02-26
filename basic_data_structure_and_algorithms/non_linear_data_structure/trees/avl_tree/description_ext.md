AVL Tree adalah pohon pencarian biner (BST) yang seimbang secara mandiri. Nama AVL diambil dari penemunya, Adelson-Velsky dan Landis. Pohon ini menjaga keseimbangan dengan memastikan bahwa selisih tinggi antara subpohon kiri dan kanan tidak lebih dari 1 untuk setiap node.

---

### **Mengapa AVL Tree?**
BST biasa dapat menjadi tidak seimbang jika elemen dimasukkan dalam urutan tertentu, menyebabkan kompleksitas operasi menjadi O(n) dalam kasus terburuk. AVL tree mengatasi masalah ini dengan menyeimbangkan dirinya menggunakan rotasi setelah setiap operasi **insert** atau **delete**, sehingga tetap menjaga kompleksitas O(log n).

---

### **Properti AVL Tree**
- **Balance Factor (BF)**: Perbedaan tinggi antara subpohon kiri dan kanan.  
  \[
  BF = \text{Height(left subtree)} - \text{Height(right subtree)}
  \]
  - Jika **BF = -1, 0, atau 1**, pohon dalam keadaan seimbang.
  - Jika **BF < -1 atau BF > 1**, pohon tidak seimbang dan perlu dilakukan rotasi.

---

### **Rotasi pada AVL Tree**
Untuk menyeimbangkan AVL Tree setelah **insert** atau **delete**, kita menggunakan rotasi:

1. **Single Rotations:**
   - **Right Rotation (LL Case)** → terjadi jika elemen baru ditambahkan ke subpohon kiri dari anak kiri.
   - **Left Rotation (RR Case)** → terjadi jika elemen baru ditambahkan ke subpohon kanan dari anak kanan.

2. **Double Rotations:**
   - **Left-Right Rotation (LR Case)** → terjadi jika elemen baru ditambahkan ke subpohon kanan dari anak kiri.
   - **Right-Left Rotation (RL Case)** → terjadi jika elemen baru ditambahkan ke subpohon kiri dari anak kanan.

---

### **Mari Kita Implementasikan AVL Tree dalam Go**
Kita akan membuat AVL Tree dengan fitur dasar seperti **insert** dan **rotasi** untuk menjaga keseimbangan.

Langsung saja kita buat implementasinya! 😊

Ini adalah implementasi dasar AVL Tree dalam Go, dengan fitur **insert**, **rotasi (LL, RR, LR, RL)**, dan **in-order traversal** untuk melihat hasilnya.

Setelah ini, kita bisa melanjutkan dengan implementasi fitur **delete** dalam AVL Tree. Mau lanjut ke sana? 😊

```
package main

import (
	"fmt"
)

type TreeNode struct {
	Key    int
	Height int
	Left   *TreeNode
	Right  *TreeNode
}

type AVLTree struct {
	Root *TreeNode
}

// Get height of a node
func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// Get balance factor
func getBalanceFactor(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

// Right Rotation (LL Case)
func rightRotate(y *TreeNode) *TreeNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	return x
}

// Left Rotation (RR Case)
func leftRotate(x *TreeNode) *TreeNode {
	y := x.Right
	T2 := y.Left

y.Left = x
	x.Right = T2
	x.Height = max(height(x.Left), height(x.Right)) + 1

y.Height = max(height(y.Left), height(y.Right)) + 1

	return y
}

// Insert a key in AVL Tree
func insert(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return &TreeNode{Key: key, Height: 1}
	}
	if key < node.Key {
		node.Left = insert(node.Left, key)
	} else if key > node.Key {
		node.Right = insert(node.Right, key)
	} else {
		return node // Duplicate keys not allowed
	}

	node.Height = 1 + max(height(node.Left), height(node.Right))
	balance := getBalanceFactor(node)

	// Left Left (LL) Case
	if balance > 1 && key < node.Left.Key {
		return rightRotate(node)
	}

	// Right Right (RR) Case
	if balance < -1 && key > node.Right.Key {
		return leftRotate(node)
	}

	// Left Right (LR) Case
	if balance > 1 && key > node.Left.Key {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	// Right Left (RL) Case
	if balance < -1 && key < node.Right.Key {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

// Wrapper function for inserting into AVL Tree
func (tree *AVLTree) Insert(key int) {
	tree.Root = insert(tree.Root, key)
}

// In-Order Traversal
func inOrder(node *TreeNode) {
	if node != nil {
		inOrder(node.Left)
		fmt.Printf("%d ", node.Key)
		inOrder(node.Right)
	}
}

// Main function
tfunc main() {
	avl := &AVLTree{}
	avl.Insert(10)
	avl.Insert(20)
	avl.Insert(30)
	avl.Insert(40)
	avl.Insert(50)
	avl.Insert(25)

	fmt.Println("In-order Traversal of AVL Tree:")
	inOrder(avl.Root)
}
```
