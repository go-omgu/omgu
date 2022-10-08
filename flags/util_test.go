package flags

import "testing"

func TestParseRange(t *testing.T) {
	r, err := parseIntRange("1-10")
	if err != nil {
		t.Fail()
	}
	if len(r) != 10 {
		t.Fail()
	}
	if r[0] != 1 || r[9] != 10 {
		t.Fail()
	}

	r2, err := parseIntRange("9-1")
	if err != nil {
		t.Fail()
	}
	if len(r2) != 9 {
		t.Fail()
	}
	if r2[0] != 1 || r2[8] != 9 {
		t.Fail()
	}
}
