<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL II</strong></h2>
<h2 align="center"><strong> REVIEW STRUKTUR KONTROL </strong></h2>

<br>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Wahyu Hidayat / 2311102178<br>
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

## I. Dasar Teori
### Pengertian Fungsi
Fungsi adalah blok kode yang dapat dijalankan secara independen dalam sebuah program. Fungsi digunakan untuk memecah program menjadi bagian-bagian yang lebih kecil dan lebih terorganisir. Setiap fungsi dirancang untuk melakukan tugas tertentu dan dapat dipanggil dari bagian lain dari program untuk menghindari penulisan ulang kode yang sama. Dengan kata lain, fungsi memberikan cara untuk membuat program lebih modular dan efisien dalam hal pemeliharaan serta penggunaan ulang kode [1]. 

Fungsi juga meningkatkan keterbacaan program karena pengembang dapat memisahkan logika yang kompleks ke dalam beberapa fungsi yang lebih kecil dan lebih mudah dikelola [2].

### Deklarasi Fungsi
Dalam bahasa Go, deklarasi fungsi dimulai dengan kata kunci func diikuti dengan nama fungsi, parameter yang diperlukan (opsional), dan tipe pengembalian (opsional). Fungsi dideklarasikan dengan sintaks sebagai berikut:
```go
func namaFungsi(parameter1 tipe1, parameter2 tipe2) tipePengembalian {
    // blok kode yang akan dieksekusi
    return nilaiPengembalian
}

```
#### Penjelasan:
namaFungsi: Nama yang digunakan untuk memanggil fungsi.

parameter: Nilai masukan yang diterima oleh fungsi, dapat terdiri dari nol atau lebih parameter.

tipePengembalian: Tipe data yang dikembalikan oleh fungsi, bisa berupa tipe data tunggal atau lebih dari satu (multi-return values).

blok kode: Kumpulan perintah yang akan dijalankan ketika fungsi dipanggil.

return: Mengembalikan hasil dari fungsi ke pemanggil (opsional, tergantung tipe pengembalian).

## Contoh Deklarasi Fungsi dalam Golang
Contoh sederhana deklarasi fungsi dalam Golang:

```go
func sapa(nama string) string {
    return "Halo, " + nama
}

```
Pada contoh di atas, fungsi sapa menerima parameter nama bertipe string dan mengembalikan nilai bertipe string yang merupakan pesan sapaan [3].

### Pemanggilan Fungsi
Pemanggilan fungsi adalah cara untuk mengeksekusi kode yang berada di dalam sebuah fungsi. Fungsi yang telah dideklarasikan bisa dipanggil kapan saja selama berada dalam lingkup yang benar. Untuk memanggil fungsi, kita cukup menuliskan nama fungsi diikuti oleh kurung buka dan tutup beserta argumen yang diperlukan (jika ada) [1].

Contoh pemanggilan fungsi dari contoh sebelumnya:

```go
func main() {
    hasil := sapa("Andi") // Memanggil fungsi 'sapa' dengan argumen "Andi"
    fmt.Println(hasil)    // Output: Halo, Andi
}
```
Dalam contoh di atas, fungsi sapa dipanggil dalam fungsi main dengan memberikan nilai "Andi" sebagai argumen. Hasil dari pemanggilan fungsi tersebut disimpan dalam variabel hasil yang kemudian dicetak ke layar [3].

#### Fungsi dengan Nilai Kembali Lebih dari Satu
Bahasa Go mendukung fungsi dengan pengembalian lebih dari satu nilai (multi-return). Ini berguna ketika kita perlu mengembalikan beberapa hasil dari satu fungsi. Contoh fungsi dengan beberapa nilai kembali:

```go
func hitung(a int, b int) (int, int) {
    penjumlahan := a + b
    pengurangan := a - b
    return penjumlahan, pengurangan
}

```
#### Pemanggilan fungsi:
```go
func hitung(a int, b int) (int, int) {
    penjumlahan := a + b
    pengurangan := a - b
    return penjumlahan, pengurangan
}

```
Pada contoh di atas, fungsi hitung mengembalikan dua nilai, yaitu hasil penjumlahan dan pengurangan dari dua bilangan. Saat memanggil fungsi tersebut, kita bisa menangkap kedua nilai pengembaliannya sekaligus [3].






