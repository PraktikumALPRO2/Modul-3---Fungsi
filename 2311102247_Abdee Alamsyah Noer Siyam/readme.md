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
  Abdee Alamsyah Noer Siyam
  <br>
  2311102247
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


## <strong> Dasar Teori </strong>

<strong><h2>PENGERTIAN FUNCTION</h2></strong>

Function (fungsi) adalah sebuah unit dari logika program yang dirancang untuk melakukan tugas atau perhitungan tertentu dan dapat dipanggil dari bagian lain dalam program. Di dunia pemrograman, fungsi sering disebut sebagai blok bangunan utama karena membantu dalam menyederhanakan, mengorganisasi, dan meningkatkan keterbacaan kode. Dengan memecah program menjadi fungsi-fungsi kecil, pengembang dapat fokus pada satu tugas spesifik di setiap fungsi, membuat kode lebih modular dan mudah dipelihara.

Dalam konteks bahasa Go (Golang), fungsi memiliki peran yang sangat penting. Bahasa Go mendukung pendekatan prosedural, di mana program terdiri dari urutan instruksi dan alur kerja yang dipisahkan dalam fungsi-fungsi yang dapat dipanggil. Selain itu, Go mendukung konsep first-class function, yang berarti fungsi dalam Go dapat diperlakukan sebagai nilai (value). Ini memungkinkan fungsi untuk disimpan dalam variabel, dipassing sebagai argumen ke fungsi lain, atau dikembalikan dari fungsi.

Fungsi dalam Go memiliki beberapa karakteristik unik dan fitur yang memungkinkan fleksibilitas besar bagi pengembang, seperti multiple return values, fungsi anonim (anonymous function), fungsi variadik, dan closures. Mari kita lihat konsep-konsep ini secara lebih mendalam.



<strong>Struktur Fungsi
</strong>

Dalam Go, fungsi dideklarasikan menggunakan kata kunci func, diikuti oleh nama fungsi, daftar parameter (jika ada), dan tipe data yang dikembalikan (return type). Berikut adalah struktur dasar dari sebuah fungsi:

```go
func namaFungsi(parameter1 tipe1, parameter2 tipe2) returnType {
    // blok kode
    return nilai
}
```

Penjelasan : 
- namaFungsi adalah nama fungsi yang akan digunakan untuk memanggil fungsi tersebut.
- parameter1, parameter2 adalah daftar parameter (input) yang diterima oleh fungsi, diikuti dengan tipe datanya.
- returnType adalah tipe data yang akan dikembalikan oleh fungsi. Fungsi dapat mengembalikan satu atau lebih nilai.
- return digunakan untuk mengembalikan nilai dari fungsi ke pemanggil.

Contoh :

```go
package main
import "fmt"

func tambah(a int, b int) int {
    return a + b
}

func main() {
    hasil := tambah(3, 5)
    fmt.Println("Hasil penjumlahan:", hasil)
}
```

Pada contoh di atas, fungsi tambah menerima dua parameter bertipe int dan mengembalikan nilai bertipe int1

<strong><h2>Parameter dan Argumen</h2></strong>

Parameter adalah variabel yang didefinisikan dalam deklarasi fungsi, sedangkan argumen adalah nilai yang diberikan saat memanggil fungsi. Fungsi dapat memiliki nol, satu, atau beberapa parameter. Berikut adalah contoh penggunaan parameter:

```go 
func cetakPesan(pesan string) {
    fmt.Println(pesan)
}
```
Fungsi di atas menerima satu parameter bertipe string dan mencetak pesan tersebut ke layar23.

<strong><h2>Nilai Kembali (returnType)</h2></strong>

Fungsi dalam Go dapat mengembalikan nilai. Tipe data dari nilai kembali harus ditentukan saat mendeklarasikan fungsi. Jika sebuah fungsi tidak mengembalikan nilai, maka ia disebut sebagai "void function". Contoh fungsi dengan nilai kembali:

```go
func luasPersegi(panjang int, lebar int) int {
    return panjang * lebar
}
```

Fungsi ini mengembalikan hasil perkalian panjang dan lebar

<strong><h2>Fungsi Sebagai Parameter</h2></strong>

