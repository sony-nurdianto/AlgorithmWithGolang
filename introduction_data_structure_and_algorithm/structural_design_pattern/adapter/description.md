Pola adapter menyediakan sebuah pembungkus (*wrapper*) dengan antarmuka yang diperlukan oleh klien API untuk menghubungkan tipe yang tidak kompatibel dan bertindak sebagai penerjemah antara kedua tipe tersebut. Adapter menggunakan antarmuka dari suatu kelas agar dapat menjadi kelas dengan antarmuka lain yang kompatibel.  

Ketika ada perubahan kebutuhan, sering kali terdapat skenario di mana fungsionalitas kelas perlu diubah karena adanya antarmuka yang tidak kompatibel. **Prinsip Dependency Inversion** dapat diterapkan dengan menggunakan pola adapter, di mana sebuah kelas mendefinisikan antarmukanya sendiri untuk berkomunikasi dengan modul tingkat berikutnya yang diimplementasikan oleh kelas adapter. **Delegasi** adalah prinsip lain yang digunakan dalam pola adapter. Pola ini sering diterapkan dalam skenario di mana berbagai format perlu ditangani dalam transformasi dari sumber ke tujuan.  

Pola adapter terdiri dari beberapa komponen utama:  

1. **Target** → Antarmuka yang dipanggil oleh klien, yang kemudian meneruskan metode ke adapter dan adaptee.  
2. **Client** → Entitas yang ingin menggunakan antarmuka yang tidak kompatibel melalui adapter.  
3. **Adapter** → Komponen yang menerjemahkan antarmuka yang tidak kompatibel dari **Adaptee** ke dalam antarmuka yang diinginkan oleh klien.  
4. **Adaptee** → Kelas asli yang memiliki fungsionalitas tetapi tidak sesuai dengan antarmuka yang diharapkan oleh klien.  

### Contoh Implementasi  
Misalkan Anda memiliki antarmuka `IProcessor` dengan metode `process()`. Kelas `Adapter` mengimplementasikan metode `process()` dan memiliki instance `Adaptee` sebagai atributnya. Kelas `Adaptee` memiliki metode `convert()` serta variabel `adapterType`. Saat menggunakan klien API, pengembang memanggil metode `process()` pada antarmuka, yang kemudian akan memanggil `convert()` di dalam `Adaptee`.  

Kode berikut mengilustrasikan konsep ini:  
*(Kode belum disertakan dalam teks asli, tetapi dapat diimplementasikan berdasarkan deskripsi ini.)*
