package banner

import "testing"

func TestVFConstant(t *testing.T) {
	f := VFConstant("1.0")
	if f() != "1.0" {
		t.Fail()
	}
}

func TestVFHash(t *testing.T) {
	t.Log(VFFileHash())
}
