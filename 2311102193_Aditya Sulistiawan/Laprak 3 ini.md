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



#### Cara Kerja


1. **Input Pengguna**:
   - Program meminta pengguna untuk memasukkan panjang sisi persegi melalui konsol.

2. **Proses Perhitungan**:
   - Setelah menerima input, program memanggil fungsi `hitungLuas` untuk menghitung luas berdasarkan sisi yang diberikan.
   - Selanjutnya, program memanggil fungsi `hitungKeliling` untuk menghitung keliling dari persegi tersebut.

3. **Output Hasil**:
   - Program menampilkan hasil perhitungan luas dan keliling ke layar dengan format yang mudah dibaca, menggunakan dua angka di belakang koma.
     
## <strong> Unguided </strong>

### <h2> UNGUIDED 1 </h2>

#### Study Case
Minggu Ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kullah matematika diskett untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, Iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kallan membantu Jonas? (tidak tentunya ya :p)

#### Source Code

```go
// Aditya Sulistiawan
// 2311102193
// 11-05

package main

import "fmt"

func main() {
	var a, b, c, d int
5
	// fungsi untuk menginput bilangan a, b, c, dan d
	fmt.Print("Bilangan a: ")
	fmt.Scan(&a)

	fmt.Print("Bilangan b: ")
	fmt.Scan(&b)

	fmt.Print("Bilangan c: ")
	fmt.Scan(&c)

	fmt.Print("Bilangan d: ")
	fmt.Scan(&d)

	// Untuk menjelaskan syarat bahwa a>= c dan b>=d
	if a >= c && b >= d {
		// Untuk menampilkan hasil permutasi dan kombinasi dari bilangan a dan c
		fmt.Printf("\nPermutasi (a, c): %d\n", permutasi(a, c))
		fmt.Printf("Kombinasi (a, c): %d\n", kombinasi(a, c))

		// Untuk menampilkan hasil permutasi dan kombinasi dari bilangan b dan d
		fmt.Printf("\nPermutasi (b, d): %d\n", permutasi(b, d))
		fmt.Printf("Kombinasi (b, d): %d\n", kombinasi(b, d))
	} else {
		// Untuk yang tidak memenuhi syarat
		fmt.Println("Input tidak sesuai dengan syarat yang ada")
	}
}

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
```

#### Source Code
![SC](https://github.com/user-attachments/assets/31374b8d-770c-4190-bfa5-231d1488d8ae)


#### Output
![Ungui1](https://github.com/user-attachments/assets/d202d887-911b-4643-adf0-b761ef7962db)



#### Deskripsi Program : 
Program ini ditulis dalam bahasa pemrograman Go (Golang) dan bertujuan untuk menghitung permutasi dan kombinasi dari dua pasang bilangan bulat, yaitu `a` dan `c`, serta `b` dan `d`. Program ini juga memeriksa syarat bahwa `a` harus lebih besar atau sama dengan `c`, dan `b` harus lebih besar atau sama dengan `d`. Jika syarat tersebut tidak terpenuhi, program akan memberikan pesan kesalahan

#### Algoritma dan Cara Kerja Program
1. **Inisialisasi**:
   - Import paket `fmt` untuk input/output.
   
2. **Deklarasi Variabel**:
   - Deklarasikan empat variabel bertipe `int`: `a`, `b`, `c`, dan `d`.

3. **Input Pengguna**:
   - Minta pengguna untuk memasukkan nilai untuk bilangan `a`, `b`, `c`, dan `d`.

4. **Pemeriksaan Syarat**:
   - Periksa apakah `a >= c` dan `b >= d`.
     - Jika syarat terpenuhi:
       - Hitung dan tampilkan permutasi dan kombinasi untuk pasangan `(a, c)` dan `(b, d)`.
     - Jika syarat tidak terpenuhi:
       - Tampilkan pesan kesalahan.

5. **Fungsi Pendukung**:
   - **Fungsi `faktorial(n int) int`**: Menghitung faktorial dari bilangan `n`.
   - **Fungsi `permutasi(n, r int) int`**: Menghitung permutasi menggunakan rumus $$ P(n, r) = \frac{n!}{(n-r)!} $$.
   - **Fungsi `kombinasi(n, r int) int`**: Menghitung kombinasi menggunakan rumus $$ C(n, r) = \frac{n!}{r!(n-r)!} $$.


#### Cara Kerja
1. **Input Pengguna**:
   - Program meminta pengguna untuk memasukkan empat bilangan bulat: `a`, `b`, `c`, dan `d`.

2. **Pemeriksaan Syarat**:
   - Setelah menerima input, program memeriksa apakah nilai-nilai yang dimasukkan memenuhi syarat (`a >= c` dan `b >= d`).
   - Jika syarat tidak terpenuhi, program akan mencetak pesan kesalahan.

3. **Perhitungan Permutasi dan Kombinasi**:
   - Jika syarat terpenuhi, program akan menghitung permutasi dan kombinasi untuk kedua pasangan bilangan:
     - Untuk pasangan `(a, c)`:
       - Memanggil fungsi `permutasi(a, c)` untuk menghitung permutasi.
       - Memanggil fungsi `kombinasi(a, c)` untuk menghitung kombinasi.
     - Untuk pasangan `(b, d)`:
       - Memanggil fungsi `permutasi(b, d)` untuk menghitung permutasi.
       - Memanggil fungsi `kombinasi(b, d)` untuk menghitung kombinasi.

4. **Output Hasil**:
   - Program menampilkan hasil perhitungan permutasi dan kombinasi ke layar dengan format yang jelas.

Dengan struktur yang terorganisir dengan baik, program ini menyediakan cara yang efektif untuk melakukan perhitungan kombinatorial dasar dengan memanfaatkan fungsi-fungsi pendukung untuk faktorial, permutasi, dan kombinasi.

### <h2> UNGUIDED 2 </h2>

#### Study Case


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
