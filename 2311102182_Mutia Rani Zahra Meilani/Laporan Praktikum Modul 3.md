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
  <strong>FUNGSI</strong>
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
  Mutia Rani Zahra Meilani
  <br>
  2311102182
  <br>
  S1IF1105
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

## <strong> DASAR TEORI </strong>

<span style="font-size:16px"><strong> ── PENGERTIAN FUNGSI</strong></span>
<br>
Fungsi dalam Go (Golang) adalah blok kode yang dapat dipanggil berulang kali untuk menjalankan tugas tertentu. Fungsi digunakan untuk memecah program menjadi bagian-bagian yang lebih kecil, modular, dan dapat digunakan kembali, sehingga memudahkan pemeliharaan dan pengembangan. Fungsi di Golang memiliki ciri khas sebagai tipe data first-class, artinya mereka dapat disimpan dalam variabel, dilewatkan sebagai argumen, dan dikembalikan dari fungsi lain.

<span style="font-size:16px"><strong> ── JENIS JENIS FUNGSI</strong></span>

**Fungsi Biasa**


Fungsi standar adalah fungsi yang menerima sejumlah parameter dan mengembalikan nilai. Ini adalah fungsi paling umum yang digunakan untuk melakukan tugas tertentu.
Contoh Penggunaan :
```go
func greet(name string) string {
    return "Hello, " + name
}
```

**Fungsi dengan Multiple Return**


Golang memungkinkan fungsi mengembalikan lebih dari satu nilai. Ini berguna untuk mengembalikan hasil dan error secara bersamaan.
Contoh Penggunaan :
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

**Fungsi tanpa Nama**


Anonymous function adalah fungsi yang tidak memiliki nama. Fungsi ini sering digunakan sebagai callback atau untuk deklarasi inline.
Contoh Penggunaan :
```go
func main() {
    func(msg string) {
        fmt.Println(msg)
    }("Hello, Go!")
}
```

**Variadic Function**


Fungsi variadic memungkinkan untuk mengirim sejumlah parameter yang tidak terbatas dari tipe data yang sama. Tipe ini sangat berguna ketika jumlah argumennya tidak pasti.
Contoh Penggunaan :
```go
func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

```

**Method Function**


Method Function yaitu fungsi yang terkait dengan tipe struct. Ini memungkinkan untuk menambahkan perilaku pada tipe data custom (struct).
Contoh Penggunaan :
```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

```

**Defer Function**


dFungsi defer digunakan untuk menunda eksekusi suatu fungsi hingga fungsi induk selesai. Biasanya digunakan untuk memastikan pembersihan atau pelepasan sumber daya.
Contoh Penggunaan :
```go
func main() {
    defer fmt.Println("Selesai")
    fmt.Println("Mulai")
}

```

## <strong> GUIDED </strong>

### ── Guided 1

#### Source Code

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

#### Output

