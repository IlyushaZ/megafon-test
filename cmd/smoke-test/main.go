package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type request struct {
	Phrase string `json:"phrase"`
}

type response struct {
	Phrase string `json:"phrase"`
	Hash   int64  `json:"hash"`
}

type testCase struct {
	request  []request
	response []response
}

func main() {
	tc := testCase{
		request: []request{
			{Phrase: "hello"},
			{Phrase: "world"},
		},
		response: []response{
			{Hash: 3238736544897475342},
			{Hash: 5219289759600458575},
		},
	}

	reqBody, _ := json.Marshal(tc.request)
	resp, err := http.Post("http://localhost:8087/get-phrase-hash", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}

	var respBody []response
	_ = json.NewDecoder(resp.Body).Decode(&respBody)

	for i := range tc.request {
		if respBody[i].Phrase != tc.request[i].Phrase {
			log.Fatalf("expected response's phrase to be %s, %s given", tc.response[0].Phrase, respBody[0].Phrase)
		}

		if respBody[i].Hash != tc.response[i].Hash {
			log.Fatalf("expected response's hash to be %d, %d given", tc.response[0].Hash, respBody[0].Hash)
		}
	}
}
