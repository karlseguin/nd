package nd

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
)

var Guidv4 func() []byte

func init() {
	ResetGuidv4()
}

func ResetGuidv4() {
	Guidv4 = func() []byte {
		t := make([]byte, 16)
		rand.Read(t)
		t[6] = t[6]&0xF | 64
		return t
	}
}

func LockGuid() string {
	id := Guidv4String()
	ForceGuid(id)
	return id
}

func Guidv4String() string {
	t := Guidv4()
	tmp := make([]byte, 36)
	hex.Encode(tmp[:8], t[:4])
	tmp[8] = '-'
	hex.Encode(tmp[9:], t[4:6])
	tmp[13] = '-'
	hex.Encode(tmp[14:], t[6:8])
	tmp[18] = '-'
	hex.Encode(tmp[19:], t[8:10])
	tmp[23] = '-'
	hex.Encode(tmp[24:], t[10:])
	return string(tmp)
}

func ForceGuid(guid string) error {
	if len(guid) < 28 {
		return errors.New("Length of the GUID should be at least 28 hexadecimal characters")
	}
	guid = strings.Replace(guid, "-", "", -1)
	g, err := hex.DecodeString(guid)
	if err != nil {
		return err
	}
	Guidv4 = func() []byte { return g }
	return nil
}
