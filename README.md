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