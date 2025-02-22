Di frontend, **Flyweight Pattern** bisa digunakan untuk mengoptimalkan **komponen UI yang sering dipakai** (misalnya ikon, tombol, atau elemen dengan state serupa), sedangkan di backend, penggunaannya memang kurang umum tetapi tetap ada beberapa kasus di mana pattern ini berguna.  

---

### **💡 Kapan Flyweight Pattern Berguna di Backend?**

#### **1️⃣ Cache dan Object Pooling**
✅ **Kasus:** Jika ada banyak objek dengan state yang mirip dan hanya beberapa bagian kecil yang berubah.  
🚀 **Solusi:** Gunakan Flyweight agar objek yang sama bisa digunakan kembali.  

🔹 **Contoh:**  
Misalnya, dalam sistem laporan keuangan, setiap laporan memiliki banyak **kategori transaksi** yang berulang-ulang (misalnya *Food*, *Transport*, *Entertainment*). Daripada membuat objek baru setiap kali laporan dibuat, kita bisa menggunakan Flyweight Pattern untuk berbagi objek kategori.  

```go
package main

import (
	"fmt"
)

// Flyweight Interface
type Category interface {
	GetName() string
}

// ConcreteFlyweight
type TransactionCategory struct {
	name string
}

func (c *TransactionCategory) GetName() string {
	return c.name
}

// Flyweight Factory
type CategoryFactory struct {
	pool map[string]*TransactionCategory
}

func NewCategoryFactory() *CategoryFactory {
	return &CategoryFactory{pool: make(map[string]*TransactionCategory)}
}

func (cf *CategoryFactory) GetCategory(name string) *TransactionCategory {
	if _, exists := cf.pool[name]; !exists {
		fmt.Println("Creating new category:", name)
		cf.pool[name] = &TransactionCategory{name: name}
	}
	return cf.pool[name]
}

func main() {
	factory := NewCategoryFactory()

	// Menggunakan kategori yang sama tanpa membuat objek baru
	foodCategory1 := factory.GetCategory("Food")
	foodCategory2 := factory.GetCategory("Food")

	fmt.Println("Kategori 1:", foodCategory1.GetName())
	fmt.Println("Kategori 2:", foodCategory2.GetName())

	// Cek apakah objeknya sama
	fmt.Println("Apakah foodCategory1 == foodCategory2?", foodCategory1 == foodCategory2)
}
```
**Output:**
```
Creating new category: Food
Kategori 1: Food
Kategori 2: Food
Apakah foodCategory1 == foodCategory2? true
```
👉 **Kategori "Food" hanya dibuat sekali dan digunakan kembali untuk semua transaksi yang membutuhkan kategori tersebut.**  

---

#### **2️⃣ Optimasi Database Connection Pooling**
✅ **Kasus:** Jika aplikasi backend sering membuka koneksi ke database yang sama.  
🚀 **Solusi:** Alih-alih membuat koneksi baru setiap kali, kita bisa menggunakan Flyweight untuk berbagi koneksi antar request.  

🔹 **Tanpa Flyweight (Setiap request buat koneksi baru)**
```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=youruser dbname=yourdb sslmode=disable"

	db1, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db1.Close()

	db2, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db2.Close()

	fmt.Println("Apakah db1 == db2?", db1 == db2) // Akan selalu false
}
```
👉 **Setiap request membuat koneksi baru, boros resource!**  

🔹 **Dengan Flyweight (Gunakan pool untuk berbagi koneksi)**
```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

// Flyweight Factory
type DBConnectionPool struct {
	mu        sync.Mutex
	instances map[string]*sql.DB
}

func NewDBConnectionPool() *DBConnectionPool {
	return &DBConnectionPool{
		instances: make(map[string]*sql.DB),
	}
}

func (p *DBConnectionPool) GetConnection(connStr string) *sql.DB {
	p.mu.Lock()
	defer p.mu.Unlock()

	if db, exists := p.instances[connStr]; exists {
		return db
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	p.instances[connStr] = db
	return db
}

func main() {
	pool := NewDBConnectionPool()
	connStr := "user=youruser dbname=yourdb sslmode=disable"

	db1 := pool.GetConnection(connStr)
	db2 := pool.GetConnection(connStr)

	fmt.Println("Apakah db1 == db2?", db1 == db2) // Akan true
}
```
**Output:**
```
Apakah db1 == db2? true
```
👉 **Koneksi database hanya dibuat sekali dan dipakai ulang oleh semua request.**  

---

#### **3️⃣ Cache Data di Memory (Redis, In-Memory Store)**
✅ **Kasus:** Jika ada banyak objek yang sering diakses dengan nilai yang tidak sering berubah.  
🚀 **Solusi:** Gunakan Flyweight agar objek yang sama bisa digunakan kembali tanpa harus membaca dari sumber eksternal.  

🔹 **Contoh:** Misalkan kita sering mengambil data negara dari database, lebih baik menyimpannya dalam cache:  
```go
package main

import (
	"fmt"
	"sync"
)

// Flyweight Factory
type CountryFactory struct {
	mu    sync.Mutex
	cache map[string]*Country
}

type Country struct {
	name string
}

func NewCountryFactory() *CountryFactory {
	return &CountryFactory{cache: make(map[string]*Country)}
}

func (f *CountryFactory) GetCountry(name string) *Country {
	f.mu.Lock()
	defer f.mu.Unlock()

	if country, exists := f.cache[name]; exists {
		return country
	}

	fmt.Println("Fetching country from database:", name)
	newCountry := &Country{name: name}
	f.cache[name] = newCountry
	return newCountry
}

func main() {
	factory := NewCountryFactory()

	indonesia1 := factory.GetCountry("Indonesia")
	indonesia2 := factory.GetCountry("Indonesia")

	fmt.Println("Apakah indonesia1 == indonesia2?", indonesia1 == indonesia2) // true
}
```
👉 **Negara "Indonesia" hanya dibuat sekali dan digunakan ulang tanpa harus query ke database lagi.**  

---

### **Kesimpulan**
Flyweight Pattern memang tidak selalu terlihat langsung berguna di backend seperti di frontend, tetapi tetap bisa membantu dalam beberapa kasus:  
✅ **Object Caching** (misalnya kategori transaksi, negara, dsb).  
✅ **Database Connection Pooling** untuk menghindari overhead koneksi baru.  
✅ **In-Memory Data Store** seperti Redis atau cache untuk menghindari query berulang.  

Kalau butuh solusi yang lebih **otomatis** dan tidak perlu dikelola manual seperti Flyweight, biasanya backend developer lebih memilih **`sync.Pool` atau connection pooling bawaan database**. 🚀
