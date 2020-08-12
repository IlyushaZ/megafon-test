package hasher

import (
	"testing"
)

type hasherTestCase struct {
	body         Body
	expectedHash int64
}

func TestSha256_Hash(t *testing.T) {
	tcs := []hasherTestCase{
		{body: Body{Phrase: "hello"}, expectedHash: 3238736544897475342},
		{body: Body{Phrase: "world"}, expectedHash: 5219289759600458575},
		{body: Body{Phrase: "some"}, expectedHash: -6434397223967105402},
		{body: Body{Phrase: "random"}, expectedHash: -6610807752503341226},
		{body: Body{Phrase: "phrases"}, expectedHash: -46529847572792449},
	}

	for _, tc := range tcs {
		t.Run(tc.body.Phrase, func(t *testing.T) {
			Hash(&tc.body)

			if tc.expectedHash != tc.body.Hash {
				t.Errorf("expected hash to be %d, %d given", tc.expectedHash, tc.body.Hash)
			}
		})
	}
}
