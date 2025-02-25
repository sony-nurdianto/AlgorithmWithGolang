### **Pohon Pencarian Biner (Binary Search Tree - BST)**  
Pohon pencarian biner adalah **struktur data** yang memungkinkan **pencarian, penambahan, dan penghapusan elemen secara cepat**.  

- **BST menyimpan kunci (keys) dalam urutan yang terurut**, sehingga pencarian lebih cepat dibandingkan dengan struktur data yang tidak terurut.  
- Struktur data ini ditemukan oleh **P. F. Windley, A. D. Booth, A. J. T. Colin, dan T. N. Hibbard**.  
- **Penggunaan ruang BST rata-rata sebesar O(n)**, sedangkan operasi **penyisipan (insert), pencarian (search), dan penghapusan (delete) memiliki kompleksitas rata-rata O(log n)**.  

Setiap node dalam BST memiliki beberapa properti atau atribut:  
- **Key** (bilangan integer) sebagai identitas atau nilai unik.  
- **Value** (bilangan integer) sebagai data yang disimpan dalam node.  
- **leftNode** (pointer ke anak kiri) yang berisi nilai lebih kecil dari key node saat ini.  
- **rightNode** (pointer ke anak kanan) yang berisi nilai lebih besar dari key node saat ini.  

Contoh representasi BST dalam kode Go:  

```go
// TreeNode struct
type TreeNode struct {
    key       int
    value     int
    leftNode  *TreeNode 
    rightNode *TreeNode 
}
```
Pada bagian selanjutnya, implementasi lengkap **BinarySearchTree** akan dibahas lebih lanjut. Untuk kode lengkap, silakan lihat file **binary_search_tree.go**.
