package main

import (
	"flag"
	"fmt"

	"gopkg.in/omgu.v0/banner"
	"gopkg.in/omgu.v0/flags"
)

type Opts struct {
	count int
	ports *flags.MultiInt
}

func main() {
	// Banner
	b := &banner.Banner{
		Name:        "omgu",
		Description: "Lightweight useful zero-dependance go util",
		Author:      "xbol0",
		Link:        "https://github.com/xbol0",
		Email:       "xbolo@duck.com",
		VerFunc:     banner.VFConstant("1.0"),
	}

	_ = b.ShowBanner()

	// Flags
	opts := parseFlags()
	_ = flags.PrintAll()
	fmt.Printf("count=%v\n", opts.count)
	fmt.Printf("Ports=%v\n", *opts.ports)
}

func parseFlags() *Opts {
	var opts Opts
	var x flags.MultiInt
	flag.IntVar(&opts.count, "c", 0, "Count")
	flag.Var(&x, "p", "Ports")
	flag.Parse()
	opts.ports = &x
	return &opts
}
