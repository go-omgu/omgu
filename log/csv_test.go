package log

import (
	"os"
	"testing"
)

type Cat struct {
	Name   string `csv:"name,0"`
	Age    int    `csv:"age,1"`
	bio    string
	Master *Human `csv:"master,2"`
}

type Human struct {
	Name string
}

func (h *Human) GoString() string {
	return "Dr. " + h.Name
}

func TestDigHeaders(t *testing.T) {
	m1 := &Human{"John"}
	s, err := digHeaders(&Cat{"mimi", 13, "", m1})
	if err != nil {
		t.Fail()
	}

	if len(s) != 2 {
		t.Fail()
	}
	t.Log(s)
}

func TestCSVLog(t *testing.T) {
	m1 := &Human{"Lisa"}
	s, err := NewCSVLike(os.Stderr, &Cat{"mimi", 13, "", m1})
	if err != nil {
		t.Fail()
	}

	m2 := &Cat{"dudu", 2, "123", m1}
	m3 := &Cat{"di, xi", 5, "noop", m1}
	s.Append(m2)
	s.Append(m3)
	s.AppendMany([]*Cat{m2, m3})
}
