package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	if err := http.ListenAndServe(":8005", nil); err != nil {
		log.Fatal("Failed to start", err)
	}

}
