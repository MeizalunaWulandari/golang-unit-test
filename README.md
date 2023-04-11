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