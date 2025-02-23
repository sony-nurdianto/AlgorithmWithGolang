### **Kompleksitas Kubik (Cubic Complexity)**  
Dalam **kompleksitas kubik**, waktu pemrosesan algoritma **sebanding dengan pangkat tiga dari jumlah elemen input**.  

Pada contoh berikut, kompleksitas algoritma adalah **10 × 10 × 10 = 1.000**. **Tiga loop bersarang** memiliki maksimum **10 iterasi**. Oleh karena itu, kompleksitas kubik untuk pembaruan matriks adalah **O(n³)**.  

### **Contoh Kode Kompleksitas Kubik (O(n³)) dalam Go**  
```go
// Paket utama berisi contoh kode
// dari buku Hands-On Data Structures and Algorithms with Go
package main

// Mengimpor paket fmt
import (
    "fmt"
)

// Fungsi utama
func main() {
    var k, l, m int
    var arr [10][10][10]int

    for k = 0; k < 10; k++ {
        for l = 0; l < 10; l++ {
            for m = 0; m < 10; m++ {
                arr[k][l][m] = 1
                fmt.Println("Nilai elemen", k, l, m, "adalah", arr[k][l][m])
            }
        }
    }
}
```
### **Cara Menjalankan Kode:**  
Jalankan perintah berikut untuk mengeksekusi kode di atas:  
```
go run cubic_complexity.go
```

### **Penjelasan:**  
Kode ini menginisialisasi **array 3D** dengan ukuran **10 × 10 × 10** dan menetapkan setiap elemen matriks menjadi **1**. Karena ada **tiga loop bersarang**, jumlah total operasi adalah **n × n × n = n³**, yang merupakan ciri khas kompleksitas **O(n³)**.  

Pada bagian selanjutnya, kita akan membahas **kompleksitas logaritmik (O(log n))**. 🚀
