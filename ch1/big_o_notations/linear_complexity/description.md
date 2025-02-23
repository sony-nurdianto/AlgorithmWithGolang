### Kompleksitas Linear  
Sebuah algoritma memiliki **kompleksitas linear** jika **waktu pemrosesan** atau **ruang penyimpanan** berbanding lurus dengan jumlah elemen input yang diproses. Dalam **notasi Big O**, kompleksitas linear ditulis sebagai **O(n)**.  

Contoh algoritma **string matching**, seperti **Boyer-Moore** dan **Ukkonen**, memiliki kompleksitas **linear (O(n))**.  

Berikut adalah contoh implementasi kompleksitas **linear (O(n))** dalam kode Go:  

```go
// Paket utama berisi contoh kode
// dari buku Go Data Structures and Algorithms
package main

// Mengimpor paket fmt
import (
 "fmt"
)

// Fungsi utama
func main() {
 var m [10]int
 var k int

 for k = 0; k < 10; k++ {
   m[k] = k * 200
   fmt.Printf("Element[%d] = %d\n", k, m[k])
 }
}
```
Jalankan perintah berikut untuk mengeksekusi kode:  
```
go run linear_complexity.go
```
Kode di atas menunjukkan bahwa jumlah operasi yang dilakukan **bertambah seiring dengan bertambahnya jumlah elemen input**, yang merupakan ciri khas kompleksitas **O(n)**.  

Pada bagian berikutnya, kita akan melihat lebih dalam tentang **kompleksitas kuadratik (O(n²))**. 🚀
