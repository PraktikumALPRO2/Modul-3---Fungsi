<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br> 

<h2 align="center"><strong>MODUL III</strong></h2>
<h2 align="center"><strong> FUNGSI </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Fatik Nurimamah / 2311102190<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------


## Daftar Isi
1. [Dasar Teori](#dasar-teori)
2. [Guided](#guided)
3. [Unguided](#unguided)
4. [Kesimpulan](#kesimpulan)
5. [Daftar Pustaka](#daftar-pustaka)


## Dasar Teori


## Guided 

### 1.Buatlah sebuah program beserta fungsi yang digunakan untuk menghitung nilai faktorial dan permutasi

### Source Code :
```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println(permutasi(a, b))
	} else {
		fmt.Println(permutasi(b, a))
	}
}
func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
```

### Output:


### Full code Screenshot:
![Screenshot 2024-10-09 085622](https://github.com/user-attachments/assets/335d5a58-0c57-4007-babe-9345028dd8d6)

### Deskripsi Program : 

### Algoritma Program


### Cara Kerja Program:

### 2. Program untyk menghitung luas persegi dan keliling persegi

### Source Code :
```go
package main

import (
	"fmt"
)

func main() {
	// Meminta input panjang sisi persegi
	var panjangSisi float64
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&panjangSisi)

	// Menghitung luas dan keliling persegi
	LuasPersegi := panjangSisi * panjangSisi
	KelilingPersegi := 4 * panjangSisi

	// Menampilkan hasil
	fmt.Printf("Luas persegi: %g\n", LuasPersegi)
	fmt.Printf("Keliling persegi: %g\n", KelilingPersegi)
}

```

### Output:
![Screenshot 2024-10-09 090104](https://github.com/user-attachments/assets/fc3a1936-41c7-476b-8375-5f48aa345490)

### Full code Screenshot:
![Screenshot 2024-10-09 090155](https://github.com/user-attachments/assets/e53d16c7-1380-47ac-8145-62cdea270e15)

### Deskripsi Program : 

### Algoritma Program


### Cara Kerja Program


## Unguided 

### 1. Program untuk menghitung Faktorial, Permutasi dan Kombinasi
![Screenshot 2024-10-09 090531](https://github.com/user-attachments/assets/939bb468-ffa3-42f2-a6a3-1e7d44633ed1)

### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func HitungFaktorial(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi
func HitungPermutasi(n, r int) int {
	if n < r {
		return 0
	}
	return HitungFaktorial(n) / HitungFaktorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func HitungKombinasi(n, r int) int {
	if n < r {
		return 0
	}
	return HitungFaktorial(n) / (HitungFaktorial(r) * HitungFaktorial(n-r))
}


func main() {
	var a, b, c, d int

	// Meminta pengguna untuk memasukkan input
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	fmt.Print("Masukkan nilai d: ")
	fmt.Scan(&d)

	// Menghitung permutasi dan kombinasi
	hasilPermutasi1 := HitungPermutasi(a, c)
	hasilKombinasi1 := HitungKombinasi(a, c)
	hasilPermutasi2 := HitungPermutasi(b, d)
	hasilKombinasi2 := HitungKombinasi(b, d)

	// Menampilkan hasil
	fmt.Printf("\nPermutasi dan Kombinasi untuk %d dan %d: %d %d\n", a, c, hasilPermutasi1, hasilKombinasi1)
	fmt.Printf("Permutasi dan Kombinasi untuk %d dan %d: %d %d\n", b, d, hasilPermutasi2, hasilKombinasi2)
}

```
### Output:
![Screenshot 2024-10-09 090843](https://github.com/user-attachments/assets/5f071abd-c8a3-4f53-a3d5-1c33eb557bbf)

### Full code Screenshot:
![Screenshot 2024-10-09 090935](https://github.com/user-attachments/assets/ac54511f-5101-446f-8952-fe9dab2635d3)

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program


### 2. ![Screenshot 2024-10-09 091205](https://github.com/user-attachments/assets/0613877e-afcf-448f-a46a-3fff064d9891)

### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi f(x) = x^2
func fungsiKuadrat(nilai int) int {
	return nilai * nilai
}

// Fungsi g(x) = x - 2
func fungsiKurang(nilai int) int {
	return nilai - 2
}

// Fungsi h(x) = x + 1
func fungsiTambah(nilai int) int {
	return nilai + 1
}


// Fungsi komposisi fogoh(a)
func komposisiFogoh(nilai int) int {
	return fungsiKuadrat(fungsiKurang(fungsiTambah(nilai))) // f(g(h(x)))
}

// Fungsi komposisi gohof(b)
func komposisiGohof(nilai int) int {
	return fungsiKurang(fungsiTambah(fungsiKuadrat(nilai))) // g(h(f(x)))
}

// Fungsi komposisi hofog(c)
func komposisiHofog(nilai int) int {
	return fungsiTambah(fungsiKuadrat(fungsiKurang(nilai))) // h(f(g(x)))
}

func main() {
	// Meminta input dari pengguna
	var a, b, c int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	// Menghitung hasil komposisi
	hasilFogoh := komposisiFogoh(a)
	hasilGohof := komposisiGohof(b)
	hasilHofog := komposisiHofog(c)

	// Menampilkan hasil
	fmt.Printf("\nHasil komposisi (fogoh)(%d): %d\n", a, hasilFogoh)
	fmt.Printf("Hasil komposisi (gohof)(%d): %d\n", b, hasilGohof)
	fmt.Printf("Hasil komposisi (hofog)(%d): %d\n", c, hasilHofog)
}

```
### Output:
![Screenshot 2024-10-09 091414](https://github.com/user-attachments/assets/039e442d-eace-414a-b8db-3f24f986662a)

### Full code Screenshot:

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program



