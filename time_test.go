package nd

import (
	"testing"
	"time"
)

func TestNowsDefault(t *testing.T) {
	assertNowIsNow(t)
}

func TestCanForceNow(t *testing.T) {
	expected := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	ForceNow(expected)
	actual := Now()
	if actual.Equal(expected) == false {
		t.Errorf("Now should equal %q but equals %q", expected, actual)
	}
}

func TestCanForceNowTimestamp(t *testing.T) {
	expected := time.Date(2010, time.December, 11, 24, 1, 2, 0, time.UTC)
	ForceNowTimestamp(expected.Unix())
	actual := Now()
	if actual.Equal(expected) == false {
		t.Errorf("Now should equal %q but equals %q", expected, actual)
	}
}

func TestCanResetNow(t *testing.T) {
	ForceNow(time.Date(2010, time.December, 11, 24, 1, 2, 3, time.UTC))
	ResetNow()
	assertNowIsNow(t)
}

func assertNowIsNow(t *testing.T) {
	start := time.Now()
	actual := Now()
	end := time.Now()
	if actual.Before(start) {
		t.Errorf("time should not be before %q", start)
	}
	if actual.After(end) {
		t.Errorf("time should not be after %q", end)
	}
}
