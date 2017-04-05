package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	//http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello explorer.\n")
	})
}

func main() {
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
