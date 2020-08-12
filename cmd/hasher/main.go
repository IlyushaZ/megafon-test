package main

import (
	"github.com/IlyushaZ/megafon-test/internal/hasher"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-phrase-hash", hasher.HashPhrases)

	s := http.Server{
		Handler: mux,
		Addr:    ":80",
	}

	log.Println("starting server on port 80")
	log.Fatal(s.ListenAndServe())
}
