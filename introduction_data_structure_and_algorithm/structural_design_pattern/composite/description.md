### **Composite Pattern - Penjelasan dalam Bahasa Indonesia**  

**Composite** adalah pola desain yang memungkinkan sekelompok objek yang serupa untuk dikelola dalam satu objek tunggal. Dalam pola ini, objek-objek disimpan dalam struktur pohon untuk mempertahankan hierarki keseluruhannya. Pola *composite* sering digunakan untuk mengelola koleksi objek yang bersifat hierarkis.  

Pola ini dirancang untuk menangani koleksi objek yang heterogen, memungkinkan penambahan tipe objek baru tanpa mengubah antarmuka (*interface*) atau kode klien (*client code*). Beberapa contoh penggunaan pola ini meliputi:  
- Struktur UI pada web (misalnya layout dengan elemen-elemen bertingkat).  
- Struktur pohon direktori.  
- Manajemen karyawan dalam departemen yang berbeda.  

Pola *composite* memberikan mekanisme untuk mengakses objek individual maupun kelompok dengan cara yang seragam.  

---

### **Komponen dalam Composite Pattern**  
Pola ini terdiri dari beberapa elemen utama:  
1. **Component Interface**  
   - Mendefinisikan perilaku dasar dari semua objek dalam pola ini.  
   - Menyediakan metode untuk mengakses elemen-elemen dalam komposit.  
2. **Component Class**  
   - Mengimplementasikan antarmuka *component*.  
3. **Composite**  
   - Kelas yang mengandung beberapa *component*, baik dalam bentuk daun (*leaf*) maupun cabang (*branch*).  
   - Biasanya memiliki hubungan *one-to-many* dengan elemen-elemen lainnya.  
4. **Client**  
   - Berinteraksi dengan *component interface* untuk menjalankan metode dalam komposit.  

---

### **Contoh Implementasi dalam Golang**  
Berikut adalah contoh implementasi pola *composite* dalam bahasa Go:

```go
package main

import (
	"fmt"
)

// IComposite interface
type IComposite interface {
	perform()
}

// Leaflet struct (sebagai "leaf" dalam pola Composite)
type Leaflet struct {
	name string
}

// Method perform untuk Leaflet
func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet: " + leaf.name)
}

// Branch struct (sebagai "composite" dalam pola Composite)
type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

// Method perform untuk Branch, yang memanggil perform pada leaf dan sub-branch
func (branch *Branch) perform() {
	fmt.Println("Branch: " + branch.name)

	// Menjalankan perform pada semua leaf dalam branch ini
	for _, leaf := range branch.leafs {
		leaf.perform()
	}

	// Menjalankan perform pada semua sub-branch dalam branch ini
	for _, subBranch := range branch.branches {
		subBranch.perform()
	}
}

// Method untuk menambahkan leaf ke dalam branch
func (branch *Branch) addLeaf(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

// Method untuk menambahkan sub-branch ke dalam branch
func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

// Method untuk mendapatkan daftar leaf dalam branch
func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

// Fungsi utama (main)
func main() {
	// Membuat branch utama
	var branch1 = &Branch{name: "branch 1"}

	// Membuat leaf
	var leaf1 = Leaflet{name: "leaf 1"}
	var leaf2 = Leaflet{name: "leaf 2"}

	// Membuat sub-branch
	var branch2 = Branch{name: "branch 2"}

	// Menambahkan leaf ke branch1
	branch1.addLeaf(leaf1)
	branch1.addLeaf(leaf2)

	// Menambahkan branch2 ke branch1
	branch1.addBranch(branch2)

	// Menjalankan perform pada branch1 (akan memanggil perform pada semua komponennya)
	branch1.perform()
}
```

---

### **Penjelasan Kode**  
1. **`IComposite`** adalah antarmuka yang memiliki metode `perform()`, yang akan digunakan oleh setiap elemen dalam pola ini.  
2. **`Leaflet`** adalah kelas yang mewakili elemen terkecil dalam hierarki (daun dalam pohon).  
3. **`Branch`** adalah kelas yang bisa berisi beberapa `Leaflet` dan `Branch` lainnya (sub-branch).  
4. **Metode `perform()`** dalam `Branch` akan menjalankan `perform()` pada setiap `Leaflet` dan `Branch` di dalamnya secara rekursif.  
5. **Metode `addLeaf()`** digunakan untuk menambahkan *leaf* ke dalam branch.  
6. **Metode `addBranch()`** digunakan untuk menambahkan sub-branch ke dalam branch utama.  
7. **Di fungsi `main()`**, kita membuat struktur pohon dengan satu *branch* utama (`branch1`), yang memiliki dua *leaf* (`leaf1`, `leaf2`) dan satu sub-branch (`branch2`).  
8. Ketika kita memanggil `branch1.perform()`, program akan mencetak semua *branch* dan *leaf* yang ada dalam hierarki tersebut.  

---

### **Output yang Dihasilkan**  
Ketika program dijalankan, output yang dihasilkan adalah:

```
Branch: branch 1
Leaflet: leaf 1
Leaflet: leaf 2
Branch: branch 2
```

Ini menunjukkan bahwa pola *composite* berhasil mengelola hierarki objek dengan baik.

---

### **Kapan Menggunakan Composite Pattern?**  
Gunakan *Composite Pattern* ketika:  
✅ Anda memiliki struktur data yang bersifat hierarkis (*tree-like structure*).  
✅ Anda ingin memperlakukan objek tunggal dan kumpulan objek dengan cara yang seragam.  
✅ Anda ingin memungkinkan ekspansi sistem tanpa perlu banyak perubahan pada kode yang sudah ada.  

Beberapa contoh nyata dari penggunaan *Composite Pattern* termasuk:  
- **Sistem manajemen file**, di mana folder dapat berisi file maupun folder lainnya.  
- **Struktur organisasi perusahaan**, di mana departemen bisa memiliki sub-departemen dan karyawan.  
- **Pembuatan UI/UX**, di mana komponen bisa terdiri dari beberapa sub-komponen.  

---

### **Kesimpulan**  
Pola *Composite* memungkinkan kita untuk menangani objek tunggal dan kumpulan objek secara seragam dengan menggunakan struktur pohon. Dengan menggunakan pola ini, kita bisa memperlakukan objek individu dan grup secara serupa tanpa harus menangani mereka dengan cara yang berbeda.  

Pola ini sangat berguna untuk berbagai kasus seperti UI Layouts, sistem manajemen file, dan hierarki organisasi.  

Selanjutnya, kita bisa melihat **Decorator Pattern** untuk memahami bagaimana cara menambahkan perilaku secara dinamis pada objek. 🚀
