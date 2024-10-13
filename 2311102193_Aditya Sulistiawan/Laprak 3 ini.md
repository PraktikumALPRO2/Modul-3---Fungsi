<p align="center">
  <strong>LAPORAN PRAKTIKUM</strong>
  <br>
  <strong>ALGORITMA DAN PEMROGRAMAN 2</strong>
  <br>
</p>

<br>

<p align="center">
  <strong>MODUL III</strong>
  <br>
  <strong>Fungsi</strong>
  <br>
</p>

<br>

<p align="center">
  <img src="https://github.com/user-attachments/assets/296eb3c2-bd6b-401f-8256-3661150ec39e" alt="Logo" width="200"/>
</p>

<br>

<p align="center">
  <strong>Disusun Oleh :</strong>
  <br>
 Aditya Sulistiaawan
  <br>
  2311102193
  <br>
  IF-11-05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu :</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong>
  <br>
  <strong>FAKULTAS INFORMATIKA</strong>
  <br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong>
  <br>
  <strong>2024</strong>
</p>


## <strong> Dasar Teori </strong>

Fungsi adalah salah satu konsep fundamental dalam pemrograman yang memungkinkan programmer untuk mengorganisir dan mengelompokkan serangkaian instruksi menjadi unit yang dapat digunakan kembali. Berikut adalah beberapa aspek penting dari fungsi:

### 1. Definisi:
   Fungsi adalah blok kode yang dirancang untuk melakukan tugas tertentu. Fungsi hanya akan dijalankan ketika dipanggil dan dapat menerima data (parameter), memprosesnya, dan mengembalikan hasil.

### 2. Tujuan Penggunaan Fungsi:
   - Meningkatkan modularitas kode
   - Mengurangi redundansi dengan memungkinkan penggunaan kembali kode
   - Meningkatkan keterbacaan dengan memecah program menjadi bagian-bagian yang lebih kecil dan terkelola
   - Mempermudah pemeliharaan dan debugging kode

### 3. Struktur Umum Fungsi:
   - Nama fungsi: Identifier unik untuk fungsi
   - Parameter: Data yang diterima fungsi (opsional)
   - Body: Serangkaian instruksi yang mendefinisikan operasi fungsi
   - Return value: Nilai yang dikembalikan oleh fungsi (opsional)

### 4. Jenis-jenis Fungsi:
   - Built-in functions: Fungsi yang sudah tersedia dalam bahasa pemrograman
   - User-defined functions: Fungsi yang dibuat oleh programmer
   - Anonymous functions / Lambda functions: Fungsi tanpa nama yang biasanya digunakan untuk operasi singkat

### 5. Konsep Penting dalam Fungsi:
   - Pemanggilan fungsi (Function call)
   - Passing parameter
   - Scope variabel (local dan global)
   - Rekursi: Fungsi yang memanggil dirinya sendiri

### 6. Manfaat Penggunaan Fungsi:
   - Abstraksi: Menyembunyikan kompleksitas implementasi
   - Reusabilitas: Kode dapat digunakan kembali di berbagai bagian program
   - Testabilitas: Memudahkan pengujian unit kode
   - Kolaborasi: Memungkinkan pembagian tugas dalam tim pengembangan

### 7. Best Practices dalam Penulisan Fungsi:
   - Fungsi sebaiknya melakukan satu tugas spesifik (Single Responsibility Principle)
   - Nama fungsi harus deskriptif dan mencerminkan tujuannya
   - Batasi jumlah parameter untuk meningkatkan keterbacaan
   - Dokumentasikan fungsi dengan komentar atau docstrings

Fungsi merupakan komponen penting dalam paradigma pemrograman terstruktur dan berorientasi objek. Pemahaman yang baik tentang fungsi sangat penting untuk mengembangkan kode yang efisien, mudah dipelihara, dan dapat digunakan kembali.


<strong> Deklarasi Fungsi
</strong>

Deklarasi fungsi di Golang menggunakan kata kunci func, diikuti oleh nama fungsi, daftar parameter (opsional), tipe data hasil pengembalian (opsional), dan badan fungsi yang ditulis dalam kurung kurawal {}.

