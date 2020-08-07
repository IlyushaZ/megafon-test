package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	h := http.NewServeMux()
	h.HandleFunc("/get-phrase-hash", handle)

	s := http.Server{
		Handler: h,
		Addr:    ":80",
	}

	fmt.Println("started server")
	log.Fatal(s.ListenAndServe())
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var body []struct {
		Phrase json.RawMessage `json:"phrase"`
		Hash   int64           `json:"hash"`
	}

	if json.NewDecoder(r.Body).Decode(&body) != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	sha := sha256.New()
	for i := range body {
		sha.Write(body[i].Phrase)
		h := hex.EncodeToString(sha.Sum(nil)[0:8])
		body[i].Hash, _ = strconv.ParseInt(h, 16, 64)

		sha.Reset()
	}

	resp, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
