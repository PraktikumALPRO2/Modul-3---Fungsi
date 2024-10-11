<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL III</strong></h2>
<h2 align="center"><strong> FUNGSI </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Andreas Besar Wibowo / 2311102198<br>
  S1 IF11-05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh,S.Kom.,M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------
## I. Dasar Teori
### Definisi Function
Fungsi dalam pemrograman adalah blok kode yang menjalankan tugas tertentu. Fungsi ini membantu memisahkan logika program menjadi bagian-bagian yang lebih kecil dan mudah dipahami. Dengan fungsi, kode bisa digunakan kembali (reusable) tanpa perlu menulis ulang, sehingga program menjadi lebih terstruktur dan efisien. Misalnya, ketika kita butuh mencetak nama dan pesan, kita cukup membuat satu fungsi untuk melakukan itu dan memanggilnya berkali-kali dengan parameter yang berbeda[1].

### Deklarasi Fungsi
Deklarasi fungsi dimulai dengan menentukan nama, parameter, dan tipe datanya. Parameter adalah variabel lokal yang menerima nilai dari pemanggil fungsi. Fungsi bisa mengembalikan nilai atau tidak, tergantung kebutuhan. Jika tidak mengembalikan apa-apa, kita cukup mendeklarasikannya tanpa bagian return.
Contoh kode 
```go
package main

import “fmt”
func main(){
  var nama string = “adipati”
  cetaknama(“Selamat siang”, nama)
}

func cetaknama(message string, data string){
  fmt.Println(message, data)
}
```
Jika kita memiliki fungsi cetakNama(pesan string, nama string), saat kita memanggil fungsi ini, kita akan memberikan nilai ke dua parameter tersebut, dan hasilnya akan dicetak.

### Fungsi Variadic
Fungsi variadik adalah fungsi yang bisa menerima jumlah argumen yang bervariasi, bukan jumlah yang tetap. Misalnya, fungsi fmt.Printf() di Go yang bisa menerima beberapa parameter. Kita dapat mendeklarasikan fungsi variadik dengan menambahkan tiga titik (...) di depan tipe parameter.

Contoh kode 
```go
package main

import “fmt”

func main(){
  var nilai = hitung(3, 3, 3, 5, 2, 4, 2, 3, 5, 1)
  var rata = fmt.Sprintf(“Rata-rata nilai adalah %.2f”,
  nilai)
  fmt.Println(rata)
}

func hitung (numbers...int)float64{
  var total int = 0
  for _, number := range numbers{
        total += number
  }
  var rata = float64(total)/float64(len(numbers))
  return rata
}
```
Jika kita punya fungsi hitung(numbers ...int), kita bisa memberikan banyak angka sebagai argumen saat memanggil fungsi ini, dan semua argumen tersebut akan diproses di dalam fungsi.

### Fungsi Closure
Closure adalah fungsi yang disimpan dalam variabel. Ini memungkinkan kita untuk mendefinisikan fungsi di dalam fungsi lain, dan fungsi tersebut tetap bisa mengakses variabel dari lingkup luarnya. 

Contoh kode 
```go
package main

import “fmt”
func main(){
  hitung := func(x, y int)int{
        return x + y
  }
  fmt.Println(hitung(5, 10))
}
```
Kita bisa membuat fungsi penjumlahan dalam variabel, lalu memanggil fungsi itu berkali-kali untuk menghitung nilai yang berbeda-beda. Closure berguna ketika kita ingin mempertahankan konteks variabel tertentu saat fungsi dijalankan[2].

## II. GUIDED
**1.Buatlah sebuah program beserta fungsi yang digunakan untuk menghitung nilai faktorial dan permutasi.**

*Masukan : Terdiri dari dua buah bilangan positif a dan b*

*Keluaran : Berupa sebuah bilangan bulat yang menyatakan nilai a permutasi b apabila >= b atau b permutasi a untuk kemungkinan yang lain*

