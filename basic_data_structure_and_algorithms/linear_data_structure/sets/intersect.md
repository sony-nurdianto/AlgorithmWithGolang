### **Penjelasan Metode `Intersect` pada Set**  
Metode **`Intersect`** digunakan untuk mendapatkan **himpunan irisan** dari dua set. Artinya, hasilnya adalah elemen-elemen yang **ada di kedua set**.

---

### **Bagaimana Cara Kerjanya?**
1. **Membuat Set Baru (`intersectSet`)**  
   - `intersectSet` akan menyimpan elemen-elemen yang muncul di **kedua** set.
  
2. **Melakukan Iterasi pada Set Awal (`set.integerMap`)**  
   - Untuk setiap elemen di **set pertama**, dicek apakah elemen tersebut ada di **set kedua** (`anotherSet`).

3. **Menambahkan Elemen yang Cocok ke `intersectSet`**  
   - Jika elemen tersebut juga ada di **anotherSet**, maka elemen itu dimasukkan ke dalam `intersectSet`.

4. **Mengembalikan `intersectSet` sebagai hasilnya**.

---

### **Kode dengan Penjelasan**
```go
func (set *Set) Intersect(anotherSet *Set) *Set {
    var intersectSet = &Set{} // Membuat Set baru untuk menyimpan hasil irisan
    intersectSet.New()        // Menginisialisasi map di dalamnya

    for value := range set.integerMap { // Iterasi setiap elemen di set pertama
        if anotherSet.ContainsElement(value) { // Jika elemen ada di anotherSet
            intersectSet.AddElement(value)    // Tambahkan ke intersectSet
        }
    }

    return intersectSet // Kembalikan set hasil irisan
}
```

---

### **Contoh Penggunaan**
```go
package main

import "fmt"

type Set struct {
    integerMap map[int]bool
}

func (set *Set) New() {
    set.integerMap = make(map[int]bool)
}

func (set *Set) AddElement(value int) {
    set.integerMap[value] = true
}

func (set *Set) ContainsElement(value int) bool {
    return set.integerMap[value]
}

func (set *Set) Intersect(anotherSet *Set) *Set {
    var intersectSet = &Set{}
    intersectSet.New()

    for value := range set.integerMap {
        if anotherSet.ContainsElement(value) {
            intersectSet.AddElement(value)
        }
    }

    return intersectSet
}

func main() {
    setA := &Set{}
    setA.New()
    setA.AddElement(1)
    setA.AddElement(2)
    setA.AddElement(3)

    setB := &Set{}
    setB.New()
    setB.AddElement(2)
    setB.AddElement(3)
    setB.AddElement(4)

    result := setA.Intersect(setB)

    fmt.Println("Irisan dari Set A dan Set B:")
    for key := range result.integerMap {
        fmt.Println(key)
    }
}
```
**Output:**  
```
Irisan dari Set A dan Set B:
2
3
```

---

### **Kesimpulan**
- **`Intersect`** digunakan untuk menemukan elemen yang **sama** dalam dua set.
- Dalam contoh di atas, **set A** `{1, 2, 3}` dan **set B** `{2, 3, 4}` memiliki irisan `{2, 3}`.
- Metode ini berguna dalam berbagai kasus, seperti:
  - **Mencari user yang memiliki minat yang sama**
  - **Menemukan item yang ada di dua daftar**
  - **Menganalisis data statistik atau database**
