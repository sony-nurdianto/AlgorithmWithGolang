## **Apa Itu Himpunan? (Set)**
Himpunan (set) adalah kumpulan elemen unik tanpa urutan tertentu. Dalam berbagai konteks, himpunan memiliki peran yang berbeda. Mari kita bahas dari beberapa perspektif.

---

### **1. Himpunan dalam Matematika**  
Dalam **matematika**, himpunan adalah kumpulan objek yang tidak berulang dan dikelompokkan berdasarkan kriteria tertentu.  

**Contoh:**
- Himpunan bilangan genap kurang dari 10:  
  **S = {2, 4, 6, 8}**
- Himpunan huruf vokal:  
  **V = {a, e, i, o, u}**

**Operasi pada Himpunan:**
1. **Union (∪)** → Menggabungkan dua himpunan  
   `{1, 2, 3} ∪ {3, 4, 5} = {1, 2, 3, 4, 5}`
2. **Intersection (∩)** → Irisan atau elemen yang sama  
   `{1, 2, 3} ∩ {3, 4, 5} = {3}`
3. **Difference (-)** → Selisih antara dua himpunan  
   `{1, 2, 3} - {3, 4, 5} = {1, 2}`
4. **Subset (⊆)** → Cek apakah himpunan adalah bagian dari himpunan lain  
   `{1, 2} ⊆ {1, 2, 3}` ✅

---

### **2. Himpunan dalam Ilmu Komputer (Data Structure)**
Dalam pemrograman, **himpunan** adalah tipe data yang menyimpan elemen **unik**. Biasanya digunakan untuk menyimpan data tanpa duplikasi dan melakukan pencarian cepat.

**Implementasi dalam Go:**
Go tidak memiliki tipe data bawaan untuk **Set**, tetapi bisa menggunakan **map** untuk meniru perilaku Set.

```go
package main

import "fmt"

type Set struct {
    data map[int]bool
}

func (s *Set) New() {
    s.data = make(map[int]bool)
}

func (s *Set) Add(value int) {
    s.data[value] = true
}

func (s *Set) Remove(value int) {
    delete(s.data, value)
}

func (s *Set) Contains(value int) bool {
    return s.data[value]
}

func main() {
    s := &Set{}
    s.New()
    s.Add(1)
    s.Add(2)
    s.Add(3)
    fmt.Println("Apakah 2 ada dalam set?", s.Contains(2)) // true
    s.Remove(2)
    fmt.Println("Apakah 2 ada dalam set setelah dihapus?", s.Contains(2)) // false
}
```
**Keunggulan Set dalam Pemrograman:**
- **Tidak ada duplikasi** (misalnya, untuk menyimpan daftar ID unik)
- **Pencarian cepat** (O(1) dalam map)
- **Operasi matematis seperti Union, Intersect, dan Difference lebih efisien**

---

### **3. Himpunan dalam Basis Data**
Di **SQL dan basis data**, himpunan digunakan dalam operasi pencarian dan penyaringan data.

**Contoh Query SQL dengan Operasi Himpunan:**
```sql
SELECT name FROM employees
INTERSECT
SELECT name FROM managers;
```
Query ini mencari nama karyawan yang juga merupakan manajer.

---

### **4. Himpunan dalam Dunia Nyata**
Di kehidupan sehari-hari, konsep himpunan digunakan dalam berbagai aplikasi:

- **Tag dalam Blog atau Media Sosial:**  
  Setiap artikel atau postingan bisa memiliki **tag unik** seperti `{Golang, Programming, Data Structure}`.
  
- **Daftar Teman Bersama di Media Sosial:**  
  Jika A berteman dengan `{B, C, D}` dan B berteman dengan `{C, D, E}`, maka teman bersama mereka adalah `{C, D}`.

- **Analisis Pasar dalam E-commerce:**  
  Produk yang sering dibeli bersama oleh pelanggan bisa dianalisis menggunakan **set intersection**.

---

## **Kesimpulan**
Himpunan memiliki berbagai kegunaan di **matematika, pemrograman, basis data, dan dunia nyata**. Dalam **Go**, meskipun tidak ada tipe bawaan untuk Set, kita bisa menggunakan **map[int]bool** untuk menyimpan data unik dan melakukan operasi himpunan dengan cepat.