#### Source Code
```go
package main 

package main 

import "fmt"
 
func main(){

	var a,b int // mendeklarasikan 2 variabel 
	fmt.Scan(&a, &b) // untuk input dari pengguna
	if a >= b { // jika a >= b, akan ada permutasi
		fmt.Println(permutasi(a,b))
	}else{
		fmt.Println(permutasi(b,a)) // jika b >= a, akan ada permutasi
	}
}
// fungsi untuk menghitung faktorial
func faktorial(n int) int{ 
	var hasil int = 1
	var i int
	
	//perulangan untuk faktorial
	for i = 1; i <= n; i++ { 
		hasil = hasil * i
	}

	return hasil
}
// fungsi untuk menghitung permutasi
func permutasi(n,r int) int {
	return faktorial(n) / faktorial(n-r)
}
```
#### Screenshoot Source Code
![Guided1 carbon](https://github.com/user-attachments/assets/0468f198-eaea-41c8-ac30-34756bf4b709)
#### Screenshoot Output
![Guided1 go](https://github.com/user-attachments/assets/790f9c70-9762-4a88-bb06-221e02b5e68e)
#### Deskripsi Program
Program diatas merupakan program untuk menghitung permutasi dari dua bilangan yang diinputkan oleh pengguna. Didalam perhitungan permutasi terdapat perhitungan faktorial juga.

Algoritma dari program tersebut sebagai berikut :
1. Pengguna menginputkan dua bilangan untuk mengisi variabel
2. Program mengecek manakah diantara 2 variabel itu yang terbesar.
3. Program akan menghitung permutasi dengan rumus yang sudah tertera
4. Pada output akan menampilkan hasil dari permutasi

Cara kerja dari program tersebut yaitu program menyimpan dua variabel untuk menyimpan dua bilangan yang diinputkan pengguna. Fungsi 'faktorial' digunakan untuk menghitung nilai faktorial dan fungsi 'permutasi' digunakan untuk menghitung permutasi. Diakhir program akan menampilkan hasil permutasi berdasarkan 2 angka mana yang dinggap lebih besar

**2.Buatlah sebuah program beserta fungsi yang digunakan untuk menghitung luas dan keliling persegi.**

*Masukan : Terdiri dari bilangan yang menjelaskan sisi persegi*

*Keluaran : Berupa sebuah bilangan bulat yang menyatakan luas dan keliling persegi sesuai sisi yang diinputkan user*

#### Source Code
```go
package main

import (
	"fmt"
)

// fungsi untuk menghitung luas persegi dengan rumus s x s
func hitungLuas(sisi float64) float64 {
	return sisi * sisi
}

// fungsi untuk menghitung keliling persegi dengan rumus 4 x s
func hitungKeliling(sisi float64) float64 {
	return 4 * sisi
}

func main() {
	var sisi float64

	// untuk inputan dari user
	fmt.Print("Sisi Persegi: ")
	fmt.Scan(&sisi)

	// menghitung luas persegi dengan fungsi hitungLuas
	luas := hitungLuas(sisi)
	// menghitung keliling persegi dengan fungsi hitungKeliling
	keliling := hitungKeliling(sisi)

	// menampilkan hasil dari perhitungan luas dan keliling peregi
	fmt.Printf("Luas persegi: %.2f\n", luas)
	fmt.Printf("Keliling persegi: %.2f\n", keliling)
}
```
#### Screenshoot Source Code
![Guided2 carbon](https://github.com/user-attachments/assets/b68d6fa7-ded1-45c8-9384-471ed6108755)
#### Screenshoot Output
![Guided2 go](https://github.com/user-attachments/assets/5538e29b-c851-4429-8867-ed01733b776f)
#### Deskripsi Program
Program diatas merupakan program untuk menghitung luas dan keliling persegi berdasarkan sisi yang diinputkan oleh user. Pada output program akan menampilkan luas dan keliling persegi sesuai dengan sisi yang diinginkan oleh user.

Algoritma dari program tersebut sebagai berikut :
1. Pengguna menginputkan suatu bilangan untuk mengisi sisi dari persegi
2. Program menghitung luas persegi dengan fungsi 'hitungLuas' dan menyimpan hasilnya
3. Program menghitung keliling persegi dengan fungsi 'hitungKeliling' dan menyimpan hasilnya
4. Tampilkan hasil dari hitung luas dan keliling

Cara kerja dari program tersebut yaitu program meminta user untuk menginputkan sisi dari persegi. Setelah itu program akan menghitung luas dan keliling persegi sesuai dengan rumus dan inputan yang ada. Di akhir program akan menampilkan hasil dari hitungan itu.

## III. UNGUIDED
**1.Minggu Ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kullah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, Iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakan kaltan membantu Jonas? (tidak tentunya ya :p)**

*Masukan : Terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a >= c dan b >= d*

*Keluaran : Terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.*

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / S1IF11-05

package main

import "fmt"

func main() {
	var a, b, c, d int

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

#### Screenshoot Source Code
![Unguided1 carbon](https://github.com/user-attachments/assets/e3c582d7-6e95-4683-ac8e-3e90813521e7)
#### Screenshoot Output
![Unguided1 go](https://github.com/user-attachments/assets/1812fa63-5739-46bc-bf93-8543367e7828)
#### Deskripsi Program
Program diatas merupakan program untuk menghitung permutasi dan kombinasi dari 4 angka yang diinputkan oleh user dengan syarat yang ada. Program akan menampilkan permutasi dan kombinasi dengan pasangan bilangan (a, c) dan (b, d)

Algoritma dari program tersebut sebagai berikut :
1. Pengguna mengisi bilangan a, b, c, dan d 
2. Ambil inputan dari user untuk bilangan a, b, c, dan d 
3. Periksa bilangan dengan syarat yang tersedia
4. Jika syarat terpenuhi program akan menghitung dan menampilkan hasil
5. Jika syarat tidak terpenuhi program tidak akan menampilkan hasil

Cara kerja dari program tersebut yaitu program meminta pengguna untuk memasukkan 4 bilang (a, b, c, dan d). Program akan berjalan jika syaratnya terpenuhi (a>= c) dan (b>= d). Jika syaratnya tidak terpenuhi, program tidak akan berjalan. Pada layar output akan menampilkan hasil dari permutasi dan kombinasi.

**2.Diberikan tiga buah fungsi matematika yaitu f(x) = x ^ 2 , g(x) = x - 2 h(x) = x + 1 Fungsi komposisi (f goh )(x) artinya adalah f(g(h(x))) Tuliskan f(x) g(x) dan h(x) dalam bentuk function.**



*Masukan :  Terdiri dari sebuah bilangan bulat a, b dan a yang dipisahkan oleh spasi.*

*Keluaran : terdiri dari tiga baris. Barls pertama adalah (f g h) (a) , baris kedua (gohof) (b) , dan baris ketiga adalah (hofog) (c).*

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / S1IF11-05

package main

import "fmt"

// Untuk deklarasi fungsi f(x) = x^2
func f(x int) int {
	return x * x
}

// Untuk deklarasi fungsi g(x) = x-2
func g(x int) int {
	return x - 2
}

// Untuk deklarasi fungsi h(x) = x+1
func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	// Inputan bilangan a,b, dan c
	fmt.Print("Bilangan a: ")
	fmt.Scan(&a)

	fmt.Print("Bilangan b: ")
	fmt.Scan(&b)

	fmt.Print("Bilangan c: ")
	fmt.Scan(&c)

	// Rumus menghitung f(g(h(a)))
	f_g_h_a := f(g(h(a)))
	// Rumus menghtiung g(h(f(b)))
	g_h_f_b := g(h(f(b)))
	// Rumus menghitung h(f(g(c)))
	h_f_g_c := h(f(g(c)))

	// Menampilkan hasil dari perhitungan
	fmt.Printf("\n(f ∘ g ∘ h)(%d) = %d\n", a, f_g_h_a)
	fmt.Printf("(g ∘ h ∘ f)(%d) = %d\n", b, g_h_f_b)
	fmt.Printf("(h ∘ f ∘ g)(%d) = %d\n", c, h_f_g_c)
}
```

#### Screenshoot Source Code
![Unguided2 carbon](https://github.com/user-attachments/assets/329424ea-9ce6-4353-ab0f-39b7e8d45357)
#### Screenshoot Output
![Unguided2 go](https://github.com/user-attachments/assets/43d54124-b442-4a44-afe0-72b02ae8e4d0)
#### Deskripsi Program
Program diatas merupakan program untuk menghitung (f ∘ g ∘ h), (g ∘ h ∘ f), dan (h ∘ f ∘ g) . Program meminta user untuk memasukkan nilai f, g, dan h. Hasil dari setiap perhitungan akan ditampilkan di output

Algoritma dari program tersebut sebagai berikut :
1. Deklarasikan fungsi f, g, dan h dengan rumus masing-masing sesuai perintah
2. Deklarasikan variabel untuk menyimpan input dari user
3. Ambil input dari user untuk 3 bilangan
4. Menghitung fungsi (f ∘ g ∘ h), (g ∘ h ∘ f), dan (h ∘ f ∘ g)
5. Tampilkan hasil dari fungsi-fungsinya

Cara kerja dari program tersebut yaitu program meminta user untuk memasukkan tiga bilangan untuk disimpan di variabel. Program menghitung hasil dari fungsi yang tersedia. Pada layar output akan menampilkan hasil dari hitungannya

**3.[Lingkaran) Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius 7. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut.**



*Masukan : Terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.*

*Keluaran : Berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1 ^ * , "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".*

#### Source Code
```go
// Andreas Besar Wibowo
// 2311102198 / S1IF11-05

package main

import "fmt"

// Fungsi untuk menghitung kuadrat jarak dari (x1, y1) dam (x2, y2)
func hitungKuadratJarak(x1, y1, x2, y2 int) int {
	return (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)
}

// Fungsi untuk menentukan apakah titiknya berada di dalam lingkaran
func dalamLingkaran(cx, cy, r, x, y int) bool {
	return hitungKuadratJarak(cx, cy, x, y) <= r*r
}

func main() {
	// Untuk menginputkan koor pusat dan radius pada lingkaran 1
	var cx1, cy1, r1 int
	fmt.Print("Koordinat pusat dan radius pada lingkaran 1 ")
	fmt.Println()
	fmt.Print("Koordinat pusat x: ")
	fmt.Scan(&cx1)
	fmt.Print("Koordinat pusat y: ")
	fmt.Scan(&cy1)
	fmt.Print("Radius lingkaran 1: ")
	fmt.Scan(&r1)

	// Untuk menginputkan koor pusat dan radius pada lingkaran 2
	var cx2, cy2, r2 int
	fmt.Print("Koordinat pusat dan radius pada lingkaran 2 ")
	fmt.Println()
	fmt.Print("Koordinat pusat x: ")
	fmt.Scan(&cx2)
	fmt.Print("Koordinat pusat y: ")
	fmt.Scan(&cy2)
	fmt.Print("Radius lingkaran 2: ")
	fmt.Scan(&r2)

	// Untuk menginputkan koor titk sembarang
	var x, y int
	fmt.Print("Koordinat titik sembarang (x, y)")
	fmt.Println()
	fmt.Print("titik x: ")
	fmt.Scan(&x)
	fmt.Print("titik y: ")
	fmt.Scan(&y)

	// Untuk pengecekan titik
	diLingkaran1 := dalamLingkaran(cx1, cy1, r1, x, y)
	diLingkaran2 := dalamLingkaran(cx2, cy2, r2, x, y)

	// Menampilkan posisi titkk
	fmt.Print("\nPosisi titik: ")
	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik tidak berada di dalam lingkaran 1 dan 2")
	}
}
```

#### Screenshoot Source Code
![Unguided3 carbon](https://github.com/user-attachments/assets/7fbb8c98-f17c-4cea-8745-bec20df11827)
#### Screenshoot Output
![Unguided3(1) go](https://github.com/user-attachments/assets/4c487d49-230c-40c8-b606-a9c897695e05)
![Unguided3(2) go](https://github.com/user-attachments/assets/bda26447-922a-49dc-8062-9b64b5bde6ee)
![Unguided3(3) go](https://github.com/user-attachments/assets/01371583-ffa2-444f-a1e3-1062eaaa0a81)
![Unguided3(4) go](https://github.com/user-attachments/assets/d5820a45-6489-4f21-86a0-7634b9599f8b)
#### Deskripsi Program
Program diatas merupakan program untuk mencari tahu titik yang berada di dalam dua lingkaran yang berbeda. Program meminta user untuk menginputkan koordinat pusat dan radius dari 2 lingkaran, dan koordinat titik sembarang. Program akan memeriksa dan menampilkan posisi titiknya

Algoritma dari program tersebut sebagai berikut :
1. Deklarasikan fungsi untuk menghitung jarak dua titik dan menentukan titik berada dalam lingkaran atau tidak
2. Ambil input dari pengguna untuk koordinat pusat dan radius dari 2 lingkaran
3. Ambil input dari pengguna untuk koordinat titik sembarang
4. Periksa titik berada di dalam lingkaran 1 / 2
5. Tampilkan hasil posisi titik

Cara kerja dari program tersebut yaitu program menghitung jarak kedua titik dan mengecek titik tersebut ada didalam lingkaran atau tidak. Program meminta user untuk menginputkan koordinat pusat dan radius pada 2 lingkaran, dan koordinat titik sembarang. Pada layar output akan menampilkan titik berada di dalam lingkaran atau tidak

## Referensi
[1] MySkill ID. (2021). Function on Golang. Diakses dari https://medium.com/@myskill.id/function-on-golang-27b6577e9cbe

[2] Ahmad Faisal. (2021). Dasar-Dasar Bahasa Pemrograman Golang. Diakses dari https://pta.pilkommedia.org/progress/upload/AhmadFaisal_A1C615001_Dasar-DasarBahasaPemrogramanGolang.pdf

