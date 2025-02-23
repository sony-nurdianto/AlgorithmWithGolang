### **Pola Desain Facade**  

**Facade** digunakan untuk mengabstraksi antarmuka subsistem dengan bantuan kelas perantara (*helper*). Pola desain **Facade** diterapkan ketika jumlah antarmuka dalam suatu sistem meningkat dan menyebabkan sistem menjadi lebih kompleks. Facade bertindak sebagai titik masuk (*entry point*) ke berbagai subsistem dan menyederhanakan ketergantungan antara sistem yang berbeda.  

Pola ini menyediakan antarmuka yang menyembunyikan detail implementasi dari kode di baliknya.  

---

### **Karakteristik Pola Facade**  
- **Mendukung Prinsip Loose Coupling**  
  - Mengurangi ketergantungan langsung antara klien dan subsistem.  
- **Menyederhanakan API yang Buruk**  
  - Facade dapat digunakan untuk memperbaiki API yang kurang terstruktur dengan memberikan lapisan abstraksi.  
- **Dapat Digunakan dalam SOA**  
  - Dalam *Service-Oriented Architecture* (SOA), *service facade* dapat digunakan untuk menangani perubahan kontrak dan implementasi layanan.  

---

### **Komponen dalam Pola Facade**  
1. **Facade Class**  
   - Berperan sebagai perantara antara klien dan subsistem.  
   - Meneruskan permintaan dari klien ke kelas modul (*module classes*).  
   - Menyembunyikan kompleksitas logika dan aturan dalam subsistem.  

2. **Module Classes**  
   - Mengimplementasikan perilaku dan fungsionalitas spesifik dari subsistem.  

3. **Client**  
   - Menggunakan metode yang disediakan oleh *Facade Class*.  
   - Tidak berinteraksi langsung dengan subsistem.  

---

### **Contoh Implementasi Facade dalam Go**  
Misalkan kita memiliki tiga kelas: **Account**, **Customer**, dan **Transaction** yang masing-masing menangani pembuatan akun, pelanggan, dan transaksi.  
Sebagai contoh, kita akan membuat **BranchManagerFacade** untuk menyederhanakan interaksi klien dengan sistem ini.

#### **1. Struktur Kelas Modul (Subsystem)**
```go
package main

import "fmt"

// Account struct
type Account struct {
    id          string
    accountType string
}

// Method untuk membuat akun
func (account *Account) create(accountType string) *Account {
    fmt.Println("Membuat akun dengan tipe:", accountType)
    account.accountType = accountType
    return account
}

// Method untuk mendapatkan akun berdasarkan ID
func (account *Account) getById(id string) *Account {
    fmt.Println("Mengambil akun berdasarkan ID:", id)
    return account
}

// Method untuk menghapus akun berdasarkan ID
func (account *Account) deleteById(id string) {
    fmt.Println("Menghapus akun dengan ID:", id)
}

// Customer struct
type Customer struct {
    name string
    id   int
}

// Method untuk membuat pelanggan baru
func (customer *Customer) create(name string) *Customer {
    fmt.Println("Membuat pelanggan dengan nama:", name)
    customer.name = name
    return customer
}

// Transaction struct
type Transaction struct {
    id            string
    amount        float32
    srcAccountId  string
    destAccountId string
}

// Method untuk membuat transaksi
func (transaction *Transaction) create(srcAccountId, destAccountId string, amount float32) *Transaction {
    fmt.Println("Membuat transaksi")
    transaction.srcAccountId = srcAccountId
    transaction.destAccountId = destAccountId
    transaction.amount = amount
    return transaction
}
```

---

#### **2. Kelas Facade (Abstraksi untuk Modul)**
```go
// BranchManagerFacade struct
type BranchManagerFacade struct {
    account     *Account
    customer    *Customer
    transaction *Transaction
}

// Constructor untuk Facade
func NewBranchManagerFacade() *BranchManagerFacade {
    return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

// Method untuk membuat pelanggan dan akun secara bersamaan
func (facade *BranchManagerFacade) createCustomerAccount(customerName, accountType string) (*Customer, *Account) {
    customer := facade.customer.create(customerName)
    account := facade.account.create(accountType)
    return customer, account
}

// Method untuk membuat transaksi
func (facade *BranchManagerFacade) createTransaction(srcAccountId, destAccountId string, amount float32) *Transaction {
    transaction := facade.transaction.create(srcAccountId, destAccountId, amount)
    return transaction
}
```

---

#### **3. Kode dalam `main.go`**
```go
func main() {
    facade := NewBranchManagerFacade()

    // Membuat pelanggan dan akun baru
    customer, account := facade.createCustomerAccount("Thomas Smith", "Savings")
    fmt.Println("Pelanggan:", customer.name)
    fmt.Println("Tipe Akun:", account.accountType)

    // Membuat transaksi
    transaction := facade.createTransaction("21456", "87345", 1000)
    fmt.Println("Jumlah Transaksi:", transaction.amount)
}
```

---

### **Cara Menjalankan Program**
1. Simpan kode di file `facade.go`
2. Jalankan perintah:
   ```sh
   go run facade.go
   ```

---

### **Kesimpulan**
- Pola **Facade** membantu menyederhanakan interaksi dengan subsistem yang kompleks.
- Klien hanya berkomunikasi dengan *Facade Class* tanpa harus memahami logika dari subsistem.
- Berguna untuk meningkatkan desain API dan menjaga keterpisahan antara klien dan sistem backend.
- Cocok digunakan dalam sistem besar, termasuk **SOA** dan **Microservices**.

Di bagian selanjutnya, kita akan membahas **Flyweight Pattern**. 🚀
