package nd

import (
	"bytes"
	"testing"
)

func TestCryptRandReturnsRandomIds(t *testing.T) {
	assertCryptRandIsRandom(t)
}

func TestCanForceACryptRand(t *testing.T) {
	expected := []byte{170, 170, 170, 170, 187, 187, 204, 204, 221, 221, 187, 187, 187, 187}
	ForceCryptRand(expected)

	actual := make([]byte, 18)
	n, err := CryptRand(actual)
	if err != nil {
		t.Errorf("CryptRand returned an error %v", err)
	}
	if n != len(expected) {
		t.Errorf("CryptRand should have copied %d bytes, but got %d", len(expected), n)
	}
	if bytes.Compare(actual[0:n], expected) != 0 {
		t.Errorf("CryptRand should be %q, got %q", expected, actual[0:n])
	}
}

func TestCanResetCryptRand(t *testing.T) {
	ForceCryptRand([]byte{170, 170, 170, 170, 187, 187, 204, 204, 221, 221, 187, 187, 187, 187})
	ResetCryptRand()
	assertCryptRandIsRandom(t)
}

func TestIntRandReturnsRandomIds(t *testing.T) {
	assertIntRandIsRandom(t)
}

func TestCanForceAnIntRand(t *testing.T) {
	ForceIntRand(178)
	if i := IntRand(); i != 178 {
		t.Errorf("IntRand should be 178, got %d", i)
	}
}

func TestCanResetIntRand(t *testing.T) {
	ForceIntRand(178)
	ResetIntRand()
	assertIntRandIsRandom(t)
}

func TestIntnRandReturnsRandomIdsWithInLimits(t *testing.T) {
	assertIntnRandIsRandomWithinLimits(t)
}

func TestCanForceAnIntnRand(t *testing.T) {
	ForceIntnRand(42)
	if i := IntnRand(10); i != 42 {
		t.Errorf("IntRand should be 42, got %d", i)
	}
}

func TestCanResetIntnRand(t *testing.T) {
	ForceIntnRand(43)
	ResetIntnRand()
	assertIntnRandIsRandomWithinLimits(t)
}

func TestCanSeedRandomGeneration(t *testing.T) {
	Seed(42)
	if n := IntRand(); n != 3440579354231278675 {
		t.Errorf("IntRand should return 3440579354231278675 when seeded with 42, got %d", n)
	}
	if n := IntnRand(10); n != 7 {
		t.Errorf("IntnRand should return 7 when seeded with 42, got %d", n)
	}
}

func assertCryptRandIsRandom(t *testing.T) {
	seen := make(map[string]bool, 500)
	b := make([]byte, 18)
	for i := 0; i < 500; i++ {
		CryptRand(b)
		seen[string(b)] = true
	}
	if n := len(seen); n != 500 {
		t.Errorf("Should have seen 500 unique values, got %d", n)
	}
}

func assertIntRandIsRandom(t *testing.T) {
	seen := make(map[int]bool, 500)
	for i := 0; i < 500; i++ {
		seen[IntRand()] = true
	}
	if n := len(seen); n != 500 {
		t.Errorf("Should have seen 500 unique values, got %d", n)
	}
}

func assertIntnRandIsRandomWithinLimits(t *testing.T) {
	seen := make(map[int]bool, 5000)
	for i := 0; i < 5000; i++ {
		seen[IntnRand(50)] = true
	}
	if n := len(seen); n != 50 {
		t.Errorf("Should have seen 50 unique values, got %d", n)
	}
}
