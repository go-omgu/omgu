package scanner

import (
	"bufio"
	"bytes"
	"io"
)

type Scanner struct {
	SkipPrefix [][]byte
	r          *bufio.Reader
}

var DefaultPrefixes = [][]byte{
	[]byte("#"),
	[]byte("//"),
}

func NewScanner(r io.Reader) *Scanner {
	s := new(Scanner)
	s.SkipPrefix = DefaultPrefixes
	s.r = bufio.NewReader(r)
	return s
}

func (s *Scanner) Scan() ([]byte, error) {
	l, err := s.r.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	l = l[:len(l)-1]

	if len(l) == 0 {
		return s.Scan()
	}

	for _, p := range s.SkipPrefix {
		if bytes.HasPrefix(l, p) {
			return s.Scan()
		}
	}

	return l, nil
}

func (s *Scanner) ScanString() (string, error) {
	b, err := s.Scan()
	if err != nil {
		return "", err
	}
	return string(b), err
}