## II. GUIDED
## 1. Program Perhitungan Permutasi Dua Bilangan Bulat

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
#### Screenshoot Source Code
![Screenshot 2024-10-13 165455](https://github.com/user-attachments/assets/876a0dc8-2ef3-44db-a9ae-11e12592d7e3)



#### Screenshoot Output
![Screenshot 2024-10-13 165444](https://github.com/user-attachments/assets/757e7796-17de-4103-ad21-3c2386e5038c)



#### Deskripsi Program
Program ini menghitung permutasi dua bilangan bulat positif, yaitu jumlah cara untuk menyusun sejumlah objek dari total objek yang tersedia. Program meminta dua input bilangan dan menghitung permutasi berdasarkan rumus matematika permutasi. Jika nilai bilangan pertama lebih kecil dari yang kedua, maka keduanya dibalik agar perhitungan tetap valid.

#### Algoritma Program
1. Input dua bilangan bulat.
2. Periksa apakah bilangan pertama lebih besar atau sama dengan bilangan kedua.
3. Jika benar, hitung permutasi menggunakan rumus permutasi.
4. Jika salah, tukar urutan bilangan dan hitung permutasi.
5. Hitung faktorial dari bilangan pertama dan faktorial dari selisih bilangan pertama dan kedua.
6. Bagikan hasil faktorial untuk mendapatkan nilai permutasi.
7. Cetak hasil perhitungan.

#### Cara Kerja
- Fungsi faktorial(n int): Fungsi ini menghitung faktorial dengan mengalikan bilangan dari satu hingga bilangan tersebut.
- Fungsi permutasi(n, r int): Fungsi ini menghitung permutasi dengan membagi faktorial bilangan pertama dengan faktorial selisih antara bilangan pertama dan kedua.
- Proses Input: Program menerima dua bilangan bulat dari pengguna, kemudian memeriksa apakah bilangan pertama lebih besar dari atau sama dengan bilangan kedua.
- Proses Output: Hasil perhitungan permutasi dicetak ke layar berdasarkan input pengguna.


## 2. Pengecekan Urutan Warna 

#### Source Code
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Urutan warna yang benar
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	// Membaca input untuk 5 percobaan
	reader := bufio.NewReader(os.Stdin)
	success := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)

		// Membaca input dari pengguna
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Memisahkan input berdasarkan spasi
		colors := strings.Split(input, " ")

		// Mengecek apakah urutan warna sesuai
		for j := 0; j < 4; j++ {
			if colors[j] != correctOrder[j] {
				success = false
				break
			}
		}

		// Jika ada percobaan yang tidak sesuai, keluar dari loop
		if !success {
			break
		}
	}

	// Menampilkan hasil
	if success {
		fmt.Println("BERHASIL : true")
	} else {
		fmt.Println("BERHASIL : false")
	}
}

```
#### Screenshoot Source Code

![Screenshot 2024-10-05 191216](https://github.com/user-attachments/assets/c1329543-1985-43e3-bf0d-e6f276fe7448)



#### Screenshoot Output

![Screenshot 2024-10-05 191311](https://github.com/user-attachments/assets/7c48ef60-8008-4add-8fdb-835ae82420ef)



#### Deskripsi Program

Kode di atas meminta pengguna untuk memasukkan urutan warna selama lima percobaan dan memeriksa apakah urutan yang dimasukkan sesuai dengan urutan yang benar, yaitu "merah", "kuning", "hijau", dan "ungu". Dengan menggunakan bufio untuk membaca input dari pengguna, program memisahkan input berdasarkan spasi dan membandingkannya dengan urutan warna yang benar. Jika semua percobaan berhasil sesuai urutan, program akan mencetak "BERHASIL : true"; jika tidak, program akan mencetak "BERHASIL : false" dan menghentikan proses. Program ini mengilustrasikan penggunaan kontrol perulangan, manipulasi string, dan pengecekan kondisi dalam pemrograman Go.

## 3.  Penjumlahan Lima Angka 

#### Source Code
```go
package main 

import "fmt"

func main (){

	var a, b, c, d, e int
	var hasil int
	fmt.Scanln(&a, &b, &c, &d, &e)

	hasil = a + b + c + d + e
	fmt.Println("Hasil Penjumlahan ", a,b,c,d,e, "Adalah = ", hasil)
}

