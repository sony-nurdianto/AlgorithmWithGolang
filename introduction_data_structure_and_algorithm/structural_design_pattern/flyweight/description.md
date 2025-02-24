**Flyweight**  

Flyweight digunakan untuk mengelola keadaan objek yang memiliki variasi tinggi. Pola ini memungkinkan kita untuk berbagi bagian umum dari keadaan objek di antara banyak objek, daripada menyimpan semuanya di setiap objek secara individual. Data objek yang bervariasi disebut sebagai **extrinsic state** (keadaan ekstrinsik), sedangkan sisanya disebut **intrinsic state** (keadaan intrinsik). Data ekstrinsik diberikan ke metode flyweight dan tidak pernah disimpan di dalamnya.  

Pola Flyweight membantu mengurangi penggunaan memori secara keseluruhan dan overhead saat menginisialisasi objek. Pola ini juga memungkinkan hubungan antar kelas dan membantu menekan penggunaan memori ke tingkat yang lebih terkendali.  

Flyweight bersifat **immutable** (tidak bisa diubah). Objek dengan nilai tetap (**value objects**) adalah contoh yang baik dari pola Flyweight. Objek flyweight dapat dibuat dalam mode **single-thread**, memastikan hanya ada satu instance untuk setiap nilai. Dalam skenario multi-thread, beberapa instance bisa dibuat, tergantung pada kriteria kesamaan dari objek flyweight.  

---

### Komponen dalam Pola Flyweight  

1. **FlyWeight Interface** → Menyediakan metode yang memungkinkan flyweight untuk mendapatkan dan memproses keadaan ekstrinsik.  
2. **ConcreteFlyWeight** → Mengimplementasikan antarmuka FlyWeight dan merepresentasikan objek flyweight.  
3. **FlyWeightFactory** → Digunakan untuk membuat dan mengelola objek flyweight. Klien memanggil FlyWeightFactory untuk mendapatkan objek flyweight.  
4. **Client Classes** → Kelas-kelas yang menggunakan flyweight.  

---

### Contoh Implementasi  

Misalkan kita memiliki **DataTransferObject** sebagai antarmuka dengan metode `getId()`. **DataTransferObjectFactory** digunakan untuk membuat objek DTO berdasarkan jenisnya, seperti customer, employee, manager, dan address.  

#### Implementasi `DataTransferObjectFactory`  

```go
package main

import (
	"fmt"
)

// DataTransferObjectFactory struct
type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

// Method getDataTransferObject pada DataTransferObjectFactory
func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	var dto = factory.pool[dtoType]

	if dto == nil {
		fmt.Println("Membuat DTO baru untuk tipe: " + dtoType)
		switch dtoType {
		case "customer":
			factory.pool[dtoType] = Customer{id: "1"}
		case "employee":
			factory.pool[dtoType] = Employee{id: "2"}
		case "manager":
			factory.pool[dtoType] = Manager{id: "3"}
		case "address":
			factory.pool[dtoType] = Address{id: "4"}
		}
		dto = factory.pool[dtoType]
	}

	return dto
}
```

---

### Implementasi `DataTransferObject` dan Struktur Lainnya  

```go
// Interface DataTransferObject
type DataTransferObject interface {
	getId() string
}

// Struct Customer
type Customer struct {
	id   string // ID unik
	name string
	ssn  string
}

// Method getId pada Customer
func (customer Customer) getId() string {
	return customer.id
}

// Struct Employee
type Employee struct {
	id   string
	name string
}

// Method getId pada Employee
func (employee Employee) getId() string {
	return employee.id
}

// Struct Manager
type Manager struct {
	id   string
	name string
	dept string
}

// Method getId pada Manager
func (manager Manager) getId() string {
	return manager.id
}

// Struct Address
type Address struct {
	id         string
	streetLine1 string
	streetLine2 string
	state      string
	city       string
}

// Method getId pada Address
func (address Address) getId() string {
	return address.id
}
```

---

### Fungsi `main()`  

```go
func main() {
	var factory = DataTransferObjectFactory{make(map[string]DataTransferObject)}

	var customer DataTransferObject = factory.getDataTransferObject("customer")
	fmt.Println("Customer", customer.getId())

	var employee DataTransferObject = factory.getDataTransferObject("employee")
	fmt.Println("Employee", employee.getId())

	var manager DataTransferObject = factory.getDataTransferObject("manager")
	fmt.Println("Manager", manager.getId())

	var address DataTransferObject = factory.getDataTransferObject("address")
	fmt.Println("Address", address.getId())
}
```

---

### Cara Menjalankan  

Kompilasi dan jalankan program dengan perintah berikut:

```sh
go run flyweight.go
```

---

### Output  

```
Membuat DTO baru untuk tipe: customer
Customer 1
Membuat DTO baru untuk tipe: employee
Employee 2
Membuat DTO baru untuk tipe: manager
Manager 3
Membuat DTO baru untuk tipe: address
Address 4
```

---

### Kesimpulan  

Pola Flyweight berguna untuk menghemat penggunaan memori dengan berbagi data yang sama di antara beberapa objek, daripada menyimpannya di setiap instance secara terpisah. Ini sangat bermanfaat untuk aplikasi yang membutuhkan banyak objek dengan properti yang serupa, seperti sistem caching atau rendering grafis.  

Selanjutnya, kita bisa mempelajari pola **Private Class** dan **Proxy Data Patterns** untuk mengelola akses terhadap objek dengan lebih aman dan efisien. 🚀
