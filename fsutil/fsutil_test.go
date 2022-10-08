package fsutil

import (
	"testing"
)

func TestWalk(t *testing.T) {
	all, err := DirWalk("./")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", all)
}
