### **Algoritma Brute Force**  

Sebuah **algoritma brute force** menyelesaikan masalah berdasarkan pernyataan dan definisi masalah secara langsung. Algoritma brute force yang umum digunakan untuk pencarian dan pengurutan adalah **pencarian sekuensial (sequential search)** dan **selection sort**.  

**Pencarian menyeluruh (exhaustive search)** juga merupakan salah satu jenis algoritma brute force, di mana solusi diperoleh dari **sekumpulan kandidat solusi** dengan properti tertentu. Ruang pencarian disebut **ruang keadaan (state space)** atau **ruang kombinatorial (combinatorial space)**, yang terdiri dari **permutasi, kombinasi, atau subset**.  

#### **Karakteristik Algoritma Brute Force**  
✅ **Dapat diterapkan secara luas** dan **sederhana** untuk menyelesaikan masalah yang kompleks.  
✅ **Digunakan dalam pencarian, string matching, dan perkalian matriks**.  
✅ **Dapat menyelesaikan tugas komputasi tunggal** dengan cara langsung.  
❌ **Tidak efisien** dalam banyak kasus.  
❌ **Cenderung lambat dan tidak optimal** dibandingkan algoritma yang lebih canggih.  

---

### **Contoh Implementasi Brute Force dalam Go**  
Kode berikut menunjukkan **pencarian brute force** dalam array menggunakan metode **linear search**:  

```go
// main package untuk contoh brute force
package main

// Mengimpor fmt package
import (
    "fmt"
)

// Fungsi findElement untuk mencari elemen k dalam array
func findElement(arr [10]int, k int) bool {
    for i := 0; i < 10; i++ {
        if arr[i] == k {
            return true
        }
    }
    return false
}

// Fungsi utama (main)
func main() {
    var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
    
    // Mencari angka 10 dalam array (hasilnya false)
    var check bool = findElement(arr, 10)
    fmt.Println(check)

    // Mencari angka 9 dalam array (hasilnya true)
    var check2 bool = findElement(arr, 9)
    fmt.Println(check2)
}
```

#### **Analisis Kompleksitas**  
- Algoritma ini menggunakan **pencarian linear (linear search)**.  
- Kompleksitas waktu dalam kasus terburuk adalah **O(n)**, karena kita harus mengecek setiap elemen dalam array satu per satu.  

---

### **Kesimpulan**  
- **Algoritma brute force** bekerja dengan mencoba **semua kemungkinan solusi** tanpa optimasi tambahan.  
- **Kelebihan**: Sederhana dan mudah diterapkan.  
- **Kekurangan**: Tidak efisien, lambat dalam menangani dataset besar.  
- **Digunakan dalam skenario** seperti pencarian dalam array, pencocokan string, dan perkalian matriks sederhana.  

📌 **Selanjutnya**, kita akan membahas **algoritma Divide and Conquer**, yang jauh lebih efisien dalam menyelesaikan berbagai masalah! 🚀