![image](https://github.com/user-attachments/assets/f34b8bb2-8119-4f91-a348-a27a94e7b064)

Deskripsi Program :
Program ini digunakan untuk menghitung permutasi dari dua bilangan bulat yang diberikan. Program ini akan meminta pengguna untuk memasukkan dua bilangan bulat, kemudian akan menghitung permutasi dari bilangan yang lebih besar terhadap bilangan yang lebih kecil. Program ini menggunakan fungsi permutasi yang akan menghitung faktorial, sehingga hasil akhirnya berupa jumlah permutasi dari kedua bilangan yang sudah diinputkan oleh pengguna.

### ── Guided 2

#### Source Code

```go
package main

import "fmt"

func luas(a int)int{
	return a*a
}

func keliling(a int)int{
	return 4*a
}

func main(){
	var input int
	fmt.Scan(&input)
	fmt.Println("Luas Persegi : ", luas(input))
	fmt.Println("Keliling Persegi : ", keliling(input))
}
```

#### Output

![image](https://github.com/user-attachments/assets/a865daf7-1cd6-4340-8c43-93d38f0c3746)

Deskripsi Program :
Program ini digunakan untuk menghitung luas dan keliling dari sebuah persegi. Program ini akan meminta pengguna untuk memasukan panjang sisi dari sebuah persegi. Program ini memiliki 2 fungsi yang berguna untuk menghitung luas dan keliling persegi dari panjang sisi yang dimasukan oleh pengguna.

## <strong>  UNGUIDED </strong>

### ── Unguided 1

#### Study Case

Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)

**Masukan** terdiri dari empat buah bilangan asli `a`, `b`, `c`, dan `d` yang dipisahkan oleh spasi, dengan syarat `a ≥ c` dan `b ≥ d`.

**Keluaran** terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi `a` terhadap `c`, sedangkan baris kedua adalah hasil permutasi dan kombinasi `b` terhadap `d`.

**Catatan**: permutasi (P) dan kombinasi (C) dari `n` terhadap `r` (`n ≥ r`) dapat dihitung dengan menggunakan persamaan berikut!

$P(n, r) = \frac{n!}{(n-r)!}$ ,  Sedangkan  $C(n, r) = \frac{n!}{r!(n-r)!}$

#### Source Code

```go
package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Print("Input Nilai : ")
	fmt.Scan( &a, &b, &c, &d)
	if a >= c && b >= d {
		fmt.Println("Permutasi A dan C : ", permutasi(a, c))
		fmt.Println("Kombinasi A dan C : ", kombinasi(a, c))
		fmt.Println("Permutasi B dan D : ", permutasi(b, d))
		fmt.Println("Kombinasi B dan D : ", kombinasi(b, d))
	} else {
		fmt.Println("Angka yang dimasukan tidak valid")
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

func kombinasi(n, r int) int {
    return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
```

#### Output

![image](https://github.com/user-attachments/assets/c0a733e3-5d5e-47b5-a581-1dd4b8430b06)

Deskripsi Program :
Program ini bertujuan untuk menghitung permutasi dan kombinasi dari angka yang dimasukan oleh pengguna. Program akan menghitung permutasi dan kombinasi dari setiap 2 pasang angka. Dalam program ini nilai n yang dimasukan harus lebih besar atau sama dengan r, jika tidak maka program akan memberi tahu bahwa angka yang dimasukan tidak valid. Hasil akhir dari program ini adalah menampilkan hasil permutasi dan kombinasi dari setiap 2 pasang angka yang sudah diinputkan pengguna.

### ── Unguided 2

#### Study Case

Diberikan tiga buah fungsi matematika yaitu $ f (x) = x^2, g (x) = x - 2 $  dan $ h (x) = x + 1 $.
Fungsi komposisi $(fogoh)(x)$ artinya adalah $f(g(h(x)))$. Tuliskan $f(x), g(x)$ dan $h(x)$ dalam bentuk function.

**Masukan** terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi.

**Keluaran** terdiri dari tiga baris. Baris pertama adalah $(fogoh)(a)$, baris kedua $(gohof)(b)$, dan baris ketiga adalah $(hofog)(c)$!

#### Source Code

```go
package main

import (
	"fmt"
)

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int
	fmt.Print("Masukan 3 Nilai : ")
	fmt.Scan(&a, &b, &c)

	fogoh := f(g(h(a)))
	gohof := g(h(f(b)))
	hofog := h(f(g(c)))

	fmt.Printf("Nilai (fogoh)(%d) : %d\n", a, fogoh)
	fmt.Printf("Nilai (gohof)(%d) : %d\n", b, gohof)
	fmt.Printf("Nilai (hofog) %d) : %d\n", c, hofog)
}
```

#### Output

![image](https://github.com/user-attachments/assets/1761deb3-a233-4f70-9990-548aa97facee)

Deskripsi Program :
Program ini bertujuan untuk menghitung nilai dari fungsi f, g, dan h yang sudah ditentukan. Fungsi f mengkuadratkan bilangan, g mengurangi bilangan dengan 2, dan h menambahkan 1 pada bilangan. Nilai fogoh, gohof, dan hofog dihitung dengan mengaplikasikan fungsi-fungsi yang telah dibuat sesuai dengan rumus yang tertera. Program kemudian akan menampilkan hasil dari nilai fogoh, gohof, dan hofog sesuai dengan nilai yang diinputkan oleh pengguna.

### ── Unguided 3

#### Study Case

[Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat $(cx, cy)$ dengan radius $r$. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang $(x, y)$ berdasarkan dua lingkaran tersebut.

**Masukan** terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsikan sumbu $x$ dan $y$ dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.

**Keluaran** berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

Fungsi untuk menghitung jarak titik (a, b) dan (c, d) dimana rumus jarak adalah:

$jarak = √((a-c)² + (b-d)²)$

dan juga fungsi untuk menentukan posisi sebuah titik sembarang berada di dalam suatu lingkaran atau tidak.

function jarak $(a, b, c, d : real)$ -> real
{Mengembalikan jarak antara titik $(a, b)$ dan titik $(c, d)$}

function didalam $(cx, cy, r, x, y : real)$ -> boolean
{Mengembalikan true apabila titik $(x, y)$ berada di dalam lingkaran yang memiliki titik pusat $(cx, cy)$ dan radius $r$}

#### Source Code

```go
package main

import (
	"fmt"
)

func jarak(a, b, c, d int) int {
	return (a-c)*(a-c) + (b-d)*(b-d)
}

func in_circle(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= r*r
}

func main() {
	var cx1, cy1, r1 int
	fmt.Print("Masukan Koordinat dan Radius (Lingkaran 1) : ")
	fmt.Scan(&cx1, &cy1, &r1)

	var cx2, cy2, r2 int
	fmt.Print("Masukan Koordinat dan Radius (Lingkaran 2) : ")
	fmt.Scan(&cx2, &cy2, &r2)

	var x, y int
	fmt.Print("Masukkan Koordinat Titik Sembarang : ")
	fmt.Scan(&x, &y)

	in_circle1 := in_circle(cx1, cy1, r1, x, y)
	in_circle2 := in_circle(cx2, cy2, r2, x, y)


	if in_circle1 && in_circle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in_circle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in_circle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```

#### Output

![image](https://github.com/user-attachments/assets/7c30e703-2433-45a6-8559-68bd64a607ea)

Deskripsi Program :
Program ini digunakan untuk menentukan apakah sebuah titik berada di dalam, luar atau pada lingkaran yang diberikan. Pengguna diminta untuk memasukkan koordinat pusat dan jari-jari dari dua lingkaran, serta koordinat dari sebuah titik. Program kemudian menghitung jarak antara titik dengan pusat masing-masing lingkaran dan membandingkannya dengan jari-jari lingkaran tersebut. Jika jaraknya lebih kecil atau sama dengan jari-jari, maka titik berada di dalam lingkaran. Jika tidak, maka titik berada di luar lingkaran. Hasil akhirnya adalah akan diberikan posisi dari titik yang telah diinputkan pengguna.

## <strong> REFERENSI </strong>

#### [1] Singh, Hiten Pratap. “Understanding Functions in Golang: A Comprehensive Guide.” LinkedIn, 2023. https://www.linkedin.com/pulse/understanding-functions-golang-comprehensive-guide-hiten-pratap-singh.

#### [2] Udacity Team. “Functions in Go (Golang).” Udacity Blog, 2023. https://www.udacity.com/blog/2023/05/functions-in-go-golang.html.

#### [3] MySkill.ID. “Function on Golang.” Medium, 2023. https://medium.com/@myskill.id/function-on-golang-27b6577e9cbe.
