package nd

import (
	. "github.com/karlseguin/expect"
	"regexp"
	"testing"
)

type GuidTests struct{}

func Test_Guid(t *testing.T) {
	Expectify(new(GuidTests), t)
}

func (_ GuidTests) Guidv4ReturnsRandomGuids() {
	assertGuidv4IsRandom()
}

func (_ GuidTests) Guidv4StringLooksOk() {
	guid := Guidv4String()
	matched, _ := regexp.Match("[\\da-f]{8}\\-[\\da-f]{4}\\-[\\da-f]{4}\\-[\\da-f]{4}\\-[\\da-f]{8}", []byte(guid))
	Expect(matched).To.Equal(true)
}

func (_ GuidTests) CanForceAGuidv4() {
	expected := "aaaaaaaa-bbbb-cccc-dddd-bbbbbbbb"
	expectedBytes := []byte{170, 170, 170, 170, 187, 187, 204, 204, 221, 221, 187, 187, 187, 187}
	ForceGuid(expected)

	Expect(Guidv4()).To.Equal(expectedBytes)
	Expect(Guidv4String()).To.Equal("aaaaaaaa-bbbb-cccc-dddd-bbbbbbbb")
}

func (_ GuidTests) CanResetGuidv4() {
	ForceGuid("aaaaaaaa-bbbb-cccc-dddd-bbbbbbbb")
	ResetGuidv4()
	assertGuidv4IsRandom()
}

func assertGuidv4IsRandom() {
	seen := make(map[string]bool, 500)
	for i := 0; i < 500; i++ {
		guid := Guidv4String()
		if guid[14:15] != "4" {
			Fail("GUID v4 marker is missing")
		}
		seen[guid] = true
	}
	Expect(len(seen)).To.Equal(500)
}
