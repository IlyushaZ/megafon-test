package hasher

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type handlerTestCase struct {
	name         string
	method       string
	body         []byte
	expectedCode int
	expectedBody string
}

func TestHandler_HashPhrases(t *testing.T) {
	tcs := []handlerTestCase{
		{
			name:         "correct request 1",
			method:       http.MethodPost,
			body:         []byte(`[{"phrase":"hello"}]`),
			expectedCode: http.StatusCreated,
			expectedBody: `[{"phrase":"hello","hash":3238736544897475342}]`,
		},
		{
			name:         "correct request 2",
			method:       http.MethodPost,
			body:         []byte(`[{"phrase":"world"}]`),
			expectedCode: http.StatusCreated,
			expectedBody: `[{"phrase":"world","hash":5219289759600458575}]`,
		},
		{
			name:         "incorrect method 1",
			method:       http.MethodGet,
			body:         []byte(`[{"phrase":"hello"}]`),
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
		{
			name:         "incorrect method 2",
			method:       http.MethodPut,
			body:         []byte(`[{"phrase":"hello"}]`),
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "",
		},
		{
			name:         "incorrect body 1",
			method:       http.MethodPost,
			body:         nil,
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: "",
		},
		{
			name:         "incorrect body 2",
			method:       http.MethodPost,
			body:         []byte("hello world"),
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: "",
		},
	}

	handler := http.HandlerFunc(HashPhrases)

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(tc.method, "/get-phrase-hash", bytes.NewBuffer(tc.body))

			handler.ServeHTTP(rec, req)

			if tc.expectedCode != rec.Code {
				t.Errorf(
					"expected response code to be %d, got %d",
					tc.expectedCode, rec.Code,
				)
			}

			if tc.expectedBody != rec.Body.String() {
				t.Errorf(
					"expected response body to be %s, got %s",
					tc.expectedBody, rec.Body.String(),
				)
			}
		})
	}
}
