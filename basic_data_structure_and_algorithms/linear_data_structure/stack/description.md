**Stack**  
Stack adalah struktur data *last in, first out* (LIFO), di mana elemen terakhir yang dimasukkan akan menjadi elemen pertama yang dikeluarkan. Stack sering digunakan dalam *parser* dan algoritma pencarian jalan (*maze-solving algorithms*). Operasi umum yang dilakukan pada stack meliputi:  

- `Push` (menambahkan elemen ke atas stack)  
- `Pop` (menghapus elemen dari atas stack)  
- `Top` (melihat elemen teratas tanpa menghapusnya)  
- `GetSize` (mendapatkan ukuran stack)  

Beberapa contoh penggunaan stack dalam dunia nyata adalah dalam *syntax parsing*, *backtracking*, dan manajemen memori saat *compiling*.  

Berikut adalah contoh implementasi stack dalam Go (`stack.go`):  

```go
// Paket utama berisi contoh dari buku Hands-On Data Structures and Algorithms with Go
package main

// Mengimpor paket fmt dan strconv
import (
	"fmt"
	"strconv"
)

// Struct Element
type Element struct {
	elementValue int
}

// Method String untuk struct Element
func (element *Element) String() string {
	return strconv.Itoa(element.elementValue)
}
```

Struct `Element` memiliki atribut `elementValue`. Method `String` digunakan untuk mengembalikan nilai `elementValue` dalam bentuk string.  

Metode-metode stack seperti `New`, `Push`, `Pop`, dan `main` akan dijelaskan di bagian selanjutnya.
