### **Metode Union (Union Method)**

**Union** pada sebuah himpunan (Set) adalah operasi yang menghasilkan himpunan baru yang berisi **semua elemen unik** dari dua himpunan yang digabungkan. Dengan kata lain, union menggabungkan semua elemen dari **set** dan **anotherSet** sehingga setiap elemen muncul hanya sekali, tanpa duplikasi.

---

### **Untuk Apa Union Digunakan?**
- **Menggabungkan Data:** Misalnya, jika kita memiliki satu set tag dari satu artikel dan satu set tag dari artikel lain, operasi union akan memberikan semua tag yang muncul di kedua artikel tersebut.
- **Analisis Data:** Dalam pemrosesan data, union dapat digunakan untuk menggabungkan hasil dari dua sumber data yang berbeda.
- **Optimasi Query Database:** Dalam konteks basis data, union digunakan untuk menggabungkan hasil dari dua query tanpa duplikasi.

---

### **Penjelasan Kode Union**
Berikut adalah kode yang diberikan, dengan penjelasan:

```go
// Metode Union mengembalikan sebuah unionSet yang merupakan gabungan dari set dan anotherSet.
func (set *Set) Union(anotherSet *Set) *Set {
    var unionSet = &Set{}
    unionSet.New() // Inisialisasi unionSet (misal, membuat map baru)

    // Iterasi melalui setiap key pada set pertama (set.integerMap)
    for value, _ := range set.integerMap {
        unionSet.AddElement(value) // Menambahkan setiap nilai ke unionSet
    }
    
    // Iterasi melalui setiap key pada set kedua (anotherSet.integerMap)
    for value, _ := range anotherSet.integerMap {
        unionSet.AddElement(value) // Menambahkan nilai ke unionSet (jika sudah ada, tidak ditambahkan lagi)
    }
    
    return unionSet
}
```

#### **Bagaimana Cara Kerjanya:**
1. **Membuat Himpunan Baru:**  
   - Fungsi `New()` digunakan untuk menginisialisasi `unionSet` sehingga siap menampung elemen.
   
2. **Menggabungkan Elemen dari Set Pertama:**  
   - Melalui loop, setiap elemen pada `set.integerMap` ditambahkan ke `unionSet`.
   
3. **Menggabungkan Elemen dari Set Kedua:**  
   - Kemudian, loop kedua menambahkan setiap elemen dari `anotherSet.integerMap` ke `unionSet`.
   - Karena Set tidak mengizinkan duplikasi, jika sebuah elemen sudah ada di `unionSet`, penambahan ulang tidak akan terjadi.
   
4. **Mengembalikan Union Set:**  
   - Hasilnya adalah sebuah himpunan baru yang berisi semua elemen dari kedua set.

---

### **Contoh Hasil Output:**
Setelah memanggil metode `Union` dengan parameter `anotherSet`, sebuah himpunan baru (`unionSet`) akan terbentuk.  
Misalnya, jika:
- **Set A:** {1, 2, 3}
- **Set B:** {2, 3, 4}

Maka **union** dari keduanya adalah: {1, 2, 3, 4}.

---

### **Kesimpulan:**
- **Union** adalah operasi untuk menggabungkan dua himpunan sehingga menghasilkan himpunan yang memuat **semua elemen unik** dari kedua himpunan.
- **Kode Union** bekerja dengan mengiterasi kedua himpunan dan menambahkan setiap elemen ke himpunan baru.
- **Aplikasi:** Digunakan untuk menggabungkan data dari berbagai sumber, analisis data, dan optimasi query.

Semoga penjelasan dan terjemahan ini membantu! Jika ada pertanyaan lebih lanjut, silakan tanya.