*Contoh*
``` go
func functionName(parameter1 type, parameter2 type, ...) returnType {
    // Badan fungsi
}
```
- functionName: Nama fungsi yang akan digunakan untuk memanggilnya.
- parameter: Nilai yang diterima fungsi sebagai input (opsional).
- returnType: Tipe data hasil pengembalian (opsional).
- Badan fungsi: Kode yang dijalankan saat fungsi dipanggil.

<strong> Cara Memanggil Function</strong>

Setelah dideklarasikan, fungsi dapat dipanggil di dalam kode utama program. Pemanggilan fungsi dilakukan dengan menuliskan nama fungsi diikuti oleh tanda kurung `()`. Jika fungsi memiliki parameter, maka argumen yang sesuai harus diberikan di dalam tanda kurung.

*Contoh*
```go
package main

import "fmt"

// Fungsi dengan parameter
func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    // Pemanggilan fungsi dengan argumen
    greet("John")  // Hello, John
    greet("Alice") // Hello, Alice
}
```


## <strong> Guided </strong>

### <h2>GUIDED 1</h2>

#### Source Code

```go
package main 

import "fmt"
 
func main(){
	var a,b int 
	fmt.Scan(&a, &b) 
	if a >= b {
		fmt.Println(permutasi(a,b))
	}else{
		fmt.Println(permutasi(b,a))
	}
}
func faktorial(n int) int{ 
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ { 
		hasil = hasil * i
	}

	return hasil
}
func permutasi(n,r int) int {
	return faktorial(n) / faktorial(n-r)
}
```