Go juga mendukung penggunaan fungsi sebagai parameter. Ini memungkinkan pengembang untuk membuat fungsi yang lebih fleksibel dan reusable. Berikut adalah contoh penerapan fungsi sebagai parameter:

```go
func filter(data []string, callback func(string) bool) []string {
    var result []string
    for _, each := range data {
        if callback(each) {
            result = append(result, each)
        }
    }
    return result
}
```

Dalam contoh ini, callback adalah fungsi yang diteruskan sebagai parameter untuk memfilter data berdasarkan kondisi tertentu41.

## <strong> Guided </strong>

### <h2>GUIDED 1</h2>

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

![image](https://github.com/user-attachments/assets/a8945484-f117-4d2c-b34f-5a08a4211bc5)


Deskripsi Program : 
Program di atas adalah sebuah aplikasi sederhana yang menghitung permutasi dari dua bilangan yang dimasukkan oleh pengguna. Permutasi dalam matematika adalah cara mengatur sejumlah item dalam urutan tertentu, tanpa pengulangan. Program ini menggunakan fungsi `fmt.Scan` untuk membaca dua bilangan integer dari pengguna, kemudian membandingkan kedua angka tersebut. Jika bilangan pertama (`a`) lebih besar atau sama dengan bilangan kedua (`b`), maka program akan menghitung permutasi `P(a, b)` menggunakan fungsi khusus yang telah dibuat. Namun, jika bilangan kedua lebih besar, maka program akan menghitung `P(b, a)`. Program ini memiliki dua fungsi utama: **fungsi faktorial (`faktorial`)** dan **fungsi permutasi (`permutasi`)**. Fungsi `faktorial` menghitung hasil faktorial dari bilangan yang diberikan, di mana faktorial adalah hasil perkalian semua bilangan positif hingga angka tersebut. Fungsi `permutasi` kemudian menggunakan hasil faktorial ini untuk menghitung permutasi dengan rumus $$(P(n, r) = \frac{n!}{(n - r)!})$$, di mana `n` adalah bilangan yang lebih besar, dan `r` adalah bilangan yang lebih kecil. Hasil dari perhitungan permutasi tersebut dicetak di layar.

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

#### Output

![image](https://github.com/user-attachments/assets/8323bf85-bae5-49e7-82f7-198f11f9f90d)

Deskripsi Program : 

Program ini merupakan aplikasi sederhana yang menghitung **luas** dan **keliling** dari sebuah persegi berdasarkan input sisi yang dimasukkan oleh pengguna. Program memanfaatkan dua fungsi utama, yaitu `hitungLuas` dan `hitungKeliling`, yang masing-masing digunakan untuk menghitung luas dan keliling persegi. Fungsi `hitungLuas` menerima parameter berupa panjang sisi dalam tipe data `float64` dan mengembalikan hasil perhitungan luas dengan rumus **luas = sisi × sisi**. Fungsi kedua, `hitungKeliling`, juga menerima parameter sisi dalam tipe data yang sama dan menghitung keliling persegi dengan rumus **keliling = 4 × sisi**. Di dalam fungsi `main`, program meminta pengguna untuk memasukkan nilai sisi persegi melalui `fmt.Scanln`. Setelah itu, program memanggil fungsi `hitungLuas` dan `hitungKeliling` untuk menghitung luas dan keliling, lalu mencetak hasilnya dengan format desimal dua angka di belakang koma menggunakan `fmt.Printf`. Sebagai contoh, jika pengguna memasukkan sisi persegi sebesar `8`, program akan menghitung luas sebagai **8 × 8 = 64** dan keliling sebagai **4 × 8 = 32**. Program ini sangat sederhana namun menunjukkan bagaimana fungsi dapat digunakan untuk memecah tugas perhitungan dalam program menjadi lebih modular dan mudah dipahami.

## <strong> Unguided </strong>

### <h2>UNGUIDED 1</h2>

#### Question

Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)

**Masukan** terdiri dari empat buah bilangan asli 𝑎, 𝑏, 𝑐, dan 𝑑 yang dipisahkan oleh spasi, dengan syarat 𝑎 ≥ 𝑐 dan 𝑏 ≥ 𝑑. 

**Keluaran** terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi 𝒂 terhadap 𝑐, sedangkan baris kedua adalah hasil permutasi dan kombinasi 𝑏 terhadap 𝑑.

**Catatan** permutasi (𝑃) dan kombinasi (𝐶) dari 𝑛 terhadap 𝑟 (𝑛 ≥ 𝑟) dapat dihitung dengan 
menggunakan persamaan berikut!

$$P(n,r) =\frac{n!}{(n-r)!} $$ 
sedangkan $$C(n,r) =\frac{n!}{(n-r)!}$$

#### Source Code

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a, b, c, d int
	a, b, c, d = ambilInput()

	if !validasiInput(a, b, c, d) {
		fmt.Println("Input tidak valid. Pastikan a >= c dan b >= d.")
		return
	}

	tampilkanHasil(a, b, c, d)
}

