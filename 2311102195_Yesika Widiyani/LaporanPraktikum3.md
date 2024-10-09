<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>
<h2 align="center"><strong>MODUL III</strong></h2>
<h2 align="center"><strong> FUNGSI </strong></h2>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Yesika Widiyani<br>
  2311102195<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

# DASAR TEORI
## FUNGSI
### Dasar Teori: Fungsi dalam Bahasa Pemrograman Go

Fungsi adalah blok kode yang dapat dipanggil untuk menjalankan tugas tertentu. Fungsi digunakan untuk mengelompokkan kode yang bisa diulang, mengurangi redundansi, dan meningkatkan keterbacaan serta pemeliharaan kode. Dalam bahasa pemrograman Go (Golang), fungsi adalah salah satu elemen fundamental. Berikut adalah penjelasan dasar tentang fungsi dalam Go.

#### 1. Definisi Fungsi
Fungsi dalam Go dideklarasikan menggunakan kata kunci `func`, diikuti dengan nama fungsi, parameter (jika ada), tipe nilai kembalian (jika ada), dan tubuh fungsi yang berisi pernyataan-pernyataan yang akan dijalankan ketika fungsi dipanggil.

**Sintaks Umum:**
```go
func namaFungsi(parameter1 tipeData, parameter2 tipeData) tipeKembalian {
    // blok kode
    return nilai
}
```

Contoh fungsi sederhana yang tidak menerima parameter dan tidak mengembalikan nilai:
```go
func sayHello() {
    fmt.Println("Hello, World!")
}
```

#### 2. Fungsi dengan Parameter
Fungsi dapat menerima satu atau lebih parameter yang akan digunakan di dalam blok kode fungsi. Parameter adalah variabel yang dideklarasikan dalam tanda kurung saat mendefinisikan fungsi.

**Contoh:**
```go
func tambah(a int, b int) int {
    return a + b
}
```
Fungsi `tambah` menerima dua parameter `a` dan `b` yang bertipe `int`, dan mengembalikan nilai berupa hasil penjumlahan `a + b`.

#### 3. Fungsi dengan Nilai Kembali
Fungsi dalam Go dapat mengembalikan satu atau lebih nilai menggunakan perintah `return`. Tipe dari nilai yang dikembalikan harus disebutkan setelah tanda kurung parameter dalam deklarasi fungsi.

**Contoh:**
```go
func kali(a int, b int) int {
    return a * b
}
```
Fungsi `kali` menerima dua parameter `a` dan `b`, dan mengembalikan hasil perkalian keduanya dalam bentuk integer (`int`).

#### 4. Fungsi dengan Multiple Return Values
Go mendukung fungsi yang dapat mengembalikan lebih dari satu nilai. Fitur ini berguna untuk mengembalikan nilai utama bersama dengan nilai tambahan seperti error atau status.

**Contoh:**
```go
func bagi(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("tidak bisa dibagi dengan nol")
    }
    return a / b, nil
}
```
Fungsi `bagi` mengembalikan dua nilai: hasil pembagian dan error. Jika `b == 0`, maka fungsi mengembalikan error; jika tidak, mengembalikan hasil bagi.

#### 5. Fungsi Variadic
Fungsi variadic adalah fungsi yang dapat menerima sejumlah parameter yang tidak terbatas. Parameter ini ditulis dengan tiga titik (`...`) di depan tipe datanya. Biasanya digunakan untuk operasi seperti penjumlahan banyak angka.

**Contoh:**
```go
func jumlahSemua(angka ...int) int {
    total := 0
    for _, n := range angka {
        total += n
    }
    return total
}
```
Fungsi `jumlahSemua` dapat menerima beberapa nilai `int` sebagai parameter, dan menjumlahkan semuanya.

#### 6. Fungsi sebagai Nilai
Dalam Go, fungsi adalah "first-class citizens," yang artinya fungsi dapat disimpan dalam variabel, dilewatkan sebagai parameter, atau dikembalikan sebagai nilai dari fungsi lain.

**Contoh:**
```go
func operasi(f func(int, int) int, a int, b int) int {
    return f(a, b)
}

func main() {
    hasil := operasi(tambah, 3, 4)
    fmt.Println(hasil) // Output: 7
}
```
Fungsi `operasi` menerima fungsi lain sebagai parameter (`f`), bersama dengan dua integer, dan menerapkannya.

#### 7. Fungsi Anonim (Anonymous Functions)
Fungsi anonim tidak memiliki nama dan sering digunakan untuk tugas-tugas sementara. Mereka dapat dideklarasikan dan dieksekusi langsung di tempat yang sama.

**Contoh:**
```go
func main() {
    func(nama string) {
        fmt.Println("Halo", nama)
    }("Go")
}
```
Fungsi anonim ini langsung dipanggil setelah dideklarasikan dengan melewatkan parameter `"Go"`.

