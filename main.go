package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello worldz")
		d, _ := io.ReadAll(r.Body)
		log.Printf("Data %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye!")
	})
	http.ListenAndServe(":9090", nil)
}
