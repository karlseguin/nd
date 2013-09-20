package nd

import (
  "crypto/rand"
)

var CryptRand func([]byte) (int, error)

func init() {
  ResetCryptRand()
}

func ResetCryptRand() {
  CryptRand = func(b []byte) (int, error) { return rand.Read(b) }
}

func ForceCryptRand(forced []byte) {
  CryptRand = func(b []byte) (int, error) {
    n := copy(b, forced)
    return n, nil
  }
}
