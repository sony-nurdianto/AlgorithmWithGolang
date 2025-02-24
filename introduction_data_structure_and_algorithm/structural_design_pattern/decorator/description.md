### **Decorator Pattern**  

Pola **Decorator** digunakan dalam skenario di mana kita ingin menambahkan atau menghapus tanggung jawab dari suatu kelas tanpa menggunakan pewarisan statis.  

Alih-alih membuat subclass baru untuk menambah fungsionalitas, **Decorator Pattern** memungkinkan kita menambahkan perilaku baru secara dinamis saat runtime. Pola ini juga membantu menerapkan **Single Responsibility Principle (SRP)** dengan cara yang lebih fleksibel.  

---

### **Kapan Menggunakan Decorator Pattern?**  
✔ **Menambahkan fungsionalitas baru** ke objek tanpa mengubah struktur dasarnya.  
✔ **Mengurangi kebutuhan subclassing** dengan memungkinkan perubahan dinamis.  
✔ **Cocok untuk komponen UI**, seperti jendela dan elemen grafis yang membutuhkan berbagai fitur tambahan.  

---

### **Komponen dalam Decorator Pattern**  

1. **Component Interface** – Mendefinisikan metode yang akan didekorasi.  
2. **Concrete Component** – Implementasi konkret dari antarmuka komponen.  
3. **Decorator Class** – Implementasi dari antarmuka komponen yang menambahkan fungsionalitas baru tanpa mengubah struktur utama.  

---

### **Implementasi Decorator dalam Golang**  

Misalkan kita memiliki antarmuka **IProcess** dengan metode `process()`.  
- **ProcessClass** adalah implementasi dari antarmuka ini.  
- **ProcessDecorator** adalah kelas dekorator yang menambahkan fungsionalitas tambahan di atas **ProcessClass**.  

```go
package main

import (
	"fmt"
)

// IProcess Interface
type IProcess interface {
	process()
}

// ProcessClass struct (Implementasi dari IProcess)
type ProcessClass struct{}

// Implementasi metode process() untuk ProcessClass
func (process *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

// ProcessDecorator struct (Dekorator)
type ProcessDecorator struct {
	processInstance *ProcessClass
}

// Implementasi metode process() untuk ProcessDecorator
func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Printf("ProcessDecorator process and ")
		decorator.processInstance.process()
	}
}

// Fungsi main()
func main() {
	// Membuat instance ProcessClass
	var process = &ProcessClass{}

	// Membuat instance ProcessDecorator tanpa mendekorasi apa pun
	var decorator = &ProcessDecorator{}
	decorator.process() // Output: ProcessDecorator process

	// Menambahkan instance ProcessClass ke dalam ProcessDecorator
	decorator.processInstance = process
	decorator.process() // Output: ProcessDecorator process and ProcessClass process
}
```

---

### **Penjelasan Kode**  

✔ **Step 1:**  
- `IProcess` adalah antarmuka dengan metode `process()`.  

✔ **Step 2:**  
- `ProcessClass` mengimplementasikan `IProcess` dan mencetak `"ProcessClass process"`.  

✔ **Step 3:**  
- `ProcessDecorator` memiliki referensi ke `ProcessClass`.  
- Jika tidak memiliki referensi, ia hanya mencetak `"ProcessDecorator process"`.  
- Jika memiliki referensi ke `ProcessClass`, ia mencetak `"ProcessDecorator process and ProcessClass process"`.  

✔ **Step 4:**  
- Di dalam `main()`, kita membuat **ProcessDecorator** tanpa dekorasi (hasilnya hanya `"ProcessDecorator process"`).  
- Setelah kita menambahkan **ProcessClass** ke dalam **ProcessDecorator**, kita mendapatkan hasil `"ProcessDecorator process and ProcessClass process"`.  

---

### **Cara Menjalankan Program**  

Kompilasi dan jalankan dengan perintah:  

```sh
go run decorator.go
```

---

### **Output yang Dihasilkan**  

```
ProcessDecorator process
ProcessDecorator process and ProcessClass process
```

---

### **Kesimpulan**  

✔ **Decorator Pattern** memungkinkan kita **menambah atau menghapus fungsionalitas** tanpa mengubah struktur kelas utama.  
✔ **Lebih fleksibel dibandingkan pewarisan**, karena memungkinkan modifikasi dinamis.  
✔ **Cocok untuk banyak skenario**, seperti pengembangan UI, middleware, logging, dan validasi.  

Selanjutnya, kita akan membahas **Facade Pattern**, yang menyederhanakan akses ke sistem yang kompleks. 🚀
