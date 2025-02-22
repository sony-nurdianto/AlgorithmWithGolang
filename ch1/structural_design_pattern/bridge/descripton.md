### **Pola Desain Bridge (Bridge Pattern)**  

Pola desain **Bridge** bertujuan untuk **memisahkan (decouple) implementasi dari abstraksi**, sehingga implementasi dapat diubah tanpa memengaruhi abstraksinya.  

Dalam pola ini:  
✅ **Abstraksi** → Kelas dasar yang dapat diturunkan untuk berbagai implementasi.  
✅ **Implementasi** → Detail dari cara kerja abstraksi, yang bisa diubah dengan mudah.  
✅ **Jembatan (Bridge)** → Menghubungkan antara **abstraksi** dan **implementasi**, sehingga kelas-kelas konkret dapat bekerja secara independen.  
✅ **Komposisi lebih diutamakan daripada pewarisan** → Daripada menggunakan banyak subclass untuk berbagai variasi, lebih baik menggunakan **komposisi (has-a relationship)** agar lebih fleksibel.  

---  

### **📌 Kapan Menggunakan Bridge Pattern?**
🔹 **Ketika ada dua hierarki yang orthogonal (independen satu sama lain)**, misalnya:  
   - **Bentuk** (Shape) dan **Metode Gambar** (DrawMethod).  
   - **Perangkat Lunak** dan **Platform OS**.  
🔹 **Ketika implementasi perlu berubah di runtime**.  
🔹 **Ketika ingin menjaga kode tetap fleksibel** untuk perubahan di masa depan.  

---

### **📌 Komponen dalam Bridge Pattern**
1. **Abstraction** → Interface atau kelas abstrak yang digunakan oleh klien.  
2. **Refined Abstraction** → Variasi dari Abstraction yang lebih spesifik.  
3. **Implementer** → Interface yang berisi metode implementasi.  
4. **Concrete Implementer** → Implementasi konkret dari Implementer.  

---

### **📌 Contoh Implementasi Bridge Pattern**  
Misalkan kita punya **IDrawShape** sebagai interface dengan metode `drawShape()`.  
- **DrawShape** mengimplementasikan **IDrawShape**.  
- **IContour** adalah interface jembatan yang memiliki metode `drawContour()`.  
- **Contour** mengimplementasikan **IContour**.  
- **Ellipse** memiliki properti `a, b, r` serta referensi ke `drawShape` untuk menggambar bentuk.  
- **Ellipse** juga mengimplementasikan interface jembatan `IContour`.  

Saat metode `drawContour()` dipanggil, metode ini akan memanggil `drawShape()` pada instance `DrawShape`.  

🚀 **Kesimpulan:**  
Dengan **Bridge Pattern**, kita bisa mengubah **cara menggambar bentuk** tanpa harus mengubah struktur dari bentuk itu sendiri! 🎨✨
