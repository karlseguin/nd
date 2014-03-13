package nd

import (
	"crypto/rand"
	mr "math/rand"
)

var IntRand func() int
var IntnRand func(int) int
var CryptRand func([]byte) (int, error)

func init() {
	ResetIntRand()
	ResetIntnRand()
	ResetCryptRand()
}

func Seed(s int64) {
	mr.Seed(s)
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

func ResetIntnRand() {
	IntnRand = func(n int) int { return mr.Intn(n) }
}

func ForceIntnRand(forced int) {
	IntnRand = func(_ int) int { return forced }
}
