Pohon Adelson, Velski, dan Landis (AVL)  
Adelson, Velski, dan Landis adalah pionir dalam struktur data pohon AVL, sehingga nama pohon ini diambil dari mereka. Pohon AVL merupakan pohon pencarian biner yang menyesuaikan tinggi secara otomatis. Balance factor (faktor keseimbangan) dihitung dengan mencari selisih antara tinggi subpohon kiri dan subpohon kanan. Penyeimbangan dilakukan menggunakan teknik rotasi. Jika balance factor lebih dari satu, rotasi akan menggeser node ke arah yang berlawanan dari subpohon kiri atau kanan. Operasi pencarian, penambahan, dan penghapusan pada pohon ini berjalan dalam kompleksitas waktu O(log n).  

Bagian berikutnya membahas tentang definisi antarmuka *KeyValue* dan kelas *TreeNode*. Untuk bagian ini, silakan merujuk ke file *avl_tree.go*.
