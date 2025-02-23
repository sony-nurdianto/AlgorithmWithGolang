## **1️⃣ O(log n) - Logaritmik Complexity**
### **Apa itu O(log n)?**
O(log n) berarti **jumlah operasi bertambah secara lambat saat input `n` membesar**.  

➡ **Semakin besar `n`, pertambahan operasi berkurang drastis** karena setiap langkah **mengurangi jumlah pekerjaan yang tersisa**.

📌 **Intinya:**  
- Jika setiap langkah **membagi** masalah menjadi dua (atau bagian tetap lainnya), maka kompleksitasnya **O(log n)**.
- Biasanya muncul dalam **Binary Search** atau **Divide & Conquer Algorithms**.

---

### **🔹 Contoh: Binary Search**
Misalkan kita punya **array terurut** dan ingin mencari angka **X**.  
Alih-alih mengecek satu per satu (O(n)), kita bisa **menghapus setengah array setiap langkah**.

#### **Contoh kode:**
```go
func binarySearch(arr []int, target int) bool {
    left, right := 0, len(arr)-1
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] == target {
            return true // Ketemu!
        } else if arr[mid] < target {
            left = mid + 1 // Cari di bagian kanan
        } else {
            right = mid - 1 // Cari di bagian kiri
        }
    }
    return false // Tidak ketemu
}
```

---

### **🔹 Cara kerja O(log n)**
Jika ada **n = 16** elemen:  
1. **Langkah 1** → Periksa elemen tengah → **16 jadi 8**  
2. **Langkah 2** → Periksa setengahnya lagi → **8 jadi 4**  
3. **Langkah 3** → Periksa setengahnya lagi → **4 jadi 2**  
4. **Langkah 4** → Periksa setengahnya lagi → **2 jadi 1**  
5. **Langkah 5** → Selesai  

➡ **Jumlah langkah = 5** (log₂(16) = 4).  
➡ **Secara umum, log₂(n) = jumlah iterasi yang dibutuhkan untuk mengurangi `n` ke 1**.  

**Perbandingan O(n) vs O(log n)**
| `n`    | **O(n) (Linear Search)** | **O(log n) (Binary Search)** |
|--------|----------------|------------------|
| 10     | 10 langkah    | 4 langkah        |
| 100    | 100 langkah   | 7 langkah        |
| 1000   | 1000 langkah  | 10 langkah       |
| 1M     | 1M langkah    | 20 langkah       |

📌 **Kesimpulan:**  
- **O(log n) jauh lebih cepat daripada O(n)** untuk input besar.  
- **Penerapan umum:** Binary Search, Heap, Tree Traversal, dan algoritma Divide & Conquer.  

---

## **2️⃣ O(n²) - Quadratic Complexity**
### **Apa itu O(n²)?**
O(n²) berarti **jumlah operasi tumbuh sangat cepat seiring bertambahnya `n`**.  

➡ **Jika setiap elemen berinteraksi dengan semua elemen lainnya, kita mendapatkan O(n²)**.  
➡ Biasanya terjadi dalam **nested loop (loop di dalam loop)**.

---

### **🔹 Contoh: Bubble Sort**
Algoritma sorting seperti **Bubble Sort** melakukan **perbandingan setiap elemen dengan elemen lainnya**.

#### **Contoh kode:**
```go
func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {        // Loop pertama (n)
        for j := 0; j < n-i-1; j++ {  // Loop kedua (n)
            if arr[j] > arr[j+1] {     // Swap jika perlu
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}
```

---

### **🔹 Cara kerja O(n²)**
Misalkan kita punya **n = 5** elemen:  
1. **Loop pertama (`i=0`)** → 4 perbandingan  
2. **Loop kedua (`i=1`)** → 3 perbandingan  
3. **Loop ketiga (`i=2`)** → 2 perbandingan  
4. **Loop keempat (`i=3`)** → 1 perbandingan  

Total operasi ≈ **(n-1) + (n-2) + ... + 2 + 1 = (n² - n)/2**  
➡ **Dalam notasi Big O, kita abaikan konstanta → O(n²).**  

**Perbandingan O(n) vs O(n²)**
| `n`    | **O(n) (Linear Search)** | **O(n²) (Bubble Sort)** |
|--------|----------------|------------------|
| 10     | 10 langkah    | 100 langkah     |
| 100    | 100 langkah   | 10.000 langkah  |
| 1000   | 1000 langkah  | 1.000.000 langkah |

📌 **Kesimpulan:**  
- **O(n²) sangat lambat untuk input besar**.  
- **Hindari O(n²) jika memungkinkan, ganti dengan algoritma lebih efisien (O(n log n) atau O(n))**.  
- **Penerapan umum:** Sorting sederhana, perbandingan pasangan dalam dataset, brute-force algorithm.  

---

## **Kesimpulan Besar**
| Kompleksitas | Contoh Algoritma | Pertumbuhan Operasi |
|-------------|----------------|----------------|
| **O(1)** (konstan) | Hash Table lookup | Tidak berubah |
| **O(log n)** (logaritmik) | Binary Search | Lambat bertambah |
| **O(n)** (linear) | Looping array | Bertambah seiring `n` |
| **O(n log n)** | Quick Sort, Merge Sort | Lebih cepat dari O(n²) |
| **O(n²)** (kuadratik) | Bubble Sort, Nested Loops | Cepat memburuk |
| **O(2^n)** (eksponensial) | Rekursi Fibonacci | **Sangat lambat!** |

⚡ **Intinya:**
- **O(log n) sangat efisien** (contoh: Binary Search).  
- **O(n²) sangat lambat**, sebaiknya **dihindari** untuk input besar.  

Semoga makin paham! Ada bagian yang masih membingungkan? 😃
