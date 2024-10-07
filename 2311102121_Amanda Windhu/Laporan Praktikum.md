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

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas, salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) <br/> Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. <br/> Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. <br/> Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut!<br/> ![image](https://github.com/user-attachments/assets/c5b20217-7f5f-4d3a-bdcc-09dcfe46b89f) <br/> 

```go
package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Fungsi untuk menghitung permutasi
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	// Input empat bilangan: a, b, c, d
	var a, b, c, d int
	fmt.Print("Masukkan 4 bilangan: ")
	fmt.Scan(&a, &b, &c, &d)

	// Baris pertama: permutasi dan kombinasi a terhadap c
	p1 := permutation(a, c)
	c1 := combination(a, c)
	fmt.Printf("%d %d\n", p1, c1)

	// Baris kedua: permutasi dan kombinasi b terhadap d
	p2 := permutation(b, d)
	c2 := combination(b, d)
	fmt.Printf("%d %d\n", p2, c2)
}
```

## Output: ![image](https://github.com/user-attachments/assets/d49c1caa-89c0-4df2-959b-14d42d03ae4d) 
![image](https://github.com/user-attachments/assets/87df1016-0855-4158-a42e-e832e26f6430)

Kode di atas

### 2. Diberikan tiga buah fungsi matematika yaitu f(x) = x², g(x) = x - 2 dan h(x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function.<br/> Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi. Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!<br/> Contoh
![image](https://github.com/user-attachments/assets/ceaf99df-bba0-4106-85d4-84f072ad8c99)

```go
package main

import (
	"fmt"
)

// Definisi fungsi f(x) = x^2
func f(x int) int {
	return x * x
}

// Definisi fungsi g(x) = x - 2
func g(x int) int {
	return x - 2
}

// Definisi fungsi h(x) = x + 1
func h(x int) int {
	return x + 1
}

// Komposisi f(g(h(x)))
func fogoh(x int) int {
	return f(g(h(x)))
}

// Komposisi g(h(f(x)))
func gohof(x int) int {
	return g(h(f(x)))
}

// Komposisi h(f(g(x)))
func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	// Input tiga bilangan: a, b, c
	var a, b, c int
	fmt.Print("Masukkan 3 bilangan: ")
	fmt.Scan(&a, &b, &c)

	// Output komposisi fungsi sesuai instruksi
	fmt.Println(fogoh(a))
	fmt.Println(gohof(b))
	fmt.Println(hofog(c))
}
```

## Output: ![image](https://github.com/user-attachments/assets/8f15405f-390d-4cc1-af61-8e17e6675519) ![image](https://github.com/user-attachments/assets/6f26976c-9048-4cf1-a90c-65be0a68cc37)

Kode di atas dirancang untuk menghitung komposisi dari tiga fungsi matematika yang diberikan, yaitu Fungsi f(x) = x², g(x)=x−2, dan h(x)=x+1. Setiap fungsi
diimplementasikan sebagai fungsi terpisah dalam kode: `f(x)` mengkuadratkan input, `g(x)` mengurangi input dengan 2, dan `h(x)` menambah 1 pada input. Selain itu, terdapat tiga fungsi komposisi yang menghitung kombinasi dari ketiga fungsi tersebut. Fungsi `fogoh(x)` menghitung
f(g(h(x))), `gohof(x)` menghitung g(h(f(x))), dan `hofog(x)` menghitung h(f(g(x))). Program menerima tiga bilangan bulat a, b, dan sebagai masukan dari pengguna. Setelah itu, program menghitung dan menampilkan hasil dari masing-masing komposisi fungsi tersebut untuk nilai a, b, dan c sesuai urutan yang diminta.











