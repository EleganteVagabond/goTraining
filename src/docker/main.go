package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("go docker tutorial")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