```
#### Screenshoot Source Code

![Screenshot 2024-10-05 191628](https://github.com/user-attachments/assets/48e530c2-8c25-439e-a479-c5532f3491f9)



#### Screenshoot Output

![Screenshot 2024-10-05 191833](https://github.com/user-attachments/assets/627573d4-b5ec-4463-b3ca-1e009f5e8f27)



#### Deskripsi Program

Kode di atas menghitung jumlah dari lima angka yang dimasukkan oleh pengguna. Dalam fungsi main, variabel a, b, c, d, dan e dideklarasikan untuk menyimpan input angka, sementara variabel hasil digunakan untuk menyimpan hasil penjumlahan. Program ini menggunakan fmt.Scanln untuk membaca lima angka yang dimasukkan dalam satu baris, kemudian menjumlahkan angka-angka tersebut dan mencetak hasil penjumlahan bersama dengan angka-angka yang dimasukkan oleh pengguna. Program ini menunjukkan penggunaan input, variabel, dan operasi dasar dalam Go.

## 4. Penilaian Indeks Nilai 

#### Source Code
```go
package main

import "fmt"

func main() {
	var nam float32
	var nmk string
	var lagi string

	for {
		//meminta input nilai
		fmt.Printf("Masukkan Nilai: ")
		fmt.Scan(&nam)

		//logika penentuan nilai huruf berdasarkan nilai numerik
		if nam > 80 {
			nmk = "A"
		} else if nam > 72.5 {
			nmk = "B"
		} else if nam > 65 {
			nmk = "C"
		} else if nam > 50 {
			nmk = "D"
		} else if nam > 40 {
			nmk = "E"
		} else {
			nmk = "F"
		}

		//mENAMPILKAN HASIL
		fmt.Printf("Nilai Indeks untuk nilai %.2f adalah  %s\n", nam, nmk)
		fmt.Printf("Jika ingin mengulang (y/n) : ")
		fmt.Scan(&lagi)

		if lagi != "y" {
			break
		}
	}
}

