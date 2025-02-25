### **Tuple dalam Pemrograman dan Basis Data**  

**Tuple** adalah urutan objek yang berhingga dan terurut. Tuple dapat menyimpan berbagai tipe data yang berbeda dalam satu struktur. Biasanya, tuple digunakan untuk mengelompokkan data yang saling berkaitan.  

Dalam **basis data relasional**, **tuple** merepresentasikan **satu baris dalam tabel**. Sekumpulan tuple dalam sebuah tabel disebut sebagai **instance relasi**.  

Dibandingkan dengan **list**, tuple memiliki ukuran yang tetap (immutable) sehingga lebih cepat dalam eksekusi. Tuple juga dapat digunakan dalam satu pernyataan, yang berguna untuk **menukar nilai (swapping values)**.  

**Perbedaan utama antara list dan tuple:**
- **List** biasanya menyimpan elemen dengan tipe data yang sama.
- **Tuple** dapat menyimpan elemen dengan tipe data yang berbeda.

Sebagai contoh, kita bisa menyimpan **nama, umur, dan warna favorit seseorang** dalam satu tuple.

---

### **Contoh Tuple dalam Go**
Di Go, tuple bisa direpresentasikan menggunakan **multiple return values** pada fungsi.  

Kode berikut menunjukkan bagaimana kita bisa mengembalikan lebih dari satu nilai dari fungsi:

#### **Kode:**
```go
package main

import (
	"fmt"
)

// Fungsi h mengembalikan hasil perkalian dari parameter x dan y
func h(x int, y int) int {
	return x * y
}

// Fungsi g mengembalikan dua nilai x dan y setelah modifikasi
func g(l int, m int) (x int, y int) {
	x = 2 * l
	y = 4 * m
	return
}

// Fungsi utama (main)
func main() {
	fmt.Println(h(g(2, 3))) // Output: 96
}
```

---

### **Penjelasan Kode:**
1. **Fungsi `g(l, m)`:**  
   - Menerima dua parameter `l` dan `m`.
   - Mengembalikan dua nilai `x = 2*l` dan `y = 4*m`.

2. **Fungsi `h(x, y)`:**  
   - Menerima dua angka (`x` dan `y`).
   - Mengembalikan hasil perkalian `x * y`.

3. **Di dalam `main()`:**
   - `g(2,3)` menghasilkan `(4,12)`, karena `x = 2*2 = 4` dan `y = 4*3 = 12`.
   - Nilai `(4,12)` ini digunakan sebagai parameter untuk `h(4,12)`, sehingga hasilnya adalah `4 * 12 = 48`.

---

### **Kesimpulan**
- Tuple di Go direpresentasikan dengan **multiple return values** dalam fungsi.
- Hal ini sangat berguna untuk mengembalikan beberapa nilai sekaligus, seperti dalam operasi **swapping values**, **modifikasi data**, atau **fungsi matematika**.
- Konsep tuple sering digunakan dalam basis data relasional untuk merepresentasikan satu baris data dalam tabel.

Bagian selanjutnya dalam buku akan membahas **Queue**, yang merupakan salah satu struktur data linear.
