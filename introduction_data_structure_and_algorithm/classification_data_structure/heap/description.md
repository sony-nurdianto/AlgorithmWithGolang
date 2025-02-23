## **Heap**  
Heap adalah **struktur data berbasis heap property** yang sering digunakan dalam **algoritma seleksi, graf**, dan **k-way merge**. Beberapa operasi yang dapat dilakukan pada heap meliputi:  
✅ **Pencarian (Finding)**  
✅ **Penggabungan (Merging)**  
✅ **Penyisipan (Insertion)**  
✅ **Perubahan nilai kunci (Key changes)**  
✅ **Penghapusan (Deleting)**  

Di Go, **heap** tersedia dalam package `container/heap`.  

---

## **📌 Jenis Heap**  
Heap memiliki dua jenis berdasarkan **heap order property**:  

1. **Max Heap** (Heap Maksimum)  
   - Setiap node memiliki nilai **lebih besar atau sama** dengan nilai anak-anaknya.  
   - Urutan **descending** (menurun).  

2. **Min Heap** (Heap Minimum)  
   - Setiap node memiliki nilai **lebih kecil atau sama** dengan nilai anak-anaknya.  
   - Urutan **ascending** (menaik).  

Heap pertama kali diperkenalkan oleh **J.W.J. Williams pada tahun 1964** sebagai bagian dari **Heap Sort Algorithm**. Heap **bukan struktur data yang sepenuhnya terurut**, tetapi memiliki **partial order**.

---

## **📌 Contoh Implementasi Heap di Go**
Berikut adalah contoh penggunaan `container/heap` untuk membuat **min-heap** di Go:

```go
// Paket utama berisi contoh kode dari buku Hands-On Data Structures and Algorithms with Go
package main

// Import package yang dibutuhkan
import (
	"container/heap"
	"fmt"
)

// IntegerHeap adalah tipe data berbentuk slice dari int
type IntegerHeap []int

// Len mengembalikan panjang heap
func (iheap IntegerHeap) Len() int { return len(iheap) }

// Less menentukan aturan min-heap (elemen lebih kecil punya prioritas lebih tinggi)
func (iheap IntegerHeap) Less(i, j int) bool { return iheap[i] < iheap[j] }

// Swap menukar dua elemen dalam heap
func (iheap IntegerHeap) Swap(i, j int) { iheap[i], iheap[j] = iheap[j], iheap[i] }

// Push menambahkan elemen baru ke dalam heap
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

// Pop menghapus elemen terkecil dari heap
func (iheap *IntegerHeap) Pop() interface{} {
	old := *iheap
	n := len(old)
	x := old[n-1]
	*iheap = old[0 : n-1]
	return x
}

// Fungsi utama (main)
func main() {
	// Inisialisasi heap dengan beberapa elemen awal
	intHeap := &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)

	// Menambahkan elemen baru ke dalam heap
	heap.Push(intHeap, 2)
	fmt.Printf("Minimum: %d\n", (*intHeap)[0]) // Menampilkan elemen terkecil dalam heap

	// Menghapus dan mencetak elemen dari heap satu per satu
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
```

---

## **📌 Cara Menjalankan Program**
Jalankan perintah berikut di terminal:
```sh
go run heap.go
```
Output yang dihasilkan:
```
Minimum: 1
1
2
4
5
```

---

## **📌 Kesimpulan**  
✅ **Heap adalah struktur data yang digunakan untuk pencarian, pengurutan, dan pemrosesan antrian prioritas.**  
✅ **Go memiliki package `container/heap` yang memudahkan implementasi heap.**  
✅ **Min Heap memprioritaskan elemen terkecil, sedangkan Max Heap memprioritaskan elemen terbesar.**  
✅ **Heap bukan sepenuhnya terurut, tetapi memiliki sifat *partial order*.** 🚀
