package flags

import (
	"strconv"
	"strings"
)

type MultiString []string
type MultiWord []string
type MultiInt []int
type MultiIntRange []int

func (s *MultiString) Set(val string) error {
	*s = append(*s, val)
	return nil
}

func (s *MultiString) String() string {
	return ""
}

func (s *MultiWord) Set(val string) error {
	l := strings.Split(val, ",")
	*s = append(*s, l...)
	return nil
}

func (s *MultiWord) String() string {
	return ""
}

func (s *MultiInt) Set(val string) error {
	l := strings.Split(val, ",")
	for _, i := range l {
		n, err := strconv.Atoi(i)
		if err != nil {
			return err
		}

		*s = append(*s, n)
	}
	return nil
}

func (s *MultiInt) String() string {
	return ""
}

func (s *MultiIntRange) Set(val string) error {
	l := strings.Split(val, ",")
	for _, i := range l {
		if strings.Index(i, "-") > 0 {
			ls, err := parseIntRange(i)
			if err != nil {
				return err
			}
			*s = append(*s, ls...)
		} else {
			n, err := strconv.Atoi(i)
			if err != nil {
				return err
			}

			*s = append(*s, n)
		}
	}
	return nil
}

func (s *MultiIntRange) String() string {
	return ""
}