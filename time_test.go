package nd

import (
	. "github.com/karlseguin/expect"
	"testing"
	"time"
)

type TimeTests struct{}

func Time_Rand(t *testing.T) {
	Expectify(new(TimeTests), t)
}

func (_ TimeTests) NowsDefault() {
	assertNowIsNow()
}

func (_ TimeTests) CanForceNow() {
	expected := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	ForceNow(expected)
	Expect(Now()).To.Equal(expected)
}

func (_ TimeTests) CanForceNowTimestamp() {
	expected := time.Date(2010, time.December, 11, 24, 1, 2, 0, time.UTC)
	ForceNowTimestamp(expected.Unix())
	Expect(Now()).To.Equal(expected)
}

func (_ TimeTests) UTCFollowsNow() {
	loc, _ := time.LoadLocation("EST")
	expected := time.Date(2010, time.December, 11, 24, 1, 2, 0, loc)
	ForceNowTimestamp(expected.Unix())
	actual := UTC()
	Expect(actual).To.Equal(expected)
	Expect(actual.Location().String()).To.Equal("UTC")
}

func (_ TimeTests) CanResetNow() {
	ForceNow(time.Date(2010, time.December, 11, 24, 1, 2, 3, time.UTC))
	ResetNow()
	assertNowIsNow()
}

func assertNowIsNow() {
	start := time.Now()
	actual := Now()
	end := time.Now()
	Expect(actual.Before(start)).To.Equal(false)
	Expect(actual.After(end)).To.Equal(false)
}
