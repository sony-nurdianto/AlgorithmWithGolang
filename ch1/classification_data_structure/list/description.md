### **Daftar (Lists)**  
Sebuah *list* adalah urutan elemen yang dapat terhubung satu sama lain melalui sebuah tautan ke arah maju atau mundur. Setiap elemen dalam *list* juga dapat memiliki properti tambahan (*payload*). Struktur data ini merupakan salah satu tipe dasar dari *container*.  

*List* memiliki panjang yang bervariasi, dan pengembang dapat dengan mudah menambah atau menghapus elemen dibandingkan dengan *array*. Elemen-elemen dalam *list* tidak harus bersebelahan dalam memori atau disk. *Linked list* pertama kali diperkenalkan oleh **Allen Newell, Cliff Shaw, dan Herbert A. Simon** di RAND Corporation.  

Di Go, *list* dapat digunakan dengan cara berikut. Elemen dapat ditambahkan menggunakan metode `PushBack`, yang tersedia dalam paket `container/list`:  

```go
// Paket utama berisi contoh penggunaan list
// dari buku Hands-On Data Structures and Algorithms with Go
package main

// Mengimpor paket fmt dan container/list
import (
   "fmt"
   "container/list"
)

// Fungsi utama
func main() {
    var intList list.List
    intList.PushBack(11)
    intList.PushBack(23)
    intList.PushBack(34)

    for element := intList.Front(); element != nil; element = element.Next() {
        fmt.Println(element.Value.(int))
    }
}
```
Pada contoh di atas, daftar *list* diiterasi menggunakan perulangan `for`, dan nilai setiap elemen diakses menggunakan metode `Value`.  

Untuk menjalankan program, gunakan perintah berikut:  
```sh
go run list.go
```
Hasil output akan ditampilkan seperti pada screenshot berikut:  

---

Bagian selanjutnya akan membahas tentang *Tuples*.
