package nd

import (
	"time"
)

var Now func() time.Time

func init() {
	ResetNow()
}

func ResetNow() {
	Now = func() time.Time { return time.Now() }
}

func ForceNow(t time.Time) {
	Now = func() time.Time { return t }
}

func ForceNowTimestamp(timestamp int64) {
	ForceNow(time.Unix(timestamp, 0))
}
