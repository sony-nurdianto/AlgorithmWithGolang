### **Algoritma Backtracking**  

**Algoritma Backtracking** menyelesaikan masalah dengan **membangun solusi secara bertahap**. Algoritma ini mengevaluasi berbagai kemungkinan dan menggunakan **rekursi** untuk menentukan langkah berikutnya dalam solusi.  

Pendekatan **Backtracking** dapat berupa:  
✅ **Chronological Backtracking** → Mengembalikan langkah sebelumnya secara berurutan jika menemukan jalan buntu.  
✅ **Graph Traversal Backtracking** → Menelusuri berbagai jalur yang memungkinkan, seperti dalam graf atau pohon keputusan.  

---

### **Bagaimana Backtracking Bekerja?**  
1️⃣ **Menentukan kandidat solusi yang memungkinkan**  
2️⃣ **Memeriksa validitas solusi sementara**  
3️⃣ **Jika solusi valid → lanjutkan ke langkah berikutnya**  
4️⃣ **Jika solusi tidak valid → mundur (backtrack) dan coba pilihan lain**  

Teknik ini digunakan dalam berbagai **masalah kombinatorial** dan **optimasi**, seperti:  
✔ **Mencari nilai dalam tabel tidak terurut**  
✔ **Parsing dan pemrosesan aturan**  
✔ **Masalah Knapsack**  
✔ **Optimasi kombinatorial (contoh: Sudoku, N-Queens, Traveling Salesman Problem)**  

---

### **Contoh Implementasi Backtracking dalam Go**  
Berikut adalah contoh algoritma **Backtracking** yang mencari kombinasi elemen dalam array yang jumlahnya sama dengan **18**. Jika jumlahnya melebihi **18**, algoritma akan **backtrack**:  

```go
// main package untuk contoh Backtracking
package main

// Mengimpor fmt package
import (
    "fmt"
)

// Fungsi rekursif untuk mencari kombinasi elemen yang jumlahnya sama dengan k
func findElementsWithSum(arr [10]int, combinations [19]int, size int, k int, addValue int, l int, m int) int {
    var num int = 0
    if addValue > k {
        return -1
    }
    if addValue == k {
        num = num + 1
        var p int = 0
        for p = 0; p < m; p++ {
            fmt.Printf("%d,", arr[combinations[p]])
        }
        fmt.Println(" ")
    }
    var i int
    for i = l; i < size; i++ {
        combinations[m] = l
        findElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
        l = l + 1
    }
    return num
}

// Fungsi utama (main)
func main() {
    var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
    var addedSum int = 18
    var combinations [19]int
    findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)
}
```

---

### **Analisis Kompleksitas**  
🔹 **Worst Case Complexity**: **O(2ⁿ)** → Semua kombinasi elemen dievaluasi.  
🔹 **Best Case Complexity**: **O(n)** → Jika solusi ditemukan dengan cepat.  

**Backtracking lebih cepat dari Brute Force**, karena secara **efisien mengeliminasi banyak solusi yang tidak mungkin** sebelum mencoba semua kemungkinan.  

---

### **Kesimpulan**  
📌 **Backtracking sangat berguna** untuk menyelesaikan masalah optimasi dan kombinatorial.  
📌 **Lebih efisien dibanding Brute Force**, tetapi masih memiliki **kompleksitas eksponensial** dalam kasus terburuk.  
📌 **Digunakan dalam masalah seperti Sudoku, N-Queens, dan Traveling Salesman Problem (TSP).**  

📢 **Selanjutnya**, kita akan membahas **algoritma Greedy**, yang sering digunakan dalam optimasi keputusan cepat! 🚀
