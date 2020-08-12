package hasher

import (
	"github.com/mailru/easyjson"
	"net/http"
	"sync"
)

func HashPhrases(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var reqBody ArrBody
	if err := easyjson.UnmarshalFromReader(r.Body, &reqBody); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(reqBody))

	for i := range reqBody {
		go func(body *Body) {
			defer wg.Done()
			Hash(body)
		}(&reqBody[i])
	}

	wg.Wait()

	resp, err := easyjson.Marshal(reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
