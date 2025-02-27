### **Apa Itu B-Tree?**

**B-Tree** adalah struktur data pohon seimbang (balanced tree) yang dirancang untuk bekerja secara efisien pada penyimpanan sekunder seperti hard disk atau SSD. B-Tree memungkinkan operasi penyisipan, pencarian, dan penghapusan dalam waktu **O(log n)**, di mana `n` adalah jumlah elemen dalam pohon.

B-Tree sangat cocok untuk sistem yang membutuhkan akses data secara efisien, terutama ketika data disimpan di disk (bukan di memori utama), karena B-Tree meminimalkan jumlah akses disk.

---

### **Karakteristik B-Tree**

1. **Seimbang (Balanced)**:
   - Semua leaf node berada pada level yang sama.
   - Ini menjamin bahwa operasi pencarian, penyisipan, dan penghapusan selalu membutuhkan waktu **O(log n)**.

2. **Node dengan Banyak Anak**:
   - Setiap node dalam B-Tree dapat memiliki banyak anak (biasanya lebih dari 2).
   - Ini berbeda dengan binary search tree (BST) yang hanya memiliki maksimal 2 anak per node.

3. **Derajat Minimum (t)**:
   - Setiap node (kecuali root) harus memiliki setidaknya `t-1` key.
   - Setiap node dapat memiliki maksimal `2t-1` key.
   - Nilai `t` disebut derajat minimum B-Tree.

4. **Penyimpanan Data di Setiap Node**:
   - Setiap node menyimpan array key yang terurut.
   - Key-key ini membagi subtree menjadi interval-interval.

5. **Leaf Node**:
   - Semua leaf node berada pada level yang sama.
   - Leaf node tidak memiliki anak.

---

### **Struktur Node B-Tree**

Setiap node dalam B-Tree memiliki:
- Array key yang terurut: `[key1, key2, ..., keyN]`.
- Array pointer ke child node: `[child1, child2, ..., childN+1]`.
- Jumlah key dalam node: `n` (di mana `t-1 <= n <= 2t-1`).

---

### **Operasi pada B-Tree**

1. **Pencarian (Search)**:
   - Mirip dengan pencarian di binary search tree (BST).
   - Mulai dari root, bandingkan key yang dicari dengan key dalam node.
   - Jika key ditemukan, kembalikan nilai.
   - Jika tidak, lanjutkan ke child yang sesuai.

2. **Penyisipan (Insert)**:
   - Sisipkan key ke leaf node.
   - Jika leaf node penuh, lakukan split node dan naikkan key tengah ke parent.
   - Jika root penuh, split root dan buat root baru.

3. **Penghapusan (Delete)**:
   - Hapus key dari node.
   - Jika node memiliki terlalu sedikit key setelah penghapusan, lakukan penggabungan (merge) atau pinjam key dari sibling.

---

### **Contoh B-Tree**

Misalkan kita memiliki B-Tree dengan derajat minimum `t = 2`. Setiap node dapat memiliki:
- Minimal 1 key (`t-1 = 1`).
- Maksimal 3 key (`2t-1 = 3`).

Contoh B-Tree:

```
        [10, 20]
       /    |    \
[5]  [15]  [25, 30]
```

- Root node memiliki 2 key: `10` dan `20`.
- Child node pertama memiliki 1 key: `5`.
- Child node kedua memiliki 1 key: `15`.
- Child node ketiga memiliki 2 key: `25` dan `30`.

---

### **Keuntungan B-Tree**

1. **Efisiensi Disk**:
   - B-Tree dirancang untuk meminimalkan jumlah akses disk.
   - Setiap node dapat menyimpan banyak key, sehingga mengurangi tinggi pohon dan jumlah akses disk.

2. **Operasi yang Cepat**:
   - Pencarian, penyisipan, dan penghapusan dalam waktu **O(log n)**.

3. **Cocok untuk Database**:
   - B-Tree digunakan secara luas dalam sistem database (misalnya, MySQL, PostgreSQL) untuk mengimplementasikan indeks.

---

### **Perbandingan dengan Struktur Data Lain**

| Fitur                  | B-Tree               | Binary Search Tree (BST) | AVL Tree              |
|------------------------|----------------------|--------------------------|-----------------------|
| **Jumlah Anak per Node** | Banyak (biasanya > 2) | Maksimal 2               | Maksimal 2            |
| **Keseimbangan**        | Selalu seimbang      | Tidak selalu seimbang    | Selalu seimbang       |
| **Penggunaan**          | Database, file system | Memori utama             | Memori utama          |
| **Kompleksitas Operasi** | O(log n)             | O(log n) (jika seimbang) | O(log n)              |

---

### **Implementasi B-Tree di Go**

Berikut adalah contoh implementasi sederhana B-Tree di Go:

```go
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
	fmt.Println("Cari key 6:", tree.Search(6)) // true
	fmt.Println("Cari key 15:", tree.Search(15)) // false
}
```

---

### **Output Program**

```
B-Tree setelah penyisipan:
5 6 7 10 12 17 20 30 

Cari key 6: true
Cari key 15: false
```

---

### **Kesimpulan**
- **B-Tree** adalah struktur data yang sangat efisien untuk sistem yang membutuhkan akses data secara cepat, terutama pada penyimpanan sekunder.
- B-Tree digunakan secara luas dalam database dan sistem file.
- Implementasi di atas adalah contoh sederhana B-Tree dengan derajat minimum `t = 2`.

Jika Anda memiliki pertanyaan lebih lanjut, beri tahu saya! 🚀