func ambilInput() (int, int, int, int) {
	var inputA, inputB, inputC, inputD string
	fmt.Scan(&inputA)
	fmt.Scan(&inputB)
	fmt.Scan(&inputC)
	fmt.Scan(&inputD)

	a, _ := strconv.Atoi(inputA)
	b, _ := strconv.Atoi(inputB)
	c, _ := strconv.Atoi(inputC)
	d, _ := strconv.Atoi(inputD)

	return a, b, c, d
}

func validasiInput(a, b, c, d int) bool {
	return a >= c && b >= d
}

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func hitungPermutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func hitungKombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func tampilkanHasil(a, b, c, d int) {
	fmt.Println("Hasil Perhitungan:")
	fmt.Printf("Permutasi P(%d, %d) = %d!/(%d-%d)! = %d\n", a, c, a, a, c, hitungPermutasi(a, c))
	fmt.Printf("Kombinasi C(%d, %d) = %d!/((%d!)*(%d-%d)!) = %d\n", a, c, a, c, a, c, hitungKombinasi(a, c))
	fmt.Printf("Permutasi P(%d, %d) = %d!/(%d-%d)! = %d\n", b, d, b, b, d, hitungPermutasi(b, d))
	fmt.Printf("Kombinasi C(%d, %d) = %d!/((%d!)*(%d-%d)!) = %d\n", b, d, b, d, b, d, hitungKombinasi(b, d))
}
```

#### Output

![image](https://github.com/user-attachments/assets/76ec6470-a5eb-4abe-afa0-748b06ac7924)


Deskripsi Program : 

Program ini adalah aplikasi yang dirancang untuk menerima empat input dari pengguna dan menghitung **permutasi** serta **kombinasi** dari dua pasang angka. Permutasi dan kombinasi adalah konsep dasar dalam matematika kombinatorik yang digunakan untuk menghitung kemungkinan pengaturan dan pemilihan item. Program dimulai dengan fungsi `main`, yang meminta pengguna untuk memasukkan empat angka. Fungsi `ambilInput` digunakan untuk membaca input ini dan mengonversinya dari string ke integer. Setelah input diambil, fungsi `validasiInput` memeriksa apakah angka pertama (`a`) lebih besar atau sama dengan angka ketiga (`c`), dan angka kedua (`b`) lebih besar atau sama dengan angka keempat (`d`). Jika validasi gagal, program memberikan pesan bahwa input tidak valid. Jika validasi berhasil, program melanjutkan untuk menampilkan hasil perhitungan permutasi dan kombinasi melalui fungsi `tampilkanHasil`. Fungsi `faktorial` menghitung hasil faktorial dari suatu bilangan dengan menggunakan loop untuk mengalikan semua bilangan bulat positif hingga bilangan tersebut. Dua fungsi tambahan, `hitungPermutasi` dan `hitungKombinasi`, menggunakan rumus matematika untuk menghitung permutasi dan kombinasi berdasarkan nilai faktorial yang dihitung sebelumnya. Permutasi dihitung menggunakan rumus $$P(n, r) = \frac{n!}{(n - r)!} $$ sementara kombinasi dihitung dengan rumus $$C(n, r) = \frac{n!}{r! \cdot (n - r)!}$$ Hasil dari perhitungan ini ditampilkan di layar dalam format yang jelas, menunjukkan rumus yang digunakan dan hasilnya. Program ini mengilustrasikan penggunaan fungsi dalam Go untuk membuat kode yang modular dan mudah dibaca.

### <h2>UNGUIDED 2</h2>

#### Question

Diberikan tiga buah fungsi matematika yaitu `𝑓 (𝑥) = 𝑥<sup>2</sup>` , `𝑔 (𝑥) = 𝑥 − 2` dan `ℎ (𝑥) = 𝑥 + 1`. Fungsi komposisi `(𝑓 𝑜 𝑔 𝑜 ℎ)(𝑥)` artinya adalah `𝑓(𝑔(ℎ(𝑥)))`.Tuliskan `𝑓 (𝑥), 𝑔 (𝑥), ℎ (𝑥)` bentuk function.
Masukan terdiri dari sebuah bilangan bulat 𝑎 , 𝑏 dan 𝑐 yang dipisahkan oleh spasi.

Keluaran terdiri dari tiga baris. Baris pertama adalah `(𝑓 𝑜 𝑔 𝑜 ℎ)(𝑎)`, baris kedua` (𝑔 𝑜 ℎ 𝑜 𝑓)(𝑏)`, baris ketiga adalah` (ℎ 𝑜𝑓 𝑜 𝑔)(𝑐)` !

#### Source Code

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func fgh(nilai int) int {
	nilai = h(nilai)
	nilai = g(nilai)
	nilai = f(nilai)

	return nilai
}

func ghf(nilai int) int {
	nilai = f(nilai)
	nilai = h(nilai)
	nilai = g(nilai)

	return nilai
}

func hfg(nilai int) int {
	nilai = g(nilai)
	nilai = f(nilai)
	nilai = h(nilai)

	return nilai
}

func main() {
	x := [3]int{}
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")
	
	if len(parts) != 3 {
		fmt.Println("Angka yang dimasukkan tidak sesuai")
		return
	}
	var err1, err2, err3 error
	x[0], err1 = strconv.Atoi(parts[0])
	x[1], err2 = strconv.Atoi(parts[1])
	x[2], err3 = strconv.Atoi(parts[2])

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Angka yang dimasukkan harus berupa angka")
		return
	}

	fgh := fgh(x[0])
	ghf := ghf(x[1])
	hfg := hfg(x[2])

	fmt.Println(fgh)
	fmt.Println(ghf)
	fmt.Println(hfg)
}
```