#### Screenshot Source Code
![SC](https://github.com/user-attachments/assets/e977240d-e959-481f-b9a8-77818f00196c)


#### Output
![Gui1](https://github.com/user-attachments/assets/51f1bee9-d0b0-402a-9aca-fbe34751e7be)


#### Deskripsi Program : 
Program ini ditulis dalam bahasa Go (Golang) dan berfungsi untuk menghitung permutasi dari dua angka yang diinputkan oleh pengguna.

#### Algoritma dan Cara Kerja Program

1. Mulai program
2. Deklarasikan variabel a dan b sebagai integer
3. Minta input dari pengguna untuk nilai a dan b
4. Jika a lebih besar atau sama dengan b, maka:
   - Panggil fungsi permutasi(a, b)
   - Tampilkan hasil permutasi
5. Jika tidak (a lebih kecil dari b), maka:
   - Panggil fungsi permutasi(b, a)
   - Tampilkan hasil permutasi
6. Definisikan fungsi faktorial(n):
   a. Inisialisasi variabel hasil = 1
   b. Untuk i dari 1 sampai n, lakukan:
      - hasil = hasil * i
   c. Kembalikan nilai hasil
7. Definisikan fungsi permutasi(n, r):
   a. Hitung faktorial(n)
   b. Hitung faktorial(n-r)
   c. Hitung hasil = faktorial(n) / faktorial(n-r)
   d. Kembalikan nilai hasil
8. Selesai


#### Cara Kerja*

1. Program meminta pengguna memasukkan dua angka.
2. Program menentukan angka mana yang lebih besar.
3. Fungsi `permutasi()` dipanggil dengan angka yang lebih besar sebagai `n` dan yang lebih kecil sebagai `r`.
4. Fungsi `permutasi()` memanggil fungsi `faktorial()` untuk menghitung faktorial yang diperlukan.
5. Hasil perhitungan permutasi dicetak ke layar.

Program ini mendemonstrasikan penggunaan fungsi, pengambilan input dari pengguna, pengambilan keputusan dengan `if-else`, dan perhitungan matematika dasar dalam pemrograman Go.

### <h2>GUIDED 2</h2>

#### Source Code

```go
package main

import "fmt"

func hitungLuas(sisi float64) float64 {
	return sisi * sisi
}

func hitungKeliling(sisi float64) float64 {
	return 4 * sisi
}

func main() {
	var sisi float64

	fmt.Print("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi)

	luas := hitungLuas(sisi)
	keliling := hitungKeliling(sisi)

	fmt.Printf("Luas persegi adalah: %.2f\n", luas)
	fmt.Printf("Keliling persegi adalah: %.2f\n", keliling)
}
```

#### Source Code
![SC](https://github.com/user-attachments/assets/65d918b5-dc30-48a3-b647-9983985d9245)


#### Output
![Gui2](https://github.com/user-attachments/assets/d61e062a-ebaa-435e-b3a8-25ec5a2f3b1e)



#### Deskripsi Program : 
Program ini ditulis dalam bahasa pemrograman Go (Golang) dan berfungsi untuk menghitung luas dan keliling dari sebuah persegi berdasarkan panjang sisi yang dimasukkan oleh pengguna. Program ini menggunakan dua fungsi terpisah: satu untuk menghitung luas dan satu lagi untuk menghitung keliling. Setelah pengguna memasukkan nilai sisi, program akan menampilkan hasil perhitungan tersebut dengan format yang rapi.

#### Algoritma 
1. **Inisialisasi**:
   - Import paket `fmt` untuk input/output.
   
2. **Definisi Fungsi**:
   - `hitungLuas(sisi float64)`: Menghitung luas persegi dengan rumus $$ \text{luas} = \text{sisi} \times \text{sisi} $$.
   - `hitungKeliling(sisi float64)`: Menghitung keliling persegi dengan rumus $$ \text{keliling} = 4 \times \text{sisi} $$.

3. **Fungsi `main`**:
   - Deklarasi variabel `sisi` bertipe `float64`.
   - Mencetak pesan untuk meminta pengguna memasukkan panjang sisi persegi.
   - Menggunakan `fmt.Scanln(&sisi)` untuk membaca input dari pengguna.
   - Memanggil fungsi `hitungLuas` dan menyimpan hasilnya dalam variabel `luas`.
   - Memanggil fungsi `hitungKeliling` dan menyimpan hasilnya dalam variabel `keliling`.
   - Mencetak hasil luas dan keliling dengan format dua desimal.



#### Cara Kerja*


1. **Input Pengguna**:
   - Program meminta pengguna untuk memasukkan panjang sisi persegi melalui konsol.

2. **Proses Perhitungan**:
   - Setelah menerima input, program memanggil fungsi `hitungLuas` untuk menghitung luas berdasarkan sisi yang diberikan.
   - Selanjutnya, program memanggil fungsi `hitungKeliling` untuk menghitung keliling dari persegi tersebut.

3. **Output Hasil**:
   - Program menampilkan hasil perhitungan luas dan keliling ke layar dengan format yang mudah dibaca, menggunakan dua angka di belakang koma.
     
## <strong> Unguided </strong>
### <h2> UNGUIDED 1 </h2>

#### Source Code

```go
package main

import "fmt"

func f(x int) int {
	fungsi_f := x * x
	return fungsi_f
}

func g(x int) int {
	fungsi_g := x - 2
	return fungsi_g
}

func h(x int) int {
	fungsi_h := x + 1
	return fungsi_h
}

func main() {
	fmt.Print("Masukkan bilangan: ")

	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println("Hasil fungsi fogoh dari bilangan pertama adalah ", f(g(h(a))))

	fmt.Println("Hasil fungsi gohof dari bilangan pertama adalah ", g(h(f(b))))

	fmt.Println("Hasil fungsi gohof dari bilangan pertama adalah ", h(f(g(c))))
}
```

#### Source Code
![Unguided1 carbon](https://github.com/user-attachments/assets/c69d6d1e-7c19-4aeb-8c7f-232eb78fabae)

#### Output
![Screenshot 2024-10-10 102109](https://github.com/user-attachments/assets/d568a814-9508-4796-8e59-5ee16e5d26b0)


#### Deskripsi Program : 
Program ini adalah sebuah aplikasi Go yang mendemonstrasikan konsep komposisi fungsi matematika. Program meminta input tiga bilangan dari pengguna dan menerapkan berbagai kombinasi dari tiga fungsi matematika sederhana pada bilangan-bilangan tersebut.

#### Algoritma dan Cara Kerja Program
*Algoritma*
1. Program meminta pengguna memasukkan tiga bilangan (a, b, c)
2. Program menghitung tiga komposisi fungsi berbeda:
3. f(g(h(a))) - disebut sebagai "fogoh"
g(h(f(b))) - disebut sebagai "gohof"
h(f(g(c))) - juga disebut sebagai "hofg"


Program menampilkan hasil ketiga perhitungan tersebut

*Cara Kerja*
Input:

1. Program meminta pengguna memasukkan tiga bilangan
Bilangan-bilangan ini disimpan dalam variabel a, b, dan c


2. Proses untuk setiap komposisi fungsi:
- Program menjalankan fungsi-fungsi secara berurutan dari dalam ke luar
- Hasil dari satu fungsi menjadi input untuk fungsi berikutnya


3. Output:
Program menampilkan hasil ketiga komposisi fungsi


### <h2> UNGUIDED 2 </h2>

#### Source Code

```go
//Rangga Pradarrell Fathi
//2311102200
//IF-11
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Println("Masukkan koordinat titik pusat lingkaran 1 (cx, cy) dan radius (r):")
	fmt.Scanln(&cx1, &cy1, &r1)

	fmt.Println("Masukkan koordinat titik pusat lingkaran 2 (cx, cy) dan radius (r):")
	fmt.Scanln(&cx2, &cy2, &r2)

	fmt.Println("Masukkan koordinat titik (x, y):")
	fmt.Scanln(&x, &y)

	if didalam(cx1, cy1, r1, x, y) && didalam(cx2, cy2, r2, x, y) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if didalam(cx1, cy1, r1, x, y) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if didalam(cx2, cy2, r2, x, y) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran ")
	}
}

```

#### Source Code
![Unguided 2 carbon](https://github.com/user-attachments/assets/0d6da1f4-765b-4ed8-9b39-8ea7c54ed471)


#### Output
![Screenshot 2024-10-10 103140](https://github.com/user-attachments/assets/aec71ebf-875f-4da0-a222-d807c009230f)
![Screenshot 2024-10-10 103122](https://github.com/user-attachments/assets/278399f3-b9b3-4640-888e-78c7c4231f85)
![Screenshot 2024-10-10 103059](https://github.com/user-attachments/assets/1013be5a-cb61-41a8-a5e5-f4830014c5d1)
![Screenshot 2024-10-10 103348](https://github.com/user-attachments/assets/ac9380b6-b184-4d51-8ea6-0c00934e206f)


#### Deskripsi Program : 
Program ini adalah aplikasi Go yang menentukan posisi sebuah titik relatif terhadap dua buah lingkaran. Program meminta input koordinat pusat dan radius untuk dua lingkaran, serta koordinat sebuah titik, kemudian menentukan apakah titik tersebut berada di dalam salah satu, kedua, atau di luar kedua lingkaran.

#### Algoritma dan Cara Kerja Program
*Algoritma*
1. Input Data

Minta dan simpan koordinat pusat dan radius lingkaran 1
Minta dan simpan koordinat pusat dan radius lingkaran 2
Minta dan simpan koordinat titik yang akan dicek


2. Proses Pengecekan

Cek apakah titik berada di dalam lingkaran 1
Cek apakah titik berada di dalam lingkaran 2
Tentukan output berdasarkan hasil pengecekan


3. Output

Tampilkan hasil pengecekan sesuai kondisi yang terpenuhi

*Cara Kerja*
1. Perhitungan Jarak

- Program menggunakan teorema Pythagoras untuk menghitung jarak
- Implementasi dalam fungsi jarak() menggunakan math.Sqrt dan math.Pow


2. Penentuan Posisi Titik

- Titik dianggap di dalam lingkaran jika jaraknya ke pusat ≤ radius
- Fungsi didalam() mengembalikan true jika kondisi ini terpenuhi


3. Logika Pengecekan. Program memeriksa 4 kemungkinan:

- Titik di dalam kedua lingkaran
- Titik hanya di dalam lingkaran 1
- Titik hanya di dalam lingkaran 2
- Titik di luar kedua lingkaran
