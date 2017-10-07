package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