```
#### Screenshoot Source Code
![Screenshot 2024-10-05 192125](https://github.com/user-attachments/assets/d8faad57-8dc7-4edf-b59e-7c42420c538e)



#### Screenshoot Output

![Screenshot 2024-10-05 192245](https://github.com/user-attachments/assets/0035c998-0a16-4605-89e8-a6ef483b5c8f)

#### Deskripsi Program

Kode di atas menentukan indeks nilai huruf berdasarkan input nilai numerik dari pengguna. Dalam fungsi main, program meminta pengguna untuk memasukkan nilai dengan tipe data float32 dan menentukan huruf indeks berdasarkan rentang nilai yang telah ditentukan: "A" untuk nilai di atas 80, "B" untuk 72.5 hingga 80, "C" untuk 65 hingga 72.5, "D" untuk 50 hingga 65, "E" untuk 40 hingga 50, dan "F" untuk nilai di bawah 40. Setelah menentukan indeks, program mencetak hasilnya dan menanyakan pengguna apakah ingin mengulangi proses tersebut. Pengguna dapat memasukkan "y" untuk melanjutkan atau karakter lain untuk menghentikan program. Ini menunjukkan penggunaan struktur kontrol, pengulangan, dan pengolahan input dalam Go.




## III. UNGUIDED
## 1. Program Input Bunga dengan Kata Kunci

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

func main() {
	var (
		nBunga int
		namaBunga string
		pita string
		jumlahBunga int
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("N: ")
	inputN, _ := reader.ReadString('\n')
	inputN = strings.TrimSpace(inputN)
	nBunga, _ = strconv.Atoi(inputN)

	for i := 1; i <= nBunga; i++ {
		fmt.Printf("Bunga %d: ", i)
		namaBunga, _ = reader.ReadString('\n')
		namaBunga = strings.TrimSpace(namaBunga)

		// Hentikan input jika user menginputkan 'SELESAI'
		if namaBunga == "SELESAI" {
			break
		}

		// Tambahkan nama bunga ke pita
		if pita == "" {
			pita = namaBunga
		} else {
			pita = pita + " - " + namaBunga
		}

		// Hitung jumlah bunga
		jumlahBunga++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlahBunga)
}

```
#### Screenshoot Source Code
![Screenshot 2024-10-05 194914](https://github.com/user-attachments/assets/52d86e31-2bba-425a-97f6-53a799f29b00)



#### Screenshoot Output

![Screenshot 2024-10-05 195144](https://github.com/user-attachments/assets/312f87a2-535b-472e-9a2b-6d744e2c2375)


#### Deskripsi Program

Kode di atas menerima input bilangan bulat positif 'N'. Kemudian akan meminta input nama bunga sebanyak 'N' kali. Input nama bunga tersebut akan digabungkan ke dalam variabel 'pita' dengan pemisah spasi dan '-'. Program dimodifikasi agar proses input berhenti ketika user menginput 'SELESAI'. Setelah itu, program akan menampilkan isi pita dan jumlah bunga yang ada di dalamnya.


#### Algoritma Program
1. Jumlah Bunga (N): User memasukkan bilangan bulat positif (selain nol) yang menunjukkan berapa banyak nama bunga yang akan diinput.
2. Nama-nama Bunga:
User memasukkan nama bunga satu per satu sebanyak N kali.
Setiap nama bunga akan digabung menjadi sebuah string (pita) dengan pemisah " - ".
Contoh 1:

Input:
N: 3
Bunga 1: Kertas
Bunga 2: Mawar
Bunga 3: Tulip
Output:
Pita: Kertas - Mawar - Tulip -
Modifikasi Program:

Program dimodifikasi agar proses input nama bunga berhenti ketika user memasukkan kata kunci 'SELESAI'. Setelah itu, program akan menampilkan:

Isi Pita: String yang berisi semua nama bunga yang telah diinput.
Jumlah Bunga: Banyaknya bunga di dalam pita.
Contoh 2:

Input:
Bunga 1: Kertas
Bunga 2: Mawar
Bunga 3: Tulip
Bunga 4: SELESAI
Output:
Pita: Kertas - Mawar - Tulip -
Bunga: 3

#### Cara Kerja Program
1. Program meminta pengguna memasukkan bilangan bulat positif N yang menentukan berapa kali input nama bunga akan diminta.
2. Program kemudian menjalankan loop for sebanyak N kali.
3. Di dalam loop, program meminta pengguna untuk memasukkan nama bunga.
4. Jika pengguna memasukkan 'SELESAI', loop akan dihentikan.
5. Jika tidak, nama bunga yang diinputkan akan ditambahkan ke variabel pita yang akan membentuk string berisi kumpulan nama-nama bunga.
6. Setelah loop selesai, program akan mencetak isi variabel pita dan jumlah_bunga.

## 2. Program Keseimbangan Berat Belanjaan di Dua Kantong

#### Source Code
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var kiri, kanan, total float64
	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&kiri, &kanan)

		total = kiri + kanan
		selisih := math.Abs(kiri - kanan)

		if total > 150 || kiri < 0 || kanan < 0 {
			break
		}

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
	fmt.Println("Proses selesai.")
}

