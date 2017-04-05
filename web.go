package main

import (
	"log"
	"net/http"
	"os"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir("static")))
}

func main() {
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
