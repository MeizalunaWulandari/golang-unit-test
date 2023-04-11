# Golang Unit Test

## Membuat Unit Test
	Untuk membuat unit test nama file harus berakhiran `_test`
	contohnya `helloWorld_test.go`
	Selain Nama file harus ada `_test` penamaan function juga harus diawali dengan `Test`
	contohnya TestHelloWorld
	function unit test memerlukan parameter pointer `testing.T`
	`TestNamaFunction(t *testing.T)`

## Menjalankan Unit Test
	Untuk mejalankan unit test hanya perlu menjalankan perintah `go test`
	Untuk mejalankan unit test dengan disertain unit test yang dijalankan bisa bisa dengan perintah `go test -v`
	Dan Untuk menjalan spesifik function unit test bisa dengan perintah `go test -v -run=TestFunction`

## Menggagalkan Unit Test
	Untuk menggagalkan Unit Test bisa dengan: 
		t.Fail = Jika terjadi gagal akan melanjutkan ke kode progmaram selanjutnya
		t.FailNow = Akan Menghentikan program
		t.Error = Menggagalkan Unit test dengan log dan secara otomatis memanggil t.Fail
		t.Fatal = Menggagalkan Unit test dengan log dan secara otomatis memanggil t.FailNow

## Assertion
	Golang tidak menyediakan package Assertion, sehingga diperlukan library tambahan
	library yang digunakan yaitu Testify
	https://github.com/stretchr/testify
	go get -u github.com/stretchr/testify

	Testify menyediakan package untuk assertion, yaitu assert dan require
	saat menggunakan assert, jika pengecekan gagal, maka assert akan memanggil Fail(), artinya eksikusi unit test akan tetap dilanjutkan
	Sedangkan require jika pengecekkan gagal dia akan, maka require akan memanggil FailNow, sehingga eksekusi unit test tidak dilanjutkan

## Membatalkan Unit Test
	Terkadang unit test tidak dijalankan di kondisi tertentu
	untuk membatalkan unit test bisa dengan function `t.Skip`

## Before dan After Test
	Biasanya dalam unit test, kadang kita ingin melakukan sesuatu sebelum dan sesudah unit test di eksekusi
	di golang terdapat fitur bernama testing.M
	fitur ini digunakan untuk mengeksekusi unit test, namun hal ini juga bisa digunakan sebagai before dan after di unit test

	Untuk mengeksekusi unit test, kita cukup membuat sebuah function bernama TestMain
	denan Parameter testing.M

	Jika terdapat function TestMain, maka secara otomatis golang akan mengeksekusi 
	function  ini setiap kali akan menjalankan unit test disebuah package
	dengan function ini kita bisa mengatur Before dan After unit test sesuai yang kita mau
	perlu diingat function TestMain itu hanya di eksekusi sekali per golang package, bukan tiap function unit test

## Sub-Test
	Golang mendukung fitur pembuatan function unit test di dalam function unit test
	hal ini disebut sub test
	Untuk membuat sub test kita bisa menggunakan function Run()
### Menjalankan Hanya Sub Test
	Untuk hanya menjalankan salah satu sub test bisa dengan perintah go test -v -run=TestNamaFunction/NamaSubTest
	Untuk menjalankan semua sub test bisa dengan perintah go test -run=/NamaSubTest

## Table Test
	Dengan Sub Test kita dapat membuat test secara dinamis 
	dan fitur sub test ini biasa gunakan untuk membuat test dengan konsep table test
	Table test yaitu dimana kita menyediakan data berupa slice yang berisi parameter dan ekspektasi dari hasil unit test
	Lalu slice tersebut kita iterasi menggunakan sub test

## Mock
	Mock adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil , dia akan menghasilkan data data yang sudah kita program diawal
	Mock adalah salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari suatu object yang memang sulit untuk di testing
	Misalnya kita ingin unit testing, namun ternyata ada kode program yang harus memanggil API Call ke third party service, Hal ini sangat sulit untuk di test, karena unit testing kita harus selalu memanggil third party service, dan belum tentu responsenya sesuai dengan apa yang kita mau
	Pada kasus seperti ini, cocok sekali untuk menggunakan mock object
	Golang sendiri tidak menyediakan mock, Namun kita dapat menggunakan Testify untuk membuat mock object
	Perlu diperhatikan , jika desain kode program kita jelek, akan sulit untuk melakukan mocking, jadi pastikan kita melakukan pembuatan desain kode program dengan baik

## Benchmark
	Selain Unit test, golang testing package juga mendukung melakukan benchmark
	Benchmark adalah mekanisme menghitung kecepatan performa kode aplikasi kita
	Benchmark di golang dilakukan dengan cara secara otomatis melakukan iterasi kode yang kita panggil berkali-kali
	kita tidak perlu melakukan jumlah iterasi dan lamanya, karena itu sudah diatur oleh testing.B
	bawaan dari testing package
### testing.B
	testing.B adalah struct yang digunakan untuk melakukan benchmark
	testing.B mirip dengan testing.T, terdapat function Fail(), FailNow(), Error(), Fatal() dll
	yang membedakan, ada atribut dan function tambahan yang digunakan untuk melakukan benchmark 
	Salah satunya adalah atribut N, ini ini digunakan untuk melakukan total iterasi sebuah benchmark
### Cara kerja Benchmark
	Cara kerja benchmark sangat sederhana yaitu hanya perlu membuat perulangan sejumlah atribut N
	nanti secara otomatis golang akan melakukan eksekusi sejumlah perulangan yang ditentukan secara otomatis
	, lalu mendeteksi berapa lama proses tersebut berjalan, dan simpulkan perforna benchmarknya dalam satuan waktu
### Membuat Benchmark
	Mirip seperti unit test, untuk benchmark pun, di golang sudah di tentukan nama functionnya
	harus diawali dengan kata Benchmark, misalnya BenchmarkHelloWorld
	Selain itu harus memiliki parameter (b *testing.B)
	Dan tidak boleh mengembalikan return value
	untuk nama file benchmark sama seperti unit test, diakhiri dengan _test misalnya hello_world_test.go
### Menjalankan Benchmark
	Untuk menjalankan seluruh benchmark di module, kita bisa menggunakan perintah 
	go test -v -bench=.
	Namun jika menggukan perintah tersebut unit test juga akan dijalankan 
	Jika hanya ingin menjalankan benchmark tanpa unit test bisa menggunakan perintah 
	go test -v -run=NotMatchUntiTest -bench=.
	kode diatas selain menjalankan benchmark, akan menjalankan semua benchmark
	Jika hanya ingin menjalankan salah satu benchmark dapat menggunakan perintah
	go test -v -run=NotMatchUntiTest -bench=BenchmarkTest
	Jika menjalankan benchmark di root module dan ingin semua module dijalankan , bisa menggunakan perintah
	go test -v -bench=. ./...

## Sub Benchmark
	Sama seperti tesing.T, di testing.B juga kita bisa membuat sub benchmark menggunakan function Run()

## Table Benchmark
	Table benchmark digunakan agar bisa memudahkan untuk melakukan performance test dengan kombinasi data berbeda-beda tanpa harus membuat banyak benchmark function