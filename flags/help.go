package flags

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type FlagSection struct {
	Name  string
	Flags []string
}

func PrintAll() error {
	if !flag.Parsed() {
		return fmt.Errorf("Flags did not parsed")
	}

	fmt.Fprintf(os.Stderr, "ARGUMENTS:\n")
	flag.VisitAll(func(g *flag.Flag) {
		fmt.Fprintf(os.Stderr, "  -%s\t\t%s\n", g.Name, g.Usage)
	})

	return nil
}

func PrintHelp(list []*FlagSection) error {
	if !flag.Parsed() {
		return fmt.Errorf("Flags did not parsed")
	}

	for _, s := range list {
		fmt.Fprintf(os.Stderr, "%s:\n", strings.ToUpper(s.Name))
		for _, f := range s.Flags {
			flag.VisitAll(func(g *flag.Flag) {
				if g.Name == f {
					fmt.Fprintf(os.Stderr, "  -%s\t\t%s\n", f, g.Usage)
				}
			})
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
	return nil
}