**Pseudo Code**  

*Pseudo code* adalah desain tingkat tinggi dari sebuah program atau algoritma. Dua konstruksi utama yang digunakan dalam *pseudo code* adalah **sequence** (urutan) dan **selection** (pemilihan).  

Dibandingkan dengan *flowchart*, *pseudo code* lebih mudah untuk dimodifikasi dan diperbarui. Kesalahan dalam desain dapat dideteksi lebih awal dengan *pseudo code*, sehingga menghemat biaya dalam memperbaiki cacat (*defects*) di tahap selanjutnya.  

Sebagai contoh, kita ingin mencari nilai maksimum dalam sebuah array dengan panjang *n*. *Pseudo code*-nya dapat ditulis sebagai berikut:  

```plaintext
maximum(arr) {
    n <- len(arr)
    max <- arr[0]
    for k <- 0 to n do {
        if arr[k] > max {
            max <- arr[k]
        }
    }
    return max
}
```  

Sekarang setelah kita memahami berbagai cara merepresentasikan algoritma, pada bagian selanjutnya kita akan melihat bagaimana cara memantau kompleksitas dan performanya.
