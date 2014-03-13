package nd

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
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

func Guidv4String() string {
	t := Guidv4()
	return fmt.Sprintf("%x-%x-%x-%x-%x", t[0:4], t[4:6], t[6:8], t[8:10], t[10:])
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
