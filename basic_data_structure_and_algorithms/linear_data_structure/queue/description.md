### **Queue (Antrian) dalam Struktur Data**  

**Queue (antrian)** adalah struktur data linear yang menyimpan elemen dalam urutan tertentu atau berdasarkan **prioritas**. Elemen dalam queue diproses dengan urutan **First-In-First-Out (FIFO)**, kecuali jika menggunakan **Priority Queue**, di mana elemen dengan prioritas lebih tinggi akan diproses lebih dahulu.  

---

### **Karakteristik Queue:**
1. **Linear & Sequential** → Elemen ditambahkan di **akhir** dan dihapus dari **awal**.
2. **Operasi utama:**
   - **Enqueue** → Menambahkan elemen ke belakang antrian.
   - **Dequeue** → Menghapus elemen dari depan antrian.
   - **Peek** → Melihat elemen terdepan tanpa menghapusnya.
3. **Penggunaan dalam dunia nyata:**  
   - **Task scheduling** dalam CPU atau sistem real-time.
   - **Manajemen permintaan HTTP** dalam server.
   - **Sistem antrian pelanggan** seperti call center atau antrean bank.

---

### **Contoh Implementasi Queue dalam Go**
Berikut ini adalah contoh **Queue untuk Order (Pesanan)** yang didefinisikan sebagai **array of pointers** ke objek `Order`:

#### **Definisi Struct Queue dan Order:**
```go
// Queue adalah array berisi pointer ke Order
type Queue []*Order

// Struct Order berisi detail pesanan
type Order struct {
    priority     int
    quantity     int
    product      string
    customerName string
}
```
---

### **Penjelasan Kode**
1. **`Queue` dideklarasikan sebagai slice dari pointer ke `Order` (`Queue []*Order`)**  
   - Ini memungkinkan queue untuk menyimpan daftar pesanan (`Order`) dalam antrian.

2. **`Order` struct berisi:**
   - `priority` → Prioritas pesanan.
   - `quantity` → Jumlah item dalam pesanan.
   - `product` → Nama produk yang dipesan.
   - `customerName` → Nama pelanggan.

---

Bagian selanjutnya dalam buku akan membahas fungsi **New, Add, dan Main** untuk mengelola antrian pesanan.
