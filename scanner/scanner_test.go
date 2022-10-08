package scanner

import (
	"bytes"
	"testing"
)

var buf1 = []byte("line1\nline2\nline3\n")
var buf2 = []byte("line1\n#line2\n// line3\n")
var buf3 = []byte("line1\n\n#line2\nline3\n")

func TestScanner(t *testing.T) {
	var count int

	s := NewScanner(bytes.NewBuffer(buf1))
	for _, err := s.Scan(); err == nil; _, err = s.Scan() {
		count += 1
	}
	if count != 3 {
		t.Fatal(count)
	}
	count = 0

	s = NewScanner(bytes.NewBuffer(buf2))
	for _, err := s.Scan(); err == nil; _, err = s.Scan() {
		count += 1
	}
	if count != 1 {
		t.Fatal(count)
	}
	count = 0

	s = NewScanner(bytes.NewBuffer(buf3))
	for _, err := s.Scan(); err == nil; _, err = s.Scan() {
		count += 1
	}
	if count != 2 {
		t.Fatal(count)
	}
	count = 0
}
