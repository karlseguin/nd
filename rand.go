package nd

import (
  mr "math/rand"
  "crypto/rand"
)

var CryptRand func([]byte) (int, error)
var IntRand func() int

func init() {
  ResetCryptRand()
  ResetIntRand()
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

func ResetIntRand() {
  IntRand = func() int { return mr.Int() }
}

func ForceIntRand(forced int) {
  IntRand = func() int { return forced }
}
