<h2 align="center">LAPORAN PAKTIKUM ALGORITMA DAN PEMROGRAMAN 2</h2>
<h2 align="center">MODUL 3</h2>
<h2 align="center">FUNGSI</h2>

<p align="center"><img src=https://github.com/user-attachments/assets/37e9c953-078b-4ef4-97e7-a5d25344e50b alt="Logo" width="300"/><p>

<p align="center">Disusun Oleh : </p>
<p align="center">Irshad Benaya Fardeca / 2311102199</p>
<p align="center">IF-11-05</p>
<br></br>
<p align="center">Dosen Pengampu : </p>
<p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>
<br></br>
<h3 align="center">PROGRAM STUDI S1 TEKNIK INFORMATIKA</h3>
<h3 align="center">FAKULTAS INFORMATIKA</h3>
<h3 align="center">TELKOM UNIVERSITY PURWOKERTO</h3>
<h3 align="center">2024</p>

<br></br>

#### I. DASAR TEORI
##### Fungsi
Dalam konteks pemrograman, fungsi adalah sekumpulan blok kode yang dibungkus dengan nama tertentu. Penerapan fungsi yang tepat akan menjadikan kode lebih modular dan juga dry (singkatan dari don't repeat yourself) yang artinya kita tidak perlu menuliskan banyak kode untuk kegunaan yang sama berulang kali. Cukup deklarasikan sekali saja blok kode sebagai suatu fungsi, lalu panggil sesuai kebutuhan.

Fungsi main() sendiri merupakan fungsi utama pada program Go, yang akan dieksekusi ketika program dijalankan. Selain fungsi main() , kita juga bisa membuat fungsi lainnya. Dan caranya cukup mudah, yaitu dengan menuliskan keyword func kemudian diikuti nama fungsi, lalu kurung () (yang bisa diisi parameter), dan diakhiri dengan kurung kurawal untuk membungkus blok kode. Parameter merupakan variabel yang menempel di fungsi yang nilainya ditentukan saat pemanggilan fungsi tersebut. Parameter sifatnya opsional, suatu fungsi bisa tidak memiliki parameter, atau bisa saja memeliki satu atau banyak parameter (tergantung kebutuhan).

Deklarasi fungsi: 
```go
func name (parameter-list) (result-list) {
	body
}
```

Selain parameter, fungsi bisa memiliki attribute return value atau nilai balik. Fungsi yang memiliki return value, saat deklarasinya harus ditentukan terlebih dahulu tipe data dari nilai baliknya.
Contoh fungsi yang memiliki attribute return value :
```go
func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}

```

##### Return Multiple Value
Sebuah fungsi dapat mengembalikan nilai lebih dari satu.
Contohnya : 
```go
func f() (int, int) {
  return 5, 6
}

func main() {
  x, y := f()
}
```
Multiple value sering diganakan untuk mengembalikan nilai error bersama dengan hasil atau boolean untuk menandakan succes

<br></br>


#### II. GUIDED
##### 1. Buatlah sebuah program beserta fungsi yang digunakan untuk menghitung nilai faktorial dan permutasi. Masukan terdiri dari dua buah bilangan positif a dan b. Keluaran berupa sebuah bilangan bulat yang menyatakan nilai a permutasi b apabila a≥ b atau b pemutasi a untuk kemungkinan yang lain.
##### Source code
```go
package main

import "fmt"

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

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println(permutasi(a, b))
	} else {
		fmt.Println(permutasi(b, a))
	}

}
```
##### Screenshoot Output
![Screenshot 2024-10-13 200124](https://github.com/user-attachments/assets/18792a59-5136-46f3-a251-f91183e70a12)

##### Deskripsi Program
Program ini menghitung dan mencetak permutasi dari dua bilangan bulat yang diinput oleh pengguna. Fungsi faktorial menghitung faktorial dari sebuah bilangan bulat dengan menggunakan perulangan for. Fungsi permutasi menghitung permutasi dari dua bilangan bulat menggunakan rumus permutasi: nPr = n! / (n-r)!. Dalam fungsi main, pengguna diminta untuk memasukkan dua bilangan bulat. Kemudian, fungsi permutasi dipanggil dengan argumen yang sesuai untuk menghitung permutasi. Hasil permutasi kemudian dicetak ke layar.


##### 2. Program menghitung luas dan keliling persegi
##### Source code
```go
package main

import "fmt"

func Luas(s int) int {
	var luas int
	luas = s * s
	return luas
}
func Keliling(s int) int {
	var keliling int
	keliling = s * 4
	return keliling
}
func main() {
	var s int
	fmt.Scan(&s)
	fmt.Println("Luas Persegi", Luas(s))
	fmt.Println("Keliling Persegi", Keliling(s))
}
```
##### Screenshoot Output
![Screenshot 2024-10-13 200511](https://github.com/user-attachments/assets/60c562bf-6294-443f-90de-489c3cba9d42)

##### Deskripsi Program
Program ini dirancang untuk menghitung luas dan keliling sebuah persegi. Program dimulai dengan mendefinisikan dua fungsi utama, yaitu Luas dan Keliling. Fungsi Luas menerima panjang sisi persegi sebagai input (dalam variabel s) dan mengembalikan nilai luas persegi yang dihitung dengan rumus sisi * sisi. Fungsi Keliling juga menerima panjang sisi sebagai input dan mengembalikan nilai keliling persegi yang dihitung dengan rumus 4 * sisi. Di dalam fungsi main, program meminta user untuk memasukkan panjang sisi persegi. Nilai yang dimasukkan user kemudian disimpan dalam variabel s. Setelah itu, program memanggil fungsi Luas dan Keliling dengan nilai s sebagai argumen, lalu mencetak hasil perhitungan luas dan keliling.

<br></br>


#### III. UNGUIDED

##### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p). Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
```go
package main

import "fmt"

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

func main() {
	var a, b, c, d int
	fmt.Scanln(&a, &b, &c, &d)
	if a >= c && b >= d {
		fmt.Println(permutasi(a, c))
		fmt.Println(kombinasi(a, c))
		fmt.Println(permutasi(b, d))
		fmt.Println(kombinasi(b, d))
	} else {
		fmt.Println("Tidak memenuhi kondisi")
	}
}
```
##### Screenshoot Output
![Screenshot 2024-10-13 200850](https://github.com/user-attachments/assets/1c0cc7ba-4578-4357-8175-bccfe2808d5a)

##### Deskripsi Program
Program ini dirancang untuk menghitung permutasi dan kombinasi dari dua pasang bilangan bulat yang diberikan oleh user. Program akan mengecek terlebih dahulu apakah nilai a lebih besar sama dengan c dan b lebih besar sama dengan d. Jika kondisi ini terpenuhi, maka perhitungan permutasi dan kombinasi akan dilakukan.

##### 2. Diberikan tiga buah fungsi matematika yaitu f(x) = x ^ 2 g(x) = x - 2 danh h(x) = x + 1 Fungsi komposisi (fogoh) (x) artinya adalah f(g(h(x))) Tuliskan f(x) g(x) dan h(x) dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasl. Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh) (a) baris kedua (gohof) (b) , dan baris ketiga adalah (hofog) (c)!
```go
package main

import "fmt"

func fx(bil int) int {
	return bil * bil
}

func gx(bil int) int {
	return bil - 2
}

func hx(bil int) int {
	return bil + 1
}

func fogog(bil int) int {
	return fx(gx(hx(bil)))
}
func gohof(bil int) int {
	return gx(hx(fx(bil)))
}
func hofog(bil int) int {
	return hx(fx(gx(bil)))
}
func main() {
	var bil1 int
	var bil2 int
	var bil3 int

	fmt.Print("Masukkan Bilangan : ")
	fmt.Scan(&bil1, &bil2, &bil3)

	fmt.Printf("f(g(h(x))) (%d) = %d\n", bil1, fogog(bil1))
	fmt.Printf("g(h(f(x))) (%d) = %d\n", bil2, gohof(bil2))
	fmt.Printf("h(f(g(x))) (%d) = %d\n", bil3, hofog(bil3))
}
```
##### Screenshoot Output
![Screenshot 2024-10-13 201249](https://github.com/user-attachments/assets/637eb9ba-dace-43eb-bd8f-8b2a34fd1158)

##### Deskripsi Program
Program ini adalah program untuk menghitung nilai dari tiga fungsi komposisi: f(g(h(x))), g(h(f(x))), dan h(f(g(x))). Fungsi-fungsi dasar yang digunakan adalah fx, gx, dan hx yang masing-masing melakukan operasi kuadrat, pengurangan 2, dan penambahan 1 terhadap bilangan yang diberikan.Program meminta user untuk memasukkan tiga bilangan, kemudian menghitung nilai dari ketiga fungsi komposisi menggunakan bilangan-bilangan tersebut. Hasil perhitungan kemudian ditampilkan.

##### 3. [Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return (math.Sqrt((a-c)*(a-c) + (b-d)*(b-d)))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, cx2, cy2, x, y, r1, r2 float64

	fmt.Println("Masukkan Koordinat Lingkaran dan Radius Lingkaran 1 :")
	fmt.Println("cx cy r")
	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Println("Masukkan Koordin at Lingkaran dan Radius Lingkaran 2 :")
	fmt.Println("cx cy r")
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Println("Masukkan Koordinat Titik Sembarang:")
	fmt.Println("x y")
	fmt.Scanln(&x, &y)

	if didalam(cx1, cy1, r1, x, y) && didalam(cx2, cy2, r2, x, y) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if didalam(cx1, cy1, r1, x, y) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if didalam(cx2, cy2, r2, x, y) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
##### Screenshoot Output
![Screenshot 2024-10-13 212054](https://github.com/user-attachments/assets/38134f2c-b8f7-4b1c-bc21-cb9e64d5e8da)

##### Deskripsi Program
Program ini akan meminta user untuk memasukkan koordinat titik pusat dan jari-jari dari dua lingkaran, serta koordinat titik sembarang yang ingin diuji posisinya. Kemudian, program akan menghitung jarak antara titik sembarang dengan titik pusat masing-masing lingkaran dan membandingkannya dengan jari-jari lingkaran tersebut. Berdasarkan hasil perbandingan ini, program akan menentukan apakah titik tersebut berada di dalam lingkaran pertama, lingkaran kedua, kedua-duanya, atau di luar kedua lingkaran.


### Referensi
[1] Donovan, A., Kernighan, B. (2015). The Go Programming Language. United Kingdom: Pearson Education.
[2] Agung, Noval (2024, 30 Agustus). Dasar Pemrograman Golang. Diakses pada 13 Oktober 2024, dari https://dasarpemrogramangolang.novalagung.com/