#### Output

![image](https://github.com/user-attachments/assets/76ec6470-a5eb-4abe-afa0-748b06ac7924)

Deskripsi Program :

Program ini merupakan aplikasi sederhana dalam bahasa Go yang menerima tiga input angka dari pengguna, lalu menerapkan tiga fungsi matematika yang berbeda pada setiap input. Fungsi-fungsi tersebut adalah `f(x)`, `g(x)`, dan `h(x)`, yang masing-masing mengubah nilai input berdasarkan aturan tertentu. Fungsi `f(x)` menghitung kuadrat dari nilai `x` *f(x) = x<sup>2</sup>*, fungsi `g(x)` mengurangi nilai `x` dengan 2  *g(x) = x - 2*, dan fungsi `h(x)` menambahkan 1 ke nilai `x` *h(x) = x + 1*. Program ini juga memiliki tiga fungsi gabungan, yaitu `fgh`, `ghf`, dan `hfg`, yang menerapkan urutan berbeda dari ketiga fungsi tersebut terhadap input yang diberikan. Misalnya, fungsi `fgh(nilai)` menerapkan fungsi `h()`, kemudian `g()`, lalu `f()` terhadap nilai input. Begitu juga untuk fungsi `ghf(nilai)` dan `hfg(nilai)` yang mengaplikasikan fungsi-fungsi ini dalam urutan yang berbeda. Pengguna diharuskan untuk memasukkan tiga angka secara bersamaan, yang kemudian diproses menggunakan fungsi `bufio.NewReader` untuk membaca input dari baris, lalu dikonversi dari string ke integer. Setelah itu, hasil dari setiap kombinasi fungsi diterapkan pada angka input pertama, kedua, dan ketiga, kemudian hasilnya ditampilkan. Jika ada kesalahan input seperti memasukkan non-angka, program akan mengeluarkan pesan kesalahan.

### <h2>UNGUIDED 3</h2>

#### Question

Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦) berdasarkan dua lingkaran tersebut.

**Masukan** terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. 

**Keluaran** berupa string yang menyatakan posisi titik `Titik di dalam lingkaran 1 dan 2`, `Titik di dalam lingkaran 1`, `Titik di dalam lingkaran 2`, atau `Titik di luar lingkaran 1 dan 2`.

