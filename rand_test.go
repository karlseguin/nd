package nd

import (
  "bytes"
  "testing"
)

func TestCryptRandReturnsRandomIds(t *testing.T) {
  assertCryptRandIsRandom(t)
}

func TestCanForceACryptRand(t *testing.T) {
  expected := []byte{170,170,170,170,187,187,204,204,221,221,187,187,187,187}
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
  ForceCryptRand([]byte{170,170,170,170,187,187,204,204,221,221,187,187,187,187})
  ResetCryptRand()
  assertCryptRandIsRandom(t)
}

func assertCryptRandIsRandom(t *testing.T) {
  seen := make(map[string]bool, 500)
  b := make([]byte, 18)
  for i := 0; i < 500; i++ {
    CryptRand(b)
    seen[string(b)] = true
  }
  if len(seen) != 500 { t.Error("Should have seen 500 unique values") }
}
