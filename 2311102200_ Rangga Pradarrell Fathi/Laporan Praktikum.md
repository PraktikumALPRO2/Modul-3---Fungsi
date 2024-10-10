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
 Rangga Pradarrell Fathi
  <br>
  2311102200
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

<strong><h2>Definisi Fungsi</h2></strong>

Dalam pemrograman, function adalah sekumpulan instruksi yang dikelompokkan untuk melakukan tugas tertentu. Di Golang, fungsi sangat penting karena membantu dalam memecah kode yang kompleks menjadi bagian-bagian yang lebih sederhana dan mudah dikelola. Fungsi juga digunakan untuk menghindari pengulangan kode serta meningkatkan modularitas dan keterbacaan.

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

func main()  {
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
func permutasi (n,r int) int {
	return faktorial(n) / faktorial (n-r)
}
```

#### Source Code
![Guided1carbon](https://github.com/user-attachments/assets/556ae404-6957-4844-b171-f728259fcb9f)

#### Output
![Screenshot 2024-10-10 080802](https://github.com/user-attachments/assets/a35ffa17-5ab6-47be-8ef5-bbaefdd28054)


#### Deskripsi Program : 
Program di atas merupakan program yang menghitung permutasi dari dua bilangan yang dimasukkan oleh user.

#### Algoritma dan Cara Kerja Program
*Algoritma*

1. **Input dari Pengguna:**
   - Program pertama kali meminta dua bilangan `a` dan `b` dari pengguna.
   - Jika `a >= b`, program akan menghitung permutasi dengan `a` sebagai `n` dan `b` sebagai `r`. Jika tidak, posisi `a` dan `b` akan ditukar sehingga `b` dianggap sebagai `n` dan `a` sebagai `r`.

2. **Perhitungan Faktorial:**
   - Fungsi `faktorial(n int) int` menghitung faktorial dari `n`. 
   - Fungsi ini menggunakan sebuah loop dari 1 hingga `n`, di mana setiap iterasi mengalikan nilai sebelumnya dengan `i` untuk menghitung faktorial.

3. **Perhitungan Permutasi:**
   - Fungsi `permutasi(n, r int) int` menghitung permutasi menggunakan rumus : Permutasi = faktorial(n)/faktorial(n-r).
   - Fungsi ini menggunakan nilai faktorial dari `n` dan `n - r`, lalu membaginya untuk mendapatkan hasil permutasi.

4. **Output:**
   - Hasil dari perhitungan permutasi ditampilkan ke layar menggunakan `fmt.Println`.

*Cara Kerja*

1. Program dimulai dengan mendeklarasikan dua variabel `a` dan `b` sebagai bilangan bulat (`int`).
2. Program meminta input dua bilangan dari pengguna dengan menggunakan `fmt.Scan(&a, &b)`.
3. Program mengecek apakah nilai `a` lebih besar atau sama dengan `b`. Jika ya, permutasi dihitung dengan `a` sebagai `n` dan `b` sebagai `r`. Jika tidak, posisi `a` dan `b` ditukar.
4. Fungsi `permutasi` dipanggil dengan parameter `n` dan `r`, dan hasilnya dihitung dengan membagi faktorial `n` dengan faktorial `(n - r)`.
5. Setelah hasil permutasi diperoleh, hasil tersebut dicetak ke layar.

### <h2>GUIDED 2</h2>

#### Source Code

```go
package main

import "fmt"

func main() {
	var sisi float64
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&sisi)

	luas := sisi * sisi
	keliling := 4 * sisi

	fmt.Printf("Luas persegi adalah: %.2f\n", luas)
	fmt.Printf("Keliling persegi adalah: %.2f\n", keliling)
}
```

#### Source Code
![Guided2carbon](https://github.com/user-attachments/assets/38105dcd-cf9b-4ba8-96d7-21df79513e6c)


#### Output
![Screenshot 2024-10-10 084550](https://github.com/user-attachments/assets/d3962c46-49ae-43a0-81dc-ce4e6d69712b)


#### Deskripsi Program : 
Program di atas merupakan program yang menghitung permutasi dari dua bilangan yang dimasukkan oleh user.

#### Algoritma dan Cara Kerja Program
*Algoritma*
1. Program pertama-tama meminta input dari pengguna berupa panjang sisi persegi. Input ini disimpan dalam variabel `sisi` yang bertipe `float64` untuk mendukung angka desimal.
2. Setelah input diterima, program menghitung **luas** persegi dengan rumus:
   Luas = sisi x sisi
3. Kemudian, program menghitung **keliling** persegi dengan rumus:
  Keliling = 4 x sisi
4. Setelah perhitungan selesai, hasil **luas** dan **keliling** ditampilkan ke layar dengan format dua angka desimal menggunakan fungsi `fmt.Printf`.


*Cara Kerja*

1. Program dimulai dengan mendeklarasikan variabel `sisi` yang akan menyimpan input dari pengguna. 
2. Program menampilkan pesan "Masukkan panjang sisi persegi:" untuk meminta pengguna memasukkan nilai panjang sisi persegi. Fungsi `fmt.Scan(&sisi)` digunakan untuk membaca input dari pengguna.
3. Setelah nilai sisi dimasukkan, program menghitung luas persegi dengan mengalikan nilai `sisi` dengan dirinya sendiri (sisi * sisi). Hasil perhitungan ini disimpan dalam variabel `luas`.
4. Program juga menghitung keliling persegi dengan mengalikan nilai `sisi` dengan 4 (4 * sisi). Hasilnya disimpan dalam variabel `keliling`.
5. Setelah perhitungan selesai, program menampilkan hasil perhitungan luas dan keliling menggunakan fungsi `fmt.Printf` untuk memformat hasil ke dua angka di belakang koma.
6. Output menampilkan nilai luas dan keliling yang sudah dihitung berdasarkan input panjang sisi yang diberikan oleh pengguna.