```
#### Screenshoot Source Code

![Screenshot 2024-10-05 201457](https://github.com/user-attachments/assets/fa0c2086-089f-4754-bef3-9056bebbb52f)


#### Screenshoot Output
![Screenshot 2024-10-05 201604](https://github.com/user-attachments/assets/38959b5f-ad57-4ca5-9159-a9c221fa0152)

#### Deskripsi Program

Kode di atas berfungsi untuk mengevaluasi berat belanjaan di dua kantong. Di dalam loop tak terbatas, pengguna diminta untuk memasukkan berat dari dua kantong (kiri dan kanan). Total berat dihitung dengan menjumlahkan kedua kantong, dan selisih antara keduanya dihitung menggunakan fungsi math.Abs. Jika total berat melebihi 150 atau jika salah satu kantong memiliki berat negatif, program akan menghentikan proses. Selain itu, jika selisih berat antara kedua kantong lebih besar atau sama dengan 9, program akan mencetak bahwa sepeda motor Pak Andi akan oleng; jika tidak, akan mencetak bahwa sepeda motor tersebut tidak akan oleng. Setelah keluar dari loop, program mencetak "Proses selesai."


#### Algoritma Program
1. Inisialisasi:

Tetapkan berat_kiri = 0
Tetapkan berat_kanan = 0
Tetapkan total_berat = 0

2. Looping Input dan Pemeriksaan:

Ulangi langkah berikut hingga salah satu kondisi terpenuhi:
Minta pengguna memasukkan berat_kiri dan berat_kanan.
Hitung total_berat = berat_kiri + berat_kanan.
Jika berat_kiri < 0 atau berat_kanan < 0 atau total_berat > 150, maka:
Keluar dari loop.
Jika berat_kiri >= 9 atau berat_kanan >= 9, maka:
Keluar dari loop.
Jika selisih absolut antara berat_kiri dan berat_kanan >= 9, maka:
Tampilkan "true" (motor oleng)
Lainnya:
Tampilkan "false" (motor tidak oleng)
3. Tampilkan Pesan Selesai:

Tampilkan "Proses selesai."

#### Cara Kerja Program
1. Input: Program meminta pengguna memasukkan berat belanjaan untuk kedua kantong terpal. Input berupa bilangan real positif.
2. Perhitungan Selisih: Program menghitung selisih berat antara kedua kantong.
3. Evaluasi:
Jika selisih lebih dari atau sama dengan 9 kg:
Program menampilkan output: "Sepeda motor pak Andi akan oleng: true"
Program kembali ke langkah 1 (meminta input berat lagi).
Jika selisih kurang dari 9 kg:
Program menampilkan output: "Sepeda motor pak Andi akan oleng: false"
Program kembali ke langkah 1 (meminta input berat lagi).
4. Kondisi Berhenti:
Jika pada suatu saat total berat kedua kantong melebihi 150 kg, atau salah satu kantong beratnya negatif, program langsung berhenti dengan pesan "Proses selesai."

## 3. Program Perhitungan Fungsi Matematika dengan Input Variabel

#### Source Code
```go
package main

import (
        "fmt"
)

func main() {
        var k int
        fmt.Print("Nilai K = ")
        fmt.Scan(&k)

        // Hitung nilai f(k)
        fk := float64((4*k+2)*(4*k+2)) / float64((4*k+1)*(4*k+3))

        fmt.Printf("Nilai f(K) = %.10f\n",fk)
}

```
#### Screenshoot Source Code

![Screenshot 2024-10-05 204841](https://github.com/user-attachments/assets/974a672c-6e21-4b38-ac78-8ebe092abe27)


#### Screenshoot Output
![Screenshot 2024-10-05 204914](https://github.com/user-attachments/assets/171c08e3-b6bc-490e-b0d2-60c93055bc7e)



#### Deskripsi Program

Kode di atas menghitung serta menampilkan hasil dari suatu fungsi matematika berdasarkan nilai masukan yang diberikan oleh pengguna. Program akan meminta pengguna untuk memasukkan nilai variabel k, lalu menghitung nilai f berdasarkan rumus tertentu yang melibatkan operasi aritmatika seperti perkalian, penjumlahan, pengkuadratan, dan pembagian. Setelah nilai f dihitung, program menampilkannya dengan presisi hingga 10 angka desimal. Fungsi ini memanfaatkan tipe data float64 untuk memastikan perhitungan dilakukan secara akurat dengan angka desimal.


#### Algoritma Program
1. Mulai
2. Minta pengguna memasukkan nilai k
3. Hitung empat kali k ditambah dua
4. Kuadratkan hasil dari langkah tiga
5. Hitung empat kali k ditambah satu
6. Hitung empat kali k ditambah tiga
7. Kalikan hasil dari langkah lima dengan hasil dari langkah enam
8. Bagi hasil dari langkah empat dengan hasil dari langkah tujuh untuk mendapatkan nilai f
9. Tampilkan hasil nilai f dengan presisi sepuluh angka di belakang koma
10. Selesai

#### Cara Kerja Program
1. Program meminta pengguna memasukkan nilai k dan menyimpannya.
2. Setelah menerima nilai k, program mulai melakukan perhitungan. Langkah pertama adalah mengalikan nilai k dengan empat dan menambahkan dua. Hasilnya kemudian dikuadratkan.
3. Program melanjutkan dengan menghitung empat kali k ditambah satu, dan menghitung juga empat kali k ditambah tiga.
4. Program mengalikan kedua hasil perhitungan tersebut.
5. Hasil kuadrat dari langkah awal kemudian dibagi dengan hasil kali dari langkah sebelumnya.
6. Program menampilkan hasil perhitungan f kepada pengguna dalam bentuk desimal dengan sepuluh angka di belakang koma.
7. Setelah hasil ditampilkan, program selesai menjalankan tugasnya.

## 4. Program Perhitungan Biaya Pengiriman

#### Source Code
```go
package main

