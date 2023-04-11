# Golang Unit Test

## Membuat Unit Test
	Untuk membuat unit test nama file harus berakhiran `_test`
	contohnya `helloWorld_test.go`

	Selain Nama file harus ada `_test` penamaan function juga harus diawali dengan `Test`
	contohnya TestHelloWorl

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