Fungsi untuk menghitung jarak titik (a,b) dan (c,d) dimana rumus jarak adalah:
$$ jarak = \sqrt{(a-c)^2 + (b-d)^2} $$

#### Source Code

```go
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

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	l1 := didalam(cx1, cy1, r1, x, y)
	l2 := didalam(cx2, cy2, r2, x, y)

	if l1 && l2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if l1 {
		fmt.Println("Titik di dalam lingkaran 1")
	}else if l2 {
		fmt.Println("Titik di dalam lingkaran 2")
	}else{
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```

#### Output

![image](https://github.com/user-attachments/assets/813ddbcc-91c8-45f4-a786-28b2d2959386)

Deskripsi Program :

Program ini adalah aplikasi sederhana dalam Go yang bertujuan untuk menentukan apakah suatu titik berada di dalam satu atau dua lingkaran, atau di luar keduanya. Program ini bekerja dengan menerima input koordinat pusat dan jari-jari dari dua lingkaran, serta koordinat titik yang ingin diperiksa. Fungsi utama yang digunakan dalam program adalah `jarak` dan `didalam`. Fungsi `jarak` menghitung jarak Euclidean antara dua titik menggunakan rumus $$\sqrt{(a - c)^2 + (b - d)^2}$$ yang diperlukan untuk mengukur jarak antara pusat lingkaran dan titik yang diuji. Sementara itu, fungsi `didalam` memeriksa apakah jarak antara pusat lingkaran dan titik tersebut lebih kecil atau sama dengan jari-jari lingkaran, menandakan bahwa titik tersebut berada di dalam lingkaran. Dalam fungsi `main`, program menerima input berupa koordinat pusat dan jari-jari untuk dua lingkaran, serta koordinat titik yang akan diperiksa. Kemudian, program menggunakan fungsi `didalam` untuk menentukan apakah titik tersebut berada di dalam lingkaran pertama, lingkaran kedua, keduanya, atau di luar kedua lingkaran. Berdasarkan hasil ini, program akan mencetak salah satu dari empat kemungkinan: "Titik di dalam lingkaran 1 dan 2," "Titik di dalam lingkaran 1," "Titik di dalam lingkaran 2," atau "Titik di luar lingkaran 1 dan 2." Program ini memberikan cara sederhana namun efektif untuk menentukan posisi titik relatif terhadap dua lingkaran.

## <strong> Kesimpulan </strong>

Bahasa pemrograman Go, atau Golang, dikenal dengan sintaksis yang sederhana dan efisiensi dalam pengembangan perangkat lunak. Salah satu fitur kunci dari Go adalah penggunaan fungsi (function) yang memungkinkan pengembang untuk menulis kode yang lebih terstruktur dan terorganisir. Fungsi dalam bahasa Go merupakan elemen penting yang membantu dalam pengorganisasian kode serta meningkatkan keefisienan dan keterbacaan. Dengan memahami cara mendeklarasikan dan menggunakan fungsi, serta konsep seperti parameter dan nilai return, programmer dapat menulis aplikasi Go yang lebih efisien dan mudah *easy to maintenance*. Penggunaan fungsi sebagai parameter juga membuka peluang untuk menciptakan kode yang lebih dinamis dan fleksibel.

## <strong> Referensi </strong>

#### [1] Web Development. [2021]. Golang : Pengertian, Fungsi & Rekomendasi Tempat Belajarnya. Diakses melalui website https://majapahit.id/blog/2021/10/21/golang-adalah/

#### [2] Yonata, Jefri. [2021]. Golang: Bahasa Pemrograman Fleksibel dari Google. Diakses melalui website https://www.dewaweb.com/blog/apa-itu-golang/

#### [3] Farid, Achmad. [2024]. Apa Itu Golang? Panduan Singkat untuk Pemula | Exabytes. Diakses melalui website https://www.exabytes.co.id/blog/apa-itu-golang-adalah/

#### [4] Pasar Trainer Blog. [2024]. Apa itu Golang? Ini Pengertian, Fungsi, dan Keunggulannya. Diakses melalui website https://pasartrainer.com/blog/apa-itu-golang-ini-pengertian-fungsi-dan-keunggulannya 
