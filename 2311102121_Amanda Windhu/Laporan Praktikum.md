# <h2 align="center">LAPORAN PRAKTIKUM</h2>
# <h2 align="center">ALGORITMA DAN PEMROGRAMAN 2</h2>
# <h2 align="center">MODUL 3</h2>
# <h2 align="center"> FUNGSI</h2>
<p align="center">
    <img src="https://github.com/user-attachments/assets/3ccfed0b-72d1-4349-ac08-c4dce379c827" alt="Gambar">
</p>
 <h3  align="center" >Disusun Oleh : </h3>
<p align="center">Amanda Windhu Gustyas - 2311102121</p>
<p align="center">IF-11-05</p>
 <h3 <p align="center" >Dosen Pengampu : </h3> </p>
 <p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>

# <h3 align="center"> PROGRAM STUDI S1 TEKNIK INFORMATIKA </h3>
# <h3 align="center"> FAKULTAS INFORMATIKA </h3>
# <h3 align="center"> TELKOM UNIVERSITY PURWOKERTO </h3>
# <h3 align="center"> 2024 </h3>

## I. DASAR TEORI

## II. GUIDED

### 1. Contoh Program dengan Function

```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b) // Membaca input dari pengguna untuk dua nilai integer a dan b
	if a >= b { 
		// Jika a lebih besar atau sama dengan b, panggil fungsi permutasi dengan parameter (a, b)
		fmt.Println(permutasi(a, b))
	} else { 
		// Jika b lebih besar dari a, panggil fungsi permutasi dengan parameter (b, a)
		fmt.Println(permutasi(b, a))
	}
}

func faktorial(n int) int {
	var hasil int = 1
	var i int
	// Loop untuk menghitung faktorial dari n
	for i = 1; i <= n; i++ {
		hasil = hasil * i // Mengalikan hasil dengan nilai i pada setiap iterasi
	}
	return hasil // Mengembalikan hasil faktorial
}

func permutasi(n, r int) int {
	// Menghitung permutasi nPr dengan membagi faktorial n dengan faktorial (n-r)
	return faktorial(n) / faktorial(n-r)
}
```
## Output: ![image](https://github.com/user-attachments/assets/8c8113f5-eece-425e-9c3f-e5f8e8dfd2d2)

Kode di atas adalah program Go yang berfungsi untuk menghitung nilai permutasi dari dua bilangan yang diinputkan oleh pengguna. Program meminta dua bilangan integer sebagai input dan menentukan mana yang lebih besar di antara keduanya. Jika bilangan pertama lebih besar atau sama dengan bilangan kedua, program akan menghitung permutasi dengan menggunakan nilai tersebut sebagai parameter untuk menghitung `nPr`, di mana n adalah bilangan yang lebih besar dan 𝑟
adalah bilangan yang lebih kecil. Fungsi `faktorial` digunakan untuk menghitung faktorial dari suatu bilangan, sementara fungsi `permutasi` menghitung nilai permutasi dengan membagi faktorial 𝑛 dengan faktorial (𝑛−𝑟). Hasil perhitungan kemudian ditampilkan kepada pengguna.

### 2. Menghitung Luas dan Keliling Persegi

```go
package main

import "fmt"

// Fungsi untuk menghitung luas persegi
func hitungLuas(sisi float64) float64 {
    return sisi * sisi
}

// Fungsi untuk menghitung keliling persegi
func hitungKeliling(sisi float64) float64 {
    return 4 * sisi
}

func main() {
    var sisi float64

    fmt.Print("Masukkan panjang sisi persegi: ")
    fmt.Scan(&sisi)

    luas := hitungLuas(sisi)
    keliling := hitungKeliling(sisi)

    fmt.Printf("Luas persegi: %.2f\n", luas)
    fmt.Printf("Keliling persegi: %.2f\n", keliling)
}
```

## Output: ![image](https://github.com/user-attachments/assets/7f538d50-83d0-4378-87a7-e7249d1a8de5)

Kode di atas program Go yang digunakan untuk menghitung luas dan keliling sebuah persegi. Program ini meminta pengguna untuk memasukkan panjang sisi persegi, kemudian menghitung luas dengan menggunakan rumus Luas = 𝑠𝑖𝑠𝑖 × 𝑠𝑖𝑠𝑖 dan menghitung keliling dengan rumus Keliling = 4 × 𝑠𝑖𝑠𝑖. Dua fungsi terpisah, `hitungLuas` dan `hitungKeliling`, digunakan untuk menghitung masing-masing nilai tersebut. Setelah menghitung, program menampilkan hasil luas dan keliling dengan format dua angka di belakang koma.

## III. UNGUIDED

### 1. 
