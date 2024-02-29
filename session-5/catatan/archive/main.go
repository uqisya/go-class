package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8080"

func main() {
	// routing
	http.HandleFunc("/", greet)

	// jalankan server dari app yg kita buat menggunakan ListenAndServe
	// hanya menggunakan port aja, belum pakai http.Handler
	http.ListenAndServe(PORT, nil)
}

// responseWrite -> interface dgn berbagai macam method yg bisa digunakan untuk mengirim response balik kepada client yg mengirim request
// Request -> struct yg akan digunakan untuk berbagai macam data request seperti form value, headers, url parameter, dll
func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello!"
	fmt.Fprintf(w, msg)
}
