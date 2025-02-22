### **Private Class Data**  
Pola *Private Class Data* digunakan untuk mengamankan data dalam sebuah kelas dengan cara *encapsulation*. Pola ini memastikan bahwa data hanya dapat diinisialisasi saat objek dibuat, dan hak akses tulis (*write privileges*) terhadap data dalam kelas dibatasi.  

Pola ini mencegah kebocoran informasi dengan menyimpannya dalam kelas yang mempertahankan *state*. Salah satu kasus penggunaan utama pola ini adalah saat ingin mengamankan inisialisasi data dalam suatu kelas.  

Sebagai contoh, kita memiliki kelas `Account` yang berisi informasi akun dan nama pelanggan. Atribut `AccountDetails` bersifat privat (hanya bisa diakses dalam *package* yang sama), sedangkan `CustomerName` bersifat publik. Saat dilakukan *JSON marshaling* terhadap `Account`, hanya `CustomerName` yang dapat diekspor, sementara `AccountDetails` tetap tersembunyi.

---

### **Contoh Implementasi dalam Go**  
```go
package main

import (
	"encoding/json"
	"fmt"
)

// AccountDetails struct (data privat)
type AccountDetails struct {
	id          string
	accountType string
}

// Account struct (kelas utama)
type Account struct {
	details      *AccountDetails
	CustomerName string // atribut publik
}

// Metode untuk mengatur detail akun
func (account *Account) setDetails(id string, accountType string) {
	account.details = &AccountDetails{id, accountType}
}

// Metode untuk mendapatkan ID akun
func (account *Account) getId() string {
	return account.details.id
}

// Metode untuk mendapatkan tipe akun
func (account *Account) getAccountType() string {
	return account.details.accountType
}

func main() {
	// Inisialisasi akun dengan nama pelanggan
	var account *Account = &Account{CustomerName: "John Smith"}
	account.setDetails("4532", "current")

	// Melakukan JSON marshaling (hanya CustomerName yang diekspor)
	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Private Class hidden:", string(jsonAccount))

	// Mengakses data privat melalui metode getter
	fmt.Println("Account Id:", account.getId())
	fmt.Println("Account Type:", account.getAccountType())
}
```
---

### **Penjelasan:**
1. **`AccountDetails`** menyimpan data akun seperti `id` dan `accountType`. Struktur ini tidak diekspor secara langsung karena tidak ada *exported field*.
2. **`Account`** memiliki atribut `CustomerName` yang bersifat publik, sementara `details` adalah referensi ke `AccountDetails`.
3. **Metode `setDetails`** digunakan untuk menginisialisasi data privat `AccountDetails`.
4. **Metode `getId` dan `getAccountType`** bertindak sebagai *getter* untuk mengambil nilai `id` dan `accountType`, karena atribut ini tidak dapat diakses langsung dari luar.
5. Saat dilakukan **JSON marshaling**, hanya `CustomerName` yang muncul dalam hasil JSON karena `details` tidak memiliki *exported field*.

---

### **Output yang Diharapkan:**
```json
Private Class hidden: {"CustomerName":"John Smith"}
Account Id: 4532
Account Type: current
```

---

### **Kesimpulan**  
Pola *Private Class Data* sangat berguna untuk menjaga data tetap privat dalam sebuah kelas. Dalam Go, kita bisa mencapainya dengan:  
✅ **Menggunakan struktur privat untuk menyimpan data sensitif**  
✅ **Mengontrol akses melalui metode getter**  
✅ **Menghindari kebocoran data saat proses serialisasi**  

Pola ini sering digunakan untuk meningkatkan keamanan dan memastikan bahwa data hanya dapat diubah melalui metode yang telah ditentukan. 🚀
