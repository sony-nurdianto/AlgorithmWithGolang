#### **Tuple**  
Tuple adalah daftar elemen yang tersusun secara **terurut dan berukuran tetap (finite)**. Ini adalah struktur data yang digunakan untuk mengelompokkan beberapa nilai menjadi satu entitas. Tuple biasanya bersifat **immutable** (tidak dapat diubah setelah dibuat) dan digunakan untuk menyimpan koleksi data yang **berurutan**.  

Setiap elemen dalam tuple dapat memiliki **tipe data yang berbeda**. Satu-satunya cara untuk memodifikasi tuple adalah dengan mengubah isi dari setiap bidangnya (field). **Operator** seperti `+` dan `*` bisa digunakan pada tuple. Dalam konteks **basis data**, sebuah **record (baris dalam tabel)** juga disebut sebagai **tuple**.  

Berikut contoh kode dalam Go yang menghitung **pangkat dua (square) dan pangkat tiga (cube)** dari sebuah bilangan bulat dan mengembalikannya sebagai tuple:  

---

### **📌 Contoh Program Tuple di Go**
```go
// Paket utama yang berisi contoh kode
// dari buku Hands-On Data Structures and Algorithms with Go
package main

// Import package fmt
import (
  "fmt"
)

// Fungsi powerSeries menghitung kuadrat dan kubus dari bilangan `a`
// dan mengembalikan dua nilai sekaligus sebagai tuple.
func powerSeries(a int) (int, int) {
  return a * a, a * a * a
}

// Fungsi utama (main)
func main() {
  var square int
  var cube int

  // Memanggil fungsi powerSeries dengan parameter 3
  square, cube = powerSeries(3)

  // Menampilkan hasil perhitungan
  fmt.Println("Square:", square, "Cube:", cube)
}
```
### **📌 Cara Menjalankan Program**
Jalankan perintah berikut di terminal:  
```sh
go run tuples.go
```
Output yang dihasilkan:  
```
Square: 9 Cube: 27
```
---

### **📌 Tuple dengan Nama Variabel di Go**
Dalam Go, kita bisa memberikan **nama** pada nilai yang dikembalikan oleh fungsi untuk meningkatkan keterbacaan kode. Contohnya:  
```go
func powerSeries(a int) (square int, cube int) {
  square = a * a
  cube = square * a
  return
}
```
📌 **Kenapa tidak pakai `return square, cube`?**  
Karena variabel `square` dan `cube` sudah diberi nama, Go secara otomatis mengembalikannya saat `return` dipanggil tanpa argumen.

---

### **📌 Mengembalikan Error dalam Tuple**
Go mendukung pengembalian **error** sebagai bagian dari tuple. Contoh berikut menunjukkan bagaimana kita bisa mengembalikan error jika diperlukan:  
```go
import "errors"

func powerSeries(a int) (int, int, error) {
  if a < 0 {
    return 0, 0, errors.New("Input tidak boleh negatif")
  }

  square := a * a
  cube := square * a

  return square, cube, nil
}
```
📌 **`nil` berarti tidak ada error.** Jika ada kesalahan, kita bisa mengembalikan pesan error menggunakan `errors.New()`.  

---

### **📌 Kesimpulan**
✅ **Tuple di Go digunakan dengan return multiple values dalam fungsi.**  
✅ **Tuple bisa diberi nama untuk meningkatkan keterbacaan kode.**  
✅ **Error handling bisa dilakukan dengan menambahkan error sebagai nilai kembalian dalam tuple.** 🚀