#### 8. Fungsi Rekursif
Fungsi rekursif adalah fungsi yang memanggil dirinya sendiri. Fungsi ini berguna dalam menyelesaikan masalah yang bisa dipecah menjadi sub-masalah yang lebih kecil, seperti perhitungan faktorial atau Fibonacci.

**Contoh:**
```go
func faktorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * faktorial(n-1)
}
```
Fungsi `faktorial` memanggil dirinya sendiri hingga mencapai kondisi dasar (`n == 0`).

#### 9. Fungsi Defer
Kata kunci `defer` digunakan untuk menunda eksekusi fungsi sampai fungsi yang mengandung perintah `defer` selesai dieksekusi. Ini sering digunakan untuk memastikan pembersihan sumber daya seperti file atau koneksi jaringan.

**Contoh:**
```go
func main() {
    fmt.Println("Mulai")
    defer fmt.Println("Akan dieksekusi terakhir")
    fmt.Println("Selesai")
}
```
Outputnya adalah:
```
Mulai
Selesai
Akan dieksekusi terakhir
```

#### 10. Fungsi Penanganan Error
Dalam Go, penanganan error sering dilakukan dengan menggunakan return value, bukan menggunakan mekanisme exception seperti di bahasa lain. Fungsi sering mengembalikan nilai kedua berupa `error` untuk menunjukkan adanya kesalahan.

**Contoh:**
```go
func bacaFile(namaFile string) (string, error) {
    if namaFile == "" {
        return "", fmt.Errorf("nama file tidak boleh kosong")
    }
    // simulasi pembacaan file
    return "isi file", nil
}
```

------

# GUIDED
## Guided - 1 - 3.1
### Study Case
Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)
Masukan terdiri dari empat buah bilangan asli 𝑎, 𝑏, 𝑐, dan 𝑑 yang dipisahkan oleh spasi, dengan syarat 𝑎 ≥ 𝑐 dan 𝑏 ≥ 𝑑.
Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi 𝒂 terhadap
𝑐, sedangkan baris kedua adalah hasil permutasi dan kombinasi 𝑏 terhadap 𝑑.