import (
	"fmt"
)

// Fungsi hitungBiayaPengiriman mengembalikan 5 nilai
func hitungBiayaPengiriman(beratGrams int) (int, int, int, int, int) {
	// Konversi ke kilogram dan sisa gram
	kg := beratGrams / 1000
	sisaGram := beratGrams % 1000

	// Biaya per kilogram
	biayaPerKg := 10000
	biayaTotal := kg * biayaPerKg

	// Hitung biaya sisa gram
	var biayaSisaGram int
	if sisaGram >= 500 {
		biayaSisaGram = sisaGram * 5 // jika sisa berat lebih dari atau sama dengan 500 gram
	} else if sisaGram > 0 && sisaGram < 500 {
		biayaSisaGram = sisaGram * 15 // jika sisa berat kurang dari 500 gram
	}

	// Total biaya termasuk biaya sisa gram
	totalBiaya := biayaTotal + biayaSisaGram

	// Jika total berat lebih dari 10kg, gratis biaya sisa gram
	if kg >= 10 {
		biayaSisaGram = 0
		totalBiaya = kg * biayaPerKg
	}

	// Mengembalikan 5 nilai: kg, sisaGram, biayaTotal, biayaSisaGram, totalBiaya
	return kg, sisaGram, biayaTotal, biayaSisaGram, totalBiaya
}

func main() {
	// Input berat parsel (dalam gram)
	var beratGrams int
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&beratGrams)

	// Hitung detail dan total biaya
	kg, sisaGram, biayaTotal, biayaSisaGram, totalBiaya := hitungBiayaPengiriman(beratGrams)

	// Tampilkan hasil sesuai format yang diminta
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaTotal, biayaSisaGram)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}

```
#### Screenshoot Source Code
![Screenshot 2024-10-06 153826](https://github.com/user-attachments/assets/93bb5e15-63c9-4dc9-99d4-d3a08e3d3566)



#### Screenshoot Output
![Screenshot 2024-10-06 153844](https://github.com/user-attachments/assets/2817ff3d-b84b-4878-a442-b0177ec995ac)

#### Deskripsi Program

Kode ini menghitung biaya pengiriman berdasarkan berat parsel dalam gram. Pertama, berat dikonversi menjadi kilogram dan sisa gram. Biaya dasar pengiriman dihitung berdasarkan berat dalam kilogram dengan tarif Rp. 10.000 per kilogram. Jika sisa gram lebih dari atau sama dengan 500, biaya tambahan dihitung Rp. 5 per gram; jika kurang dari 500 gram, dikenakan Rp. 15 per gram. Jika berat total melebihi 10 kg, biaya sisa gram dibebaskan. Program menampilkan detail berat, rincian biaya, dan total biaya pengiriman sesuai format yang diminta.


#### Algoritma Program
1. Input: Masukkan berat parsel dalam gram.
2. Konversi berat:
Hitung berat dalam kilogram (kg) dengan membagi berat parsel (gram) dengan 1000.
Hitung sisa gram dengan mengambil modulus 1000 dari berat parsel.
3. Hitung biaya pengiriman:
Setiap 1 kg dikenakan biaya Rp. 10.000.
Jika sisa gram lebih dari atau sama dengan 500:
Tambahkan biaya Rp. 5 per gram untuk sisa gram.
Jika sisa gram kurang dari 500:
Tambahkan biaya Rp. 15 per gram untuk sisa gram.
4. Kondisi khusus:
Jika total berat (kg) lebih dari atau sama dengan 10 kg:
Biaya sisa gram gratis.
5. Output:
Tampilkan berat dalam format kg dan gram.
Tampilkan biaya dasar (untuk kg) dan biaya tambahan (untuk sisa gram).
Tampilkan total biaya pengiriman.

#### Cara Kerja Program
1. Input berat parsel dalam satuan gram
2. Konversi berat
Kilogram dihitung dengan membagi berat parsel dengan 1000
Sisa gram dihitung dengan mengambil sisa pembagian dari 1000
3. Hitung biaya pengiriman
Biaya per kilogram adalah Rp. 10.000
Jika sisa gram lebih dari atau sama dengan 500, biaya tambahan adalah Rp. 5 per gram
Jika sisa gram kurang dari 500, biaya tambahan adalah Rp. 15 per gram
4. Total biaya pengiriman dihitung sebagai penjumlahan biaya dasar dan biaya sisa gram
5. Program menampilkan hasil berupa detail berat dalam kg dan gram, biaya dasar dan tambahan, serta total biaya
## 5. Program mencari faktor dari sebuah bilangan bulat positif dan prima

#### Source Code
```go
package main

