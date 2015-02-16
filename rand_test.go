package nd

import (
	. "github.com/karlseguin/expect"
	"testing"
)

type RandTests struct{}

func Test_Rand(t *testing.T) {
	Expectify(new(RandTests), t)
}

func (_ RandTests) CryptRandReturnsRandomIds() {
	assertCryptRandIsRandom()
}

func (_ RandTests) CanForceACryptRand() {
	expected := []byte{170, 170, 170, 170, 187, 187, 204, 204, 221, 221, 187, 187, 187, 187}
	ForceCryptRand(expected)

	actual := make([]byte, 18)
	n, err := CryptRand(actual)
	Expect(err).To.Equal(nil)
	Expect(n).To.Equal(len(expected))
	Expect(actual[:n]).To.Equal(expected)
}

func (_ RandTests) CanResetCryptRand() {
	ForceCryptRand([]byte{170, 170, 170, 170, 187, 187, 204, 204, 221, 221, 187, 187, 187, 187})
	ResetCryptRand()
	assertCryptRandIsRandom()
}

func (_ RandTests) IntRandReturnsRandomIds() {
	assertIntRandIsRandom()
}

func (_ RandTests) CanForceAnIntRand() {
	ForceIntRand(178)
	Expect(IntRand()).To.Equal(178)
}

func (_ RandTests) CanResetIntRand() {
	ForceIntRand(178)
	ResetIntRand()
	assertIntRandIsRandom()
}

func (_ RandTests) IntnRandReturnsRandomIdsWithInLimits() {
	assertIntnRandIsRandomWithinLimits()
}

func (_ RandTests) CanForceAnIntnRand() {
	ForceIntnRand(42)
	Expect(IntnRand(10)).To.Equal(42)
}

func (_ RandTests) CanResetIntnRand() {
	ForceIntnRand(43)
	ResetIntnRand()
	assertIntnRandIsRandomWithinLimits()
}

func (_ RandTests) CanSeedRandomGeneration() {
	Seed(42)
	Expect(IntRand()).To.Equal(3440579354231278675)
	Expect(IntnRand(10)).To.Equal(7)
}

func assertCryptRandIsRandom() {
	seen := make(map[string]bool, 500)
	b := make([]byte, 18)
	for i := 0; i < 500; i++ {
		CryptRand(b)
		seen[string(b)] = true
	}
	Expect(len(seen)).To.Equal(500)
}

func assertIntRandIsRandom() {
	seen := make(map[int]bool, 500)
	for i := 0; i < 500; i++ {
		seen[IntRand()] = true
	}
	Expect(len(seen)).To.Equal(500)
}

func assertIntnRandIsRandomWithinLimits() {
	seen := make(map[int]bool, 5000)
	for i := 0; i < 5000; i++ {
		seen[IntnRand(50)] = true
	}
	Expect(len(seen)).To.Equal(50)
}
