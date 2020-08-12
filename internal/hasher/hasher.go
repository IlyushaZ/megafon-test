package hasher

import (
	"crypto/sha256"
	"encoding/binary"
)

//easyjson:json
type Body struct {
	Phrase string `json:"phrase"`
	Hash   int64  `json:"hash"`
}

//easyjson:json
type ArrBody []Body

func Hash(body *Body) {
	sha := sha256.Sum256([]byte(body.Phrase))
	body.Hash = int64(binary.BigEndian.Uint64(sha[:8]))
}
