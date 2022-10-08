package flags

import (
	"flag"
	"testing"
)

func TestFlags(t *testing.T) {
	flag.String("n", "", "the name")
	flag.Parse()
	g := &FlagSection{Name: "general", Flags: []string{"n"}}
	g2 := &FlagSection{Name: "scan", Flags: []string{"n"}}
	PrintHelp([]*FlagSection{g, g2})
}
