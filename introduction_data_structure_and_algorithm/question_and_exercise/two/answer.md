## **1. Kompleksitas GetMin dan GetMax**
Fungsi **GetMin** dan **GetMax** menggunakan pendekatan **divide and conquer** dengan **rekursi**:

```go
func GetMin(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	arrBorder := len(arr) / 2
	leftMin := GetMin(arr[:arrBorder])
	rightMin := GetMin(arr[arrBorder:])

	if leftMin < rightMin {
		return leftMin
	}
	return rightMin
}
```
💡 **Analisis kompleksitasnya**:
- Pada setiap langkah, array dibagi menjadi **2 bagian**.
- Setiap bagian diproses secara **rekursif**, sampai mencapai ukuran 1.
- **Perbandingan** antara dua nilai dilakukan dalam setiap pemanggilan rekursif.

Jumlah pemanggilan rekursif mengikuti pohon biner:
- **Jumlah level rekursi** = log₂(n) (karena setiap langkah membagi array menjadi 2)
- **Jumlah total operasi per level** ≈ n (karena ada perbandingan dalam setiap langkah)

🔢 **Kompleksitas waktu (Time Complexity) = O(n)**  
Karena setiap elemen hanya diperiksa **sekali per level rekursi**, total operasinya tetap **O(n)**.

---

## **2. Kompleksitas TeeMinMax**
Fungsi **TeeMinMax** bekerja dengan mengalirkan array ke dua goroutine yang masing-masing memanggil `GetMin` dan `GetMax`.

```go
func TeeMinMax(done <-chan struct{}, in <-chan []int) (_, _ <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)

	go func() {
		defer close(out1)
		defer close(out2)

		for val := range orDone(done, in) {
			ch1, ch2 := out1, out2

			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case ch1 <- GetMin(val):
					ch1 = nil
				case ch2 <- GetMax(val):
					ch2 = nil
				}
			}
		}
	}()

	return out1, out2
}
```
💡 **Analisis kompleksitasnya**:
- `orDone(done, in)` membaca array yang sudah dihasilkan dari channel.
- Kemudian, setiap array diproses oleh `GetMin` dan `GetMax` secara paralel.
- Karena masing-masing fungsi `GetMin` dan `GetMax` memiliki kompleksitas **O(n)**, maka **tidak ada tambahan overhead besar di sini**.
  
🔢 **Kompleksitas waktu tetap O(n)**, karena pekerjaan utama dilakukan oleh `GetMin` dan `GetMax`.

---

## **3. Kompleksitas GenNum dan GenArr**
Fungsi **GenNum** dan **GenArr** digunakan untuk menghasilkan array secara dinamis.

- **GenNum** hanya menghasilkan **n = 10** angka acak dalam sebuah channel.
- **GenArr** membaca angka-angka tersebut dan membentuk sebuah array.

```go
func GenNum(done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			select {
			case <-done:
				return
			case out <- rand.Intn(10):
			}
		}
	}()
	return out
}
```
💡 **Analisis kompleksitasnya**:
- **GenNum** menjalankan loop **sebanyak 10 kali** → O(10) ≈ **O(1)** (karena jumlah iterasi tetap)
- **GenArr** hanya membaca **10 elemen** dari channel dan membentuk array → O(10) ≈ **O(1)**

🔢 **Kompleksitas waktu = O(1)** karena arraynya kecil dan tetap.

---

## **4. Kompleksitas Keseluruhan**
Dari semua analisis di atas:

| Fungsi | Kompleksitas |
|--------|-------------|
| `GenNum` | O(1) |
| `GenArr` | O(1) |
| `TeeMinMax` | O(n) |
| `GetMin` | O(n) |
| `GetMax` | O(n) |

Karena `TeeMinMax` menjalankan `GetMin` dan `GetMax` secara paralel, kompleksitas keseluruhan **tetap O(n)** karena dua O(n) yang berjalan bersamaan tetap O(n).

---

## **Kesimpulan**
🔹 **Kompleksitas keseluruhan algoritma = O(n)**  
🔹 Fungsi `GetMin` dan `GetMax` menggunakan **rekursi dengan divide and conquer**, tetapi tetap **O(n)**.  
🔹 `TeeMinMax` memproses hasil secara **paralel**, tetapi tidak menambah kompleksitas.  
🔹 `GenNum` dan `GenArr` tidak berkontribusi terhadap kompleksitas karena hanya bekerja dalam **O(1)**.  

Semoga penjelasan ini jelas! 🚀🔥
