### **Algoritma Divide and Conquer**  

**Algoritma Divide and Conquer** bekerja dengan cara **memecah masalah kompleks menjadi masalah-masalah kecil**, lalu menyelesaikan masing-masing masalah kecil tersebut. Proses ini berlanjut hingga masalah yang dipecah menjadi **masalah yang sudah diketahui solusinya**. Kemudian, solusi dari sub-masalah tersebut digabungkan kembali untuk mendapatkan solusi akhir.  

Pendekatan ini umumnya menggunakan **rekursi** untuk menyelesaikan sub-masalah dan menggabungkan hasilnya.  

---

### **Contoh Algoritma Divide and Conquer**  
Beberapa contoh algoritma yang menggunakan strategi **Divide and Conquer**:  
✅ **Binary Search** (Pencarian biner)  
✅ **Quick Sort** (Pengurutan cepat)  
✅ **Merge Sort** (Pengurutan dengan penggabungan)  
✅ **Fast Fourier Transform (FFT)**  
✅ **Rekursi Fibonacci**  

**Kelebihan:**  
✔ Menggunakan **memori secara efisien**.  
✔ **Cocok untuk eksekusi paralel** pada **multiprocessor machines** karena sub-masalah dapat diproses secara independen.  

**Kekurangan:**  
❌ **Performanya dapat menurun** jika rekursi tidak dioptimalkan.  
❌ **Memerlukan strategi optimasi**, seperti **memoization** atau **dynamic programming**, untuk menghindari perhitungan berulang yang tidak perlu.  

---

### **Implementasi Fibonacci dengan Divide and Conquer dalam Go**  
Kode berikut menghitung bilangan Fibonacci menggunakan **rekursi**, yang merupakan contoh dasar dari Divide and Conquer:  

```go
// main package untuk contoh Divide and Conquer
package main

// Mengimpor fmt package
import (
    "fmt"
)

// Fungsi fibonacci menggunakan rekursi
func fibonacci(k int) int {
    if k <= 1 {
        return 1
    }
    return fibonacci(k-1) + fibonacci(k-2)
}

// Fungsi utama (main)
func main() {
    var m int
    for m = 0; m < 8; m++ {
        var fib = fibonacci(m)
        fmt.Println(fib)
    }
}
```

---

### **Analisis Kompleksitas**  
- Algoritma Fibonacci rekursif ini memiliki kompleksitas waktu **O(2ⁿ)**, karena setiap pemanggilan rekursi bercabang menjadi dua pemanggilan lainnya.  
- **Kurang efisien** untuk nilai **n** yang besar karena banyak perhitungan berulang.  
- **Dapat dioptimalkan** menggunakan **Memoization** atau **Dynamic Programming** untuk meningkatkan performa ke **O(n)**.  

---

### **Kesimpulan**  
📌 **Divide and Conquer** adalah teknik yang kuat untuk menyelesaikan berbagai masalah komputasi.  
📌 **Kelebihan utama**: Dapat dieksekusi secara paralel dan efisien dalam penggunaan memori.  
📌 **Kelemahan utama**: Bisa menjadi lambat jika tidak dioptimalkan dengan baik.  

📢 **Selanjutnya**, kita akan membahas **algoritma Backtracking**, yang sering digunakan untuk menyelesaikan masalah kombinatorial! 🚀
