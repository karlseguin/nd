package nd

import (
  "fmt"
  "strings"
  "crypto/rand"
  "encoding/hex"
)

var Guidv4 func() []byte

func init() {
  ResetGuidv4()
}

func ResetGuidv4() {
  Guidv4 = func() []byte {
    t := make([]byte, 16)
    rand.Read(t)
    return t
  }
}

func Guidv4String() string {
  t := Guidv4()
  return fmt.Sprintf("%x-%x-%x-%x-%x", t[0:4], t[4:6], t[6:8], t[8:10], t[10:])
}

func ForceGuid(guid string) {
  guid = strings.Replace(guid, "-", "", -1)
  g, _ := hex.DecodeString(guid)
  Guidv4 = func() []byte { return g }
}