import (
	"fmt"
)

// Function to get all factors of a number
func getFactors(n int) []int {
	var factors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

// Function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	// A number is prime if it has only two divisors: 1 and itself
	return len(getFactors(n)) == 2
}

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scanln(&b)

	// Get the factors
	factors := getFactors(b)
	fmt.Printf("Faktor: ")
	for _, f := range factors {
		fmt.Printf("%d ", f)
	}
	fmt.Println()

	// Check if the number is prime
	if isPrime(b) {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}


```
#### Screenshoot Source Code

![Screenshot 2024-10-06 155431](https://github.com/user-attachments/assets/08024ab4-0d18-431e-9b32-9654ce1dfb36)



#### Screenshoot Output

![Screenshot 2024-10-06 155455](https://github.com/user-attachments/assets/0754e9f3-9a37-48b5-8d8c-036d5cd1612f)


#### Deskripsi Program

Kode di atas menerima input berupa bilangan bulat positif dan mencari semua faktornya. Setelah itu, program akan menentukan apakah bilangan tersebut adalah bilangan prima, yaitu memiliki tepat dua faktor: 1 dan dirinya sendiri. Dengan menggunakan fungsi untuk menghitung faktor dan memeriksa keprimaan, program ini memberikan output yang jelas, menampilkan faktor-faktor dan status keprimaan dari bilangan yang dimasukkan.


#### Algoritma Program
1. Mulai
2. Input: Minta pengguna untuk memasukkan sebuah bilangan bulat positif b.
3. Inisialisasi:
Buat sebuah daftar kosong untuk menyimpan faktor-faktor dari b.
Inisialisasi variabel jumlahFaktor dengan 0.
4. Cari Faktor:
Untuk setiap bilangan i dari 1 hingga b:
Jika b habis dibagi i (yaitu b % i == 0):
Tambahkan i ke daftar faktor.
Tambahkan 1 ke jumlahFaktor.
5. Tampilkan Faktor:
Cetak semua faktor yang terdapat dalam daftar.
6. Periksa Keprimaan:
Jika jumlahFaktor sama dengan 2, maka:
Bilangan b adalah bilangan prima.
Jika tidak:
Bilangan b bukan bilangan prima.
7. Tampilkan Hasil:
Cetak status keprimaan dari b.
8. Selesai.
#### Cara Kerja Program
1. Input: Pengguna memasukkan 12.
2. Inisialisasi: Program menyiapkan daftar kosong untuk faktor dan variabel untuk menghitung jumlah faktor.
3. Mencari Faktor:
Program memeriksa angka dari 1 hingga 12:
Menemukan faktor: 1, 2, 3, 4, 6, 12.
4. Menampilkan Faktor: Program mencetak semua faktor yang ditemukan.
5. Memeriksa Keprimaan:
Mengetahui bahwa 12 memiliki 6 faktor, program menyimpulkan bahwa 12 bukan bilangan prima.
6. Output: Program mencetak faktor dan status prima: Faktor: 1 2 3 4 6 12, Prima: false.

## Referensi 
[1] A. A. Demailly, "Efficient Function Calls and Modular Programming in Go," Journal of Software Engineering, vol. 34, no. 2, pp. 145-153, 2022.
[2] R. C. Martin, Clean Code: A Handbook of Agile Software Craftsmanship, Pearson Education, 2009.
[3] A. Donovan and B. W. Kernighan, The Go Programming Language, Addison-Wesley, 2015.
------
