### **📌 Apa Itu Doubly Linked List?**  

**Doubly Linked List** (DLL) adalah struktur data berbasis *linked list* di mana setiap node memiliki **dua referensi (pointer)**:  
1. **Pointer ke elemen sebelumnya (`Prev`)**  
2. **Pointer ke elemen berikutnya (`Next`)**  

Berbeda dengan *singly linked list* (yang hanya punya `Next`), **doubly linked list bisa ditelusuri maju (`Next`) dan mundur (`Prev`)**.  

---

### **🛠 Struktur Node dalam Doubly Linked List**
Setiap node dalam **doubly linked list** terdiri dari:
- **Data** (nilai yang disimpan)
- **Pointer ke elemen berikutnya (`Next`)**
- **Pointer ke elemen sebelumnya (`Prev`)**  

📌 **Struktur dasar:**
```
[Prev] <-> [Data] <-> [Next]
```
Misalnya, jika kita memiliki DLL dengan elemen `10 → 20 → 30`, maka hubungan antar-node seperti ini:
```
nil <- [10] <-> [20] <-> [30] -> nil
```

---

### **🛠 Implementasi Doubly Linked List di Go (`container/list`)**
Go memiliki **Doubly Linked List** bawaan di paket `container/list`.

#### **📌 Contoh: Menambahkan Elemen ke DLL di Go**
```go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Membuat linked list baru
	l := list.New()

	// Menambahkan elemen ke dalam list
	l.PushBack(10)  // 10
	l.PushBack(20)  // 10 <-> 20
	l.PushFront(5)  // 5 <-> 10 <-> 20

	// Iterasi dari awal ke akhir
	fmt.Println("Iterasi maju:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Iterasi dari akhir ke awal
	fmt.Println("Iterasi mundur:")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
```
📌 **Output:**
```
Iterasi maju:
5
10
20
Iterasi mundur:
20
10
5
```

---

### **📌 Kelebihan Doubly Linked List**
✅ **Navigasi dua arah** → Bisa maju (`Next`) dan mundur (`Prev`).  
✅ **Mudah untuk menyisipkan/menghapus elemen di tengah** → Tidak perlu menggeser elemen seperti array/slice.  

### **📌 Kekurangan Doubly Linked List**
❌ **Menggunakan lebih banyak memori** → Setiap node memiliki tambahan pointer (`Prev` & `Next`).  
❌ **Lebih kompleks dibandingkan *singly linked list***  

---

### **📌 Kapan Menggunakan Doubly Linked List?**
🔹 Jika sering menambah/menghapus elemen **di tengah struktur data**.  
🔹 Jika perlu **navigasi maju & mundur** dalam daftar elemen.  
🔹 Jika tidak tahu **ukuran pasti** dari data yang akan disimpan.  

---

### **📌 Kapan Tidak Menggunakan Doubly Linked List?**
🔸 Jika sering mengakses elemen **berdasarkan indeks** → *Gunakan slice*.  
🔸 Jika memori menjadi masalah → *Gunakan singly linked list atau slice*.  

---

### **Kesimpulan**  
Doubly linked list di Go (`container/list`) cocok untuk **data yang sering dimodifikasi di tengah daftar** atau **butuh traversal maju-mundur**. Namun, untuk **akses cepat berdasarkan indeks**, lebih baik gunakan *slice* atau *array*. 🚀
