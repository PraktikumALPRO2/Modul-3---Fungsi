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

Dalam pemrograman, function adalah sekumpulan instruksi yang dikelompokkan untuk melakukan tugas tertentu. Di Golang, fungsi sangat penting karena membantu dalam memecah kode yang kompleks menjadi bagian-bagian yang lebih sederhana dan mudah dikelola. Fungsi juga digunakan untuk menghindari pengulangan kode serta meningkatkan modularitas dan keterbacaan.[1]

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

Setelah dideklarasikan, fungsi dapat dipanggil di dalam kode utama program. Pemanggilan fungsi dilakukan dengan menuliskan nama fungsi diikuti oleh tanda kurung `()`. Jika fungsi memiliki parameter, maka argumen yang sesuai harus diberikan di dalam tanda kurung.[2]

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

Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli 𝑎, 𝑏, 𝑐, dan 𝑑 yang dipisahkan oleh spasi, dengan syarat 𝑎 ≥ 𝑐 dan 𝑏 ≥ 𝑑.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi 𝒂 terhadap
𝑐, sedangkan baris kedua adalah hasil permutasi dan kombinasi 𝑏 terhadap 𝑑.

Catatan: permutasi (P) dan kombinasi (C) dari 𝑛 terhadap 𝑟 (𝑛 ≥ 𝑟) dapat dihitung dengan menggunakan persamaan berikut!
<br>
![soalgu1](https://github.com/user-attachments/assets/2e18508b-42ec-48f9-888e-0f4d4cc674c4)


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

Membuat program dengan bahasa go untuk mencari sebuah Luas dan Keliling Persegi.

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


## <strong> Unguided </strong>
### <h2> UNGUIDED 1 </h2>

Diberikan tiga buah fungsi matematika yaitu 𝑓 (𝑥) = 𝑥2 , 𝑔 (𝑥) = 𝑥 − 2 dan ℎ (𝑥) = 𝑥 + 1. Fungsi komposisi (𝑓𝑜𝑔𝑜ℎ)(𝑥) artinya adalah 𝑓(𝑔Gℎ(𝑥)H). Tuliskan 𝑓(𝑥), 𝑔(𝑥) dan ℎ(𝑥) dalam bentuk function.
Masukan terdiri dari sebuah bilangan bulat 𝑎, 𝑏 dan 𝑐 yang dipisahkan oleh spasi.

Keluaran terdiri dari tiga baris. Baris pertama adalah (𝑓𝑜𝑔𝑜ℎ)(𝑎), baris kedua (𝑔𝑜ℎ𝑜𝑓)(𝑏), dan baris ketiga adalah (ℎ𝑜𝑓𝑜𝑔)(𝑐)!

#### Source Code

```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05
package main

import "fmt"

// Fungsi f menerima sebuah integer x dan mengembalikan kuadrat dari x
func f(x int) int {
    fungsi_f := x * x
    return fungsi_f
}

// Fungsi g menerima sebuah integer x dan mengembalikan x dikurangi 2
func g(x int) int {
    fungsi_g := x - 2
    return fungsi_g
}

// Fungsi h menerima sebuah integer x dan mengembalikan x ditambah 1
func h(x int) int {
    fungsi_h := x + 1
    return fungsi_h
}

func main() {
    // Meminta input dari pengguna
    fmt.Print("Masukkan bilangan: ")
    var a, b, c int
    fmt.Scan(&a, &b, &c)

    // Menghitung dan mencetak hasil komposisi fungsi f(g(h(a)))
    fmt.Println("Hasil fungsi fogoh dari bilangan pertama adalah ", f(g(h(a))))

    // Menghitung dan mencetak hasil komposisi fungsi g(h(f(b)))
    fmt.Println("Hasil fungsi gohof dari bilangan kedua adalah ", g(h(f(b))))

    // Menghitung dan mencetak hasil komposisi fungsi h(f(g(c)))
    fmt.Println("Hasil fungsi hofog dari bilangan ketiga adalah ", h(f(g(c))))
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

Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius
𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦)
berdasarkan dua lingkaran tersebut.

Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

Fungsi untuk menghitung jarak titik (a, b) dan (c, d) dimana rumus jarak adalah:

![soalun3](https://github.com/user-attachments/assets/4e55ec18-1092-433e-91b6-83e4153fb9d8)

#### Source Code

```go
// Rangga Pradarrell Fathi
// 2311102200
// IF-11-05

package main

import (
    "fmt"
    "math"
)

// Fungsi jarak menghitung jarak antara dua titik (a,b) dan (c,d) dalam koordinat kartesian
// menggunakan rumus jarak Euclidean: √((a-c)² + (b-d)²)
func jarak(a, b, c, d float64) float64 {
    return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

// Fungsi didalam memeriksa apakah suatu titik (x,y) berada di dalam lingkaran
// dengan pusat (cx,cy) dan radius r
func didalam(cx, cy, r, x, y float64) bool {
    return jarak(cx, cy, x, y) <= r
}

func main() {
    // Deklarasi variabel untuk koordinat dan radius dua lingkaran, serta titik yang akan diperiksa
    var cx1, cy1, r1, cx2, cy2, r2, x, y float64

    // Meminta input dari pengguna untuk lingkaran 1
    fmt.Println("Masukkan koordinat titik pusat lingkaran 1 (cx, cy) dan radius (r):")
    fmt.Scanln(&cx1, &cy1, &r1)

    // Meminta input dari pengguna untuk lingkaran 2
    fmt.Println("Masukkan koordinat titik pusat lingkaran 2 (cx, cy) dan radius (r):")
    fmt.Scanln(&cx2, &cy2, &r2)

    // Meminta input koordinat titik yang akan diperiksa
    fmt.Println("Masukkan koordinat titik (x, y):")
    fmt.Scanln(&x, &y)

    // Memeriksa posisi titik terhadap kedua lingkaran dan mencetak hasilnya
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

## Daftar Pustaka
[1] A. A. A. Donovan and B. W. Kernighan, *The Go Programming Language*. Boston, MA: Addison-Wesley, 2015.

[2] W. Kennedy, B. Ketelsen, and E. St. Martin, *Go in Action*. Greenwich, CT: Manning Publications, 2016.