Catatan: permutasi (P) dan kombinasi (C) dari 𝑛 terhadap 𝑟 (𝑛 ≥ 𝑟) dapat dihitung dengan menggunakan persamaan berikut!
<br>
![soalgu1](https://github.com/user-attachments/assets/2e18508b-42ec-48f9-888e-0f4d4cc674c4)

### Source Code
![guided1](https://github.com/user-attachments/assets/e38664cb-7165-44b2-8dd9-d114b77074f6)

### ScreenShot Output
![ssg1](https://github.com/user-attachments/assets/0b701ee2-db04-4a4b-b27a-8a735feca48d)

### Deskripsi Program
Program di atas dibuat untuk menghitung **permutasi** dari dua bilangan `a` dan `b`. Permutasi adalah operasi yang menghitung banyaknya cara menyusun `r` objek dari `n` objek yang tersedia, di mana urutan sangat diperhatikan. Program ini meminta input dua bilangan (`a` dan `b`) dari pengguna dan kemudian menghitung permutasi dengan formula:

P(n,r) = n!/(n-r)!

di mana `n!` adalah faktorial dari `n`, yaitu hasil perkalian dari semua bilangan bulat positif dari 1 hingga `n`.

### Algoritma Program

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

### Cara Kerja Program

1. Program dimulai dengan mendeklarasikan dua variabel `a` dan `b` sebagai bilangan bulat (`int`).
2. Program meminta input dua bilangan dari pengguna dengan menggunakan `fmt.Scan(&a, &b)`.
3. Program mengecek apakah nilai `a` lebih besar atau sama dengan `b`. Jika ya, permutasi dihitung dengan `a` sebagai `n` dan `b` sebagai `r`. Jika tidak, posisi `a` dan `b` ditukar.
4. Fungsi `permutasi` dipanggil dengan parameter `n` dan `r`, dan hasilnya dihitung dengan membagi faktorial `n` dengan faktorial `(n - r)`.
5. Setelah hasil permutasi diperoleh, hasil tersebut dicetak ke layar.

## Guided - 2
### Study Case
Membuat program dengan bahasa go untuk mencari sebuah Luas dan Keliling Persegi.
### Source Code
![guided2](https://github.com/user-attachments/assets/308dcbe7-3433-4a16-9268-d8ce2ee56269)

### ScreenShot Output
![guided2](https://github.com/user-attachments/assets/223e57e9-656a-47ae-9550-1297cdca2834)

### Deskripsi Program
Program yang dibuat adalah program sederhana untuk menghitung **luas** dan **keliling persegi**. Pengguna diminta untuk memasukkan panjang sisi persegi, dan program akan menghitung serta menampilkan luas dan keliling berdasarkan panjang sisi tersebut.

**Algoritma Program:**
1. Program pertama-tama meminta input dari pengguna berupa panjang sisi persegi. Input ini disimpan dalam variabel `sisi` yang bertipe `float64` untuk mendukung angka desimal.
2. Setelah input diterima, program menghitung **luas** persegi dengan rumus:
   Luas = sisi x sisi
3. Kemudian, program menghitung **keliling** persegi dengan rumus:
  Keliling = 4 x sisi
4. Setelah perhitungan selesai, hasil **luas** dan **keliling** ditampilkan ke layar dengan format dua angka desimal menggunakan fungsi `fmt.Printf`.

**Cara Kerja Program:**
1. Program dimulai dengan mendeklarasikan variabel `sisi` yang akan menyimpan input dari pengguna. 
2. Program menampilkan pesan "Masukkan panjang sisi persegi:" untuk meminta pengguna memasukkan nilai panjang sisi persegi. Fungsi `fmt.Scan(&sisi)` digunakan untuk membaca input dari pengguna.
3. Setelah nilai sisi dimasukkan, program menghitung luas persegi dengan mengalikan nilai `sisi` dengan dirinya sendiri (sisi * sisi). Hasil perhitungan ini disimpan dalam variabel `luas`.
4. Program juga menghitung keliling persegi dengan mengalikan nilai `sisi` dengan 4 (4 * sisi). Hasilnya disimpan dalam variabel `keliling`.
5. Setelah perhitungan selesai, program menampilkan hasil perhitungan luas dan keliling menggunakan fungsi `fmt.Printf` untuk memformat hasil ke dua angka di belakang koma.
6. Output menampilkan nilai luas dan keliling yang sudah dihitung berdasarkan input panjang sisi yang diberikan oleh pengguna.

------

# UNGUIDED
## Unguided - 1 - 3.2
### Study Case
Diberikan tiga buah fungsi matematika yaitu 𝑓 (𝑥) = 𝑥2 , 𝑔 (𝑥) = 𝑥 − 2 dan ℎ (𝑥) = 𝑥 + 1. Fungsi komposisi (𝑓𝑜𝑔𝑜ℎ)(𝑥) artinya adalah 𝑓(𝑔Gℎ(𝑥)H). Tuliskan 𝑓(𝑥), 𝑔(𝑥) dan ℎ(𝑥) dalam bentuk function.
Masukan terdiri dari sebuah bilangan bulat 𝑎, 𝑏 dan 𝑐 yang dipisahkan oleh spasi.

Keluaran terdiri dari tiga baris. Baris pertama adalah (𝑓𝑜𝑔𝑜ℎ)(𝑎), baris kedua (𝑔𝑜ℎ𝑜𝑓)(𝑏), dan baris ketiga adalah (ℎ𝑜𝑓𝑜𝑔)(𝑐)!

### Source Code
![un1](https://github.com/user-attachments/assets/24d00985-d7a0-4fa1-a5d9-d4bafcdfa767)

### ScreenShot Output
![un1ss](https://github.com/user-attachments/assets/4b63c26c-9701-4ed6-9a8a-83dada1b5d22)

### Deskripsi Program
Program ini bertujuan untuk menghitung komposisi tiga fungsi matematika, yaitu:
- ( f(x) = x^2) (fungsi kuadrat),
- ( g(x) = x - 2) (fungsi pengurangan 2),
- ( h(x) = x + 1) (fungsi penambahan 1).

Diberikan tiga bilangan bulat `a`, `b`, dan `c`, program akan menghitung hasil dari:
1. (f g h)(a)
2. (g h f)(b)
3. (h f g)(c)

Program ini membaca masukan berupa tiga bilangan bulat dan mencetak hasil komposisi fungsi sesuai urutan yang diminta.

**Algoritma:**
1. Program mendefinisikan tiga fungsi matematika:
   - Fungsi `f(x)` yang mengembalikan ( x^2 ).
   - Fungsi `g(x)` yang mengembalikan ( x - 2 ).
   - Fungsi `h(x)` yang mengembalikan ( x + 1 ).
   
2. Program juga mendefinisikan tiga fungsi komposisi:
   - `fgh(x)` untuk menghitung \( f(g(h(x))) \).
   - `ghf(x)` untuk menghitung \( g(h(f(x))) \).
   - `hfg(x)` untuk menghitung \( h(f(g(x))) \).
   
3. Program menerima tiga masukan: `a`, `b`, dan `c`.

4. Untuk setiap nilai `a`, `b`, dan `c`, program menghitung nilai komposisi fungsi sebagai berikut:
   - (f g h)(a) : Komposisi fungsinya adalah f(x) setelah diterapkan g(h(x)).
   - (g h f)(b) : Komposisi fungsinya adalah g(x) setelah diterapkan h(f(x)).
   - (h f g)(c) : Komposisi fungsinya adalah h(x) setelah diterapkan f(g(x)).

5. Hasil komposisi dari masing-masing fungsi dicetak satu per satu sebagai output.

**Cara Kerja Program:**
1. **Inisialisasi fungsi**: Program mendefinisikan tiga fungsi matematika `f(x)`, `g(x)`, dan `h(x)` sebagai function di Go.
   
2. **Input dari pengguna**: Program meminta input tiga bilangan bulat `a`, `b`, dan `c` dari pengguna.

3. **Perhitungan komposisi fungsi**: 
   - Pertama, program menghitung hasil (f g h)(a), di mana `h(x)` ditambahkan dulu, lalu hasilnya dimasukkan ke `g(x)`, dan terakhir dimasukkan ke `f(x)`.
   - Kedua, program menghitung (g h f)(b), di mana urutan fungsinya dimulai dari `f(x)`, lalu `h(x)`, dan terakhir `g(x)`.
   - Ketiga, program menghitung (h f g)(c), dengan urutan komposisi dari `g(x)`, lalu `f(x)`, dan terakhir `h(x)`.

4. **Output hasil komposisi**: Program mencetak hasil dari tiga komposisi fungsi tersebut ke layar.

## Unguided - 2 - 3.3
### Study Case
Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius
𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦)
berdasarkan dua lingkaran tersebut.

Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.
Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

Fungsi untuk menghitung jarak titik (a, b) dan (c, d) dimana rumus jarak adalah:

![soalun3](https://github.com/user-attachments/assets/4e55ec18-1092-433e-91b6-83e4153fb9d8)

dan juga fungsi untuk menentukan posisi sebuah titik sembarang berada di dalam suatu lingkaran atau tidak.

### Source Code
![ungu2](https://github.com/user-attachments/assets/10eb05e4-c6fc-402d-b6e5-fb3fdb66b9d0)

### ScreenShot Output
![ssun2](https://github.com/user-attachments/assets/ff07438b-2d42-489e-aaac-f428a3b56651)

### Deskripsi Program
Program yang dibuat bertujuan untuk menentukan posisi sebuah titik sembarang berdasarkan dua buah lingkaran dengan pusat dan radius tertentu. Posisi titik bisa berada di dalam lingkaran 1, lingkaran 2, kedua lingkaran, atau di luar keduanya.

### Algoritma dan Cara Kerja Program:
1. **Input**: 
   Program akan meminta pengguna memasukkan tiga set data:
   - Koordinat pusat (`cx1, cy1`) dan radius (`r1`) untuk lingkaran 1.
   - Koordinat pusat (`cx2, cy2`) dan radius (`r2`) untuk lingkaran 2.
   - Koordinat titik sembarang (`x, y`).

2. **Fungsi `hitungJarak`**:
   - Fungsi ini menghitung jarak antara dua titik, yaitu titik sembarang dan pusat lingkaran, menggunakan rumus jarak Euclidean:
     Jarak = \sqrt{(x2 - x1)^2 + (y2 - y1)^2}
   - Fungsi ini mengembalikan nilai jarak sebagai float64.

3. **Fungsi `dalamLingkaran`**:
   - Fungsi ini menerima parameter berupa titik sembarang `(x, y)`, pusat lingkaran `(cx, cy)`, dan radius lingkaran `r`.
   - Fungsi ini menggunakan hasil dari fungsi `hitungJarak` untuk mengecek apakah titik berada di dalam lingkaran. Jika jarak antara titik dan pusat lingkaran lebih kecil atau sama dengan radius, titik tersebut berada di dalam lingkaran, dan fungsi mengembalikan nilai `true`.

4. **Logika Utama**:
   - Setelah pengguna memasukkan input, program memanggil fungsi `dalamLingkaran` dua kali:
     - Pertama untuk memeriksa apakah titik berada di dalam lingkaran 1.
     - Kedua untuk memeriksa apakah titik berada di dalam lingkaran 2.
   - Berdasarkan hasil dari kedua pengecekan tersebut, program menentukan posisi titik:
     - Jika titik berada di dalam kedua lingkaran, program mencetak "Titik di dalam lingkaran 1 dan 2".
     - Jika hanya di dalam lingkaran 1, program mencetak "Titik di dalam lingkaran 1".
     - Jika hanya di dalam lingkaran 2, program mencetak "Titik di dalam lingkaran 2".
     - Jika tidak di dalam kedua lingkaran, program mencetak "Titik di luar lingkaran 1 dan 2".

------
