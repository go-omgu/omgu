package log

import (
	"os"
	"testing"
)

func TestPrintJson(t *testing.T) {
	v1 := &struct {
		Name string `json:"name"`
		Size int    `json:"size"`
	}{"dog", 13}
	if PrintJSON(os.Stderr, v1) != nil {
		t.Fail()
	}
	if PrintJSONPretty(os.Stderr, v1) != nil {
		t.Fail()
	}
}
