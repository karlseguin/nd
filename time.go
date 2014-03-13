package nd

import (
	"time"
)

var Now func() time.Time
var UTC = func() time.Time {
	return Now().UTC()
}

func init() {
	ResetNow()
}

func ResetNow() {
	Now = func() time.Time { return time.Now() }
}

func ForceNow(t time.Time) {
	Now = func() time.Time { return t }
}
func ForceUTC(t time.Time) {
	ForceNow(t.UTC())
}

func ForceNowTimestamp(timestamp int64) {
	ForceNow(time.Unix(timestamp, 0))
}
